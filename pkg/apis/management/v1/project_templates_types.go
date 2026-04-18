package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +subresource-request
type ProjectTemplates struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// DefaultVirtualClusterTemplate is the default template for the project
	DefaultVirtualClusterTemplate string `json:"defaultVirtualClusterTemplate,omitempty"`

	// VirtualClusterTemplates holds all the allowed virtual cluster templates
	VirtualClusterTemplates []VirtualClusterTemplate `json:"virtualClusterTemplates,omitempty"`

	// DefaultSpaceTemplate
	DefaultSpaceTemplate string `json:"defaultSpaceTemplate,omitempty"`

	// SpaceTemplates holds all the allowed space templates
	SpaceTemplates []SpaceTemplate `json:"spaceTemplates,omitempty"`

	// DefaultDevsyWorkspaceTemplate
	DefaultDevsyWorkspaceTemplate string `json:"defaultDevPodWorkspaceTemplate,omitempty"`

	// DevsyWorkspaceTemplates holds all the allowed space templates
	DevsyWorkspaceTemplates []DevsyWorkspaceTemplate `json:"devPodWorkspaceTemplates,omitempty"`

	// DevsyEnvironmentTemplates holds all the allowed environment templates
	DevsyEnvironmentTemplates []DevsyEnvironmentTemplate `json:"devPodEnvironmentTemplates,omitempty"`

	// DevsyWorkspacePresets holds all the allowed workspace presets
	DevsyWorkspacePresets []DevsyWorkspacePreset `json:"devPodWorkspacePresets,omitempty"`

	// DefaultDevsyEnvironmentTemplate
	DefaultDevsyEnvironmentTemplate string `json:"defaultDevPodEnvironmentTemplate,omitempty"`
}
