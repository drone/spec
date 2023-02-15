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

// job:
//   secrets:
//     DATABASE_PASSWORD:  # Store the path to the secret in this CI/CD variable
// 		vault:  # Translates to secret: `ops/data/production/db`, field: `password`
// 		engine:
// 			name: kv-v2
// 			path: ops
// 		path: production/db
// 		field: password

// job:
//   secrets:
//     DATABASE_PASSWORD:  # Store the path to the secret in this CI/CD variable
//       vault: production/db/password  # Translates to secret: `kv-v2/data/production/db`, field: `password`

// job:
//   secrets:
//     DATABASE_PASSWORD:  # Store the path to the secret in this CI/CD variable
//       vault: production/db/password@ops  # Translates to secret: `ops/data/production/db`, field: `password`

// job:
//   secrets:
//     DATABASE_PASSWORD:
//       vault: production/db/password@ops
//       file: false

// job:
//   id_tokens:
//     AWS_TOKEN:
//       aud: https://aws.example.com
//     VAULT_TOKEN:
//       aud: https://vault.example.com
//   secrets:
//     DB_PASSWORD:
//       vault: gitlab/production/db
//       token: $VAULT_TOKEN

// job:
//   secrets:
//     DATABASE_PASSWORD:
//       vault: production/db/password@ops
//       file: false

// job:
//   id_tokens:
//     AWS_TOKEN:
//       aud: https://aws.example.com
//     VAULT_TOKEN:
//       aud: https://vault.example.com
//   secrets:
//     DB_PASSWORD:
//       vault: gitlab/production/db
//       token: $VAULT_TOKEN
