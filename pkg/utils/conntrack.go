// Copyright 2020 CNI authors
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

package utils

import (
	"github.com/vishvananda/netlink"
)

// Assigned Internet Protocol Numbers
// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
const (
	PROTOCOL_TCP  = 6
	PROTOCOL_UDP  = 17
	PROTOCOL_SCTP = 132
)

// getNetlinkFamily returns the Netlink IP family constant
func getNetlinkFamily(isIPv6 bool) netlink.InetFamily {
	_ = "STUB: not implemented"
	return *new(netlink.InetFamily)
}

// DeleteConntrackEntriesForDstIP delete the conntrack entries for the connections
// specified by the given destination IP and protocol
func DeleteConntrackEntriesForDstIP(dstIP string, protocol uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteConntrackEntriesForDstPort delete the conntrack entries for the connections specified
// by the given destination port, protocol and IP family
func DeleteConntrackEntriesForDstPort(port uint16, protocol uint8, family netlink.InetFamily) error {
	_ = "STUB: not implemented"
	return nil
}
