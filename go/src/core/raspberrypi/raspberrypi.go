// Copyright 2024 The ChromiumOS Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package raspberrypi

import (
	"context"
	"fmt"
	"regexp"

	"go.chromium.org/chromiumos/platform/btpeerd/core/exec"
)

const osReleaseFilePath = "/etc/os-release"

// OSVersion returns the version of the Raspberry Pi OS as specified in the
// "/etc/os-release" file as the PRETTY_NAME.
func OSVersion(ctx context.Context, runner exec.CmdRunner) (string, error) {
	osReleaseFileContents, err := runner.Output(ctx, exec.DefaultTimeout, "cat", osReleaseFilePath)
	if err != nil {
		return "", fmt.Errorf("os version: failed to read contents of file %q: %w", osReleaseFilePath, err)
	}
	matcher := regexp.MustCompile(`PRETTY_NAME="(.+)"\n`)
	match := matcher.FindStringSubmatch(string(osReleaseFileContents))
	if len(match) != 2 {
		return "", fmt.Errorf("os version: failed parse PRETTY_NAME from %q file contents %q", osReleaseFilePath, string(osReleaseFileContents))
	}
	return match[1], nil
}
