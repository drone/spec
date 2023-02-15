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
	"strings"
)

type Secret struct {
	Vault *Vault `json:"vault,omitempty"`
	File  string `json:"file,omitempty"`
}

type Vault struct {
	Engine *VaultEngine `json:"engine,omitempty"`
	Path   string       `json:"path,omitempty"`
	Field  string       `json:"field,omitempty"`
}

type VaultEngine struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

// UnmarshalJSON implements the unmarshal interface.
func (v *Vault) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Engine *VaultEngine `json:"engine,omitempty"`
		Path   string       `json:"path,omitempty"`
		Field  string       `json:"field,omitempty"`
	}{}

	if err := json.Unmarshal(data, &out1); err == nil {
		parts := strings.SplitN(out1, "@", 2)
		if len(parts) == 2 {
			v.Path = parts[0]
			v.Field = parts[1]
		} else {
			v.Path = out1
		}
		return nil
	}

	if err := json.Unmarshal(data, &out2); err == nil {
		v.Path = out2.Path
		v.Field = out2.Field
		v.Engine = out2.Engine
		return nil
	}

	return errors.New("failed to unmarshal vault")
}
