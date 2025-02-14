/*
Copyright 2025 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRoleTemplateBindings implements ClusterRoleTemplateBindingInterface
type FakeClusterRoleTemplateBindings struct {
	Fake *FakeManagementV3
	ns   string
}

var clusterroletemplatebindingsResource = v3.SchemeGroupVersion.WithResource("clusterroletemplatebindings")

var clusterroletemplatebindingsKind = v3.SchemeGroupVersion.WithKind("ClusterRoleTemplateBinding")

// Get takes name of the clusterRoleTemplateBinding, and returns the corresponding clusterRoleTemplateBinding object, and an error if there is any.
func (c *FakeClusterRoleTemplateBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterRoleTemplateBinding, err error) {
	emptyResult := &v3.ClusterRoleTemplateBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(clusterroletemplatebindingsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.ClusterRoleTemplateBinding), err
}

// List takes label and field selectors, and returns the list of ClusterRoleTemplateBindings that match those selectors.
func (c *FakeClusterRoleTemplateBindings) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterRoleTemplateBindingList, err error) {
	emptyResult := &v3.ClusterRoleTemplateBindingList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(clusterroletemplatebindingsResource, clusterroletemplatebindingsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterRoleTemplateBindingList{ListMeta: obj.(*v3.ClusterRoleTemplateBindingList).ListMeta}
	for _, item := range obj.(*v3.ClusterRoleTemplateBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRoleTemplateBindings.
func (c *FakeClusterRoleTemplateBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(clusterroletemplatebindingsResource, c.ns, opts))

}

// Create takes the representation of a clusterRoleTemplateBinding and creates it.  Returns the server's representation of the clusterRoleTemplateBinding, and an error, if there is any.
func (c *FakeClusterRoleTemplateBindings) Create(ctx context.Context, clusterRoleTemplateBinding *v3.ClusterRoleTemplateBinding, opts v1.CreateOptions) (result *v3.ClusterRoleTemplateBinding, err error) {
	emptyResult := &v3.ClusterRoleTemplateBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(clusterroletemplatebindingsResource, c.ns, clusterRoleTemplateBinding, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.ClusterRoleTemplateBinding), err
}

// Update takes the representation of a clusterRoleTemplateBinding and updates it. Returns the server's representation of the clusterRoleTemplateBinding, and an error, if there is any.
func (c *FakeClusterRoleTemplateBindings) Update(ctx context.Context, clusterRoleTemplateBinding *v3.ClusterRoleTemplateBinding, opts v1.UpdateOptions) (result *v3.ClusterRoleTemplateBinding, err error) {
	emptyResult := &v3.ClusterRoleTemplateBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(clusterroletemplatebindingsResource, c.ns, clusterRoleTemplateBinding, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.ClusterRoleTemplateBinding), err
}

// Delete takes name of the clusterRoleTemplateBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterRoleTemplateBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusterroletemplatebindingsResource, c.ns, name, opts), &v3.ClusterRoleTemplateBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRoleTemplateBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(clusterroletemplatebindingsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterRoleTemplateBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterRoleTemplateBinding.
func (c *FakeClusterRoleTemplateBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterRoleTemplateBinding, err error) {
	emptyResult := &v3.ClusterRoleTemplateBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(clusterroletemplatebindingsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v3.ClusterRoleTemplateBinding), err
}
