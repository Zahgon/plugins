// Copyright 2015-2017 CNI authors
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
	"net"

	"github.com/vishvananda/netlink"
)

// AddRoute adds a universally-scoped route to a device.
func AddRoute(ipn *net.IPNet, gw net.IP, dev netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

// AddHostRoute adds a host-scoped route to a device.
func AddHostRoute(ipn *net.IPNet, gw net.IP, dev netlink.Link) error {
	_ = "STUB: not implemented"
	return nil
}

// AddDefaultRoute sets the default route on the given gateway.
func AddDefaultRoute(gw net.IP, dev netlink.Link) error { _ = "STUB: not implemented"; return nil }

// IsIPNetZero check if the IPNet is "0.0.0.0/0" or "::/0"
// This is needed as go-netlink replaces nil Dst with a '0' IPNet since
// https://github.com/vishvananda/netlink/commit/acdc658b8613655ddb69f978e9fb4cf413e2b830
func IsIPNetZero(ipnet *net.IPNet) bool { _ = "STUB: not implemented"; return false }
