// Copyright 2018 CNI authors
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
	"github.com/vishvananda/netlink"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

const (
	maxIfbDeviceLength = 15
	ifbDevicePrefix    = "bwp"
)

// BandwidthEntry corresponds to a single entry in the bandwidth argument,
// see CONVENTIONS.md
type BandwidthEntry struct {
	IngressRate  uint64 `json:"ingressRate"`  // Bandwidth rate in bps for traffic through container. 0 for no limit. If ingressRate is set, ingressBurst must also be set
	IngressBurst uint64 `json:"ingressBurst"` // Bandwidth burst in bits for traffic through container. 0 for no limit. If ingressBurst is set, ingressRate must also be set

	EgressRate  uint64 `json:"egressRate"`  // Bandwidth rate in bps for traffic through container. 0 for no limit. If egressRate is set, egressBurst must also be set
	EgressBurst uint64 `json:"egressBurst"` // Bandwidth burst in bits for traffic through container. 0 for no limit. If egressBurst is set, egressRate must also be set
}

func (bw *BandwidthEntry) isZero() bool { _ = "STUB: not implemented"; return false }

type PluginConf struct {
	types.NetConf

	RuntimeConfig struct {
		Bandwidth *BandwidthEntry `json:"bandwidth,omitempty"`
	} `json:"runtimeConfig,omitempty"`

	*BandwidthEntry
}

// parseConfig parses the supplied configuration (and prevResult) from stdin.
func parseConfig(stdin []byte) (*PluginConf, error) { _ = "STUB: not implemented"; return nil, nil }

func getBandwidth(conf *PluginConf) *BandwidthEntry { _ = "STUB: not implemented"; return nil }

func validateRateAndBurst(rate, burst uint64) error { _ = "STUB: not implemented"; return nil }

func getIfbDeviceName(networkName string, containerID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getMTU(deviceName string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// get the veth peer of container interface in host namespace
func getHostInterface(interfaces []*current.Interface, containerIfName string, netns ns.NetNS) (*current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get veth peer index of container interface

// find host interface by index

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.VersionsStartingFrom("0.3.0"), bv.BuildString("bandwidth"))
}

func SafeQdiscList(link netlink.Link) ([]netlink.Qdisc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter out pfifo_fast qdiscs because
// older kernels don't return them

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// No bandwidth config; nothing to do.
