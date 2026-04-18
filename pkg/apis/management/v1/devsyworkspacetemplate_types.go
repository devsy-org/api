package v1

import (
	storagev1 "github.com/devsy-org/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevsyWorkspaceTemplate holds the information
// +k8s:openapi-gen=true
// +resource:path=devpodworkspacetemplates,rest=DevsyWorkspaceTemplateREST
type DevsyWorkspaceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyWorkspaceTemplateSpec   `json:"spec,omitempty"`
	Status DevsyWorkspaceTemplateStatus `json:"status,omitempty"`
}

// DevsyWorkspaceTemplateSpec holds the specification
type DevsyWorkspaceTemplateSpec struct {
	storagev1.DevsyWorkspaceTemplateSpec `json:",inline"`
}

// DevsyWorkspaceTemplateStatus holds the status
type DevsyWorkspaceTemplateStatus struct {
	storagev1.DevsyWorkspaceTemplateStatus `json:",inline"`
}

func (a *DevsyWorkspaceTemplate) GetVersions() []storagev1.VersionAccessor {
	var retVersions []storagev1.VersionAccessor
	for _, v := range a.Spec.Versions {
		b := v
		retVersions = append(retVersions, &b)
	}

	return retVersions
}

func (a *DevsyWorkspaceTemplate) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *DevsyWorkspaceTemplate) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevsyWorkspaceTemplate) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *DevsyWorkspaceTemplate) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}
