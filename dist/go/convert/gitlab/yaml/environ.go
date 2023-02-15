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

type Environment struct {
	Name           string      `json:"name,omitempty"`
	Url            string      `json:"url,omitempty"`
	OnStop         string      `json:"on_stop,omitempty"`
	Action         string      `json:"action,omitempty"` // start, prepare, stop, verify, access
	AutoStopIn     string      `json:"auto_stop_in,omitempty"`
	DeploymentTier string      `json:"deployment_tier"` // production, staging, testing, development, other
	Kubernetes     *Kubernetes `json:"kubernetes,omitempty"`
}

type Kubernetes struct {
	Namespace string `json:"namespace,omitempty"`
}

// deploy to production:
//   stage: deploy
//   script: git push production HEAD:main
//   environment: production

// deploy to production:
//   stage: deploy
//   script: git push production HEAD:main
//   environment:
//     name: production

// deploy to production:
//   stage: deploy
//   script: git push production HEAD:main
//   environment:
//     name: production
//     url: https://prod.example.com

// stop_review_app:
//   stage: deploy
//   variables:
//     GIT_STRATEGY: none
//   script: make delete-app
//   when: manual
//   environment:
//     name: review/$CI_COMMIT_REF_SLUG
//     action: stop

// review_app:
//   script: deploy-review-app
//   environment:
//     name: review/$CI_COMMIT_REF_SLUG
//     auto_stop_in: 1 day

// deploy:
//   stage: deploy
//   script: make deploy-app
//   environment:
//     name: production
//     kubernetes:
//       namespace: production

// deploy:
//   script: echo
//   environment:
//     name: customer-portal
//     deployment_tier: production
