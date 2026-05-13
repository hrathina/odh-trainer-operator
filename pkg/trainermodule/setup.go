package trainermodule

import (
	ctrl "sigs.k8s.io/controller-runtime"

	platformv1alpha1 "github.com/hrathina/odh-trainer-operator/pkg/apis/v1alpha1"
)

// SetupWithManager sets up the controller with the Manager
func (r *TrainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&platformv1alpha1.Trainer{}).
		Complete(r)
}
