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

package ip

import (
	"math/big"
	"net"
)

// NextIP returns IP incremented by 1, if IP is invalid, return nil
func NextIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// PrevIP returns IP decremented by 1, if IP is invalid, return nil
func PrevIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// Cmp compares two IPs, returning the usual ordering:
// a < b : -1
// a == b : 0
// a > b : 1
// incomparable : -2
func Cmp(a, b net.IP) int { _ = "STUB: not implemented"; return 0 }

func ipToInt(ip net.IP) *big.Int { _ = "STUB: not implemented"; return nil }

func intToIP(i *big.Int, isIPv6 bool) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// normalizeIP will normalize IP by family,
// IPv4 : 4-byte form
// IPv6 : 16-byte form
// others : nil
func normalizeIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// Network masks off the host portion of the IP, if IPNet is invalid,
// return nil
func Network(ipn *net.IPNet) *net.IPNet { _ = "STUB: not implemented"; return nil }
