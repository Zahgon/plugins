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
	"net"

	"github.com/containernetworking/cni/pkg/types"
)

// setupIPMasqIPTables is the iptables-based implementation of SetupIPMasqForNetworks
func setupIPMasqIPTables(ipns []*net.IPNet, network, _, containerID string) error {
	_ = "STUB: not implemented"
	// Note: for historical reasons, the iptables implementation ignores ifname.
	return nil
}

// SetupIPMasq installs iptables rules to masquerade traffic
// coming from ip of ipn and going outside of ipn.
//
// Deprecated: This function only supports iptables. Use SetupIPMasqForNetworks, which
// supports both iptables and nftables.
func SetupIPMasq(ipn *net.IPNet, chain string, comment string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create chain if doesn't exist

// Packets to this network should not be touched

// Don't masquerade multicast - pods should be able to talk to other pods
// on the local network via multicast.

// Packets from the specific IP of this network will hit the chain

// teardownIPMasqIPTables is the iptables-based implementation of TeardownIPMasqForNetworks
func teardownIPMasqIPTables(ipns []*net.IPNet, network, _, containerID string) error {
	_ = "STUB: not implemented"
	// Note: for historical reasons, the iptables implementation ignores ifname.
	return nil
}

// TeardownIPMasq undoes the effects of SetupIPMasq.
//
// Deprecated: This function only supports iptables. Use TeardownIPMasqForNetworks, which
// supports both iptables and nftables.
func TeardownIPMasq(ipn *net.IPNet, chain string, comment string) error {
	_ = "STUB: not implemented"
	return nil
}

// for downward compatibility

// gcIPMasqIPTables is the iptables-based implementation of GCIPMasqForNetwork
func gcIPMasqIPTables(_ string, _ []types.GCAttachment) error {
	_ = "STUB: not implemented"
	// FIXME: The iptables implementation does not support GC.
	//
	// (In theory, it _could_ backward-compatibly support it, by adding a no-op rule
	// with a comment indicating the network to each chain it creates, so that it
	// could later figure out which chains corresponded to which networks; older
	// implementations would ignore the extra rule but would still correctly delete
	// the chain on teardown (because they ClearChain() before doing DeleteChain()).
	return nil
}

// isNotExist returnst true if the error is from iptables indicating
// that the target does not exist.
func isNotExist(err error) bool { _ = "STUB: not implemented"; return false }
