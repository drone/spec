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

type AllowFailure struct {
	Value     bool  `json:"-"`
	ExitCodes []int `json:"exit_codes"`
}

// UnmarshalJSON implements the unmarshal interface.
func (v *AllowFailure) UnmarshalJSON(data []byte) error {
	var out1 bool
	var out2 = struct {
		ExitCode int `json:"exit_codes"`
	}{}
	var out3 = struct {
		ExitCodes []int `json:"exit_codes"`
	}{}

	if err := json.Unmarshal(data, &out1); err == nil {
		v.Value = out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err == nil {
		v.Value = true
		v.ExitCodes = []int{out2.ExitCode}
		return nil
	}

	if err := json.Unmarshal(data, &out3); err == nil {
		v.Value = true
		v.ExitCodes = out3.ExitCodes
		return nil
	}

	return errors.New("failed to unmarshal allow_failure")
}
