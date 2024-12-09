// Code generated by scripts/generate.js; DO NOT EDIT.

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
)

// Clone defines the stage clone behavior.
type CloneStage struct {
	Depth    int64  `json:"depth,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
	Insecure bool   `json:"insecure,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	Trace    bool   `json:"trace,omitempty"`
}

type CloneStageV1 struct {
	Disabled bool `json:"disabled"`
}

func (c *CloneStageV1) MarshalJSON() ([]byte, error) {
	// If CloneStageV1 is nil, provide a default value
	if c == nil {
		return json.Marshal(&CloneStageV1{Disabled: true})
	}

	type Alias CloneStageV1
	return json.Marshal(&struct {
		Disabled bool `json:"disabled"`
		*Alias
	}{
		Disabled: c.Disabled,
		Alias:    (*Alias)(c),
	})
}