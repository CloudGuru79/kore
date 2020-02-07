/**
 * Copyright (C) 2020 Appvia Ltd <info@appvia.io>
 *
 * This file is part of kore-apiserver.
 *
 * kore-apiserver is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * kore-apiserver is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with kore-apiserver.  If not, see <http://www.gnu.org/licenses/>.
 */

package authentication

// ContextKey is the context key
type ContextKey struct{}

// Identity provides the user
type Identity interface {
	// IsGlobalAdmin checks if the user is a global admin
	IsGlobalAdmin() bool
	// Email returns the user email
	Email() string
	// Disabled checks if the user is disabled
	Disabled() bool
	// Username is a unique username for this identity
	Username() string
	// Teams is a list of teams for the user
	Teams() []string
}