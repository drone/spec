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

type (
	// When defines execution conditions.
	When struct {
		Eval string
		Expr []map[string]*Expr
	}

	// Expr defines an execution expression.
	Expr struct {
		Eq  string `json:"eq"`
		In  string `json:"in"`
		Not *Expr  `json:"not"`
	}
)

// UnmarshalJSON implements the unmarshal interface.
func (v *When) UnmarshalJSON(data []byte) error {
	// parse the expression string.
	if err := json.Unmarshal(data, &v.Eval); err == nil {
		return nil
	}

	// parse the declarative when clause array.
	if err := json.Unmarshal(data, &v.Expr); err == nil {
		return nil
	}

	// parse the declarative when clause.
	vv := map[string]*Expr{}
	if err := json.Unmarshal(data, &vv); err == nil {
		v.Expr = append(v.Expr, vv)
		return nil
	}

	return errors.New("failed to unmarshal when clause")
}
