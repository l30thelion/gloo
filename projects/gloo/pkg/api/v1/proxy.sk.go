// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewProxy(namespace, name string) *Proxy {
	proxy := &Proxy{}
	proxy.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return proxy
}

func (r *Proxy) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Proxy) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Proxy) GroupVersionKind() schema.GroupVersionKind {
	return ProxyGVK
}

type ProxyList []*Proxy

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ProxyList) Find(namespace, name string) (*Proxy, error) {
	for _, proxy := range list {
		if proxy.GetMetadata().Name == name {
			if namespace == "" || proxy.GetMetadata().Namespace == namespace {
				return proxy, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find proxy %v.%v", namespace, name)
}

func (list ProxyList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, proxy := range list {
		ress = append(ress, proxy)
	}
	return ress
}

func (list ProxyList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, proxy := range list {
		ress = append(ress, proxy)
	}
	return ress
}

func (list ProxyList) Names() []string {
	var names []string
	for _, proxy := range list {
		names = append(names, proxy.GetMetadata().Name)
	}
	return names
}

func (list ProxyList) NamespacesDotNames() []string {
	var names []string
	for _, proxy := range list {
		names = append(names, proxy.GetMetadata().Namespace+"."+proxy.GetMetadata().Name)
	}
	return names
}

func (list ProxyList) Sort() ProxyList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list ProxyList) Clone() ProxyList {
	var proxyList ProxyList
	for _, proxy := range list {
		proxyList = append(proxyList, resources.Clone(proxy).(*Proxy))
	}
	return proxyList
}

func (list ProxyList) Each(f func(element *Proxy)) {
	for _, proxy := range list {
		f(proxy)
	}
}

func (list ProxyList) EachResource(f func(element resources.Resource)) {
	for _, proxy := range list {
		f(proxy)
	}
}

func (list ProxyList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Proxy) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Proxy

func (o *Proxy) GetObjectKind() schema.ObjectKind {
	t := ProxyCrd.TypeMeta()
	return &t
}

func (o *Proxy) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Proxy)
}

func (o *Proxy) DeepCopyInto(out *Proxy) {
	clone := resources.Clone(o).(*Proxy)
	*out = *clone
}

var (
	ProxyCrd = crd.NewCrd(
		"proxies",
		ProxyGVK.Group,
		ProxyGVK.Version,
		ProxyGVK.Kind,
		"px",
		false,
		&Proxy{})
)

func init() {
	if err := crd.AddCrd(ProxyCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	ProxyGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Proxy",
	}
)
