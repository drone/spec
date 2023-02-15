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

type (
	// Pipeline defines a gitlab pipeline.
	Pipeline struct {
		Default   *Default             `json:"default,omitempty"`
		Include   []*Include           `json:"include,omitempty"`
		Jobs      map[string]*Job      `json:"jobs,omitempty"`
		Stages    []string             `json:"stages,omitempty"`
		Variables map[string]*Variable `json:"variables,omitempty"`
		Workflow  *Workflow            `json:"workflow,omitempty"`
	}

	// Default defines global pipeline defaults.
	Default struct {
		After         Stringorslice `json:"after_script,omitempty"`
		Before        Stringorslice `json:"before_script,omitempty"`
		Artifacts     *Artifacts    `json:"artifacts,omitempty"`
		Cache         *Cache        `json:"cache,omitempty"`
		Image         *Image        `json:"image,omitempty"`
		Interruptible bool          `json:"interruptible,omitempty"`
		Retry         *Retry        `json:"retry,omitempty"`
		Services      []*Image      `json:"services,omitempty"`
		Tags          Stringorslice `json:"tags,omitempty"`
		Timeout       string        `json:"duration,omitempty"`
	}
)
