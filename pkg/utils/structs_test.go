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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestConvertToMapOK(t *testing.T) {
	s := &testStruct{
		Name: "test",
		Age:  10,
	}
	expected := map[string]interface{}{
		"name": "test",
		"age":  float64(10),
	}

	values, err := ConvertToMap(s)
	require.NoError(t, err)
	require.NotNil(t, values)
	assert.Equal(t, expected, values)
}
