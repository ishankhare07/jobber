package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	jobberv1alpha1 "github.com/ishankhare07/jobber/api/v1alpha1"
)

// CronReconciler reconciles a Cron object
type CronReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=jobber.ishankhare.dev,resources=crons,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=jobber.ishankhare.dev,resources=crons/status,verbs=get;update;patch

func (r *CronReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("cron", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *CronReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jobberv1alpha1.Cron{}).
		Complete(r)
}
