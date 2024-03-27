#!/bin/bash

# Copyright 2024 The ChromiumOS Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This runs the build.sh script in the chromiumos chroot go environment.
# Note: The btpeer environment will differ slightly from the chromiumos chroot,
# but it's close enough to test the build and run unit tests.

set -e

SCRIPT_DIR="$(dirname "$(realpath -e "${BASH_SOURCE[0]}")")"
source "${SCRIPT_DIR}/lib/project_vars.sh"

echo "Building ${PROJECT_EXE_NAME} with ChromeOS chroot GOROOT"
export GOROOT="${CHROOT_GOROOT}"
"${PROJECT_DIR}/scripts/build_local.sh"
echo "Successfully built ${PROJECT_EXE_NAME} with ChromeOS chroot GOROOT"
