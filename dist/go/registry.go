// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import (
	"encoding/json"
	"errors"
)

// Registry provides container registry details.
type Registry struct {
	Credentials []*Credentials `json:"credentials,omitempty"`
}

// Credentials provides registry credentials.
type Credentials struct {
	Name  string `json:"name,omitempty"`
	Match string `json:"match,omitempty"`
}

// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Credentials) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Name  string `json:"name,omitempty"`
		Match string `json:"match,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err == nil {
		v.Name = out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err == nil {
		v.Name = out2.Name
		v.Match = out2.Match
		return nil
	}

	return errors.New("yaml: cannot unmarshal registry credentials")
}
