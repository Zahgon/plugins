// Copyright 2016 CNI authors
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

// This is a "meta-plugin". It reads in its own netconf, it does not create
// any network interface but just changes the network sysctl.

package main

import (
	"net"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

// FirewallNetConf represents the firewall configuration.
type FirewallNetConf struct {
	types.NetConf

	// Backend is the firewall type to add rules to.  Allowed values are
	// 'iptables' and 'firewalld'.
	Backend string `json:"backend"`

	// IptablesAdminChainName is an optional name to use instead of the default
	// admin rules override chain name that includes the interface name.
	IptablesAdminChainName string `json:"iptablesAdminChainName,omitempty"`

	// FirewalldZone is an optional firewalld zone to place the interface into.  If
	// the firewalld backend is used but the zone is not given, it defaults
	// to 'trusted'
	FirewalldZone string `json:"firewalldZone,omitempty"`

	// IngressPolicy is an optional ingress policy.
	// Defaults to "open".
	IngressPolicy IngressPolicy `json:"ingressPolicy,omitempty"`
}

// IngressPolicy is an ingress policy string.
type IngressPolicy = string

const (
	// IngressPolicyOpen ("open"): all inbound connections to the container are accepted.
	// IngressPolicyOpen is the default ingress policy.
	IngressPolicyOpen IngressPolicy = "open"

	// IngressPolicySameBridge ("same-bridge"): connections from the same bridge are accepted, others are blocked.
	// This is similar to how Docker libnetwork works.
	// IngressPolicySameBridge executes `iptables` regardless to the value of `Backend`.
	// IngressPolicySameBridge may not work as expected for non-bridge networks.
	IngressPolicySameBridge IngressPolicy = "same-bridge"

	// IngressPolicyIsolated ("isolated"): similar to ingress policy "same-bridge" with the exception
	// that connections from the same bridge are also blocked.
	// This is equivalent to Docker network option "enable_icc" when set to false.
	// IngressPolicyIsolated executes `iptables` regardless to the value of `Backend`.
	// IngressPolicyIsolated may not work as expected for non-bridge networks.
	IngressPolicyIsolated IngressPolicy = "isolated"
)

type FirewallBackend interface {
	Add(*FirewallNetConf, *current.Result) error
	Del(*FirewallNetConf, *current.Result) error
	Check(*FirewallNetConf, *current.Result) error
}

func ipString(ip net.IPNet) string { _ = "STUB: not implemented"; return "" }

func parseConf(data []byte) (*FirewallNetConf, *current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Default the firewalld zone to trusted

// Parse previous result.

// return early if there was no previous result, which is allowed for DEL calls

// Parse previous result.

func getBackend(conf *FirewallNetConf) (FirewallBackend, error) {
	_ = "STUB: not implemented"
	return *new(FirewallBackend), nil
}

// Default to firewalld if it's running

// Otherwise iptables

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Runtime errors are ignored

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.VersionsStartingFrom("0.4.0"), bv.BuildString("firewall"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Ensure we have previous result.
