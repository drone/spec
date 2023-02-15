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

type Variable struct {
	Value   string        `json:"value,omitempty"`
	Desc    string        `json:"description,omitempty"`
	Options Stringorslice `json:"options,omitempty"`
	Expand  bool          `json:"expand,omitempty"`
}

// UnmarshalJSON implements the unmarshal interface.
func (v *Variable) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Value   string        `json:"value,omitempty"`
		Desc    string        `json:"description,omitempty"`
		Options Stringorslice `json:"options,omitempty"`
		Expand  *bool         `json:"expand,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err == nil {
		v.Value = out1
		v.Expand = true
		return nil
	}

	if err := json.Unmarshal(data, &out2); err == nil {
		v.Value = out2.Value
		v.Desc = out2.Desc
		v.Options = out2.Options
		v.Expand = true
		if out2.Expand != nil {
			v.Expand = *out2.Expand
		}
		return nil
	}

	return errors.New("failed to unmarshal variable")
}
