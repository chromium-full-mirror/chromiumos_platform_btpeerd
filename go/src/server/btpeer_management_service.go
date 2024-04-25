// Copyright 2024 The ChromiumOS Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package server

import (
	"context"

	"go.chromium.org/chromiumos/config/go/test/lab/api/btpeerd"
)

type BtpeerManagementServiceServer struct{}

func NewBtpeerManagementServiceServer() *BtpeerManagementServiceServer {
	return &BtpeerManagementServiceServer{}
}

func (s *BtpeerManagementServiceServer) DeviceInfo(ctx context.Context, request *btpeerd.DeviceInfoRequest) (*btpeerd.DeviceInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BtpeerManagementServiceServer) DeviceStatus(ctx context.Context, request *btpeerd.DeviceStatusRequest) (*btpeerd.DeviceStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BtpeerManagementServiceServer) Reboot(ctx context.Context, request *btpeerd.RebootRequest) (*btpeerd.RebootResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BtpeerManagementServiceServer) GetActiveBluetoothStack(ctx context.Context, request *btpeerd.GetActiveBluetoothStackRequest) (*btpeerd.GetActiveBluetoothStackResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BtpeerManagementServiceServer) SetActiveBluetoothStack(ctx context.Context, request *btpeerd.SetActiveBluetoothStackRequest) (*btpeerd.SetActiveBluetoothStackResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BtpeerManagementServiceServer) GetActiveBluetoothStackAPI(ctx context.Context, request *btpeerd.GetActiveBluetoothStackAPIRequest) (*btpeerd.GetActiveBluetoothStackAPIResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *BtpeerManagementServiceServer) SetActiveBluetoothStackAPI(ctx context.Context, request *btpeerd.SetActiveBluetoothStackAPIRequest) (*btpeerd.SetActiveBluetoothStackAPIResponse, error) {
	//TODO implement me
	panic("implement me")
}
