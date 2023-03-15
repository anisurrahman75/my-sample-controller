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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/anisurrahman75/k8s-crd-controller/pkg/apis/mycrd.k8s/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppsCodeLister helps list AppsCodes.
// All objects returned here must be treated as read-only.
type AppsCodeLister interface {
	// List lists all AppsCodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppsCode, err error)
	// AppsCodes returns an object that can list and get AppsCodes.
	AppsCodes(namespace string) AppsCodeNamespaceLister
	AppsCodeListerExpansion
}

// appsCodeLister implements the AppsCodeLister interface.
type appsCodeLister struct {
	indexer cache.Indexer
}

// NewAppsCodeLister returns a new AppsCodeLister.
func NewAppsCodeLister(indexer cache.Indexer) AppsCodeLister {
	return &appsCodeLister{indexer: indexer}
}

// List lists all AppsCodes in the indexer.
func (s *appsCodeLister) List(selector labels.Selector) (ret []*v1alpha1.AppsCode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppsCode))
	})
	return ret, err
}

// AppsCodes returns an object that can list and get AppsCodes.
func (s *appsCodeLister) AppsCodes(namespace string) AppsCodeNamespaceLister {
	return appsCodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppsCodeNamespaceLister helps list and get AppsCodes.
// All objects returned here must be treated as read-only.
type AppsCodeNamespaceLister interface {
	// List lists all AppsCodes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppsCode, err error)
	// Get retrieves the AppsCode from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AppsCode, error)
	AppsCodeNamespaceListerExpansion
}

// appsCodeNamespaceLister implements the AppsCodeNamespaceLister
// interface.
type appsCodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppsCodes in the indexer for a given namespace.
func (s appsCodeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppsCode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppsCode))
	})
	return ret, err
}

// Get retrieves the AppsCode from the indexer for a given namespace and name.
func (s appsCodeNamespaceLister) Get(name string) (*v1alpha1.AppsCode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appscode"), name)
	}
	return obj.(*v1alpha1.AppsCode), nil
}