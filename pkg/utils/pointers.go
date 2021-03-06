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

import "time"

// DurationPtr returns a poiintr to a duration
func DurationPtr(value time.Duration) *time.Duration {
	return &value
}

// StringPtr returns a pointer to a string
func StringPtr(value string) *string {
	return &value
}

// TruePtr return a pointer to a true boolean
func TruePtr() *bool {
	yes := true

	return &yes
}

// BoolPtr returns the a pointer to the boolean
func BoolPtr(v bool) *bool {
	return &v
}

// StringValue return the string or empty
func StringValue(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}

// BoolValue return the boolean value or false
func BoolValue(v *bool) bool {
	if v == nil {
		return false
	}

	return *v
}
