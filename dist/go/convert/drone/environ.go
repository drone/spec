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
	"github.com/drone/drone-go/drone"
)

// regular expression to extract the pull request number
// from the git ref (e.g. refs/pulls/{d}/head)
var re = regexp.MustCompile("\\d+")

// helper function returns repository and build environment
// variables for environment substitution in the yaml.
func toEnv(repo *drone.Repo, build *drone.Build, system *drone.System) map[string]string {
	env := map[string]string{
		"DRONE_REPO":                 repo.Slug,
		"DRONE_REPO_OWNER":           repo.Namespace,
		"DRONE_REPO_NAMESPACE":       repo.Namespace,
		"DRONE_REPO_NAME":            repo.Name,
		"DRONE_REPO_LINK":            repo.Link,
		"DRONE_REPO_BRANCH":          repo.Branch,
		"DRONE_GIT_HTTP_URL":         repo.HTTPURL,
		"DRONE_GIT_SSH_URL":          repo.SSHURL,
		"DRONE_REMOTE_URL":           repo.HTTPURL,
		"DRONE_BRANCH":               build.Target,
		"DRONE_SOURCE_BRANCH":        build.Source,
		"DRONE_TARGET_BRANCH":        build.Target,
		"DRONE_COMMIT":               build.After,
		"DRONE_COMMIT_SHA":           build.After,
		"DRONE_COMMIT_BEFORE":        build.Before,
		"DRONE_COMMIT_AFTER":         build.After,
		"DRONE_COMMIT_REF":           build.Ref,
		"DRONE_COMMIT_BRANCH":        build.Target,
		"DRONE_COMMIT_LINK":          build.Link,
		"DRONE_COMMIT_MESSAGE":       build.Message,
		"DRONE_COMMIT_AUTHOR":        build.Author,
		"DRONE_COMMIT_AUTHOR_EMAIL":  build.AuthorEmail,
		"DRONE_COMMIT_AUTHOR_AVATAR": build.AuthorAvatar,
		"DRONE_COMMIT_AUTHOR_NAME":   build.AuthorName,
		"DRONE_BUILD_NUMBER":         fmt.Sprint(build.Number),
		"DRONE_BUILD_PARENT":         fmt.Sprint(build.Parent),
		"DRONE_BUILD_EVENT":          build.Event,
		"DRONE_BUILD_ACTION":         build.Action,
		"DRONE_BUILD_TRIGGER":        build.Trigger,
		"DRONE_DEPLOY_TO":            build.Deploy,
		"DRONE_DEPLOY_ID":            fmt.Sprint(build.DeployID),
		"DRONE_PULL_REQUEST":         re.FindString(build.Ref),
		"DRONE_PULL_REQUEST_TITLE":   build.Title,
		"DRONE_BUILD_LINK": fmt.Sprintf(
			"%s://%s/%s/%d",
			system.Proto,
			system.Host,
			repo.Slug,
			build.Number,
		),
	}
	// if the pipeline is for a tag, add the tag
	// environment variables and semver variables.
	if strings.HasPrefix(build.Ref, "refs/tags/") {
		tag := strings.TrimPrefix(build.Ref, "refs/tags/")
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
