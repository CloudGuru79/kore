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

package clusterappman

import (
	"context"
	"errors"
	"fmt"
	"time"

	kcore "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	log "github.com/sirupsen/logrus"

	appsv1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
	rc "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// clusterappmanNamespace is the namespace the clusterappmanager runs in
	clusterappmanNamespace = KoreNamespace
	// clusterappmanDeployment
	clusterappmanDeployment = "kore-clusterappman"
	// DeployerServiceName is the name used for logging when deploying clusterappman
	DeployerServiceName = "clusterappman-deployer"
)

type deployerImpl struct {
	// ControllerClient is the controller runtime client for deploying clusterappman
	ControllerClient rc.Client
	// ClusterAppManImage is the container image to use for clusterappman
	ClusterAppManImage string
}

// NewDeployer will start a deployment service for clusterappman
func NewDeployer(clusterappmanImage string, cc client.Client) (Interface, error) {
	return &deployerImpl{
		ControllerClient:   cc,
		ClusterAppManImage: clusterappmanImage,
	}, nil
}

// Run is responsible for starting the deployment services and keeping them running
func (d deployerImpl) Run(ctx context.Context) error {
	logger := log.WithFields(log.Fields{
		"service": DeployerServiceName,
	})

	for {
		_, err := Deploy(ctx, d.ControllerClient, logger, d.ClusterAppManImage)
		if err != nil {
			logger.Errorf("error deploying clusterappman - %s", err)
		}
		if utils.Sleep(ctx, 1*time.Minute) {
			logger.Print("exiting as requested")
			return nil
		}
	}
}

// Stop will stop the services
func (d deployerImpl) Stop(ctx context.Context) error {
	return nil
}

// Deploy will install clusterappman in a cluster and return initial status
func Deploy(ctx context.Context, cc client.Client, logger *log.Entry, clusterAppManImage string) (*kcore.Components, error) {
	// @step: check if the cluster manager namespace exists and create it if not
	if err := EnsureNamespace(ctx, cc, clusterappmanNamespace); err != nil {
		logger.WithError(err).Errorf("trying to create the kore cluster-manager namespace %s", clusterappmanNamespace)

		return nil, err
	}

	// @step: ensure the service account
	if _, err := kubernetes.CreateOrUpdateServiceAccount(ctx, cc, &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "clusterappman",
			Namespace: clusterappmanNamespace,
			Labels: map[string]string{
				kore.Label("owner"): "true",
			},
		},
	}); err != nil {
		logger.WithError(err).Error("trying to create the clusterappman service account")

		return nil, err
	}
	// @step setup correct permissions for deployment
	if err := CreateClusterManClusterRoleBinding(ctx, cc); err != nil {
		logger.WithError(err).Error("can not create cluster-manager clusterrole")

		return nil, err
	}

	// @step: check if the kore cluster manager deployment exists
	logger.Debugf("deploying clusterappman using image %s", clusterAppManImage)
	if err := CreateOrUpdateClusterAppManDeployment(ctx, cc, clusterAppManImage); err != nil {
		logger.WithError(err).Error("trying to create the cluster manager deployment")

		return nil, err
	}
	logger.Debug("waiting for kore cluster manager deployment status to appear")

	nctx, cancel := context.WithTimeout(ctx, 4*time.Minute)
	defer cancel()

	logger.Info("waiting for kore cluster manager to complete")

	// @step: wait for the clusterappman deployment to complete
	if err := WaitOnStatus(nctx, cc); err != nil {
		logger.WithError(err).Error("failed waiting for kore cluster manager status to complete")

		return nil, err
	}

	logger.Info("kube clusterappman running, status available")

	return GetStatus(ctx, cc)

}

// CreateOrUpdateClusterAppManDeployment will reconcile the clusterappman deployment
func CreateOrUpdateClusterAppManDeployment(ctx context.Context, cc client.Client, image string) error {
	name := clusterappmanDeployment
	if _, err := kubernetes.CreateOrUpdate(ctx, cc, &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterappmanDeployment,
			Namespace: clusterappmanNamespace,
			Labels: map[string]string{
				"name": name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"name": name,
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": name,
					},
				},
				Spec: v1.PodSpec{
					ServiceAccountName: "clusterappman",
					Containers: []v1.Container{
						{
							Name:  name,
							Image: image,
							Env: []v1.EnvVar{
								{
									Name:  "IN_CLUSTER",
									Value: "true",
								},
							},
							Command: []string{
								"/bin/kore-clusterappman",
							},
						},
					},
				},
			},
		},
	}); err != nil {
		return err
	}
	return nil
}

// HasConfigMap checks if a configmap exists
func HasConfigMap(ctx context.Context, cc client.Client, name string) (bool, error) {
	return kubernetes.CheckIfExists(ctx, cc, &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: clusterappmanNamespace,
		},
	})
}

// NamespaceExists checks if the bootstrap job there
func NamespaceExists(ctx context.Context, cc client.Client) (bool, error) {
	return kubernetes.CheckIfExists(ctx, cc, &core.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterappmanNamespace,
		},
	})
}

// EnsureNamespace creates a namespace for the clusterappmanager if required
func EnsureNamespace(ctx context.Context, cc client.Client, namespace string) error {
	return kubernetes.EnsureNamespace(ctx, cc, &core.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
			Labels: map[string]string{
				kore.Label("owned"): "true",
			},
		},
	})
}

// CreateClusterManClusterRoleBinding creates (or updates) the cluster role binding required for the clusterappman
func CreateClusterManClusterRoleBinding(ctx context.Context, cc client.Client) error {
	if _, err := kubernetes.CreateOrUpdateManagedClusterRoleBinding(ctx, cc, &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "kore:clusterappman",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "clusterappman",
				Namespace: KoreNamespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "cluster-admin",
		},
	}); err != nil {
		return fmt.Errorf("error tying to apply kore clusterappman clusterrole %q", err)
	}
	return nil
}

// WaitOnStatus will wait until the status object exists
// TODO: define a status object suitabvle for overall cluster status
func WaitOnStatus(ctx context.Context, cc client.Client) error {
	// WaitOnStatus checks the status of the job and if not successful returns the error
	for {
		select {
		case <-ctx.Done():
			return errors.New("context has been cancelled")
		default:
		}

		err := func() error {
			exists, err := StatusExists(ctx, cc)
			if err != nil {
				return err
			}
			if exists {
				return nil
			}

			return errors.New("Kore cluster manager has not reported status yet")
		}()
		if err == nil {
			return nil
		}
		time.Sleep(10 * time.Second)
	}
}

// StatusExists checks if the status exists already
func StatusExists(ctx context.Context, cc client.Client) (bool, error) {
	return HasConfigMap(ctx, cc, StatusCongigMap)
}
