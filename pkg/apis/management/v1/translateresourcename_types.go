package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TranslateDevsyResourceName holds rename request and response data for given resource
// +k8s:openapi-gen=true
// +resource:path=translatedevsyresourcenames,rest=TranslateDevsyResourceNameREST
type TranslateDevsyResourceName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TranslateDevsyResourceNameSpec   `json:"spec,omitempty"`
	Status TranslateDevsyResourceNameStatus `json:"status,omitempty"`
}

// TranslateDevsyResourceNameSpec holds the specification
type TranslateDevsyResourceNameSpec struct {
	// Name is the name of resource we want to rename
	Name string `json:"name"`

	// Namespace is the name of namespace in which this resource is running
	Namespace string `json:"namespace"`

	// DevsyName is the name of vCluster in which this resource is running
	DevsyName string `json:"devsyName"`
}

// TranslateDevsyResourceNameStatus holds the status
type TranslateDevsyResourceNameStatus struct {
	// Name is the converted name of resource
	// +optional
	Name string `json:"name,omitempty"`
}
