// +build integration

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

package apiserver_test

import (
	"github.com/appvia/kore/pkg/apiclient"
	"github.com/appvia/kore/pkg/apiclient/models"
	"github.com/appvia/kore/pkg/apiclient/operations"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/audit", func() {
	var api *apiclient.AppviaKore

	BeforeEach(func() {
		api = getApi()
	})

	When("GET is performed", func() {
		It("should return 401 if not authenticated", func() {
			_, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuthAnon())
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&operations.ListAuditEventsUnauthorized{}))
		})

		It("should return 403 if not authorised", func() {
			_, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuth(TestUserTeam1))
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(&operations.ListAuditEventsForbidden{}))
		})

		It("should return a list of all audit events from the last 60 minutes by default", func() {
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams(), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			Expect(*&resp.Payload.Items).To(BeAssignableToTypeOf([]*models.V1AuditEvent{}))
		})

		It("should return a list of all audit events from a period specified by the since parameter", func() {
			since := "1s"
			resp, err := api.Operations.ListAuditEvents(operations.NewListAuditEventsParams().WithSince(&since), getAuth(TestUserAdmin))
			if err != nil {
				Expect(err).ToNot(HaveOccurred())
			}
			Expect(*&resp.Payload.Items).To(Equal([]*models.V1AuditEvent{}))
		})
	})
})