package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	PhasePending = "PENDING"
	PhaseRunning = "RUNNING"
	PhaseDone    = "DONE"
)

// AtSpec defines the desired state of At
// +k8s:openapi-gen=true
type AtSpec struct {
	// Schedule is the desired time the command is supposed to be executed.
	// Note: the format used here is UTC time https://www.utctime.net
	Schedule string `json:"schedule,omitempty"`
	// Command is the desired command (executed in a Bash shell) to be
	// executed.
	Command string `json:"command,omitempty"`
	// Important: Run "make" to regenerate code after modifying this file
}

// AtStatus defines the observed state of At
// +k8s:openapi-gen=true
type AtStatus struct {
	// Phase represents the state of the schedule: until the command is
	// executed it is PENDING, afterwards it is DONE.
	Phase string `json:"phase,omitempty"`
	// Important: Run "make" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// At is the Schema for the ats API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ats,scope=Namespaced
type At struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AtSpec   `json:"spec,omitempty"`
	Status AtStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AtList contains a list of At
type AtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []At `json:"items"`
}

func init() {
	SchemeBuilder.Register(&At{}, &AtList{})
}
