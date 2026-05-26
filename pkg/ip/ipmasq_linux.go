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

// SetupIPMasqForNetworks installs rules to masquerade traffic coming from ips of ipns and
// going outside of ipns, using a chain name based on network, ifname, and containerID. The
// backend can be either "iptables" or "nftables"; if it is nil, then a suitable default
// implementation will be used.
func SetupIPMasqForNetworks(backend *string, ipns []*net.IPNet, network, ifname, containerID string) error {
	_ = "STUB: not implemented"
	return nil

	// Prefer iptables, unless only nftables is available
}

// TeardownIPMasqForNetworks undoes the effects of SetupIPMasqForNetworks
func TeardownIPMasqForNetworks(ipns []*net.IPNet, network, ifname, containerID string) error {
	_ = "STUB: not implemented"

	// Do both the iptables and the nftables cleanup, since the pod may have been
	// created with a different version of this plugin or a different configuration.
	return nil
}

// GCIPMasqForNetwork garbage collects stale IPMasq entries for network
func GCIPMasqForNetwork(network string, attachments []types.GCAttachment) error {
	_ = "STUB: not implemented"
	return nil
}
