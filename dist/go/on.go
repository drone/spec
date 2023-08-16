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

type On struct {
	Failure []*Failure `json:"failure,omitempty"`
}

// UnmarshalJSON implements the unmarshal interface.
func (v *On) UnmarshalJSON(data []byte) error {

	// parse the failure clause array.
	if err := json.Unmarshal(data, &v.Failure); err == nil {
		return nil
	}

	// parse the simple failure clause.
	vv := new(Failure)
	if err := json.Unmarshal(data, &vv); err == nil {
		v.Failure = append(v.Failure, vv)
		return nil
	}

	return errors.New("failed to unmarshal failure clause")
}

// MarshalJSON implements the marshal interface.
func (v *On) MarshalJSON() ([]byte, error) {
	if len(v.Failure) == 1 {
		return json.Marshal(v.Failure[0])
	}
	if len(v.Failure) > 1 {
		return json.Marshal(v.Failure)
	}
	return []byte("null"), nil
}
