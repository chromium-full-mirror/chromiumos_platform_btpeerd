#!/bin/bash

# Copyright 2024 The ChromiumOS Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Configures project bash variables.

set -e

SCRIPT_DIR="$(dirname "$(realpath -e "${BASH_SOURCE[0]}")")"
PROJECT_DIR="$(realpath -e "${SCRIPT_DIR}/../..")"
PROJECT_GOPATH="${PROJECT_DIR}/go"
PROJECT_EXE_NAME="btpeerd"
PROJECT_EXE_PATH="${PROJECT_DIR}/go/bin/${PROJECT_EXE_NAME}"
export PROJECT_EXE_PATH
CHROOT_DIR="${PROJECT_DIR}/../../../chroot"
CHROOT_GOROOT="${CHROOT_DIR}/usr/lib/go"
export CHROOT_GOROOT

# Configure project go environment.
export GOPATH="${PROJECT_GOPATH}"
export GO111MODULE=on
