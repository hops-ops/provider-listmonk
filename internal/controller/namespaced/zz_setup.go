// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	user "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/identity/user"
	userrole "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/identity/userrole"
	providerconfig "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/providerconfig"
	appsettings "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/settings/appsettings"
	securitysettings "github.com/hops-ops/provider-listmonk/internal/controller/namespaced/settings/securitysettings"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		user.Setup,
		userrole.Setup,
		providerconfig.Setup,
		appsettings.Setup,
		securitysettings.Setup,
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
		user.SetupGated,
		userrole.SetupGated,
		providerconfig.SetupGated,
		appsettings.SetupGated,
		securitysettings.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
