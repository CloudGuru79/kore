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

package openservicebroker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/controllers"

	"github.com/google/uuid"
	osb "github.com/kubernetes-sigs/go-open-service-broker-client/v2"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (p *Provider) ReconcileCredentials(
	ctx context.Context,
	logger logrus.FieldLogger,
	service *servicesv1.Service,
	plan *servicesv1.ServicePlan,
	creds *servicesv1.ServiceCredentials,
) (reconcile.Result, map[string]string, error) {
	providerPlan, err := p.plan(service.Spec.Kind, plan.Name)
	if err != nil {
		return reconcile.Result{}, nil, err
	}

	component, _ := creds.Status.Components.GetComponent(ComponentBind)
	if component == nil {
		component = &corev1.Component{
			Name:   ComponentBind,
			Status: corev1.Unknown,
		}
	}
	defer func() {
		creds.Status.Components.SetCondition(*component)
	}()

	if component.Status == corev1.SuccessStatus {
		return reconcile.Result{}, nil, fmt.Errorf("ReconcileCredentials was called after successful state, this should not happen")
	}

	if component.Status == corev1.PendingStatus {
		return p.pollLastBindingOperation(ctx, logger, service, plan, creds, component)
	}

	component.Update(corev1.PendingStatus, "", "")

	if creds.Status.ProviderID == "" {
		creds.Status.ProviderID = uuid.New().String()
	}

	config := map[string]interface{}{}
	if err := json.Unmarshal(creds.Spec.Configuration.Raw, &config); err != nil {
		return reconcile.Result{}, nil, controllers.NewCriticalError(fmt.Errorf("failed to unmarshal service credentials configuration"))
	}

	logger.Debug("calling bind on service broker")

	resp, err := p.client.Bind(&osb.BindRequest{
		AcceptsIncomplete: true,
		BindingID:         creds.Status.ProviderID,
		InstanceID:        service.Status.ProviderID,
		ServiceID:         providerPlan.serviceID,
		PlanID:            providerPlan.id,
		Parameters:        config,
	})
	if err != nil {
		return reconcile.Result{}, nil, handleError(component, "failed to call bind on the service broker", err)
	}

	filteredResponse := map[string]interface{}{
		"operation": resp.OperationKey,
		"async":     resp.Async,
	}

	logger.WithField("response", filteredResponse).Debug("bind response from service broker")

	creds.Status.ProviderData, err = encodeProviderData(resp.OperationKey)
	if err != nil {
		return reconcile.Result{}, nil, err
	}

	if !resp.Async {
		bindingCredentials, err := bindingCredentialsToStringMap(resp.Credentials)
		if err != nil {
			return reconcile.Result{}, nil, controllers.NewCriticalError(fmt.Errorf("failed to encode binding credentials from the service broker: %w", err))
		}

		component.Update(corev1.SuccessStatus, "", "")
		return reconcile.Result{}, bindingCredentials, nil
	}

	return reconcile.Result{RequeueAfter: 5 * time.Second}, nil, nil
}
