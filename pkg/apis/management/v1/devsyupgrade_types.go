package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevsyUpgrade holds the upgrade information
// +k8s:openapi-gen=true
// +resource:path=devsyupgrades,rest=DevsyUpgradeREST
type DevsyUpgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyUpgradeSpec   `json:"spec,omitempty"`
	Status DevsyUpgradeStatus `json:"status,omitempty"`
}

type DevsyUpgradeSpec struct {
	// If specified, updated the release in the given namespace
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// If specified, uses this as release name
	// +optional
	Release string `json:"release,omitempty"`

	// +optional
	Version string `json:"version,omitempty"`
}

type DevsyUpgradeStatus struct{}
