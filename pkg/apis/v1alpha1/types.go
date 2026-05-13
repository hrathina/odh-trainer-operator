// +kubebuilder:object:generate=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/opendatahub-io/odh-platform-utilities/api/common"
)

const (
	TrainerKind         = "Trainer"
	TrainerInstanceName = "default"
)

// TrainerSpec defines the desired state of Trainer
type TrainerSpec struct {
	// ManagementSpec defines whether the component is Managed or Removed
	common.ManagementSpec `json:",inline"`

	// AppNamespace is the namespace where trainer resources are deployed
	// This is typically set by the platform operator based on DSCI configuration
	// +optional
	AppNamespace string `json:"appNamespace,omitempty"`

	// FeatureGates enables experimental or optional features
	// +optional
	FeatureGates FeatureGatesSpec `json:"featureGates,omitempty"`
}

// FeatureGatesSpec defines optional features that can be enabled/disabled
type FeatureGatesSpec struct {
	// ProgressTracking enables progress tracking for training jobs
	// +optional
	ProgressTracking bool `json:"progressTracking,omitempty"`
}

// TrainerStatus defines the observed state of Trainer
type TrainerStatus struct {
	common.Status                 `json:",inline"`
	common.ComponentReleaseStatus `json:",inline"`
}

// GetManagementState returns the management state from spec, defaulting to Managed.
func GetManagementState(trainer *Trainer) common.ManagementState {
	if trainer.Spec.ManagementState != "" {
		return trainer.Spec.ManagementState
	}
	return common.Managed
}

// Ensure Trainer implements PlatformObject interface
var _ common.PlatformObject = &Trainer{}

// GetStatus returns the status of the Trainer
func (t *Trainer) GetStatus() *common.Status {
	return &t.Status.Status
}

// GetConditions returns the status conditions
func (t *Trainer) GetConditions() []common.Condition {
	return t.Status.Conditions
}

// SetConditions updates the status conditions
func (t *Trainer) SetConditions(conditions []common.Condition) {
	t.Status.Conditions = conditions
}

// GetReleaseStatus returns the component release status
func (t *Trainer) GetReleaseStatus() *common.ComponentReleaseStatus {
	return &t.Status.ComponentReleaseStatus
}

// SetReleaseStatus updates the component release status
func (t *Trainer) SetReleaseStatus(status common.ComponentReleaseStatus) {
	t.Status.ComponentReleaseStatus = status
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'default'",message="Trainer name must be 'default'"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// Trainer is the Schema for the trainers API
type Trainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrainerSpec   `json:"spec,omitempty"`
	Status TrainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// TrainerList contains a list of Trainer
type TrainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trainer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Trainer{}, &TrainerList{})
}
