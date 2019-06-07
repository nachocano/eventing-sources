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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing-sources/contrib/camel/pkg/apis/sources/v1alpha1"
	scheme "github.com/knative/eventing-sources/contrib/camel/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CamelSourcesGetter has a method to return a CamelSourceInterface.
// A group's client should implement this interface.
type CamelSourcesGetter interface {
	CamelSources(namespace string) CamelSourceInterface
}

// CamelSourceInterface has methods to work with CamelSource resources.
type CamelSourceInterface interface {
	Create(*v1alpha1.CamelSource) (*v1alpha1.CamelSource, error)
	Update(*v1alpha1.CamelSource) (*v1alpha1.CamelSource, error)
	UpdateStatus(*v1alpha1.CamelSource) (*v1alpha1.CamelSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CamelSource, error)
	List(opts v1.ListOptions) (*v1alpha1.CamelSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CamelSource, err error)
	CamelSourceExpansion
}

// camelSources implements CamelSourceInterface
type camelSources struct {
	client rest.Interface
	ns     string
}

// newCamelSources returns a CamelSources
func newCamelSources(c *SourcesV1alpha1Client, namespace string) *camelSources {
	return &camelSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the camelSource, and returns the corresponding camelSource object, and an error if there is any.
func (c *camelSources) Get(name string, options v1.GetOptions) (result *v1alpha1.CamelSource, err error) {
	result = &v1alpha1.CamelSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("camelsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CamelSources that match those selectors.
func (c *camelSources) List(opts v1.ListOptions) (result *v1alpha1.CamelSourceList, err error) {
	result = &v1alpha1.CamelSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("camelsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested camelSources.
func (c *camelSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("camelsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a camelSource and creates it.  Returns the server's representation of the camelSource, and an error, if there is any.
func (c *camelSources) Create(camelSource *v1alpha1.CamelSource) (result *v1alpha1.CamelSource, err error) {
	result = &v1alpha1.CamelSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("camelsources").
		Body(camelSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a camelSource and updates it. Returns the server's representation of the camelSource, and an error, if there is any.
func (c *camelSources) Update(camelSource *v1alpha1.CamelSource) (result *v1alpha1.CamelSource, err error) {
	result = &v1alpha1.CamelSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("camelsources").
		Name(camelSource.Name).
		Body(camelSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *camelSources) UpdateStatus(camelSource *v1alpha1.CamelSource) (result *v1alpha1.CamelSource, err error) {
	result = &v1alpha1.CamelSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("camelsources").
		Name(camelSource.Name).
		SubResource("status").
		Body(camelSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the camelSource and deletes it. Returns an error if one occurs.
func (c *camelSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("camelsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *camelSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("camelsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched camelSource.
func (c *camelSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CamelSource, err error) {
	result = &v1alpha1.CamelSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("camelsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
