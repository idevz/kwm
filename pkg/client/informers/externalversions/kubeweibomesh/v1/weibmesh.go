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

package v1

import (
	time "time"

	kubeweibomeshv1 "github.com/idevz/kwm/pkg/apis/kubeweibomesh/v1"
	versioned "github.com/idevz/kwm/pkg/client/clientset/versioned"
	internalinterfaces "github.com/idevz/kwm/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/idevz/kwm/pkg/client/listers/kubeweibomesh/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WeibMeshInformer provides access to a shared informer and lister for
// WeibMeshes.
type WeibMeshInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.WeibMeshLister
}

type weibMeshInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWeibMeshInformer constructs a new informer for WeibMesh type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWeibMeshInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWeibMeshInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWeibMeshInformer constructs a new informer for WeibMesh type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWeibMeshInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WeibomeshV1().WeibMeshes(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WeibomeshV1().WeibMeshes(namespace).Watch(options)
			},
		},
		&kubeweibomeshv1.WeibMesh{},
		resyncPeriod,
		indexers,
	)
}

func (f *weibMeshInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWeibMeshInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *weibMeshInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeweibomeshv1.WeibMesh{}, f.defaultInformer)
}

func (f *weibMeshInformer) Lister() v1.WeibMeshLister {
	return v1.NewWeibMeshLister(f.Informer().GetIndexer())
}
