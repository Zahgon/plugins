// Copyright 2017 CNI authors
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

package allocator

import (
	"net"

	"github.com/containernetworking/cni/pkg/types"
)

// Canonicalize takes a given range and ensures that all information is consistent,
// filling out Start, End, and Gateway with sane values if missing
func (r *Range) Canonicalize() error { _ = "STUB: not implemented"; return nil }

// Can't create an allocator for a network with no addresses, eg
// a /32 or /31

// Ensure Subnet IP is the network address, not some other address

// If the gateway is nil, claim .1

// RangeStart: If specified, make sure it's sane (inside the subnet),
// otherwise use the first free IP (i.e. .1) - this will conflict with the
// gateway but we skip it in the iterator

// RangeEnd: If specified, verify sanity. Otherwise, add a sensible default
// (e.g. for a /24: .254 if IPv4, ::255 if IPv6)

// IsValidIP checks if a given ip is a valid, allocatable address in a given Range
func (r *Range) Contains(addr net.IP) bool { _ = "STUB: not implemented"; return false }

// Not the same address family

// Not in network

// We ignore nils here so we can use this function as we initialize the range.

// Before the range start

// After the  range end

// Overlaps returns true if there is any overlap between ranges
func (r *Range) Overlaps(r1 *Range) bool {
	_ = "STUB: not implemented"
	// different families
	return false
}

func (r *Range) String() string { _ = "STUB: not implemented"; return "" }

// canonicalizeIP makes sure a provided ip is in standard form
func canonicalizeIP(ip *net.IP) error { _ = "STUB: not implemented"; return nil }

// Determine the last IP of a subnet, excluding the broadcast if IPv4
func lastIP(subnet types.IPNet) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }
