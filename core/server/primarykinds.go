package server

import (
	"fmt"

	helmv2 "github.com/fluxcd/helm-controller/api/v2beta1"
	kustomizev1 "github.com/fluxcd/kustomize-controller/api/v1beta2"
	notificationv1 "github.com/fluxcd/notification-controller/api/v1beta1"
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type PrimaryKinds struct {
	kinds map[string]schema.GroupVersionKind
}

func New() *PrimaryKinds {
	kinds := PrimaryKinds{}
	kinds.kinds = make(map[string]schema.GroupVersionKind)

	return &kinds
}

func DefaultPrimaryKinds() *PrimaryKinds {
	kinds := New()
	kinds.kinds[kustomizev1.KustomizationKind] = kustomizev1.GroupVersion.WithKind(kustomizev1.KustomizationKind)
	kinds.kinds[helmv2.HelmReleaseKind] = helmv2.GroupVersion.WithKind(helmv2.HelmReleaseKind)
	kinds.kinds[sourcev1.GitRepositoryKind] = sourcev1.GroupVersion.WithKind(sourcev1.GitRepositoryKind)
	kinds.kinds[sourcev1.HelmChartKind] = sourcev1.GroupVersion.WithKind(sourcev1.HelmChartKind)
	kinds.kinds[sourcev1.HelmRepositoryKind] = sourcev1.GroupVersion.WithKind(sourcev1.HelmRepositoryKind)
	kinds.kinds[sourcev1.BucketKind] = sourcev1.GroupVersion.WithKind(sourcev1.BucketKind)
	kinds.kinds[sourcev1.OCIRepositoryKind] = sourcev1.GroupVersion.WithKind(sourcev1.OCIRepositoryKind)
	kinds.kinds[notificationv1.ProviderKind] = notificationv1.GroupVersion.WithKind(notificationv1.ProviderKind)
	kinds.kinds[notificationv1.AlertKind] = notificationv1.GroupVersion.WithKind(notificationv1.AlertKind)

	return kinds
}

// Add sets another kind name and gvk to resolve an object
// This errors if the kind is already set, as this likely indicates 2
// different uses for the same kind string.
func (pk *PrimaryKinds) Add(kind string, gvk schema.GroupVersionKind) error {
	_, ok := pk.kinds[kind]
	if ok {
		return fmt.Errorf("couldn't add kind %v - already added", kind)
	}

	pk.kinds[kind] = gvk

	return nil
}

// Lookup ensures that a kind name is known, white-listed, and returns
// the full GVK for that kind
func (pk *PrimaryKinds) Lookup(kind string) (*schema.GroupVersionKind, error) {
	gvk, ok := pk.kinds[kind]
	if !ok {
		return nil, fmt.Errorf("looking up objects of kind %v not supported", kind)
	}

	return &gvk, nil
}
