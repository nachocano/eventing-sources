/*
Copyright 2019 The Knative Authors

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
	v1alpha1 "github.com/knative/eventing/pkg/apis/messaging/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ChoiceLister helps list Choices.
type ChoiceLister interface {
	// List lists all Choices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Choice, err error)
	// Choices returns an object that can list and get Choices.
	Choices(namespace string) ChoiceNamespaceLister
	ChoiceListerExpansion
}

// choiceLister implements the ChoiceLister interface.
type choiceLister struct {
	indexer cache.Indexer
}

// NewChoiceLister returns a new ChoiceLister.
func NewChoiceLister(indexer cache.Indexer) ChoiceLister {
	return &choiceLister{indexer: indexer}
}

// List lists all Choices in the indexer.
func (s *choiceLister) List(selector labels.Selector) (ret []*v1alpha1.Choice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Choice))
	})
	return ret, err
}

// Choices returns an object that can list and get Choices.
func (s *choiceLister) Choices(namespace string) ChoiceNamespaceLister {
	return choiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ChoiceNamespaceLister helps list and get Choices.
type ChoiceNamespaceLister interface {
	// List lists all Choices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Choice, err error)
	// Get retrieves the Choice from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Choice, error)
	ChoiceNamespaceListerExpansion
}

// choiceNamespaceLister implements the ChoiceNamespaceLister
// interface.
type choiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Choices in the indexer for a given namespace.
func (s choiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Choice, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Choice))
	})
	return ret, err
}

// Get retrieves the Choice from the indexer for a given namespace and name.
func (s choiceNamespaceLister) Get(name string) (*v1alpha1.Choice, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("choice"), name)
	}
	return obj.(*v1alpha1.Choice), nil
}
