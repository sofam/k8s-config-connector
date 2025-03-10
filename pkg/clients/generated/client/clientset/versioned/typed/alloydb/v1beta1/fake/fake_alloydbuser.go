// Copyright 2020 Google LLC
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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/alloydb/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlloyDBUsers implements AlloyDBUserInterface
type FakeAlloyDBUsers struct {
	Fake *FakeAlloydbV1beta1
	ns   string
}

var alloydbusersResource = schema.GroupVersionResource{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Resource: "alloydbusers"}

var alloydbusersKind = schema.GroupVersionKind{Group: "alloydb.cnrm.cloud.google.com", Version: "v1beta1", Kind: "AlloyDBUser"}

// Get takes name of the alloyDBUser, and returns the corresponding alloyDBUser object, and an error if there is any.
func (c *FakeAlloyDBUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AlloyDBUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alloydbusersResource, c.ns, name), &v1beta1.AlloyDBUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBUser), err
}

// List takes label and field selectors, and returns the list of AlloyDBUsers that match those selectors.
func (c *FakeAlloyDBUsers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AlloyDBUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alloydbusersResource, alloydbusersKind, c.ns, opts), &v1beta1.AlloyDBUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AlloyDBUserList{ListMeta: obj.(*v1beta1.AlloyDBUserList).ListMeta}
	for _, item := range obj.(*v1beta1.AlloyDBUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alloyDBUsers.
func (c *FakeAlloyDBUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alloydbusersResource, c.ns, opts))

}

// Create takes the representation of a alloyDBUser and creates it.  Returns the server's representation of the alloyDBUser, and an error, if there is any.
func (c *FakeAlloyDBUsers) Create(ctx context.Context, alloyDBUser *v1beta1.AlloyDBUser, opts v1.CreateOptions) (result *v1beta1.AlloyDBUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alloydbusersResource, c.ns, alloyDBUser), &v1beta1.AlloyDBUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBUser), err
}

// Update takes the representation of a alloyDBUser and updates it. Returns the server's representation of the alloyDBUser, and an error, if there is any.
func (c *FakeAlloyDBUsers) Update(ctx context.Context, alloyDBUser *v1beta1.AlloyDBUser, opts v1.UpdateOptions) (result *v1beta1.AlloyDBUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alloydbusersResource, c.ns, alloyDBUser), &v1beta1.AlloyDBUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlloyDBUsers) UpdateStatus(ctx context.Context, alloyDBUser *v1beta1.AlloyDBUser, opts v1.UpdateOptions) (*v1beta1.AlloyDBUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alloydbusersResource, "status", c.ns, alloyDBUser), &v1beta1.AlloyDBUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBUser), err
}

// Delete takes name of the alloyDBUser and deletes it. Returns an error if one occurs.
func (c *FakeAlloyDBUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(alloydbusersResource, c.ns, name, opts), &v1beta1.AlloyDBUser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlloyDBUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alloydbusersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AlloyDBUserList{})
	return err
}

// Patch applies the patch and returns the patched alloyDBUser.
func (c *FakeAlloyDBUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AlloyDBUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alloydbusersResource, c.ns, name, pt, data, subresources...), &v1beta1.AlloyDBUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AlloyDBUser), err
}
