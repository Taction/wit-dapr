/*
Copyright 2021 The Dapr Authors
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

package host_state

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Component struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec ComponentSpec `json:"spec,omitempty"`
	// +optional
	//Auth `json:"auth,omitempty"`
}

// ComponentSpec is the spec for a component
type ComponentSpec struct {
	Type     string         `json:"type"`
	Version  string         `json:"version"`
	Metadata []MetadataItem `json:"metadata"`
}

// MetadataItem is a name/value pair for a metadata
type MetadataItem struct {
	Name         string       `json:"name"`
	Value        string       `json:"value"`
	SecretKeyRef SecretKeyRef `json:"secretKeyRef,omitempty"`
}

// SecretKeyRef is a reference to a secret holding the value for the metadata item. Name is the secret name, and key is the field in the secret.
type SecretKeyRef struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Auth represents authentication details for the component
type Auth struct {
	SecretStore string `json:"secretStore"`
}

// +kubebuilder:object:root=true

// ComponentList is a list of Dapr components
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Component `json:"items"`
}
