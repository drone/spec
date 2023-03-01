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
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v3"
)

func TestConvert(t *testing.T) {
	// tests, err := filepath.Glob("testdata/*/*.yaml")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// TODO use glob once we have more test cases complete
	tests := []string{
		"testdata/steps/example1.yaml",
		"testdata/steps/example2.yaml",
		"testdata/steps/example3.yaml",
		"testdata/steps/example4.yaml",
		"testdata/steps/example5.yaml",
		"testdata/steps/example6.yaml",
		"testdata/steps/example7.yaml",
		"testdata/steps/example8.yaml",
		"testdata/steps/example9.yaml",  // TODO trigger, deploy
		"testdata/steps/example10.yaml", // TODO oidc
		"testdata/steps/example11.yaml", // TODO trigger
		"testdata/steps/example12.yaml", // TODO artifacts
		"testdata/steps/example13.yaml", // TODO changeset
		"testdata/steps/example14.yaml", // TODO changeset
		"testdata/steps/example15.yaml",
		"testdata/steps/example16.yaml",
		"testdata/steps/example17.yaml",
		"testdata/steps/example18.yaml", // TODO artifacts
		"testdata/steps/example19.yaml", // TODO artifacts
		// "testdata/steps/example20.yaml", // TODO caches
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			// convert the yaml file from bitbucket to harness
			tmp1, err := FromFile(test)
			if err != nil {
				t.Error(err)
				return
			}

			// unmarshal the converted yaml file to a map
			got := map[string]interface{}{}
			if err := yaml.Unmarshal(tmp1, &got); err != nil {
				t.Error(err)
				return
			}

			// parse the golden yaml file
			data, err := ioutil.ReadFile(test + ".golden")
			if err != nil {
				t.Error(err)
				return
			}

			// unmarshal the golden yaml file to a map
			want := map[string]interface{}{}
			if err := yaml.Unmarshal(data, &want); err != nil {
				t.Error(err)
				return
			}

			// compare the converted yaml to the golden file
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("Unexpected conversion result")
				t.Log(diff)
			}
		})
	}
}

// func TestDump(t *testing.T) {
// 	out, err := FromFile("testdata/steps/example14.yaml")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	println(string(out))
// 	t.Fail()
// }
