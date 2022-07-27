package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HADeployment struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec HADeploymentSpec
}

type HADeploymentSpec struct {
	Image  string
	Labels map[string]string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HADeploymentList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []HADeployment
}
