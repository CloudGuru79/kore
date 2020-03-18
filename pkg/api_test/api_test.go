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

package api_test

import (
	"os"
	"testing"

	"github.com/appvia/kore/pkg/api_test_client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Bunch of constants used by the tests:
var testTeam1 string = "api-test-team-1"
var testTeam2 string = "api-test-team-2"
var testUserAdmin string = "user-admin"
var testUserTeam1 string = "user-team-1"
var testUserTeam2 string = "user-team-2"
var testUserMultiTeam string = "user-team-1-2"
var emailSuffix = "@appvia.io"

// NB: This value MUST match the local-jwt-public-key used to start the API else we can't impersonate users in the test and everything will fail:
// Priv key must match the above pub key. To generate a new pair, uncomment generateJWTKeys() below and see the console output.
var pubKey = "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAIG6XiNhkwDETU2zk0tGlI0DKlbEJcN4jxwJBqhd3neReLDnqg9SBgKepdy9Nxw5LAd1gNoBkLvdFJg9SbHlM0sCAwEAAQ=="
var privKey = "MIIBOgIBAAJBAIG6XiNhkwDETU2zk0tGlI0DKlbEJcN4jxwJBqhd3neReLDnqg9SBgKepdy9Nxw5LAd1gNoBkLvdFJg9SbHlM0sCAwEAAQJAZ/KOdek8YlPo4Ubv0lRmuar8pPOckskqWrt8wzIcDU+rcenXZ4hnlJBUXuvKiixd29yFX0xoYPPcq/ootJQcCQIhAO5hhYf3dwwFRr2Eif8OJRC0Q5Z9qPcTfCOGCAmIML9NAiEAi1D9nUW+mUPmQyZNx88FZhTsyrzStduVlqVV+Rs/IPcCIQCkUEhwvl0qxgBK5i8Qxjk6WGc2NovfM2kgO2US3PNtCQIgd0oeHvB9R1bwb0b5CsGk6ce5Cc+szLL830Uq3GYMI/kCIAKAGmltwHk/3QvtXwFu6ecDn1cL3WuZatM/U9gcTFFI"

func TestAPI(t *testing.T) {
	//generateJWTKeys()
	if os.Getenv("RUN_INTEGRATION") != "true" {
		t.Skip("To run integration tests, run (export RUN_INTEGRATION=true; go test)")
	} else {
		RegisterFailHandler(Fail)
		BeforeSuite(func() {
			setupJWT()
			setupTeamsAndUsers()
		})
		RunSpecs(t, "API")
	}
}

func getApi() *api_test_client.AppviaKore {
	uri := os.Getenv("KORE_API_HOSTPORT")
	if uri == "" {
		uri = "localhost:10080"
	}
	transport := httptransport.New(uri, "", nil)
	return api_test_client.New(transport, strfmt.Default)
}
