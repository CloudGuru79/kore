/*
 * Copyright (C) 2019 Appvia Ltd <info@appvia.io>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package allocations

import (
	"context"

	configv1 "github.com/appvia/kore/pkg/apis/config/v1"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Delete is responsible for deleting any allocations
func (a acCtrl) Delete(ctx context.Context, object *configv1.Allocation) (reconcile.Result, error) {
	logger := log.WithFields(log.Fields{
		"resource.name":      object.Name,
		"resource.namespace": object.Namespace,
		"team":               object.Namespace,
	})
	logger.Info("attempting to remove the allocation")

	return reconcile.Result{}, nil
}
