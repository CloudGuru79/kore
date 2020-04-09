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

package eks

import (
	"context"
	"time"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	eksv1alpha1 "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	"github.com/appvia/kore/pkg/utils/cloud/aws"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Delete is responsible for deleting the aws eks cluster
func (t *eksCtrl) Delete(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()

	logger := log.WithFields(log.Fields{
		"name":      request.NamespacedName.Name,
		"namespace": request.NamespacedName.Namespace,
		"team":      request.NamespacedName.Name,
	})
	logger.Info("attempting to delete eks cluster")

	// @step: retrieve the object from the api
	resource := &eksv1alpha1.EKS{}
	if err := t.mgr.GetClient().Get(ctx, request.NamespacedName, resource); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	if err := t.mgr.GetClient().Get(ctx, request.NamespacedName, resource); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}
	original := resource.DeepCopy()

	finalizer := kubernetes.NewFinalizer(t.mgr.GetClient(), finalizerName)

	result, err := func() (reconcile.Result, error) {
		creds, err := t.GetCredentials(ctx, resource, request.NamespacedName.Name)
		if err != nil {
			return reconcile.Result{}, err
		}

		// @step: create a cloud client for us
		client, err := aws.NewClient(creds, resource)
		if err != nil {
			return reconcile.Result{}, err
		}

		// @step: check if the cluster exists and if so we wait or the operation or the exit
		found, err := client.Exists()
		if err != nil {
			return reconcile.Result{}, err
		}

		// @step: lets update the status of the resource to deleting
		if resource.Status.Status != corev1.DeletingStatus {
			resource.Status.Status = corev1.DeletingStatus

			return reconcile.Result{Requeue: true}, nil
		}

		if found {
			// Find any nodegroups first:
			ngs, err := client.ListNodeGroups()
			if err != nil {
				logger.WithError(err).Error("attempting to list nodegroups first")

				return reconcile.Result{}, err
			}
			for _, ng := range ngs {
				// Get the Kore nodegroup object
				eksng := &eksv1alpha1.EKSNodeGroup{}
				key := types.NamespacedName{
					Namespace: request.Namespace,
					Name:      ng,
				}
				// Get the ng object
				if err := t.mgr.GetClient().Get(ctx, key, eksng); err != nil {
					if kerrors.IsNotFound(err) {
						// Unmanaged nodepool - time to error
						logger.WithError(err).Errorf("nodegroup %s not defined in kore cowardly not deleteing", ng)
					}
					logger.WithError(err).Errorf("error getting nodegroup %s", ng)

					return reconcile.Result{}, err
				}
				// We need to delete the nodepool first so lets do that:
				if err := t.mgr.GetClient().Delete(ctx, eksng.DeepCopy()); err != nil {
					logger.Debugf("initiated deletion of resource %s, requeing cluster deletion until gone", ng)

					return reconcile.Result{Requeue: true}, nil
				}
			}
			// We can now delete the cluster
			if _, err = client.Delete(); err != nil {
				return reconcile.Result{}, err
			}

			return reconcile.Result{RequeueAfter: 1 * time.Minute}, nil
		}

		return reconcile.Result{}, nil
	}()
	if err != nil {
		logger.WithError(err).Error("attempting to delete the cluster")

		resource.Status.Status = corev1.FailureStatus
		resource.Status.Conditions.SetCondition(corev1.Component{
			Detail:  err.Error(),
			Message: "Failed to delete the EKS cluster",
			Status:  corev1.FailureStatus,
		})
	}
	// @step: update the status of the resoruce
	if err := t.mgr.GetClient().Status().Patch(ctx, resource, client.MergeFrom(original)); err != nil {
		logger.WithError(err).Error("trying to update the resource status")

		return reconcile.Result{}, err
	}

	if err != nil {
		return result, err
	}

	if err := finalizer.Remove(resource); err != nil {
		logger.WithError(err).Error("removing the finalizer")

		return reconcile.Result{}, err
	}
	logger.Debug("successfully deleted the cluster")

	return result, nil
}
