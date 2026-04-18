package v1

import (
	clusterv1 "github.com/devsy-org/agentapi/pkg/apis/devsy/cluster/v1"
	agentstoragev1 "github.com/devsy-org/agentapi/pkg/apis/devsy/storage/v1"
	uiv1 "github.com/devsy-org/api/pkg/apis/ui/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This file is just used as a collector for kiosk objects we want to generate stuff for

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kiosk holds the kiosk types
// +k8s:openapi-gen=true
// +resource:path=kiosk,rest=KioskREST
type Kiosk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KioskSpec   `json:"spec,omitempty"`
	Status KioskStatus `json:"status,omitempty"`
}

type KioskSpec struct {
	// cluster.devsy.sh
	HelmRelease     clusterv1.HelmRelease     `json:"helmRelease,omitempty"`
	SleepModeConfig clusterv1.SleepModeConfig `json:"sleepModeConfig,omitempty"`
	ChartInfo       clusterv1.ChartInfo       `json:"chartInfo,omitempty"`

	// storage.devsy.sh
	StorageClusterQuota agentstoragev1.ClusterQuota `json:"storageClusterQuota,omitempty"`

	// ui.devsy.sh
	UISettings uiv1.UISettings `json:"UISettings,omitempty"`

	License License `json:"license,omitempty"`

	// autoscaling
	NodeProviderBCMNodeWithResources    NodeProviderBCMNodeWithResources    `json:"nodeProviderBCMNodeWithResources,omitempty"`
	NodeProviderBCMGetResourcesResult   NodeProviderBCMGetResourcesResult   `json:"nodeProviderBCMGetResourcesResult,omitempty"`
	NodeProviderBCMTestConnectionResult NodeProviderBCMTestConnectionResult `json:"nodeProviderBCMTestConnectionResult,omitempty"`
	NodeProviderCalculateCostResult     NodeProviderCalculateCostResult     `json:"nodeProviderCalculateCostResult,omitempty"`
	NodeProviderTerraformValidateResult NodeProviderTerraformValidateResult `json:"nodeProviderTerraformValidateResult,omitempty"`
	NodeProviderExecResult              NodeProviderExecResult              `json:"nodeProviderExecResult,omitempty"`
	NodeClaimData                       NodeClaimData                       `json:"nodeClaimData,omitempty"`
	NodeEnvironmentData                 NodeEnvironmentData                 `json:"nodeEnvironmentData,omitempty"`
}

type KioskStatus struct{}
