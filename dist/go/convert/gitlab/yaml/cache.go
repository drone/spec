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

type Cache struct {
	Paths     Stringorslice `json:"paths,omitempty"`
	Key       string        `json:"key,omitempty"` // TODO complex item (files []string, prefix)
	Untracked bool          `json:"untracked,omitempty"`
	Unprotect bool          `json:"unprotect,omitempty"`
	When      string        `json:"when,omitempty"`   // on_success, on_failure, always
	Policy    string        `json:"policy,omitempty"` // pull, push, pull-push
}

type CacheKey struct {
	Value  string        `json:"-,omitempty"`
	Files  Stringorslice `json:"files,omitempty"`
	Prefix string        `json:"prefix,omitempty"`
}

// UnmarshalJSON implements the unmarshal interface.
func (v *CacheKey) UnmarshalJSON(data []byte) error {
	var out1 string
	var out2 = struct {
		Files  Stringorslice `json:"files"`
		Prefix string        `json:"prefix"`
	}{}

	if err := json.Unmarshal(data, &out1); err == nil {
		v.Value = out1
		return nil
	}

	if err := json.Unmarshal(data, &out2); err == nil {
		v.Files = out2.Files
		v.Prefix = out2.Prefix
		return nil
	}

	return errors.New("failed to unmarshal cache key")
}
