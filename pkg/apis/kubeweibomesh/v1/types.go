package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WeiboMeshSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas *int32 `json:"replicas"`
}

type WeiboMeshStatus struct {
	AvailableReplicas int32 `json:"avaliableReplicas"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type WeibMesh struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WeiboMeshSpec `json:"spec"`
	Status WeiboMeshStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type WeiboMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WeibMesh `json:"items"`
}