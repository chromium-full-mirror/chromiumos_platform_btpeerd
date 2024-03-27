#!/bin/bash

# Copyright 2024 The ChromiumOS Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This builds the project using the current go environment.

set -e

SCRIPT_DIR="$(dirname "$(realpath -e "${BASH_SOURCE[0]}")")"
source "${SCRIPT_DIR}/lib/project_vars.sh"

set -x

# Get packages, format files, run tests, and build executable.
cd "${PROJECT_GOPATH}/src" && go mod tidy
cd "${PROJECT_GOPATH}/src" && go fmt ./...
cd "${PROJECT_GOPATH}/src" && go test ./...
cd "${PROJECT_GOPATH}/src" && go build -o "${PROJECT_EXE_PATH}"
