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

import (
	"encoding/json"
	"fmt"
)

type Step struct {
	Id       string                 `json:"id,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Desc     string                 `json:"desc,omitempty"`
	Type     string                 `json:"type,omitempty"`
	Timeout  string                 `json:"timeout,omitempty"`
	Strategy *Strategy              `json:"strategy,omitempty"`
	When     *When                  `json:"when,omitempty"`
	Failure  *FailureList           `json:"failure,omitempty"`
	Inputs   map[string]interface{} `json:"inputs,omitempty"`
	Spec     interface{}            `json:"spec,omitempty"`
}

type StepV1 struct {
	Name string   `json:"name,omitempty"`
	Run  *RunSpec `json:"run,omitempty"`
}

type RunSpec struct {
	Container *ContainerSpec     `json:"container,omitempty"`
	Env       map[string]string  `json:"env,omitempty"`
	Script    string             `json:"script,omitempty"` 
}

type ContainerSpec struct {
	Image     string `json:"image,omitempty"`
	Connector string `json:"connector,omitempty"`
}


// UnmarshalJSON implement the json.Unmarshaler interface.
func (v *Step) UnmarshalJSON(data []byte) error {
	type S Step
	type T struct {
		*S
		Spec json.RawMessage `json:"spec"`
	}

	obj := &T{S: (*S)(v)}
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}

	switch v.Type {
	case "action":
		v.Spec = new(StepAction)
	case "background":
		v.Spec = new(StepBackground)
	case "barrier":
		v.Spec = new(StepBarrier)
	case "bitrise":
		v.Spec = new(StepBitrise)
	case "script":
		v.Spec = new(StepExec)
	case "run":
		v.Spec = new(StepRun)
	case "test":
		v.Spec = new(StepTest)
	case "group":
		v.Spec = new(StepGroup)
	case "parallel":
		v.Spec = new(StepParallel)
	case "plugin":
		v.Spec = new(StepPlugin)
	case "template":
		v.Spec = new(StepTemplate)
	case "jenkins":
		v.Spec = new(StepJenkins)
	default:
		return fmt.Errorf("unknown type %s", v.Type)
	}

	return json.Unmarshal(obj.Spec, v.Spec)
}

func (v *StepV1) UnmarshalJSONV1(data []byte) error {
	type StepV1 struct {
		Name string          `json:"name,omitempty"`
		Run  json.RawMessage `json:"run,omitempty"`
	}

	temp := &StepV1{}
	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	v.Name = temp.Name

	if temp.Run != nil {
		var runSpec RunSpec
		runData := map[string]json.RawMessage{}

		// Unmarshal into a temporary map to process specific fields
		if err := json.Unmarshal(temp.Run, &runData); err != nil {
			return err
		}

		// Unmarshal Container field if present
		if containerData, ok := runData["container"]; ok {
			var container ContainerSpec
			if err := json.Unmarshal(containerData, &container); err != nil {
				return err
			}
			runSpec.Container = &container
		}

		// Unmarshal Env field if present
		if envData, ok := runData["env"]; ok {
			var env map[string]string
			if err := json.Unmarshal(envData, &env); err == nil {
				runSpec.Env = env
			}
		}

		// Unmarshal Script field if present
		if scriptData, ok := runData["script"]; ok {
			var script string
			if err := json.Unmarshal(scriptData, &script); err == nil {
				runSpec.Script = script
			}
		}

		v.Run = &runSpec
	}

	return nil
}
