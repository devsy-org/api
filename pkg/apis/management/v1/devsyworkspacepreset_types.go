package v1

import (
	storagev1 "github.com/devsy-org/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevsyWorkspacePreset
// +k8s:openapi-gen=true
// +resource:path=devpodworkspacepresets,rest=DevsyWorkspacePresetREST
type DevsyWorkspacePreset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyWorkspacePresetSpec   `json:"spec,omitempty"`
	Status DevsyWorkspacePresetStatus `json:"status,omitempty"`
}

// DevsyWorkspacePresetSpec holds the specification.
type DevsyWorkspacePresetSpec struct {
	storagev1.DevsyWorkspacePresetSpec `json:",inline"`
}

// DevsyWorkspacePresetSource
// +k8s:openapi-gen=true
type DevsyWorkspacePresetSource struct {
	storagev1.DevsyWorkspacePresetSource `json:",inline"`
}

func (a *DevsyWorkspacePreset) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *DevsyWorkspacePreset) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevsyWorkspacePreset) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *DevsyWorkspacePreset) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}

// DevsyWorkspacePresetStatus holds the status.
type DevsyWorkspacePresetStatus struct{}
