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

package script

import (
	"regexp"
	"strings"

	schema "github.com/drone/spec/dist/go"
)

var pattern = regexp.MustCompile(`\${{(.*)}}`)

// Expand expands the script inside the text snippet.
func Expand(code string, inputs map[string]interface{}) string {
	if !strings.Contains(code, "${{") {
		return code
	}
	return pattern.ReplaceAllStringFunc(code, func(s string) string {
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "${{")
		s = strings.TrimSuffix(s, "}}")
		out, _ := EvalStr(s, inputs)
		return out
	})
}

// ExpandStep expands scripts in the step.
func ExpandStep(step *schema.Step, inputs map[string]interface{}) {
	step.Id = Expand(step.Name, inputs)
	step.Name = Expand(step.Name, inputs)

	switch spec := step.Spec.(type) {
	case *schema.StepAction:
	case *schema.StepBackground:
		spec.Run = Expand(spec.Run, inputs)
		spec.Image = Expand(spec.Image, inputs)
		spec.Entrypoint = Expand(spec.Entrypoint, inputs)
		for i, s := range spec.Args {
			spec.Args[i] = Expand(s, inputs)
		}
	case *schema.StepBitrise:
	case *schema.StepExec:
		spec.Run = Expand(spec.Run, inputs)
		spec.Image = Expand(spec.Image, inputs)
		spec.Connector = Expand(spec.Connector, inputs)
		spec.Entrypoint = Expand(spec.Entrypoint, inputs)
		for i, s := range spec.Args {
			spec.Args[i] = Expand(s, inputs)
		}
		if spec.Reports != nil {
			for _, report := range spec.Reports {
				for i, s := range report.Path {
					report.Path[i] = Expand(s, inputs)
				}
			}
		}

	case *schema.StepGroup:
	case *schema.StepParallel:
	case *schema.StepRun:
	case *schema.StepPlugin:
	case *schema.StepTemplate:
	case *schema.StepTest:
	}
}
