// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrometheusAgentSpecApplyConfiguration represents an declarative configuration of the PrometheusAgentSpec type for use
// with apply.
type PrometheusAgentSpecApplyConfiguration struct {
	v1.CommonPrometheusFieldsApplyConfiguration `json:",inline"`
}

// PrometheusAgentSpecApplyConfiguration constructs an declarative configuration of the PrometheusAgentSpec type for use with
// apply.
func PrometheusAgentSpec() *PrometheusAgentSpecApplyConfiguration {
	return &PrometheusAgentSpecApplyConfiguration{}
}

// WithPodMetadata sets the PodMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodMetadata field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPodMetadata(value *v1.EmbeddedObjectMetadataApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	b.PodMetadata = value
	return b
}

// WithServiceMonitorSelector sets the ServiceMonitorSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceMonitorSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithServiceMonitorSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.ServiceMonitorSelector = &value
	return b
}

// WithServiceMonitorNamespaceSelector sets the ServiceMonitorNamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceMonitorNamespaceSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithServiceMonitorNamespaceSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.ServiceMonitorNamespaceSelector = &value
	return b
}

// WithPodMonitorSelector sets the PodMonitorSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodMonitorSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPodMonitorSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.PodMonitorSelector = &value
	return b
}

// WithPodMonitorNamespaceSelector sets the PodMonitorNamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodMonitorNamespaceSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPodMonitorNamespaceSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.PodMonitorNamespaceSelector = &value
	return b
}

// WithProbeSelector sets the ProbeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProbeSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithProbeSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.ProbeSelector = &value
	return b
}

// WithProbeNamespaceSelector sets the ProbeNamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProbeNamespaceSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithProbeNamespaceSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.ProbeNamespaceSelector = &value
	return b
}

// WithScrapeConfigSelector sets the ScrapeConfigSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeConfigSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithScrapeConfigSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.ScrapeConfigSelector = &value
	return b
}

// WithScrapeConfigNamespaceSelector sets the ScrapeConfigNamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeConfigNamespaceSelector field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithScrapeConfigNamespaceSelector(value metav1.LabelSelector) *PrometheusAgentSpecApplyConfiguration {
	b.ScrapeConfigNamespaceSelector = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithVersion(value string) *PrometheusAgentSpecApplyConfiguration {
	b.Version = &value
	return b
}

// WithPaused sets the Paused field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Paused field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPaused(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.Paused = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithImage(value string) *PrometheusAgentSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithImagePullPolicy(value corev1.PullPolicy) *PrometheusAgentSpecApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *PrometheusAgentSpecApplyConfiguration) WithImagePullSecrets(values ...corev1.LocalObjectReference) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.ImagePullSecrets = append(b.ImagePullSecrets, values[i])
	}
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithReplicas(value int32) *PrometheusAgentSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithShards sets the Shards field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Shards field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithShards(value int32) *PrometheusAgentSpecApplyConfiguration {
	b.Shards = &value
	return b
}

// WithReplicaExternalLabelName sets the ReplicaExternalLabelName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReplicaExternalLabelName field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithReplicaExternalLabelName(value string) *PrometheusAgentSpecApplyConfiguration {
	b.ReplicaExternalLabelName = &value
	return b
}

// WithPrometheusExternalLabelName sets the PrometheusExternalLabelName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrometheusExternalLabelName field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPrometheusExternalLabelName(value string) *PrometheusAgentSpecApplyConfiguration {
	b.PrometheusExternalLabelName = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithLogLevel(value string) *PrometheusAgentSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithLogFormat sets the LogFormat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogFormat field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithLogFormat(value string) *PrometheusAgentSpecApplyConfiguration {
	b.LogFormat = &value
	return b
}

// WithScrapeInterval sets the ScrapeInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeInterval field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithScrapeInterval(value monitoringv1.Duration) *PrometheusAgentSpecApplyConfiguration {
	b.ScrapeInterval = &value
	return b
}

// WithScrapeTimeout sets the ScrapeTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScrapeTimeout field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithScrapeTimeout(value monitoringv1.Duration) *PrometheusAgentSpecApplyConfiguration {
	b.ScrapeTimeout = &value
	return b
}

// WithExternalLabels puts the entries into the ExternalLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ExternalLabels field,
// overwriting an existing map entries in ExternalLabels field with the same key.
func (b *PrometheusAgentSpecApplyConfiguration) WithExternalLabels(entries map[string]string) *PrometheusAgentSpecApplyConfiguration {
	if b.ExternalLabels == nil && len(entries) > 0 {
		b.ExternalLabels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ExternalLabels[k] = v
	}
	return b
}

// WithEnableRemoteWriteReceiver sets the EnableRemoteWriteReceiver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableRemoteWriteReceiver field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnableRemoteWriteReceiver(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.EnableRemoteWriteReceiver = &value
	return b
}

// WithEnableFeatures adds the given value to the EnableFeatures field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EnableFeatures field.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnableFeatures(values ...string) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.EnableFeatures = append(b.EnableFeatures, values[i])
	}
	return b
}

// WithExternalURL sets the ExternalURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalURL field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithExternalURL(value string) *PrometheusAgentSpecApplyConfiguration {
	b.ExternalURL = &value
	return b
}

// WithRoutePrefix sets the RoutePrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoutePrefix field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithRoutePrefix(value string) *PrometheusAgentSpecApplyConfiguration {
	b.RoutePrefix = &value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithStorage(value *v1.StorageSpecApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	b.Storage = value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *PrometheusAgentSpecApplyConfiguration) WithVolumes(values ...corev1.Volume) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.Volumes = append(b.Volumes, values[i])
	}
	return b
}

// WithVolumeMounts adds the given value to the VolumeMounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VolumeMounts field.
func (b *PrometheusAgentSpecApplyConfiguration) WithVolumeMounts(values ...corev1.VolumeMount) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.VolumeMounts = append(b.VolumeMounts, values[i])
	}
	return b
}

// WithWeb sets the Web field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Web field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithWeb(value *v1.PrometheusWebSpecApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	b.Web = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *PrometheusAgentSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *PrometheusAgentSpecApplyConfiguration) WithNodeSelector(entries map[string]string) *PrometheusAgentSpecApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithServiceAccountName(value string) *PrometheusAgentSpecApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithSecrets adds the given value to the Secrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Secrets field.
func (b *PrometheusAgentSpecApplyConfiguration) WithSecrets(values ...string) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.Secrets = append(b.Secrets, values[i])
	}
	return b
}

// WithConfigMaps adds the given value to the ConfigMaps field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConfigMaps field.
func (b *PrometheusAgentSpecApplyConfiguration) WithConfigMaps(values ...string) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.ConfigMaps = append(b.ConfigMaps, values[i])
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithAffinity(value corev1.Affinity) *PrometheusAgentSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *PrometheusAgentSpecApplyConfiguration) WithTolerations(values ...corev1.Toleration) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *PrometheusAgentSpecApplyConfiguration) WithTopologySpreadConstraints(values ...corev1.TopologySpreadConstraint) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, values[i])
	}
	return b
}

// WithRemoteWrite adds the given value to the RemoteWrite field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RemoteWrite field.
func (b *PrometheusAgentSpecApplyConfiguration) WithRemoteWrite(values ...*v1.RemoteWriteSpecApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRemoteWrite")
		}
		b.RemoteWrite = append(b.RemoteWrite, *values[i])
	}
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithSecurityContext(value corev1.PodSecurityContext) *PrometheusAgentSpecApplyConfiguration {
	b.SecurityContext = &value
	return b
}

// WithListenLocal sets the ListenLocal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ListenLocal field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithListenLocal(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.ListenLocal = &value
	return b
}

// WithContainers adds the given value to the Containers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Containers field.
func (b *PrometheusAgentSpecApplyConfiguration) WithContainers(values ...corev1.Container) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.Containers = append(b.Containers, values[i])
	}
	return b
}

// WithInitContainers adds the given value to the InitContainers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitContainers field.
func (b *PrometheusAgentSpecApplyConfiguration) WithInitContainers(values ...corev1.Container) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.InitContainers = append(b.InitContainers, values[i])
	}
	return b
}

// WithAdditionalScrapeConfigs sets the AdditionalScrapeConfigs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AdditionalScrapeConfigs field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithAdditionalScrapeConfigs(value corev1.SecretKeySelector) *PrometheusAgentSpecApplyConfiguration {
	b.AdditionalScrapeConfigs = &value
	return b
}

// WithAPIServerConfig sets the APIServerConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerConfig field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithAPIServerConfig(value *v1.APIServerConfigApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	b.APIServerConfig = value
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPriorityClassName(value string) *PrometheusAgentSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithPortName sets the PortName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortName field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithPortName(value string) *PrometheusAgentSpecApplyConfiguration {
	b.PortName = &value
	return b
}

// WithArbitraryFSAccessThroughSMs sets the ArbitraryFSAccessThroughSMs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ArbitraryFSAccessThroughSMs field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithArbitraryFSAccessThroughSMs(value *v1.ArbitraryFSAccessThroughSMsConfigApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	b.ArbitraryFSAccessThroughSMs = value
	return b
}

// WithOverrideHonorLabels sets the OverrideHonorLabels field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OverrideHonorLabels field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithOverrideHonorLabels(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.OverrideHonorLabels = &value
	return b
}

// WithOverrideHonorTimestamps sets the OverrideHonorTimestamps field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OverrideHonorTimestamps field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithOverrideHonorTimestamps(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.OverrideHonorTimestamps = &value
	return b
}

// WithIgnoreNamespaceSelectors sets the IgnoreNamespaceSelectors field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnoreNamespaceSelectors field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithIgnoreNamespaceSelectors(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.IgnoreNamespaceSelectors = &value
	return b
}

// WithEnforcedNamespaceLabel sets the EnforcedNamespaceLabel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedNamespaceLabel field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedNamespaceLabel(value string) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedNamespaceLabel = &value
	return b
}

// WithEnforcedSampleLimit sets the EnforcedSampleLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedSampleLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedSampleLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedSampleLimit = &value
	return b
}

// WithEnforcedTargetLimit sets the EnforcedTargetLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedTargetLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedTargetLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedTargetLimit = &value
	return b
}

// WithEnforcedLabelLimit sets the EnforcedLabelLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedLabelLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedLabelLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedLabelLimit = &value
	return b
}

// WithEnforcedLabelNameLengthLimit sets the EnforcedLabelNameLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedLabelNameLengthLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedLabelNameLengthLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedLabelNameLengthLimit = &value
	return b
}

// WithEnforcedLabelValueLengthLimit sets the EnforcedLabelValueLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedLabelValueLengthLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedLabelValueLengthLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedLabelValueLengthLimit = &value
	return b
}

// WithEnforcedBodySizeLimit sets the EnforcedBodySizeLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnforcedBodySizeLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithEnforcedBodySizeLimit(value monitoringv1.ByteSize) *PrometheusAgentSpecApplyConfiguration {
	b.EnforcedBodySizeLimit = &value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithMinReadySeconds(value uint32) *PrometheusAgentSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithHostAliases adds the given value to the HostAliases field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the HostAliases field.
func (b *PrometheusAgentSpecApplyConfiguration) WithHostAliases(values ...*v1.HostAliasApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHostAliases")
		}
		b.HostAliases = append(b.HostAliases, *values[i])
	}
	return b
}

// WithAdditionalArgs adds the given value to the AdditionalArgs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalArgs field.
func (b *PrometheusAgentSpecApplyConfiguration) WithAdditionalArgs(values ...*v1.ArgumentApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAdditionalArgs")
		}
		b.AdditionalArgs = append(b.AdditionalArgs, *values[i])
	}
	return b
}

// WithWALCompression sets the WALCompression field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WALCompression field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithWALCompression(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.WALCompression = &value
	return b
}

// WithExcludedFromEnforcement adds the given value to the ExcludedFromEnforcement field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExcludedFromEnforcement field.
func (b *PrometheusAgentSpecApplyConfiguration) WithExcludedFromEnforcement(values ...*v1.ObjectReferenceApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExcludedFromEnforcement")
		}
		b.ExcludedFromEnforcement = append(b.ExcludedFromEnforcement, *values[i])
	}
	return b
}

// WithHostNetwork sets the HostNetwork field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostNetwork field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithHostNetwork(value bool) *PrometheusAgentSpecApplyConfiguration {
	b.HostNetwork = &value
	return b
}

// WithPodTargetLabels adds the given value to the PodTargetLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PodTargetLabels field.
func (b *PrometheusAgentSpecApplyConfiguration) WithPodTargetLabels(values ...string) *PrometheusAgentSpecApplyConfiguration {
	for i := range values {
		b.PodTargetLabels = append(b.PodTargetLabels, values[i])
	}
	return b
}

// WithTracingConfig sets the TracingConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TracingConfig field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithTracingConfig(value *v1.PrometheusTracingConfigApplyConfiguration) *PrometheusAgentSpecApplyConfiguration {
	b.TracingConfig = value
	return b
}

// WithBodySizeLimit sets the BodySizeLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BodySizeLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithBodySizeLimit(value monitoringv1.ByteSize) *PrometheusAgentSpecApplyConfiguration {
	b.BodySizeLimit = &value
	return b
}

// WithSampleLimit sets the SampleLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampleLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithSampleLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.SampleLimit = &value
	return b
}

// WithTargetLimit sets the TargetLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithTargetLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.TargetLimit = &value
	return b
}

// WithLabelLimit sets the LabelLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithLabelLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.LabelLimit = &value
	return b
}

// WithLabelNameLengthLimit sets the LabelNameLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelNameLengthLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithLabelNameLengthLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.LabelNameLengthLimit = &value
	return b
}

// WithLabelValueLengthLimit sets the LabelValueLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelValueLengthLimit field is set to the value of the last call.
func (b *PrometheusAgentSpecApplyConfiguration) WithLabelValueLengthLimit(value uint64) *PrometheusAgentSpecApplyConfiguration {
	b.LabelValueLengthLimit = &value
	return b
}
