// Copyright 2015 CNI authors
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

package main

import (
	"net"

	dhcp4 "github.com/insomniacslk/dhcp/dhcpv4"

	"github.com/containernetworking/cni/pkg/types"
)

var optionNameToID = map[string]dhcp4.OptionCode{
	"dhcp-client-identifier":  dhcp4.OptionClientIdentifier,
	"subnet-mask":             dhcp4.OptionSubnetMask,
	"routers":                 dhcp4.OptionRouter,
	"host-name":               dhcp4.OptionHostName,
	"user-class":              dhcp4.OptionUserClassInformation,
	"vendor-class-identifier": dhcp4.OptionClassIdentifier,
}

func parseOptionName(option string) (dhcp4.OptionCode, error) {
	_ = "STUB: not implemented"
	return *new(dhcp4.OptionCode), nil
}

func classfulSubnet(sn net.IP) net.IPNet { _ = "STUB: not implemented"; return *new(net.IPNet) }

func parseRoutes(opt []byte) []*types.Route {
	_ = "STUB: not implemented"
	// StaticRoutes format: pairs of:
	// Dest = 4 bytes; Classful IP subnet
	// Router = 4 bytes; IP address of router
	return nil
}
