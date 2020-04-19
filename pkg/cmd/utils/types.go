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

package utils

import (
	"errors"
	"io"

	"github.com/appvia/kore/pkg/client"
	"github.com/appvia/kore/pkg/client/config"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	// DualScope indicate both of the above
	DualScope = "*"
	// GlobalScope indicate a kore resource
	GlobalScope = "global"
	// TeamScope indicates a team resource
	TeamScope = "team"
)

var (
	// ErrPayloadMissing indicates the resource has not payload to inspect
	ErrPayloadMissing = errors.New("no payload in request")
	// ErrNotRuntimeObject indicates the object is not a runtime.Object
	ErrNotRuntimeObject = errors.New("object is not a runtime.Object")
	// ErrNotMetaObject indicates the object does not implement metav1.Object
	ErrNotMetaObject = errors.New("object does not implement metav1.Object")
)

// Streams defines the io.streams
type Streams struct {
	// Stdin is the input device
	Stdin io.Reader
	// Stdout is the writer
	Stdout io.Writer
	// Stderr is the error writer
	Stderr io.Writer
}

// Factory provides a factory of methods for the cli
type Factory interface {
	// CheckError handles the cli errors for us
	CheckError(error)
	// Client returns a api client
	Client() client.RestInterface
	// Config returns the runtim configuration
	Config() *config.Config
	// Println writes a message to the io.Writer
	Println(string, ...interface{})
	// Printf writes a message to the io.Writer
	Printf(string, ...interface{})
	// Resources returns the resources contract
	Resources() Resources
	// SetStdin allows you to set the stdin for the factory
	SetStdin(io.Reader)
	// Stdin return the standard input
	Stdin() io.Reader
	// Stderr returns the io.Writer for errors
	Stderr() io.Writer
	// WaitForCreation is used to wait for the resource to be created
	WaitForCreation(client.RestInterface, bool) error
	// WaitForDeletion is used to wait for the resource to be created
	WaitForDeletion(client.RestInterface, string, bool) error
	// Writer returns the io.Writer for output
	Writer() io.Writer
	// UpdateConfig is responsible for updating the configuration
	UpdateConfig() error
}

// Resources is the contract to the resource cache
type Resources interface {
	// LookResourceNamesWithFilter returns a list of resource names against a regexp
	LookResourceNamesWithFilter(string, string, string) ([]string, error)
	// LookupResourceNames returns a list of resources of a specific kind
	LookupResourceNames(string, string) ([]string, error)
	// Lookup is used to check if a resource is supported
	Lookup(string) (*Resource, error)
	// Names returns all the names of the resource types
	Names() ([]string, error)
	// List returns all the resource available
	List() ([]Resource, error)
	// ResolveShorthand is used to resolve a shorthand option
	ResolveShorthand(string) string
}

// ResourceDocument defines a read in resource
type ResourceDocument struct {
	// Resource is the resource definition
	Resource Resource
	// Object the resource to send
	Object *unstructured.Unstructured
}

// ResourceScope defines the scope of a resource e.g. team spaced or not
type ResourceScope string

// Resource defines a resource in kore
type Resource struct {
	// Name is the name of the resource
	Name string `json:"name,omitempty"`
	// ShortName is the a short name of the resource (not used yet)
	ShortName string `json:"shortName,omitempty"`
	// GroupVersion is the api group version of the resource (not used yet)
	GroupVersion string `json:"groupVersion,omitempty"`
	// Kind is the resource kind
	Kind string `json:"kind,omitempty"`
	// Scope is the resource scope
	Scope ResourceScope `json:"scope,omitempty"`
	// SubResources is a collection of subresources for this resource
	SubResources []string `json:"subResources,omitempty"`
	// Printer is printer columns for the resource
	Printer []Column `json:"printer,omitempty"`
}

// IsTeamScoped checks if a team resource
func (r *Resource) IsTeamScoped() bool {
	return r.IsScoped(TeamScope)
}

// IsScoped checks the scope of a resource
func (r *Resource) IsScoped(scope ResourceScope) bool {
	return r.Scope == scope
}

// Column is used to define column field for printing
type Column struct {
	// Name is the name of the column
	Name string `json:"name,omitempty"`
	// Path is the jsonpath of the field
	Path string `json:"path,omitempty"`
	// Format is optional formatter for the value
	Format string `json:"format,omitempty"`
}
