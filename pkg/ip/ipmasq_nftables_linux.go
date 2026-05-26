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

package ip

import (
	"net"

	"sigs.k8s.io/knftables"

	"github.com/containernetworking/cni/pkg/types"
)

const (
	ipMasqTableName = "cni_plugins_masquerade"
	ipMasqChainName = "masq_checks"
)

// The nftables ipmasq implementation is mostly like the iptables implementation, with
// minor updates to fix a bug (adding `ifname`) and to allow future GC support.
//
// We add a rule for each mapping, with a comment containing a hash of its identifiers,
// so that we can later reliably delete the rules we want. (This is important because in
// edge cases, it's possible the plugin might see "ADD container A with IP 192.168.1.3",
// followed by "ADD container B with IP 192.168.1.3" followed by "DEL container A with IP
// 192.168.1.3", and we need to make sure that the DEL causes us to delete the rule for
// container A, and not the rule for container B.)
//
// It would be more nftables-y to have a chain with a single rule doing a lookup against a
// set with an element per mapping, rather than having a chain with a rule per mapping.
// But there's no easy, non-racy way to say "delete the element 192.168.1.3 from the set,
// but only if it was added for container A, not if it was added for container B".

// hashForNetwork returns a unique hash for this network
func hashForNetwork(network string) string { _ = "STUB: not implemented"; return "" }

// hashForInstance returns a unique hash identifying the rules for this
// network/ifname/containerID
func hashForInstance(network, ifname, containerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// commentForInstance returns a comment string that begins with a unique hash and
// ends with a (possibly-truncated) human-readable description.
func commentForInstance(network, ifname, containerID string) string {
	_ = "STUB: not implemented"
	return ""
}

// setupIPMasqNFTables is the nftables-based implementation of SetupIPMasqForNetworks
func setupIPMasqNFTables(ipns []*net.IPNet, network, ifname, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

func setupIPMasqNFTablesWithInterface(nft knftables.Interface, ipns []*net.IPNet, network, ifname, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that our table and chains exist.

// Ensure that the postrouting chain exists and has the correct rules. (Has to be
// done after creating ipMasqChainName, so we can jump to it.)

// Delete stale rules, add new rules to masquerade chain

// e.g. if ipn is "192.168.1.4/24", then dstNet is "192.168.1.0/24"

// teardownIPMasqNFTables is the nftables-based implementation of TeardownIPMasqForNetworks
func teardownIPMasqNFTables(ipns []*net.IPNet, network, ifname, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

func teardownIPMasqNFTablesWithInterface(nft knftables.Interface, _ []*net.IPNet, network, ifname, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// gcIPMasqNFTables is the nftables-based implementation of GCIPMasqForNetwork
func gcIPMasqNFTables(network string, attachments []types.GCAttachment) error {
	_ = "STUB: not implemented"
	return nil
}

func gcIPMasqNFTablesWithInterface(nft knftables.Interface, network string, attachments []types.GCAttachment) error {
	_ = "STUB: not implemented"
	// Find all rules for the network
	return nil
}

// Compute the comments for all elements of attachments

// Delete anything in rules that isn't in validAttachments

// findRules finds rules with comments that start with commentPrefix.
func findRules(nft knftables.Interface, commentPrefix string) ([]*knftables.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If ipMasqChainName doesn't exist yet, that's fine
