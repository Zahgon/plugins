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

package main

import (
	"net"
)

// fmtIpPort correctly formats ip:port literals for iptables and ip6tables -
// need to wrap v6 literals in a []
func fmtIPPort(ip net.IP, port int) string { _ = "STUB: not implemented"; return "" }

// getRoutableHostIF will try and determine which interface routes the container's
// traffic. This is the one on which we disable martian filtering.
func getRoutableHostIF(containerIP net.IP) string { _ = "STUB: not implemented"; return "" }

// enableLocalnetRouting tells the kernel not to treat 127/8 as a martian,
// so that connections with a source ip of 127/8 can cross a routing boundary.
func enableLocalnetRouting(ifName string) error { _ = "STUB: not implemented"; return nil }

// groupByProto groups port numbers by protocol
func groupByProto(entries []PortMapEntry) map[string][]int { _ = "STUB: not implemented"; return nil }

// splitPortList splits a list of integers in to one or more comma-separated
// string values, for use by multiport. Multiport only allows up to 15 ports
// per entry.
func splitPortList(l []int) []string { _ = "STUB: not implemented"; return nil }

// trimComment makes sure no comment is over the iptables limit of 255 chars
func trimComment(val string) string { _ = "STUB: not implemented"; return "" }
