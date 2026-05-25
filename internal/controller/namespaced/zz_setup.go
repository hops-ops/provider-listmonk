// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	client "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/client/client"
	organization "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/organization/organization"
	ssoconfig "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/organization/ssoconfig"
	project "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/project/project"
	providerconfig "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/providerconfig"
	reference "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/reference/reference"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		client.Setup,
		organization.Setup,
		ssoconfig.Setup,
		project.Setup,
		providerconfig.Setup,
		reference.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		client.SetupGated,
		organization.SetupGated,
		ssoconfig.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
		reference.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
