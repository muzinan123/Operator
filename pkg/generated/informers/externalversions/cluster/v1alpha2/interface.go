// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/clusterpedia-io/clusterpedia/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterSyncResources returns a ClusterSyncResourcesInformer.
	ClusterSyncResources() ClusterSyncResourcesInformer
	// PediaClusters returns a PediaClusterInformer.
	PediaClusters() PediaClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterSyncResources returns a ClusterSyncResourcesInformer.
func (v *version) ClusterSyncResources() ClusterSyncResourcesInformer {
	return &clusterSyncResourcesInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PediaClusters returns a PediaClusterInformer.
func (v *version) PediaClusters() PediaClusterInformer {
	return &pediaClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
