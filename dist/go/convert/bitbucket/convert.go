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

package bitbucket

import (
	"fmt"
	"strings"
	"time"

	harness "github.com/drone/spec/dist/go"
	bitbucket "github.com/drone/spec/dist/go/convert/bitbucket/yaml"
)

func convertDefault(config *bitbucket.Config) *harness.Default {
	// if the global pipeline configuration sections
	// are empty or nil, return nil
	if config.Clone == nil &&
		config.Image == nil &&
		config.Options == nil {
		return nil
	}
	if config.Clone == nil {
		// Depth      *Depth
		// Enabled    *bool
		// LFS        bool
		// SkipVerify bool
	}
	if config.Image == nil {
		// Username
		// Password
	}
	if config.Options == nil {
		// Docker (bool)
		// MaxTime (int)
		// Size (1x, 2x, 4x, 8x)
	}
	return nil
}

func convertPipeline() {
}

func convertStage(s *state) *harness.Stage {

	// create the harness stage spec
	spec := &harness.StageCI{
		// TODO Clone
		// TODO Repository
		// TODO Delegate
		// TODO Platform
		// TODO Runtime
		// TODO Envs
	}

	// find the step with the largest size and use that
	// size. else fallback to the global size.
	if size := extractSize(s.config.Options, s.stage); size != bitbucket.SizeNone {
		spec.Runtime = &harness.Runtime{
			Type: "cloud",
			Spec: &harness.RuntimeCloud{
				Size: convertSize(size),
			},
		}
	}

	// find the unique selectors and append
	// to the stage.
	if runson := extractRunsOn(s.stage); len(runson) != 0 {
		spec.Delegate = new(harness.Delegate)
		spec.Delegate.Selectors = runson
	}

	// create the harness stage.
	stage := &harness.Stage{
		Name: "build",
		Type: "ci",
		Spec: spec,
		// TODO When
		// TODO On
	}

	for _, steps := range s.stage.Steps {
		if steps.Parallel != nil {
			// TODO parallel steps
			// TODO fast fail
			s.steps = steps // push the parallel step to the state
			step := convertParallel(s)
			spec.Steps = append(spec.Steps, step)
		}
		if steps.Step != nil {
			s.step = steps.Step // push the step to the state
			step := convertStep(s)
			spec.Steps = append(spec.Steps, step)
		}
	}

	// if the stage has a single step, and that step is a
	// group step, we can eliminate the un-necessary group
	// and add the steps directly to the stage.
	if len(spec.Steps) == 1 {
		if group, ok := spec.Steps[0].Spec.(*harness.StepGroup); ok {
			spec.Steps = group.Steps
		}
	}

	return stage
}

// helper function converts a bitbucket parallel step
// group to a Harness parallel step group.
func convertParallel(s *state) *harness.Step {

	// create the step group spec
	spec := new(harness.StepParallel)

	for _, src := range s.steps.Parallel.Steps {
		if src.Step != nil {
			s.step = src.Step
			step := convertStep(s)
			spec.Steps = append(spec.Steps, step)
		}
	}

	// else create the step group wrapper.
	return &harness.Step{
		Type: "parallel",
		Spec: spec,
		Name: s.generateName("parallel", "parallel"), // TODO can we avoid a name here?
	}
}

// helper function converts a bitbucket step
// to a harness run step or plugin step.
func convertStep(s *state) *harness.Step {
	// create the step group spec
	spec := new(harness.StepGroup)

	// loop through each script item
	for _, script := range s.step.Script {
		s.script = script

		// if a pipe step
		if script.Pipe != nil {
			step := convertPipeStep(s)
			spec.Steps = append(spec.Steps, step)
		}

		// else if a script step
		if script.Pipe == nil {
			step := convertScriptStep(s)
			spec.Steps = append(spec.Steps, step)
		}
	}

	// and loop through each after script item
	for _, script := range s.step.ScriptAfter {
		s.script = script

		// if a pipe step
		if script.Pipe != nil {
			step := convertPipeStep(s)
			spec.Steps = append(spec.Steps, step)
		}

		// else if a script step
		if script.Pipe == nil {
			step := convertScriptStep(s)
			spec.Steps = append(spec.Steps, step)
		}
	}

	// if there is only a single step, no need to
	// create a step group.
	if len(spec.Steps) == 1 {
		return spec.Steps[0]
	}

	// else create the step group wrapper.
	return &harness.Step{
		Type: "group",
		Spec: spec,
		Name: s.generateName(s.step.Name, "group"),
	}
}

// helper function converts a script step to a
// harness run step.
func convertScriptStep(s *state) *harness.Step {

	// create the run spec
	spec := &harness.StepExec{
		Run: s.script.Text,

		// TODO configure an optional connector
		// TODO configure pull policy
		// TODO configure envs
		// TODO configure volumes
		// TODO configure resources
	}

	// use the global image, if set
	if image := s.config.Image; image != nil {
		spec.Image = strings.TrimPrefix(image.Name, "docker://")
		spec.User = fmt.Sprint(image.RunAsUser)
	}

	// use the step image, if set (overrides previous)
	if image := s.step.Image; image != nil {
		spec.Image = strings.TrimPrefix(image.Name, "docker://")
		spec.User = fmt.Sprint(image.RunAsUser)
	}

	// create the run step wrapper
	step := &harness.Step{
		Type: "run",
		Spec: spec,
		Name: s.generateName(s.step.Name, "run"),
	}

	// set the timeout
	if v := int64(s.step.MaxTime); v != 0 {
		step.Timeout = minuteToDurationString(v)
	}

	return step
}

// helper function converts a pipe step to a
// harness plugin step.
func convertPipeStep(s *state) *harness.Step {
	pipe := s.script.Pipe

	// create the plugin spec
	spec := &harness.StepPlugin{
		Image: strings.TrimPrefix(pipe.Image, "docker://"),

		// TODO configure an optional connector
		// TODO configure envs
		// TODO configure volumes
	}

	// append the plugin spec variables
	spec.With = map[string]interface{}{}
	for key, val := range pipe.Variables {
		spec.With[key] = val
	}

	// create the plugin step wrapper
	step := &harness.Step{
		Type: "plugin",
		Spec: spec,
		Name: s.generateName(s.step.Name, "plugin"),
	}

	// set the timeout
	if v := int64(s.step.MaxTime); v != 0 {
		step.Timeout = minuteToDurationString(v)
	}

	return step
}

// helper function converts an integer of minutes
// to a time duration string.
func minuteToDurationString(v int64) string {
	dur := time.Duration(v) * time.Minute
	return fmt.Sprint(dur)
}

func extractSize(opts *bitbucket.Options, stage *bitbucket.Stage) bitbucket.Size {
	var size bitbucket.Size

	// start with the global size, if set
	if opts != nil {
		size = opts.Size
	}

	// loop through the steps and if a step
	// defines a size greater than the global
	// size, us the step size instead.
	for _, steps := range stage.Steps {
		if steps.Step != nil {
			if steps.Step.Size > size {
				size = steps.Step.Size
			}
		}
		// TODO look at parallel steps too
	}
	return size
}

func extractRunsOn(stage *bitbucket.Stage) []string {
	set := map[string]struct{}{}

	// loop through the steps and if a step
	// defines a size greater than the global
	// size, us the step size instead.
	for _, steps := range stage.Steps {
		if steps.Step != nil {
			for _, s := range steps.Step.RunsOn {
				set[s] = struct{}{}
			}
		}
		// TODO look at parallel steps too
	}

	// convert the map to a slice.
	var unique []string
	for k, _ := range set {
		unique = append(unique, k)
	}
	return unique
}

func convertSize(size bitbucket.Size) string {
	switch size {
	case bitbucket.Size2x: // 8GB
		return "large"
	case bitbucket.Size4x: // 16GB
		return "xlarge"
	case bitbucket.Size8x: // 32GB
		return "xxlarge"
	case bitbucket.Size1x: // 4GB
		return "standard"
	default:
		return ""
	}
}
