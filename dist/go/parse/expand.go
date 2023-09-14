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
	"github.com/drone/spec/dist/go/parse/matrix"
)

// Expand expands the matrix strategies.
func Expand(in *schema.Config) error {
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

// expand the stage
func expandStage(v *schema.Stage) {
	if v.Strategy != nil {
		vv, ok := v.Strategy.Spec.(*schema.Matrix)
		if ok {
			perms := matrix.Calc(vv.Axis)
			for _, perm := range perms {
				println(perm.String()) // TODO
			}
		}
	}

	switch v := v.Spec.(type) {
	case *schema.StageCI:
		for _, vv := range v.Steps {
			expandStep(vv)
		}
	case *schema.StageGroup:
		for _, vv := range v.Stages {
			expandStage(vv)
		}
	case *schema.StageParallel:
		for _, vv := range v.Stages {
			expandStage(vv)
		}
	}
}

// expand the step matrix
func expandStep(v *schema.Step) {
	if v.Strategy == nil {
		return
	}

	switch v := v.Spec.(type) {
	case *schema.StepGroup:
		for _, vv := range v.Steps {
			expandStep(vv)
		}
	case *schema.StepParallel:
		for _, vv := range v.Steps {
			expandStep(vv)
		}
	}
}
