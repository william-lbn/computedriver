/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type ApprovedNodeInformation struct {
	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	IP string `json:"ip,omitempty"`

	Description string `json:"description,omitempty"`
}

// HypernodeSpec defines the desired state of Hypernode
type HypernodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ApprovedNodeInformations ApprovedNodeInformation `json:"ApprovedNodeInformations,omitempty"`
}

// HypernodeStatus defines the observed state of Hypernode
type HypernodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Hypernode is the Schema for the hypernodes API
type Hypernode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HypernodeSpec   `json:"spec,omitempty"`
	Status HypernodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HypernodeList contains a list of Hypernode
type HypernodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hypernode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hypernode{}, &HypernodeList{})
}
