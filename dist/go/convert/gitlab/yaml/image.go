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

// Image defines a container image.
// https://docs.gitlab.com/ee/ci/yaml/#image
type Image struct {
	Name       string        `json:"name,omitempty"`
	Alias      string        `json:"alias,omitempty"`
	Entrypoint Stringorslice `json:"entrypoint,omitempty"`
	Command    Stringorslice `json:"command,omitempty"`
	PullPolicy string        `json:"pull_policy,omitempty"` // always, never, if-not-present
}

// UnmarshalJSON implements the unmarshal interface.
func (v *Image) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Name       string        `json:"name,omitempty"`
		Alias      string        `json:"alias,omitempty"`
		Entrypoint Stringorslice `json:"entrypoint,omitempty"`
		Command    Stringorslice `json:"command,omitempty"`
		PullPolicy string        `json:"pull_policy,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err == nil {
		v.Name = out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err == nil {
		v.Name = out2.Name
		v.Alias = out2.Alias
		v.Entrypoint = out2.Entrypoint
		v.Command = out2.Command
		v.PullPolicy = out2.PullPolicy
		return nil
	}

	return errors.New("failed to unmarshal image")
}
