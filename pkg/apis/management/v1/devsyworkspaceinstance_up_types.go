package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type DevsyWorkspaceInstanceUp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyWorkspaceInstanceUpSpec   `json:"spec,omitempty"`
	Status DevsyWorkspaceInstanceUpStatus `json:"status,omitempty"`
}

type DevsyWorkspaceInstanceUpSpec struct {
	// Debug includes debug logs.
	// +optional
	Debug bool `json:"debug,omitempty"`

	// Options are the options to pass.
	// +optional
	Options string `json:"options,omitempty"`
}

type DevsyWorkspaceInstanceUpStatus struct {
	// TaskID is the id of the task that is running
	// +optional
	TaskID string `json:"taskId,omitempty"`
}
