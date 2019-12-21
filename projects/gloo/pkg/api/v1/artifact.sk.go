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

func NewArtifact(namespace, name string) *Artifact {
	artifact := &Artifact{}
	artifact.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return artifact
}

func (r *Artifact) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Artifact) GroupVersionKind() schema.GroupVersionKind {
	return ArtifactGVK
}

type ArtifactList []*Artifact

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ArtifactList) Find(namespace, name string) (*Artifact, error) {
	for _, artifact := range list {
		if artifact.GetMetadata().Name == name {
			if namespace == "" || artifact.GetMetadata().Namespace == namespace {
				return artifact, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find artifact %v.%v", namespace, name)
}

func (list ArtifactList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, artifact := range list {
		ress = append(ress, artifact)
	}
	return ress
}

func (list ArtifactList) Names() []string {
	var names []string
	for _, artifact := range list {
		names = append(names, artifact.GetMetadata().Name)
	}
	return names
}

func (list ArtifactList) NamespacesDotNames() []string {
	var names []string
	for _, artifact := range list {
		names = append(names, artifact.GetMetadata().Namespace+"."+artifact.GetMetadata().Name)
	}
	return names
}

func (list ArtifactList) Sort() ArtifactList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list ArtifactList) Clone() ArtifactList {
	var artifactList ArtifactList
	for _, artifact := range list {
		artifactList = append(artifactList, resources.Clone(artifact).(*Artifact))
	}
	return artifactList
}

func (list ArtifactList) Each(f func(element *Artifact)) {
	for _, artifact := range list {
		f(artifact)
	}
}

func (list ArtifactList) EachResource(f func(element resources.Resource)) {
	for _, artifact := range list {
		f(artifact)
	}
}

func (list ArtifactList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Artifact) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Artifact

func (o *Artifact) GetObjectKind() schema.ObjectKind {
	t := ArtifactCrd.TypeMeta()
	return &t
}

func (o *Artifact) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Artifact)
}

func (o *Artifact) DeepCopyInto(out *Artifact) {
	clone := resources.Clone(o).(*Artifact)
	*out = *clone
}

var (
	ArtifactCrd = crd.NewCrd(
		"artifacts",
		ArtifactGVK.Group,
		ArtifactGVK.Version,
		ArtifactGVK.Kind,
		"art",
		false,
		&Artifact{})
)

func init() {
	if err := crd.AddCrd(ArtifactCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	ArtifactGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Artifact",
	}
)
