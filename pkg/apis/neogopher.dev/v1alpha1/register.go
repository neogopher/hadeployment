package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   "neogopher.dev",
		Version: "v1alpha1",
	}

	SchemeBuilder runtime.SchemeBuilder
)

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &HADeployment{}, &HADeploymentList{})

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
