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

package cluster

import (
	"context"
	"errors"
	"fmt"
	"time"

	eksv1alpha1 "github.com/appvia/kore/pkg/apis/eks/v1alpha1"

	"github.com/appvia/kore/pkg/controllers"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizerName               = "cluster.clusters.kore.appvia.io"
	labelClusterResourceVersion = "cluster.clusters.kore.appvia.io/ResourceVersion"
)

// Reconcile is the entrypoint for the reconciliation logic
func (a *Controller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()

	logger := a.logger.WithFields(log.Fields{
		"name":      request.NamespacedName.Name,
		"namespace": request.NamespacedName.Namespace,
	})
	logger.Debug("attempting to reconcile the cluster")

	// @step: retrieve the object from the api
	cluster := &clustersv1.Cluster{}
	if err := a.mgr.GetClient().Get(ctx, request.NamespacedName, cluster); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		logger.WithError(err).Error("failed to get cluster")
		return reconcile.Result{}, err
	}

	finalizer := kubernetes.NewFinalizer(a.mgr.GetClient(), finalizerName)
	if finalizer.NeedToAdd(cluster) {
		err := finalizer.Add(cluster)
		if err != nil {
			logger.WithError(err).Error("failed to set the finalizer")
		}
		return reconcile.Result{Requeue: true}, err
	}

	if finalizer.IsDeletionCandidate(cluster) {
		return a.Delete(ctx, cluster)
	}

	err := func() error {
		cluster.Status.Status = corev1.PendingStatus
		if cluster.Status.Components == nil {
			cluster.Status.Components = corev1.Components{}
		}

		components, err := createClusterComponents(cluster)
		if err != nil {
			return controllers.NewCriticalError(err)
		}

		for componentName, c := range components {
			if err := a.loadComponent(ctx, cluster, c); err != nil {
				return fmt.Errorf("failed to load %s component: %w", componentName, err)
			}
		}

		for _, c := range components {
			switch r := c.(type) {
			case *eksv1alpha1.EKSVPC:
				applyEKSVPC(r, components)
			}
		}

		for componentName, c := range components {
			if readyForReconcile(c, components) {
				if err := a.createOrUpdateComponent(ctx, cluster, c); err != nil {
					return fmt.Errorf("failed to create or update %s component: %w", componentName, err)
				}
			}

			switch r := c.(type) {
			case *clustersv1.Kubernetes:
				if r.Status.Status == corev1.SuccessStatus {
					cluster.Status.APIEndpoint = r.Status.APIEndpoint
					cluster.Status.AuthProxyEndpoint = r.Status.Endpoint
					cluster.Status.CaCertificate = r.Status.CaCertificate
				}
			}

			status, message := c.GetStatus()
			if status == "" {
				status = corev1.PendingStatus
			}
			if status.IsFailed() && message == "" {
				if err := c.GetComponents().Error(); err != nil {
					message = err.Error()
				}
			}
			component := corev1.Component{
				Name:    componentName,
				Status:  status,
				Message: message,
			}
			cluster.Status.Components.SetCondition(component)
		}

		ready := cluster.Status.Components.HasStatusForAll(corev1.SuccessStatus)
		if ready {
			cluster.Status.Status = corev1.SuccessStatus
			cluster.Status.Message = "The cluster has been created successfully"
			return nil
		} else if cluster.Status.Components.HasStatus(corev1.FailureStatus) {
			return controllers.NewCriticalError(errors.New("one or more components failed"))
		}

		return nil
	}()

	if err != nil {
		logger.WithError(err).Error("failed to reconcile the cluster")
		if controllers.IsCriticalError(err) {
			cluster.Status.Status = corev1.FailureStatus
			cluster.Status.Message = err.Error()
		}
	}

	if err := a.mgr.GetClient().Status().Update(ctx, cluster); err != nil {
		logger.WithError(err).Error("failed to update the cluster status")
		return reconcile.Result{}, err
	}

	if cluster.Status.Status == corev1.SuccessStatus || cluster.Status.Status == corev1.FailureStatus {
		return reconcile.Result{}, nil
	}

	return reconcile.Result{RequeueAfter: 5 * time.Second}, err
}

func (a *Controller) loadComponent(ctx context.Context, cluster *clustersv1.Cluster, res clustersv1.ClusterComponent) error {
	_, err := kubernetes.GetIfExists(ctx, a.mgr.GetClient(), res)
	if err != nil {
		return err
	}

	if err := res.ApplyClusterConfiguration(cluster); err != nil {
		return controllers.NewCriticalError(err)
	}

	return nil
}

func (a *Controller) createOrUpdateComponent(ctx context.Context, cluster *clustersv1.Cluster, res clustersv1.ClusterComponent) error {
	status, _ := res.GetStatus()

	if status == "" {
		setClusterResourceVersion(res, cluster.ResourceVersion)
		res.SetStatus(corev1.PendingStatus)
		if err := a.mgr.GetClient().Create(ctx, res); err != nil {
			return err
		}
	} else {
		if getClusterResourceVersion(res) != cluster.ResourceVersion {
			setClusterResourceVersion(res, cluster.ResourceVersion)
			res.SetStatus(corev1.PendingStatus)
			if err := a.mgr.GetClient().Update(ctx, res); err != nil {
				return err
			}
		}
	}

	return nil
}