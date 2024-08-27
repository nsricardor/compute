/*
Copyright 2024 the Unikorn Authors.

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

package v1alpha1

import (
	unikornv1core "github.com/unikorn-cloud/core/pkg/apis/unikorn/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MachineGeneric contains common things across all pool types, including
// Baremetal cluster manager nodes and workload pools.
type MachineGeneric struct {
	// Image is the OpenStack Glance image to deploy with.
	ImageID *string `json:"imageId"`
	// Flavor is the OpenStack Nova flavor to deploy with.
	FlavorID *string `json:"flavorId"`
	// FlavorName is the name of the flavor.
	// CAPO is broken and doesn't accept an ID, so we need to use this.
	FlavorName *string `json:"flavorName"`
	// Replicas is the initial pool size to deploy.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=3
	Replicas *int `json:"replicas,omitempty"`
}

// BaremetalWorkloadPoolSpec defines the requested machine pool
// state.
type BaremetalWorkloadPoolSpec struct {
	MachineGeneric `json:",inline"`
	// Name is the name of the pool.
	Name string `json:"name"`
}

// BaremetalClusterList is a typed list of baremetal clusters.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BaremetalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaremetalCluster `json:"items"`
}

// BaremetalCluster is an object representing a Baremetal cluster.
// For now, this is a monolith for simplicity.  In future it may reference
// a provider specific implementation e.g. if CAPI goes out of favour for
// some other new starlet.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced,categories=unikorn
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="display name",type="string",JSONPath=".metadata.labels['unikorn-cloud\\.org/name']"
// +kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.conditions[?(@.type==\"Available\")].reason"
// +kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"
type BaremetalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BaremetalClusterSpec   `json:"spec"`
	Status            BaremetalClusterStatus `json:"status,omitempty"`
}

// BaremetalClusterSpec defines the requested state of the Baremetal cluster.
type BaremetalClusterSpec struct {
	// Pause, if true, will inhibit reconciliation.
	Pause bool `json:"pause,omitempty"`
	// Region to provision the cluster in.
	RegionID string `json:"regionId"`
	// Network defines the Baremetal networking.
	Network *BaremetalClusterNetworkSpec `json:"network"`
	// WorkloadPools defines the workload cluster topology.
	WorkloadPools *BaremetalClusterWorkloadPoolsSpec `json:"workloadPools"`
}

type BaremetalClusterNetworkSpec struct {
	// NodeNetwork is the IPv4 prefix for the node network.
	NodeNetwork *unikornv1core.IPv4Prefix `json:"nodeNetwork"`
	// DNSNameservers sets the DNS nameservers for pods.
	// At present due to some technical challenges, this must contain
	// only one DNS server.
	// +listType=set
	// +kubebuilder:validation:MinItems=1
	DNSNameservers []unikornv1core.IPv4Address `json:"dnsNameservers"`
}

type BaremetalClusterWorkloadPoolsPoolSpec struct {
	BaremetalWorkloadPoolSpec `json:",inline"`
}

type BaremetalClusterWorkloadPoolsSpec struct {
	// Pools contains an inline set of pools.  This field will be ignored
	// when Selector is set.  Inline pools are expected to be used for UI
	// generated clusters.
	Pools []BaremetalClusterWorkloadPoolsPoolSpec `json:"pools,omitempty"`
}

// BaremetalClusterStatus defines the observed state of the Baremetal cluster.
type BaremetalClusterStatus struct {
	// Namespace defines the namespace a cluster resides in.
	Namespace string `json:"namespace,omitempty"`

	// Current service state of a Baremetal cluster.
	Conditions []unikornv1core.Condition `json:"conditions,omitempty"`
}