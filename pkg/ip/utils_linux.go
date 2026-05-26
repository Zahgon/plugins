//go:build linux
// +build linux

// Copyright 2016 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ip

import (
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
)

func ValidateExpectedInterfaceIPs(ifName string, resultIPs []*current.IPConfig) error {
	_ = "STUB: not implemented"
	// Ensure ips
	return nil
}

// Convert the host/prefixlen to just prefix for route lookup.

func ValidateExpectedRoute(resultRoutes []*types.Route) error {
	_ = "STUB: not implemented"
	// Ensure that each static route in prevResults is found in the routing table
	return nil
}

// Default route needs Dst set to nil

// Default route needs Dst set to nil
