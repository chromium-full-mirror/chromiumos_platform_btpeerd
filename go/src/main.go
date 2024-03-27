// Copyright 2024 The ChromiumOS Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package main includes the main function for running service as an executable.
package main

import (
	"fmt"
	"os"
	"time"

	"go.chromium.org/chromiumos/config/go/test/lab/api"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// Test usage of config repo
	chameleondConfig := api.BluetoothPeerChameleondConfig{}
	chameleondConfigJSON, err := protojson.Marshal(&chameleondConfig)
	if err != nil {
		fmt.Printf("Failed to marshall chameleond config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully used config repo: %s\n", chameleondConfigJSON)

	// Just keep printing hello world to test that service runs continuously.
	for true {
		fmt.Printf("Hello world from btpeerd! (%s)\n", time.Now())
		time.Sleep(time.Second)
	}
}
