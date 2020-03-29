/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	bsutil "sigs.k8s.io/cluster-api/bootstrap/util"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	bootstrapv1alpha1 "github.com/balchua/cluster-api-bootstrap-provider-microk8s/api/v1alpha1"
)

// Microk8sConfigReconciler reconciles a Microk8sConfig object
type Microk8sConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

type Scope struct {
	logr.Logger
	Config      *bootstrapv1alpha1.Microk8sConfig
	ConfigOwner *bsutil.ConfigOwner
	Cluster     *clusterv1.Cluster
}

// +kubebuilder:rbac:groups=bootstrap.cluster.x-k8s.io,resources=microk8sconfigs;microk8sconfigs/status,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=bootstrap.cluster.x-k8s.io,resources=microk8sconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=clusters;clusters/status;machines;machines/status,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=secrets;events;configmaps,verbs=get;list;watch;create;update;patch;delete

func (r *Microk8sConfigReconciler) Reconcile(req ctrl.Request) (_ ctrl.Result, rerr error) {
	ctx := context.Background()
	log := r.Log.WithValues("microk8sconfig", req.NamespacedName)

	// your logic here
	microk8sconfig := &bootstrapv1alpha1.Microk8sConfig{}

	if err := r.Client.Get(ctx, req.NamespacedName, microk8sconfig); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "failed to get config")
		return ctrl.Result{}, err
	}

	// Look up the owner of this KubeConfig if there is one
	configOwner, err := bsutil.GetConfigOwner(ctx, r.Client, microk8sconfig)
	if apierrors.IsNotFound(err) {
		// Could not find the owner yet, this is not an error and will rereconcile when the owner gets set.
		return ctrl.Result{}, nil
	}
	if err != nil {
		log.Error(err, "failed to get owner")
		return ctrl.Result{}, err
	}
	if configOwner == nil {
		log.Error(err, "failed to get config-owner")
		return ctrl.Result{}, nil
	}
	log = log.WithValues("kind", configOwner.GetKind(), "version", configOwner.GetResourceVersion(), "name", configOwner.GetName())

	// Initialize the patch helper
	patchHelper, err := patch.NewHelper(microk8sconfig, r)
	if err != nil {
		return ctrl.Result{}, err
	}

	if configOwner.IsControlPlaneMachine() {
		_, err := r.setupControlPlaneBootstrapData(ctx, microk8sconfig)
		if err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, patchHelper.Patch(ctx, microk8sconfig)
	} else {
		// Worker node
		log.Info("Worker node not supported yet.")
		return ctrl.Result{}, nil
	}
}

func (r *Microk8sConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bootstrapv1alpha1.Microk8sConfig{}).
		Complete(r)
}

// This method is only used to setup the control plane.
func (r *Microk8sConfigReconciler) setupControlPlaneBootstrapData(ctx context.Context, microk8sconfig *bootstrapv1alpha1.Microk8sConfig) (ctrl.Result, error) {
	switch {
	// Migrate plaintext data to secret.
	case microk8sconfig.Status.BootstrapData == nil:
		microk8sconfig.Status.BootstrapData = []byte(controllerInit)
		microk8sconfig.Status.Ready = true
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}
