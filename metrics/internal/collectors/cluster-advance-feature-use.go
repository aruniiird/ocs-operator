package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/red-hat-storage/ocs-operator/metrics/internal/options"
	cephv1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	cephv1listers "github.com/rook/rook/pkg/client/listers/ceph.rook.io/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"

	storagev1 "k8s.io/api/storage/v1"
	storagev1listers "k8s.io/client-go/listers/storage/v1"
)

type AdvancedFeatureProvider interface {
	cache.SharedIndexInformer
	AdvancedFeature(namespaces ...string) int
}

type CephClusterAdvancedFeatureProvider struct {
	cache.SharedIndexInformer
}

func (c *CephClusterAdvancedFeatureProvider) AdvancedFeature(namespaces ...string) int {
	allCephClusters := getAllCephClusters(
		cephv1listers.NewCephClusterLister(c.GetIndexer()), namespaces)
	for _, cephCluster := range allCephClusters {
		if cephCluster.Spec.External.Enable {
			return 1
		} else if cephCluster.Spec.Security.KeyManagementService.IsEnabled() {
			return 1
		}
	}
	return 0
}

type CephObjectStoreAdvancedFeatureProvider struct {
	cache.SharedIndexInformer
}

func (c *CephObjectStoreAdvancedFeatureProvider) AdvancedFeature(namespaces ...string) int {
	cephObjectStoreLister := cephv1listers.NewCephObjectStoreLister(c.GetIndexer())
	cephObjectStores := getAllObjectStores(cephObjectStoreLister, namespaces)
	for _, cephObjectStore := range cephObjectStores {
		if cephObjectStore.Spec.Security.KeyManagementService.IsEnabled() {
			return 1
		}
	}
	return 0
}

type StorageClassAdvancedFeatureProvider struct {
	cache.SharedIndexInformer
}

func (s *StorageClassAdvancedFeatureProvider) AdvancedFeature(namespaces ...string) int {
	storageClassLister := storagev1listers.NewStorageClassLister(s.GetIndexer())
	storageClasses := getAllStorageClasses(storageClassLister, namespaces)
	for _, storageClass := range storageClasses {
		if storageClass.Parameters["encrypted"] == "true" {
			return 1
		}
	}
	return 0
}

type CephRBDMirrorAdvancedFeatureProvider struct {
	cache.SharedIndexInformer
}

func (c *CephRBDMirrorAdvancedFeatureProvider) AdvancedFeature(namespaces ...string) int {
	cephRBDMirrorLister := cephv1listers.NewCephRBDMirrorLister(c.GetIndexer())
	cephRBDMirrors := getAllRBDMirrors(cephRBDMirrorLister, namespaces)
	for _, rbdM := range cephRBDMirrors {
		if rbdM.Spec.Count > 0 {
			return 1
		}
	}
	return 0
}

type ClusterAdvanceFeatureCollector struct {
	AdvancedFeature     *prometheus.Desc
	AllowedNamespaces   []string
	advFeatureProviders []AdvancedFeatureProvider
}

const (
	// component within the project/exporter
	advFeatureSubSystem = "advanced_feature"
)

var _ prometheus.Collector = &ClusterAdvanceFeatureCollector{}

// NewClusterAdvancedFeatureCollector constructs the StorageCluster's advanced-feature collector
func NewClusterAdvancedFeatureCollector(opts *options.Options, sharedIndexInformerList ...cache.SharedIndexInformer) *ClusterAdvanceFeatureCollector {
	if sharedIndexInformerList == nil {
		sharedIndexInformerList = listAllSharedIndexInformers(opts)
	}
	advFeatureProviders := []AdvancedFeatureProvider{
		&CephClusterAdvancedFeatureProvider{SharedIndexInformer: sharedIndexInformerList[CephClusterSIIAI]},
		&CephObjectStoreAdvancedFeatureProvider{SharedIndexInformer: sharedIndexInformerList[CephObjectStoreSIIAI]},
		&StorageClassAdvancedFeatureProvider{SharedIndexInformer: sharedIndexInformerList[StorageClassSIIAI]},
		&CephRBDMirrorAdvancedFeatureProvider{SharedIndexInformer: sharedIndexInformerList[CephRBDMirrorSIIAI]},
	}

	return &ClusterAdvanceFeatureCollector{
		AdvancedFeature: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, advFeatureSubSystem, "usage"),
			`Indicates whether the cluster is using any advanced features, like PV/KMS encryption or external cluster mode`,
			nil, nil,
		),
		AllowedNamespaces:   opts.AllowedNamespaces,
		advFeatureProviders: advFeatureProviders,
	}
}

// Describe implements prometheus.Collector interface
func (c *ClusterAdvanceFeatureCollector) Describe(ch chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		c.AdvancedFeature,
	}

	for _, d := range ds {
		ch <- d
	}
}

// Collect implements prometheus.Collector interface
func (c *ClusterAdvanceFeatureCollector) Collect(ch chan<- prometheus.Metric) {
	for _, advProvider := range c.advFeatureProviders {
		// if any of the provider gives an advanced feature,
		// collect it and return immediately
		if advFeature := advProvider.AdvancedFeature(c.AllowedNamespaces...); advFeature > 0 {
			c.collectAdvancedFeatureUse(ch, advFeature)
			return
		}
	}
	c.collectAdvancedFeatureUse(ch, 0)
}

func (c *ClusterAdvanceFeatureCollector) collectAdvancedFeatureUse(ch chan<- prometheus.Metric, advancedFeature int) {
	ch <- prometheus.MustNewConstMetric(
		c.AdvancedFeature,
		prometheus.GaugeValue, float64(advancedFeature),
	)
}

func getAllStorageClasses(
	lister storagev1listers.StorageClassLister,
	namespaces []string) []*storagev1.StorageClass {
	var err error
	allSCs, err := lister.List(labels.Everything())
	if err != nil {
		klog.Errorf("couldn't list StorageClasses. %v", err)
		return nil
	}
	if len(namespaces) == 0 {
		return allSCs
	}
	var namespacedSCs []*storagev1.StorageClass
	for _, namespace := range namespaces {
		for _, eachSC := range allSCs {
			if eachSC.Namespace == namespace {
				namespacedSCs = append(namespacedSCs, eachSC)
			}
		}
	}
	return namespacedSCs
}

func getAllRBDMirrors(lister cephv1listers.CephRBDMirrorLister, namespaces []string) []*cephv1.CephRBDMirror {
	var err error
	allRBDMirrors, err := lister.List(labels.Everything())
	if err != nil {
		klog.Errorf("couldn't list RBD Mirrors. %v", err)
		return nil
	}
	if len(namespaces) == 0 {
		return allRBDMirrors
	}
	var namespacedRBDMirrors []*cephv1.CephRBDMirror
	for _, namespace := range namespaces {
		for _, eachRBDMirror := range allRBDMirrors {
			if eachRBDMirror.Namespace == namespace {
				namespacedRBDMirrors = append(namespacedRBDMirrors, eachRBDMirror)
			}
		}
	}
	return namespacedRBDMirrors
}
