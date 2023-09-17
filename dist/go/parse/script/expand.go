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

// Expand expands the script inside the text snippet.
func Expand(code string, inputs map[string]interface{}) string {
	pattern := regexp.MustCompile(`\${{(.*)}}`)
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
	switch step.Spec.(type) {
	case *schema.StepAction:
	case *schema.StepBackground:
	case *schema.StepBitrise:
	case *schema.StepExec:
	case *schema.StepGroup:
	case *schema.StepParallel:
	case *schema.StepRun:
	case *schema.StepPlugin:
	case *schema.StepTemplate:
	case *schema.StepTest:
	}
}
