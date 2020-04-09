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

package kubernetes

import (
	"context"
	"time"

	"github.com/appvia/kore/pkg/utils"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClearOutCluster is responsible for deleting all resources
func ClearOutCluster(ctx context.Context, cc client.Client) error {
	var resources []runtime.Object

	// @step: retrieve a complete list of namespaces
	list := &v1.NamespaceList{}
	if err := cc.List(ctx, list); err != nil {
		return err
	}

	// @step: iterate the list, filter out those namespace we can't remove
	for _, namespace := range list.Items {
		if utils.Contains(namespace.Name, SystemNamespaces()) {
			continue
		}
		resources = append(resources, namespace.DeepCopyObject())
	}

	// @step: delete services from default namespace as well
	services := &v1.ServiceList{}
	for _, service := range services.Items {
		if utils.Contains(service.Name, []string{"kubernetes"}) {
			continue
		}

		resources = append(resources, service.DeepCopyObject())
	}

	return DeleteResources(ctx, cc, resources)
}

// DeleteResources is responsible for deleting all the resources
func DeleteResources(ctx context.Context, cc client.Client, resources []runtime.Object) error {
	// @step: issue a delete on all the resources
	for _, x := range resources {
		if err := cc.Delete(ctx, x.DeepCopyObject()); err != nil {
			return err
		}
	}

	// @step: wait for all the resources to disappear
	for _, x := range resources {
		key, err := client.ObjectKeyFromObject(x)
		if err != nil {
			return err
		}

		// we wait for the the resource to disappear
		err = utils.WaitUntilComplete(ctx, 10*time.Minute, 5*time.Second, func() (bool, error) {
			if err := cc.Get(ctx, key, x.DeepCopyObject()); err != nil {
				if errors.IsNotFound(err) {
					return true, nil
				}
			}

			return false, nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// SystemNamespaces is a collection of system namespaces which can't be removed
func SystemNamespaces() []string {
	return []string{
		"default",
		"kube-node-lease",
		"kube-public",
		"kube-system",
	}
}
