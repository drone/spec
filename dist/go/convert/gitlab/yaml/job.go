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

// Job defines a gitlab job.
// https://docs.gitlab.com/ee/ci/yaml/#job-keywords
type Job struct {
	After         Stringorslice        `json:"after_script,omitempty"`
	Artifacts     *Artifacts           `json:"artifacts,omitempty"`
	AllowFailure  *AllowFailure        `json:"allow_failure,omitempty"`
	Before        Stringorslice        `json:"before_script,omitempty"`
	Cache         *Cache               `json:"cache,omitempty"`
	Environment   *Environment         `json:"environment,omitempty"`
	Extends       Stringorslice        `json:"extends,omitempty"`
	Image         *Image               `json:"image,omitempty"`
	Inherit       interface{}          // TODO
	Interruptible bool                 `json:"interruptible,omitempty"`
	Needs         interface{}          // TODO
	Only          interface{}          // TODO
	Pages         interface{}          // TODO
	Parallel      interface{}          // TODO
	Release       interface{}          // TODO
	ResourceGroup interface{}          // TODO
	Retry         *Retry               `json:"retry,omitempty"`
	Rules         interface{}          // TODO
	Script        Stringorslice        `json:"script,omitempty"`
	Secrets       map[string]*Secret   `json:"secrets,omitempty"`
	Services      []*Image             `json:"services,omitempty"`
	Stage         string               `json:"stage,omitempty"`
	Tags          Stringorslice        `json:"tags,omitempty"`
	Timeout       string               `json:"timeout,omitempty"`
	Trigger       interface{}          // TODO
	Variables     map[string]*Variable `json:"variables,omitempty"`
	When          string               `json:"when,omitempty"` // on_success, manual, always, on_failure, delayed, never
}
