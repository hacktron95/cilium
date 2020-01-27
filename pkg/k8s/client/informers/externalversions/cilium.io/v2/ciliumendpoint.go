// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	time "time"

	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	versioned "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2 "github.com/cilium/cilium/pkg/k8s/client/listers/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumEndpointInformer provides access to a shared informer and lister for
// CiliumEndpoints.
type CiliumEndpointInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.CiliumEndpointLister
}

type ciliumEndpointInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCiliumEndpointInformer constructs a new informer for CiliumEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCiliumEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCiliumEndpointInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCiliumEndpointInformer constructs a new informer for CiliumEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCiliumEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumEndpoints(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumEndpoints(namespace).Watch(options)
			},
		},
		&ciliumiov2.CiliumEndpoint{},
		resyncPeriod,
		indexers,
	)
}

func (f *ciliumEndpointInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCiliumEndpointInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ciliumEndpointInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov2.CiliumEndpoint{}, f.defaultInformer)
}

func (f *ciliumEndpointInformer) Lister() v2.CiliumEndpointLister {
	return v2.NewCiliumEndpointLister(f.Informer().GetIndexer())
}
