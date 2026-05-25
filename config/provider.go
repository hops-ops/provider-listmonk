/*
Copyright 2024 The Crossplane Authors.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "listmonk"
	modulePath     = "github.com/hops-ops/provider-listmonk"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// Listmonk's resource ID semantics split two ways:
//
//   - `listmonk_user_role` + `listmonk_user`: server-assigned numeric id
//     at Create time. No caller-supplied external name accepted.
//   - `listmonk_security_settings` + `listmonk_app_settings`: SINGLETON
//     per-instance — sentinel string IDs (`security`, `app`).
//
// Both use IdentifierFromProvider: upjet passes the TF ID through
// unchanged so MRs get `crossplane.io/external-name` reflecting the
// server-assigned value (or sentinel for singletons).
var ExternalNameConfigs = map[string]ujconfig.ExternalName{
	"listmonk_security_settings": ujconfig.IdentifierFromProvider,
	"listmonk_app_settings":      ujconfig.IdentifierFromProvider,
	"listmonk_user_role":         ujconfig.IdentifierFromProvider,
	"listmonk_user":              ujconfig.IdentifierFromProvider,
}

// ExternalNameConfigurations + ExternalNameConfigured are defined in
// config/external_name.go so the same helpers are reused for both the
// cluster-scoped and namespaced providers.

func newProvider(rootGroup string) *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	// Settings group: singleton-per-instance settings rows (security.*,
	// app.*). The chart-side post-install hook mints the
	// crossplane-provider api user; this group is for declarative
	// reconciliation of the settings table.
	pc.AddResourceConfigurator("listmonk_security_settings", func(r *ujconfig.Resource) {
		r.ShortGroup = "settings"
		r.Kind = "SecuritySettings"
	})

	pc.AddResourceConfigurator("listmonk_app_settings", func(r *ujconfig.Resource) {
		r.ShortGroup = "settings"
		r.Kind = "AppSettings"
	})

	// Identity group: user roles + users. The bootstrap
	// `crossplane-provider` api user that THIS provider authenticates
	// AS is out-of-band (chart's post-install hook + ESO PushSecret —
	// it can't be a User MR because of the chicken-and-egg dependency
	// on a ProviderConfig to authenticate against).
	pc.AddResourceConfigurator("listmonk_user_role", func(r *ujconfig.Resource) {
		r.ShortGroup = "identity"
		r.Kind = "UserRole"
	})

	pc.AddResourceConfigurator("listmonk_user", func(r *ujconfig.Resource) {
		r.ShortGroup = "identity"
		r.Kind = "User"
		// user.user_role_id references user_role.id — upjet surfaces a
		// typed UserRoleRef on the MR alongside the literal numeric
		// field so consumers can compose via metadata.name.
		r.References["user_role_id"] = ujconfig.Reference{
			TerraformName: "listmonk_user_role",
		}
	})

	pc.ConfigureResources()
	return pc
}

// GetProvider returns cluster-scoped provider configuration.
func GetProvider() *ujconfig.Provider {
	return newProvider("listmonk.crossplane.io")
}

// GetProviderNamespaced returns namespaced MR provider configuration
// (Crossplane v2).
func GetProviderNamespaced() *ujconfig.Provider {
	return newProvider("listmonk.m.crossplane.io")
}
