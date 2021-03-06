/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

import (
	"encoding/json"
	"fmt"
	"strings"

	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ServicePlanGVK is the GroupVersionKind for ServicePlan
var ServicePlanGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "ServicePlan",
}

// ServicePlanSpec defines the desired state of Service plan
// +k8s:openapi-gen=true
type ServicePlanSpec struct {
	// Kind refers to the service type this is a plan for
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// Labels is a collection of labels for this plan
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
	// Description provides a summary of the configuration provided by this plan
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Description string `json:"description"`
	// Summary provides a short title summary for the plan
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Summary string `json:"summary"`
	// Configuration are the key+value pairs describing a service configuration
	// +kubebuilder:validation:Type=object
	// +kubebuilder:validation:Optional
	Configuration *apiextv1.JSON `json:"configuration,omitempty"`
}

func (s *ServicePlanSpec) GetConfiguration(v interface{}) error {
	if s.Configuration == nil {
		return nil
	}

	if err := json.Unmarshal(s.Configuration.Raw, v); err != nil {
		return fmt.Errorf("failed to unmarshal service plan configuration: %w", err)
	}
	return nil
}

func (s *ServicePlanSpec) SetConfiguration(v interface{}) error {
	if v == nil {
		s.Configuration = nil
		return nil
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal service plan configuration: %w", err)
	}
	s.Configuration = &apiextv1.JSON{Raw: raw}
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicePlan is a template for a service
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=serviceplans
type ServicePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServicePlanSpec `json:"spec,omitempty"`
}

func NewServicePlan(name, namespace string) *ServicePlan {
	return &ServicePlan{
		TypeMeta: metav1.TypeMeta{
			Kind:       ServicePlanGVK.Kind,
			APIVersion: GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

// PlanShortName returns the plan name without the service kind prefix
func (s ServicePlan) PlanShortName() string {
	return strings.TrimPrefix(s.Name, s.Spec.Kind+"-")
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicePlanList contains a list of service plans
type ServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePlan `json:"items"`
}
