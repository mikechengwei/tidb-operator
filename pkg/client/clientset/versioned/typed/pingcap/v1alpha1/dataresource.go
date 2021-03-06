// Copyright 2019. PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	scheme "github.com/pingcap/tidb-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DataResourcesGetter has a method to return a DataResourceInterface.
// A group's client should implement this interface.
type DataResourcesGetter interface {
	DataResources(namespace string) DataResourceInterface
}

// DataResourceInterface has methods to work with DataResource resources.
type DataResourceInterface interface {
	Create(*v1alpha1.DataResource) (*v1alpha1.DataResource, error)
	Update(*v1alpha1.DataResource) (*v1alpha1.DataResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DataResource, error)
	List(opts v1.ListOptions) (*v1alpha1.DataResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataResource, err error)
	DataResourceExpansion
}

// dataResources implements DataResourceInterface
type dataResources struct {
	client rest.Interface
	ns     string
}

// newDataResources returns a DataResources
func newDataResources(c *PingcapV1alpha1Client, namespace string) *dataResources {
	return &dataResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataResource, and returns the corresponding dataResource object, and an error if there is any.
func (c *dataResources) Get(name string, options v1.GetOptions) (result *v1alpha1.DataResource, err error) {
	result = &v1alpha1.DataResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataResources that match those selectors.
func (c *dataResources) List(opts v1.ListOptions) (result *v1alpha1.DataResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DataResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataResources.
func (c *dataResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dataresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataResource and creates it.  Returns the server's representation of the dataResource, and an error, if there is any.
func (c *dataResources) Create(dataResource *v1alpha1.DataResource) (result *v1alpha1.DataResource, err error) {
	result = &v1alpha1.DataResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dataresources").
		Body(dataResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataResource and updates it. Returns the server's representation of the dataResource, and an error, if there is any.
func (c *dataResources) Update(dataResource *v1alpha1.DataResource) (result *v1alpha1.DataResource, err error) {
	result = &v1alpha1.DataResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dataresources").
		Name(dataResource.Name).
		Body(dataResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataResource and deletes it. Returns an error if one occurs.
func (c *dataResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataResource.
func (c *dataResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataResource, err error) {
	result = &v1alpha1.DataResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dataresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
