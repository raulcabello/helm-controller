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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	helmcattleiov1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHelmCharts implements HelmChartInterface
type FakeHelmCharts struct {
	Fake *FakeHelmV1
	ns   string
}

var helmchartsResource = schema.GroupVersionResource{Group: "helm.cattle.io", Version: "v1", Resource: "helmcharts"}

var helmchartsKind = schema.GroupVersionKind{Group: "helm.cattle.io", Version: "v1", Kind: "HelmChart"}

// Get takes name of the helmChart, and returns the corresponding helmChart object, and an error if there is any.
func (c *FakeHelmCharts) Get(ctx context.Context, name string, options v1.GetOptions) (result *helmcattleiov1.HelmChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(helmchartsResource, c.ns, name), &helmcattleiov1.HelmChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*helmcattleiov1.HelmChart), err
}

// List takes label and field selectors, and returns the list of HelmCharts that match those selectors.
func (c *FakeHelmCharts) List(ctx context.Context, opts v1.ListOptions) (result *helmcattleiov1.HelmChartList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(helmchartsResource, helmchartsKind, c.ns, opts), &helmcattleiov1.HelmChartList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &helmcattleiov1.HelmChartList{ListMeta: obj.(*helmcattleiov1.HelmChartList).ListMeta}
	for _, item := range obj.(*helmcattleiov1.HelmChartList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helmCharts.
func (c *FakeHelmCharts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(helmchartsResource, c.ns, opts))

}

// Create takes the representation of a helmChart and creates it.  Returns the server's representation of the helmChart, and an error, if there is any.
func (c *FakeHelmCharts) Create(ctx context.Context, helmChart *helmcattleiov1.HelmChart, opts v1.CreateOptions) (result *helmcattleiov1.HelmChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(helmchartsResource, c.ns, helmChart), &helmcattleiov1.HelmChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*helmcattleiov1.HelmChart), err
}

// Update takes the representation of a helmChart and updates it. Returns the server's representation of the helmChart, and an error, if there is any.
func (c *FakeHelmCharts) Update(ctx context.Context, helmChart *helmcattleiov1.HelmChart, opts v1.UpdateOptions) (result *helmcattleiov1.HelmChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(helmchartsResource, c.ns, helmChart), &helmcattleiov1.HelmChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*helmcattleiov1.HelmChart), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelmCharts) UpdateStatus(ctx context.Context, helmChart *helmcattleiov1.HelmChart, opts v1.UpdateOptions) (*helmcattleiov1.HelmChart, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(helmchartsResource, "status", c.ns, helmChart), &helmcattleiov1.HelmChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*helmcattleiov1.HelmChart), err
}

// Delete takes name of the helmChart and deletes it. Returns an error if one occurs.
func (c *FakeHelmCharts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(helmchartsResource, c.ns, name), &helmcattleiov1.HelmChart{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelmCharts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(helmchartsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &helmcattleiov1.HelmChartList{})
	return err
}

// Patch applies the patch and returns the patched helmChart.
func (c *FakeHelmCharts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *helmcattleiov1.HelmChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(helmchartsResource, c.ns, name, pt, data, subresources...), &helmcattleiov1.HelmChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*helmcattleiov1.HelmChart), err
}
