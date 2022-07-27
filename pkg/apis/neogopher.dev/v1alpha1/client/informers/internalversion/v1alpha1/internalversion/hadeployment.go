/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	time "time"

	v1alpha1 "github.com/neogopher/hadeployment/pkg/apis/neogopher.dev/v1alpha1"
	internalinterfaces "github.com/neogopher/hadeployment/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/neogopher/hadeployment/pkg/client/listers/v1alpha1/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
)

// HADeploymentInformer provides access to a shared informer and lister for
// HADeployments.
type HADeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.HADeploymentLister
}

type hADeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHADeploymentInformer constructs a new informer for HADeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHADeploymentInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHADeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHADeploymentInformer constructs a new informer for HADeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHADeploymentInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.V1alpha1().HADeployments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.V1alpha1().HADeployments(namespace).Watch(context.TODO(), options)
			},
		},
		&v1alpha1.HADeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *hADeploymentInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHADeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hADeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&v1alpha1.HADeployment{}, f.defaultInformer)
}

func (f *hADeploymentInformer) Lister() internalversion.HADeploymentLister {
	return internalversion.NewHADeploymentLister(f.Informer().GetIndexer())
}