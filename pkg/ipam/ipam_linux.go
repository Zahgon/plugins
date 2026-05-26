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

package ipam

import (
	"time"

	current "github.com/containernetworking/cni/pkg/types/100"
)

const (
	// Note: use slash as separator so we can have dots in interface name (VLANs)
	DisableIPv6SysctlTemplate    = "net/ipv6/conf/%s/disable_ipv6"
	KeepAddrOnDownSysctlTemplate = "net/ipv6/conf/%s/keep_addr_on_down"

	dadSettleTimeout = 5 * time.Second
)

// ConfigureIface takes the result of IPAM plugin and
// applies to the ifName interface
func ConfigureIface(ifName string, res *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// IP address is for a different interface

// Make sure sysctl "disable_ipv6" is 0 and "keep_addr_on_down" is 1
// if we are about to add an IPv6 address to the interface

// Enabled IPv6 for loopback "lo" and the interface
// being configured

// Read current sysctl value

// Write sysctl to enable IPv6

// Enable "keep_addr_on_down" for the interface being configured
// This prevents the kernel from removing the address when the interface is brought down
