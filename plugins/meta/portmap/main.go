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

// This is a post-setup plugin that establishes port forwarding - using iptables,
// from the host's network interface(s) to a pod's network interface.
//
// It is intended to be used as a chained CNI plugin, and determines the container
// IP from the previous result. If the result includes an IPv6 address, it will
// also be configured. (IPTables will not forward cross-family).
//
// This has one notable limitation: it does not perform any kind of reservation
// of the actual host port. If there is a service on the host, it will have all
// its traffic captured by the container. If another container also claims a given
// port, it will caputure the traffic - it is last-write-wins.
package main

import (
	"net"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

type PortMapper interface {
	forwardPorts(config *PortMapConf, containerNet net.IPNet) error
	checkPorts(config *PortMapConf, containerNet net.IPNet) error
	unforwardPorts(config *PortMapConf) error
}

// These are vars rather than consts so we can "&" them
var (
	iptablesBackend = "iptables"
	nftablesBackend = "nftables"
)

// PortMapEntry corresponds to a single entry in the port_mappings argument,
// see CONVENTIONS.md
type PortMapEntry struct {
	HostPort      int    `json:"hostPort"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
	HostIP        string `json:"hostIP,omitempty"`
}

type PortMapConf struct {
	types.NetConf

	mapper PortMapper

	// Generic config
	Backend       *string   `json:"backend,omitempty"`
	SNAT          *bool     `json:"snat,omitempty"`
	ConditionsV4  *[]string `json:"conditionsV4"`
	ConditionsV6  *[]string `json:"conditionsV6"`
	MasqAll       bool      `json:"masqAll,omitempty"`
	MarkMasqBit   *int      `json:"markMasqBit"`
	RuntimeConfig struct {
		PortMaps []PortMapEntry `json:"portMappings,omitempty"`
	} `json:"runtimeConfig,omitempty"`

	// iptables-backend-specific config
	ExternalSetMarkChain *string `json:"externalSetMarkChain"`

	// These are fields parsed out of the config or the environment;
	// included here for convenience
	ContainerID string    `json:"-"`
	ContIPv4    net.IPNet `json:"-"`
	ContIPv6    net.IPNet `json:"-"`
}

// The default mark bit to signal that masquerading is required
// Kubernetes uses 14 and 15, Calico uses 20-31.
const DefaultMarkBit = 13

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Delete conntrack entries for UDP to avoid conntrack blackholing traffic
// due to stale connections. We do that after the iptables rules are set, so
// the new traffic uses them. Failures are informative only.

// Set the route_localnet bit on the host interface, so that
// 127/8 can cross a routing boundary.

// Delete conntrack entries for UDP to avoid conntrack blackholing traffic
// due to stale connections. We do that after the iptables rules are set, so
// the new traffic uses them. Failures are informative only.

// Pass through the previous result

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// We don't need to parse out whether or not we're using v6 or snat,
// deletion is idempotent

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.All, bv.BuildString("portmap"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Ensure we have previous result.

// parseConfig parses the supplied configuration (and prevResult) from stdin.
func parseConfig(stdin []byte, ifName string) (*PortMapConf, *current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Parse previous result.

// go constants are "special"

// Reject invalid port numbers

// Skip known non-sandbox interfaces

// ensureBackend validates and/or sets conf.Backend
func ensureBackend(conf *PortMapConf) error { _ = "STUB: not implemented"; return nil }

// If backend wasn't requested explicitly, default to iptables, unless it is not
// available (and nftables is). FIXME: flip this default at some point.

// Make sure we dont have config for the wrong backend

// OK

// detectBackendOfConditions returns "iptables" if conditions contains iptables
// conditions, "nftables" if it contains nftables conditions, and "" if it is empty.
func detectBackendOfConditions(conditions *[]string) string { _ = "STUB: not implemented"; return "" }

// The first character of any iptables condition would either be an hyphen
// (e.g. "-d", "--sport", "-m") or an exclamation mark.
// No nftables condition would start that way. (An nftables condition might
// include a negative number, but not as the first token.)
