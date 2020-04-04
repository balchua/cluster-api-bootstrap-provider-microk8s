/*


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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Microk8sConfigSpec defines the desired state of Microk8sConfig
type Microk8sConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Microk8sConfig. Edit Microk8sConfig_types.go to remove/update
	Tokens string `json:"token,omitempty"`

	Channel string `json:"channel,omitempty"`

	ControlplaneHost string `json:"controlPlaneHost,omitempty"`
}

// Microk8sConfigStatus defines the observed state of Microk8sConfig
type Microk8sConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DataSecretName is the name of the secret that stores the bootstrap data script.
	// +optional
	DataSecretName *string `json:"dataSecretName,omitempty"`

	// Ready indicates the BootstrapData field is ready to be consumed
	Ready bool `json:"ready,omitempty"`

	// BootstrapData will be a slice of bootstrap data
	// +optional
	BootstrapData []byte `json:"bootstrapData,omitempty"`

	// ErrorReason will be set on non-retryable errors
	// +optional
	ErrorReason string `json:"errorReason,omitempty"`

	// ErrorMessage will be set on non-retryable errors
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:path=microk8sconfigs,scope=Namespaced,categories=cluster-api
// +kubebuilder:subresource:status

// Microk8sConfig is the Schema for the microk8sconfigs API
type Microk8sConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Microk8sConfigSpec   `json:"spec,omitempty"`
	Status Microk8sConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Microk8sConfigList contains a list of Microk8sConfig
type Microk8sConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Microk8sConfig `json:"items"`
}

// Microk8sConfigTemplateResource defines the Template structure
type Microk8sConfigTemplateResource struct {
	Spec Microk8sConfigSpec `json:"spec,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Microk8sConfig{}, &Microk8sConfigList{})
}
