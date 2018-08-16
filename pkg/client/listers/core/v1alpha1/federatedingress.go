/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedIngressLister helps list FederatedIngresses.
type FederatedIngressLister interface {
	// List lists all FederatedIngresses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedIngress, err error)
	// FederatedIngresses returns an object that can list and get FederatedIngresses.
	FederatedIngresses(namespace string) FederatedIngressNamespaceLister
	FederatedIngressListerExpansion
}

// federatedIngressLister implements the FederatedIngressLister interface.
type federatedIngressLister struct {
	indexer cache.Indexer
}

// NewFederatedIngressLister returns a new FederatedIngressLister.
func NewFederatedIngressLister(indexer cache.Indexer) FederatedIngressLister {
	return &federatedIngressLister{indexer: indexer}
}

// List lists all FederatedIngresses in the indexer.
func (s *federatedIngressLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedIngress))
	})
	return ret, err
}

// FederatedIngresses returns an object that can list and get FederatedIngresses.
func (s *federatedIngressLister) FederatedIngresses(namespace string) FederatedIngressNamespaceLister {
	return federatedIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedIngressNamespaceLister helps list and get FederatedIngresses.
type FederatedIngressNamespaceLister interface {
	// List lists all FederatedIngresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedIngress, err error)
	// Get retrieves the FederatedIngress from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedIngress, error)
	FederatedIngressNamespaceListerExpansion
}

// federatedIngressNamespaceLister implements the FederatedIngressNamespaceLister
// interface.
type federatedIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedIngresses in the indexer for a given namespace.
func (s federatedIngressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedIngress))
	})
	return ret, err
}

// Get retrieves the FederatedIngress from the indexer for a given namespace and name.
func (s federatedIngressNamespaceLister) Get(name string) (*v1alpha1.FederatedIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatedingress"), name)
	}
	return obj.(*v1alpha1.FederatedIngress), nil
}
