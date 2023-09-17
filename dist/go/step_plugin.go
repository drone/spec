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

type StepPlugin struct {
	Image      string                 `json:"image,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Uses       string                 `json:"uses,omitempty"`
	Connector  string                 `json:"connector,omitempty"`
	Pull       string                 `json:"pull,omitempty"`
	Envs       map[string]string      `json:"envs,omitempty"`
	Reports    []*Report              `json:"reports,omitempty"`
	Privileged bool                   `json:"privileged,omitempty"`
	User       string                 `json:"user,omitempty"`
	Group      string                 `json:"group,omitempty"`
	Network    string                 `json:"network,omitempty"`
	With       map[string]interface{} `json:"with,omitempty"`
	Inputs     map[string]interface{} `json:"inputs,omitempty"`
	Outputs    []string               `json:"outputs,omitempty"`
	Resources  *Resources             `json:"resources,omitempty"`
	Mount      []*Mount               `json:"mount,omitempty"`
}
