package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	PhasePending    = "PENDING"
	PhaseRunning    = "RUNNING"
	PhaseRegistered = "REGISTERED"
	PhaseDone       = "DONE"
)

type SecretSpec struct {
	Name      string `json:"name,omitempty"`
	MountAs   string `json:"mount_as,omitempty"`   // volume or env
	MountPath string `json:"mount_path,omitempty"` // only considered in case of mount as volume
}

// CronSpec defines the desired state of Cron
type CronSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// crontab syntax based schedule
	Schedule string     `json:"schedule,omitempty"`
	Secret   SecretSpec `json:"secret,omitempty"`
}

// CronStatus defines the observed state of Cron
type CronStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase        string `json:"phase,omitempty"`
	LastExecuted string `json:"last_executed,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cron is the Schema for the crons API
type Cron struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronSpec   `json:"spec,omitempty"`
	Status CronStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronList contains a list of Cron
type CronList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cron `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cron{}, &CronList{})
}
