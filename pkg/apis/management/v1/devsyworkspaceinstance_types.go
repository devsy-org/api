package v1

import (
	clusterv1 "github.com/devsy-org/agentapi/pkg/apis/devsy/cluster/v1"
	agentstoragev1 "github.com/devsy-org/agentapi/pkg/apis/devsy/storage/v1"
	storagev1 "github.com/devsy-org/api/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +genclient:method=Up,verb=create,subresource=up,input=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceUp,result=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceUp
// +genclient:method=Stop,verb=create,subresource=stop,input=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceStop,result=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceStop
// +genclient:method=Troubleshoot,verb=get,subresource=troubleshoot,result=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceTroubleshoot
// +genclient:method=Cancel,verb=create,subresource=cancel,input=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceCancel,result=github.com/devsy-org/api/pkg/apis/management/v1.DevsyWorkspaceInstanceCancel
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevsyWorkspaceInstance holds the DevsyWorkspaceInstance information
// +k8s:openapi-gen=true
// +resource:path=devpodworkspaceinstances,rest=DevsyWorkspaceInstanceREST
// +subresource:request=DevsyWorkspaceInstanceUp,path=up,kind=DevsyWorkspaceInstanceUp,rest=DevsyWorkspaceInstanceUpREST
// +subresource:request=DevsyWorkspaceInstanceStop,path=stop,kind=DevsyWorkspaceInstanceStop,rest=DevsyWorkspaceInstanceStopREST
// +subresource:request=DevsyWorkspaceInstanceTroubleshoot,path=troubleshoot,kind=DevsyWorkspaceInstanceTroubleshoot,rest=DevsyWorkspaceInstanceTroubleshootREST
// +subresource:request=DevsyWorkspaceInstanceLog,path=log,kind=DevsyWorkspaceInstanceLog,rest=DevsyWorkspaceInstanceLogREST
// +subresource:request=DevsyWorkspaceInstanceTasks,path=tasks,kind=DevsyWorkspaceInstanceTasks,rest=DevsyWorkspaceInstanceTasksREST
// +subresource:request=DevsyWorkspaceInstanceCancel,path=cancel,kind=DevsyWorkspaceInstanceCancel,rest=DevsyWorkspaceInstanceCancelREST
// +subresource:request=DevsyWorkspaceInstanceDownload,path=download,kind=DevsyWorkspaceInstanceDownload,rest=DevsyWorkspaceInstanceDownloadREST
type DevsyWorkspaceInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevsyWorkspaceInstanceSpec   `json:"spec,omitempty"`
	Status DevsyWorkspaceInstanceStatus `json:"status,omitempty"`
}

// DevsyWorkspaceInstanceSpec holds the specification.
type DevsyWorkspaceInstanceSpec struct {
	storagev1.DevsyWorkspaceInstanceSpec `json:",inline"`
}

// DevsyWorkspaceInstanceStatus holds the status.
type DevsyWorkspaceInstanceStatus struct {
	storagev1.DevsyWorkspaceInstanceStatus `json:",inline"`

	// SleepModeConfig is the sleep mode config of the workspace. This will only be shown
	// in the front end.
	// +optional
	SleepModeConfig *clusterv1.SleepModeConfig `json:"sleepModeConfig,omitempty"`
}

func (a *DevsyWorkspaceInstance) GetConditions() agentstoragev1.Conditions {
	return a.Status.Conditions
}

func (a *DevsyWorkspaceInstance) SetConditions(conditions agentstoragev1.Conditions) {
	a.Status.Conditions = conditions
}

func (a *DevsyWorkspaceInstance) GetOwner() *storagev1.UserOrTeam {
	return a.Spec.Owner
}

func (a *DevsyWorkspaceInstance) SetOwner(userOrTeam *storagev1.UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevsyWorkspaceInstance) GetAccess() []storagev1.Access {
	return a.Spec.Access
}

func (a *DevsyWorkspaceInstance) SetAccess(access []storagev1.Access) {
	a.Spec.Access = access
}
