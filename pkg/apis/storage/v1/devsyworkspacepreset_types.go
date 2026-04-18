package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevsyWorkspacePreset
// +k8s:openapi-gen=true
type DevsyWorkspacePreset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyWorkspacePresetSpec   `json:"spec,omitempty"`
	Status DevsyWorkspacePresetStatus `json:"status,omitempty"`
}

func (a *DevsyWorkspacePreset) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *DevsyWorkspacePreset) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevsyWorkspacePreset) GetAccess() []Access {
	return a.Spec.Access
}

func (a *DevsyWorkspacePreset) SetAccess(access []Access) {
	a.Spec.Access = access
}

type DevsyWorkspacePresetSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Source stores inline path of project source
	Source *DevsyWorkspacePresetSource `json:"source"`

	// InfrastructureRef stores reference to DevsyWorkspaceTemplate to use
	InfrastructureRef *TemplateRef `json:"infrastructureRef"`

	// EnvironmentRef stores reference to DevsyEnvironmentTemplate
	// +optional
	EnvironmentRef *EnvironmentRef `json:"environmentRef,omitempty"`

	// UseProjectGitCredentials specifies if the project git credentials should be used instead of local ones for this environment
	// +optional
	UseProjectGitCredentials bool `json:"useProjectGitCredentials,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access to the Devsy machine instance object itself
	// +optional
	Access []Access `json:"access,omitempty"`

	// Versions are different versions of the template that can be referenced as well
	// +optional
	Versions []DevsyWorkspacePresetVersion `json:"versions,omitempty"`
}

type DevsyWorkspacePresetSource struct {
	// Git stores path to git repo to use as workspace source
	// +optional
	Git string `json:"git,omitempty"`

	// Image stores container image to use as workspace source
	// +optional
	Image string `json:"image,omitempty"`
}

type DevsyWorkspacePresetVersion struct {
	// Version is the version. Needs to be in X.X.X format.
	// +optional
	Version string `json:"version,omitempty"`

	// Source stores inline path of project source
	// +optional
	Source *DevsyWorkspacePresetSource `json:"source,omitempty"`

	// InfrastructureRef stores reference to DevsyWorkspaceTemplate to use
	// +optional
	InfrastructureRef *TemplateRef `json:"infrastructureRef,omitempty"`

	// EnvironmentRef stores reference to DevsyEnvironmentTemplate
	// +optional
	EnvironmentRef *EnvironmentRef `json:"environmentRef,omitempty"`
}

// DevsyWorkspacePresetStatus holds the status.
type DevsyWorkspacePresetStatus struct{}

type WorkspaceRef struct {
	// Name is the name of DevsyWorkspaceTemplate this references
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DevsyWorkspacePresetList contains a list of DevsyWorkspacePreset objects.
type DevsyWorkspacePresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevsyWorkspacePreset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevsyWorkspacePreset{}, &DevsyWorkspacePresetList{})
}
