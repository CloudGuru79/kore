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

package kore

import (
	"context"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	"github.com/appvia/kore/pkg/store"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
)

// NamespaceClaims is the interface to the class namespace claims
type NamespaceClaims interface {
	// Delete is used to delete a namespace claim in the kore
	Delete(context.Context, string) (*clustersv1.NamespaceClaim, error)
	// Get returns the class from the kore
	Get(context.Context, string) (*clustersv1.NamespaceClaim, error)
	// List returns a list of classes
	List(context.Context) (*clustersv1.NamespaceClaimList, error)
	// Has checks if a resource exists
	Has(context.Context, string) (bool, error)
	// Update is responsible for update a namespace claim in the kore
	Update(context.Context, *clustersv1.NamespaceClaim) (*clustersv1.NamespaceClaim, error)
}

type nsImpl struct {
	*hubImpl
	// team is the team
	team string
}

// Delete is used to delete a namespace claim in the kore
func (n *nsImpl) Delete(ctx context.Context, name string) (*clustersv1.NamespaceClaim, error) {
	original, err := n.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	if err := n.Store().Client().Delete(ctx, store.DeleteOptions.From(original)); err != nil {
		log.WithError(err).Error("trying to delete the namespace claim")

		return nil, err
	}

	return original, nil
}

// Get returns the class from the kore
func (n *nsImpl) Get(ctx context.Context, name string) (*clustersv1.NamespaceClaim, error) {
	ns := &clustersv1.NamespaceClaim{}

	return ns, n.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(n.team),
		store.GetOptions.InTo(ns),
		store.GetOptions.WithName(name),
	)
}

// List returns a list of classes
func (n *nsImpl) List(ctx context.Context) (*clustersv1.NamespaceClaimList, error) {
	list := &clustersv1.NamespaceClaimList{}

	return list, n.Store().Client().List(ctx,
		store.ListOptions.InNamespace(n.team),
		store.ListOptions.InTo(list),
	)
}

// Has checks if a resource exists
func (n *nsImpl) Has(ctx context.Context, name string) (bool, error) {
	if _, err := n.Get(ctx, name); err != nil {
		if kerrors.IsNotFound(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Update is responsible for update a namespace claim in the kore
func (n *nsImpl) Update(ctx context.Context, namespace *clustersv1.NamespaceClaim) (*clustersv1.NamespaceClaim, error) {
	// @step: ensure it's for cluster we own
	if namespace.Spec.Cluster.Namespace != n.team {
		return nil, ErrNotAllowed{message: "namespace must exist in a cluster you own"}
	}
	namespace.Namespace = n.team

	return namespace, n.Store().Client().Update(ctx,
		store.UpdateOptions.To(namespace),
		store.UpdateOptions.WithCreate(true),
	)
}
