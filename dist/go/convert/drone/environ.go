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

package drone

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// regular expression to extract the pull request number
// from the git ref (e.g. refs/pulls/{d}/head)
var re = regexp.MustCompile("\\d+")

// helper function returns repository and build environment
// variables for environment substitution in the yaml.
func toEnv(params *Params) map[string]string {
	env := map[string]string{
		"DRONE_REPO":                 params.Repo.Slug,
		"DRONE_REPO_OWNER":           params.Repo.Namespace,
		"DRONE_REPO_NAMESPACE":       params.Repo.Namespace,
		"DRONE_REPO_NAME":            params.Repo.Name,
		"DRONE_REPO_LINK":            params.Repo.Link,
		"DRONE_REPO_BRANCH":          params.Repo.Branch,
		"DRONE_GIT_HTTP_URL":         params.Repo.HTTPURL,
		"DRONE_GIT_SSH_URL":          params.Repo.SSHURL,
		"DRONE_REMOTE_URL":           params.Repo.HTTPURL,
		"DRONE_BRANCH":               params.Build.Target,
		"DRONE_SOURCE_BRANCH":        params.Build.Source,
		"DRONE_TARGET_BRANCH":        params.Build.Target,
		"DRONE_COMMIT":               params.Build.After,
		"DRONE_COMMIT_SHA":           params.Build.After,
		"DRONE_COMMIT_BEFORE":        params.Build.Before,
		"DRONE_COMMIT_AFTER":         params.Build.After,
		"DRONE_COMMIT_REF":           params.Build.Ref,
		"DRONE_COMMIT_BRANCH":        params.Build.Target,
		"DRONE_COMMIT_LINK":          params.Build.Link,
		"DRONE_COMMIT_MESSAGE":       params.Build.Message,
		"DRONE_COMMIT_AUTHOR":        params.Build.Author,
		"DRONE_COMMIT_AUTHOR_EMAIL":  params.Build.AuthorEmail,
		"DRONE_COMMIT_AUTHOR_AVATAR": params.Build.AuthorAvatar,
		"DRONE_COMMIT_AUTHOR_NAME":   params.Build.AuthorName,
		"DRONE_BUILD_NUMBER":         fmt.Sprint(params.Build.Number),
		"DRONE_BUILD_PARENT":         fmt.Sprint(params.Build.Parent),
		"DRONE_BUILD_EVENT":          params.Build.Event,
		"DRONE_BUILD_ACTION":         params.Build.Action,
		"DRONE_BUILD_TRIGGER":        params.Build.Trigger,
		"DRONE_DEPLOY_TO":            params.Build.Deploy,
		"DRONE_DEPLOY_ID":            fmt.Sprint(params.Build.DeployID),
		"DRONE_PULL_REQUEST":         re.FindString(params.Build.Ref),
		"DRONE_PULL_REQUEST_TITLE":   params.Build.Title,
		"DRONE_BUILD_LINK": fmt.Sprintf(
			"%s://%s/%s/%d",
			params.System.Proto,
			params.System.Host,
			params.Repo.Slug,
			params.Build.Number,
		),
	}
	// if the pipeline is for a tag, add the tag
	// environment variables and semver variables.
	if strings.HasPrefix(params.Build.Ref, "refs/tags/") {
		tag := strings.TrimPrefix(params.Build.Ref, "refs/tags/")
		env["DRONE_TAG"] = tag
		// generate semver variables and copy into
		// environment key values.
		for k, v := range toSemverEnv(tag) {
			env[k] = v
		}
	}
	return env
}

// helper function that generates environment variables for
// semver versions.
func toSemverEnv(s string) map[string]string {
	env := map[string]string{}
	version, err := semver.NewVersion(
		strings.TrimPrefix(s, "v"),
	)
	if err != nil {
		env["DRONE_SEMVER_ERROR"] = err.Error()
		return env
	}
	env["DRONE_SEMVER"] = version.String()
	env["DRONE_SEMVER_MAJOR"] = fmt.Sprint(version.Major)
	env["DRONE_SEMVER_MINOR"] = fmt.Sprint(version.Minor)
	env["DRONE_SEMVER_PATCH"] = fmt.Sprint(version.Patch)
	if s := string(version.PreRelease); s != "" {
		env["DRONE_SEMVER_PRERELEASE"] = s
	}
	if version.Metadata != "" {
		env["DRONE_SEMVER_BUILD"] = version.Metadata
	}
	version.Metadata = ""
	version.PreRelease = ""
	env["DRONE_SEMVER_SHORT"] = version.String()
	return env
}

// copyenv returns a copy of the environment variable map.
func copyenv(src map[string]string) map[string]string {
	dst := map[string]string{}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
