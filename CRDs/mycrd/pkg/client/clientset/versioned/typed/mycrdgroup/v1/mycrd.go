/*
   哇咔咔, 我是文件头!!!
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "mycrd/pkg/apis/mycrdgroup/v1"
	scheme "mycrd/pkg/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyCrdsGetter has a method to return a MyCrdInterface.
// A group's client should implement this interface.
type MyCrdsGetter interface {
	MyCrds(namespace string) MyCrdInterface
}

// MyCrdInterface has methods to work with MyCrd resources.
type MyCrdInterface interface {
	Create(*v1.MyCrd) (*v1.MyCrd, error)
	Update(*v1.MyCrd) (*v1.MyCrd, error)
	UpdateStatus(*v1.MyCrd) (*v1.MyCrd, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MyCrd, error)
	List(opts metav1.ListOptions) (*v1.MyCrdList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MyCrd, err error)
	MyCrdExpansion
}

// myCrds implements MyCrdInterface
type myCrds struct {
	client rest.Interface
	ns     string
}

// newMyCrds returns a MyCrds
func newMyCrds(c *MycrdgroupV1Client, namespace string) *myCrds {
	return &myCrds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myCrd, and returns the corresponding myCrd object, and an error if there is any.
func (c *myCrds) Get(name string, options metav1.GetOptions) (result *v1.MyCrd, err error) {
	result = &v1.MyCrd{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mycrds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MyCrds that match those selectors.
func (c *myCrds) List(opts metav1.ListOptions) (result *v1.MyCrdList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MyCrdList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mycrds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myCrds.
func (c *myCrds) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mycrds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a myCrd and creates it.  Returns the server's representation of the myCrd, and an error, if there is any.
func (c *myCrds) Create(myCrd *v1.MyCrd) (result *v1.MyCrd, err error) {
	result = &v1.MyCrd{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mycrds").
		Body(myCrd).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myCrd and updates it. Returns the server's representation of the myCrd, and an error, if there is any.
func (c *myCrds) Update(myCrd *v1.MyCrd) (result *v1.MyCrd, err error) {
	result = &v1.MyCrd{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mycrds").
		Name(myCrd.Name).
		Body(myCrd).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *myCrds) UpdateStatus(myCrd *v1.MyCrd) (result *v1.MyCrd, err error) {
	result = &v1.MyCrd{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mycrds").
		Name(myCrd.Name).
		SubResource("status").
		Body(myCrd).
		Do().
		Into(result)
	return
}

// Delete takes name of the myCrd and deletes it. Returns an error if one occurs.
func (c *myCrds) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mycrds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myCrds) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mycrds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myCrd.
func (c *myCrds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MyCrd, err error) {
	result = &v1.MyCrd{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mycrds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}