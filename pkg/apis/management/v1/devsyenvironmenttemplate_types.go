package v1

import (
	storagev1 "github.com/devsy-org/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevsyEnvironmentTemplate holds the DevsyEnvironmentTemplate information
// +k8s:openapi-gen=true
// +resource:path=devpodenvironmenttemplates,rest=DevsyEnvironmentTemplateREST
type DevsyEnvironmentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyEnvironmentTemplateSpec   `json:"spec,omitempty"`
	Status DevsyEnvironmentTemplateStatus `json:"status,omitempty"`
}

// DevsyEnvironmentTemplateSpec holds the specification.
type DevsyEnvironmentTemplateSpec struct {
	storagev1.DevsyEnvironmentTemplateSpec `json:",inline"`
}

// DevsyEnvironmentTemplateStatus holds the status.
type DevsyEnvironmentTemplateStatus struct{}

func (a *DevsyEnvironmentTemplate) GetVersions() []storagev1.VersionAccessor {
	var retVersions []storagev1.VersionAccessor
	for _, v := range a.Spec.Versions {
		b := v
		retVersions = append(retVersions, &b)
	}

	return retVersions
}

func (a *DevsyEnvironmentTemplate) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *DevsyEnvironmentTemplate) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevsyEnvironmentTemplate) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *DevsyEnvironmentTemplate) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}
