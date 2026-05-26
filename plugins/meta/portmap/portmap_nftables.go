// Copyright 2023 CNI authors
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

	"sigs.k8s.io/knftables"
)

const (
	tableName = "cni_hostport"

	// Intermediate chain to jump to both 'hostip_hostports' and 'hostports'
	hostPortsAllChain = "hostports_all"
	// This chain was used for rules with hostIP, we do not use it anymore,
	// but we keep it to make upgrade transparent
	hostIPHostPortsChain = "hostip_hostports"
	// Chain containing all the rules
	hostPortsChain    = "hostports"
	masqueradingChain = "masquerading"
)

// The nftables portmap implementation is fairly similar to the iptables implementation:
// we add a rule for each mapping, with a comment containing a hash of the container ID,
// so that we can later reliably delete the rules we want. (This is important because in
// edge cases, it's possible the plugin might see "ADD container A with IP 192.168.1.3",
// followed by "ADD container B with IP 192.168.1.3" followed by "DEL container A with IP
// 192.168.1.3", and we need to make sure that the DEL causes us to delete the rule for
// container A, and not the rule for container B.) This iptables implementation actually
// uses a separate chain per container but there's not really any need for that...
//
// As with pkg/ip/ipmasq_nftables_linux.go, it would be more nftables-y to have a chain
// with a single rule doing a lookup against a map with an element per mapping, rather
// than having a chain with a rule per mapping. But there's no easy, non-racy way to say
// "delete the element 192.168.1.3 from the map, but only if it was added for container A,
// not if it was added for container B".

type portMapperNFTables struct {
	ipv4 knftables.Interface
	ipv6 knftables.Interface
}

// getPortMapNFT creates an nftables.Interface for port mapping for the IP family of ipn
func (pmNFT *portMapperNFTables) getPortMapNFT(ipv6 bool) (knftables.Interface, error) {
	_ = "STUB: not implemented"
	return *new(knftables.Interface), nil
}

// forwardPorts establishes port forwarding to a given container IP.
// containerNet.IP can be either v4 or v6.
func (pmNFT *portMapperNFTables) forwardPorts(config *PortMapConf, containerNet net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure basic rule structure

// setup intermediate chain

// Set up this container

// Ignore wrong-IP-family HostIPs

// Unspecified addresses cannot be used as destination

// we add the rule to 'hostports' instead of 'hostip_hostports'
// as we want to remove 'hostip_hostports' long-term

// Add mark-to-masquerade rules for hairpin and localhost
// In theory we should validate that the original dst IP and port are as
// expected, but *any* traffic matching one of these patterns would need
// to be masqueraded to be able to work correctly anyway.

// MasqAll: match traffic from any source IP

// Default: only match traffic from container's own IP (hairpin)

// Only add localhost rule when MasqAll is false
// (when MasqAll is true, 0.0.0.0/0 already covers 127.0.0.1)

func (pmNFT *portMapperNFTables) checkPorts(config *PortMapConf, containerNet net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore wrong-IP-family HostIPs

// When MasqAll is false and IPv4, we have 2 rules:
// 1. hairpin rule (container IP -> container IP)
// 2. localhost rule (127.0.0.1 -> container IP)
// When MasqAll is true, we only have 1 rule (0.0.0.0/0 -> container IP)

func checkPortsAgainstRules(nft knftables.Interface, chain, comment string, nPorts int) error {
	_ = "STUB: not implemented"
	return nil
}

// unforwardPorts deletes any nftables rules created by this plugin.
// It should be idempotent - it will not error if the chain does not exist.
func (pmNFT *portMapperNFTables) unforwardPorts(config *PortMapConf) error {
	_ = "STUB: not implemented"
	// Always clear both IPv4 and IPv6, just to be sure
	return nil
}
