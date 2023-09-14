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

package parse

import (
	schema "github.com/drone/spec/dist/go"
)

// Normalize normalizes the yaml to ensure all stages and
// steps have unique identifiers.
func Normalize(in *schema.Config) error {
	gen := newGenerator()

	switch v := in.Spec.(type) {
	case *schema.Pipeline:
		for _, vv := range v.Stages {
			normalizeStage(vv, gen)
		}
	case *schema.PluginStep:
	case *schema.PluginStage:
	case *schema.TemplateStage:
	case *schema.TemplateStep:
	}

	return nil
}

// normalize the stage
func normalizeStage(v *schema.Stage, gen *generator) {
	v.Id = gen.generate(v.Id, v.Name, v.Type)

	switch v := v.Spec.(type) {
	case *schema.StageCI:
		for _, vv := range v.Steps {
			normalizeStep(vv, gen)
		}
	case *schema.StageGroup:
		for _, vv := range v.Stages {
			normalizeStage(vv, gen)
		}
	case *schema.StageParallel:
		for _, vv := range v.Stages {
			normalizeStage(vv, gen)
		}
	}
}

// normalize the step
func normalizeStep(v *schema.Step, gen *generator) {
	v.Id = gen.generate(v.Id, v.Name, v.Type)

	switch v := v.Spec.(type) {
	case *schema.StepGroup:
		for _, vv := range v.Steps {
			normalizeStep(vv, gen)
		}
	case *schema.StepParallel:
		for _, vv := range v.Steps {
			normalizeStep(vv, gen)
		}
	}
}
