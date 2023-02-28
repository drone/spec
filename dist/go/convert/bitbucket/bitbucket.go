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

// Package bitbucket converts Bitbucket pipelines to Harness pipelines.
package bitbucket

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	harness "github.com/drone/spec/dist/go"
	bitbucket "github.com/drone/spec/dist/go/convert/bitbucket/yaml"

	"github.com/ghodss/yaml"
)

// From converts the legacy drone yaml format to the
// unified yaml format.
func From(r io.Reader) ([]byte, error) {
	src, err := bitbucket.Parse(r)
	if err != nil {
		return nil, err
	}

	// create the harness pipeline
	dstPipeline := new(harness.Pipeline)

	dstStageSpec := &harness.StageCI{
		// Delegate: convertNode(from.Node),
		// Envs:     copyenv(from.Environment),
		// Platform: convertPlatform(from.Platform),
		// Runtime:  convertRuntime(from),
		// Steps:    convertSteps(from),
	}

	// create the harness stage.
	dstStage := &harness.Stage{
		Name: "build",
		Type: "ci",
		// When: convertCond(from.Trigger),
		Spec: dstStageSpec,
	}

	// append the stage to the pipeline
	dstPipeline.Stages = append(dstPipeline.Stages, dstStage)

	// unique map of names
	names := map[string]struct{}{}

	// iterage through named stages
	for _, srcStep := range src.Pipelines.Default {

		if srcStep.Step == nil {
			continue // HACK skip non-script steps for now
		}

		// iterate through jobs and find jobs assigned to
		// the stage. skip other stages.
		for _, scrScript := range srcStep.Step.Script {

			if scrScript.Pipe != nil {
				dstStepSpec := &harness.StepPlugin{
					Image: strings.TrimPrefix("docker://", scrScript.Pipe.Image),
				}

				// append variables
				dstStepSpec.With = map[string]interface{}{}
				for key, val := range scrScript.Pipe.Variables {
					dstStepSpec.With[key] = val
				}

				// create the step and cofigure its spec
				dstStep := &harness.Step{
					Type: "plugin",
					Spec: dstStepSpec,
				}

				// create a temp name
				name := srcStep.Step.Name
				if name == "" {
					if scrScript.Pipe.Name == "" {
						name = "pipe"
					}
				}

				// ensure the name is unique
				for i := 0; ; i++ {
					tmpname := name
					if i > 0 {
						tmpname = name + fmt.Sprint(i)
					}
					if _, ok := names[tmpname]; !ok {
						names[tmpname] = struct{}{}
						dstStep.Name = tmpname
						break
					}
				}

				// append the step to the stage.
				dstStageSpec.Steps = append(dstStageSpec.Steps, dstStep)
				continue
			}

			dstStepSpec := &harness.StepExec{
				Run: scrScript.Text,
			}

			// create the step and cofigure its spec
			dstStep := &harness.Step{
				Type: "script",
				Spec: dstStepSpec,
			}

			// create a temp name
			name := srcStep.Step.Name
			if name == "" {
				name = "run"
			}

			// ensure the name is unique
			for i := 0; ; i++ {
				tmpname := name
				if i > 0 {
					tmpname = name + fmt.Sprint(i)
				}
				if _, ok := names[tmpname]; !ok {
					names[tmpname] = struct{}{}
					dstStep.Name = tmpname
					break
				}
			}

			// append the step to the stage.
			dstStageSpec.Steps = append(dstStageSpec.Steps, dstStep)
		}
	}

	out, err := yaml.Marshal(dstPipeline)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// FromBytes converts the legacy drone yaml format to the
// unified yaml format.
func FromBytes(b []byte) ([]byte, error) {
	return From(
		bytes.NewBuffer(b),
	)
}

// FromString converts the legacy drone yaml format to the
// unified yaml format.
func FromString(s string) ([]byte, error) {
	return FromBytes(
		[]byte(s),
	)
}

// FromFile converts the legacy drone yaml format to the
// unified yaml format.
func FromFile(p string) ([]byte, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return From(f)
}
