/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	dashboard "github.com/crossplane-contrib/provider-jet-dynatrace/internal/controller/dynatrace/dashboard"
	environment "github.com/crossplane-contrib/provider-jet-dynatrace/internal/controller/dynatrace/environment"
	notification "github.com/crossplane-contrib/provider-jet-dynatrace/internal/controller/dynatrace/notification"
	zone "github.com/crossplane-contrib/provider-jet-dynatrace/internal/controller/management/zone"
	providerconfig "github.com/crossplane-contrib/provider-jet-dynatrace/internal/controller/providerconfig"
	group "github.com/crossplane-contrib/provider-jet-dynatrace/internal/controller/user/group"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dashboard.Setup,
		environment.Setup,
		notification.Setup,
		zone.Setup,
		providerconfig.Setup,
		group.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
