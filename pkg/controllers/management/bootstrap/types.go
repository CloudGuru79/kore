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

package bootstrap

// GrafanaOptions are the opts for grafana
type GrafanaOptions struct {
	// Password is the default pass for grafana
	Password string
	// ClientID is the password
	ClientID string
	// ClientSecret is the openid secret
	ClientSecret string
	// UserInfoURL is the useinfo url
	UserInfoURL string
	// AuthURL is the auth url
	AuthURL string
	// TokenURL is the sso token url
	TokenURL string
	// Hostname is the hostname for the instance
	Hostname string
	// Database
	Database DatabaseOptions
}

// DatabaseOptions are the database options
type DatabaseOptions struct {
	// Name is the database name
	Name string `json:"name,omitempty"`
	// Password is the database password
	Password string `json:"password,omitempty"`
}

// Credentials are creds for the cloud provider
type Credentials struct {
	// AWS are as credentials
	AWS AWSCredentials `json:"aws,omitempty"`
	// GKe are gke credentials
	GKE GKECredentials `json:"gke,omitempty"`
}

// AWSCredentials are the aws crdentials
type AWSCredentials struct {
	// AccountID is the aws account id
	AccountID string `json:"account_id,omitempty"`
	// AccessKey is the credentials id
	AccessKey string `json:"access_key,omitempty"`
	// Region is the AWS region
	Region string `json:"region,omitempty"`
	// SecretKey is the credential key
	SecretKey string `json:"secret_key,omitempty"`
}

// GKECredentials are the creds for gcp
type GKECredentials struct {
	// Account is the json service account
	Account string `json:"account,omitempty"`
}

// KialiOptions are options for the kiali service
type KialiOptions struct {
	// Password is the password for the service
	Password string `json:"password,omitempty"`
	// Username is the username
	Username string `json:"username,omitempty"`
}

// NamespaceOptions is a name to create
type NamespaceOptions struct {
	// Name is the name of the namespace
	Name string `json:"name,omitempty"`
}

// Parameters provides the context for the job parameters
type Parameters struct {
	// BootImage is the image we are using to bootstrap
	KoreImage string `json:"kore_image,omitempty"`
	// Credentials are creds for the providers
	Credentials Credentials `json:"credentials,omitempty"`
	// Domain is the cluster domain
	Domain string `json:"domain,omitempty"`
	// Grafana are the options for grafana
	Grafana GrafanaOptions `json:"grafana,omitempty"`
	// Kiali are options for the kiali service
	Kiali KialiOptions `json:"kiali,omitempty"`
	// Namespaces is a collection of namespaces to create
	Namespaces []NamespaceOptions `json:"namespaces,omitempty"`
	// Provider is the cloud provider
	Provider string `json:"provider,omitempty"`
	// StorageClass is the class to use when creating PVC's
	StorageClass string `json:"storage_class,omitempty"`
}
