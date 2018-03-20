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
package v1alpha1

import (
	v1alpha1 "github.com/marun/fnord/pkg/apis/federation/v1alpha1"
	scheme "github.com/marun/fnord/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PropagatedVersionsGetter has a method to return a PropagatedVersionInterface.
// A group's client should implement this interface.
type PropagatedVersionsGetter interface {
	PropagatedVersions(namespace string) PropagatedVersionInterface
}

// PropagatedVersionInterface has methods to work with PropagatedVersion resources.
type PropagatedVersionInterface interface {
	Create(*v1alpha1.PropagatedVersion) (*v1alpha1.PropagatedVersion, error)
	Update(*v1alpha1.PropagatedVersion) (*v1alpha1.PropagatedVersion, error)
	UpdateStatus(*v1alpha1.PropagatedVersion) (*v1alpha1.PropagatedVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PropagatedVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.PropagatedVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PropagatedVersion, err error)
	PropagatedVersionExpansion
}

// propagatedVersions implements PropagatedVersionInterface
type propagatedVersions struct {
	client rest.Interface
	ns     string
}

// newPropagatedVersions returns a PropagatedVersions
func newPropagatedVersions(c *FederationV1alpha1Client, namespace string) *propagatedVersions {
	return &propagatedVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the propagatedVersion, and returns the corresponding propagatedVersion object, and an error if there is any.
func (c *propagatedVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.PropagatedVersion, err error) {
	result = &v1alpha1.PropagatedVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("propagatedversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PropagatedVersions that match those selectors.
func (c *propagatedVersions) List(opts v1.ListOptions) (result *v1alpha1.PropagatedVersionList, err error) {
	result = &v1alpha1.PropagatedVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("propagatedversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested propagatedVersions.
func (c *propagatedVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("propagatedversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a propagatedVersion and creates it.  Returns the server's representation of the propagatedVersion, and an error, if there is any.
func (c *propagatedVersions) Create(propagatedVersion *v1alpha1.PropagatedVersion) (result *v1alpha1.PropagatedVersion, err error) {
	result = &v1alpha1.PropagatedVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("propagatedversions").
		Body(propagatedVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a propagatedVersion and updates it. Returns the server's representation of the propagatedVersion, and an error, if there is any.
func (c *propagatedVersions) Update(propagatedVersion *v1alpha1.PropagatedVersion) (result *v1alpha1.PropagatedVersion, err error) {
	result = &v1alpha1.PropagatedVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("propagatedversions").
		Name(propagatedVersion.Name).
		Body(propagatedVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *propagatedVersions) UpdateStatus(propagatedVersion *v1alpha1.PropagatedVersion) (result *v1alpha1.PropagatedVersion, err error) {
	result = &v1alpha1.PropagatedVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("propagatedversions").
		Name(propagatedVersion.Name).
		SubResource("status").
		Body(propagatedVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the propagatedVersion and deletes it. Returns an error if one occurs.
func (c *propagatedVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("propagatedversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *propagatedVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("propagatedversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched propagatedVersion.
func (c *propagatedVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PropagatedVersion, err error) {
	result = &v1alpha1.PropagatedVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("propagatedversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
