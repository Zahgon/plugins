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

package main

import (
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

type NetConf struct {
	types.NetConf
	IPMasq        bool    `json:"ipMasq"`
	IPMasqBackend *string `json:"ipMasqBackend,omitempty"`
	MTU           int     `json:"mtu"`
}

func setupContainerVeth(netns ns.NetNS, ifName string, mtu int, pr *current.Result) (*current.Interface, *current.Interface, error) {
	_ = "STUB: not implemented"
	// The IPAM result will be something like IP=192.168.3.5/24, GW=192.168.3.1.
	// What we want is really a point-to-point link but veth does not support IFF_POINTTOPOINT.
	// Next best thing would be to let it ARP but set interface to 192.168.3.5/32 and
	// add a route like "192.168.3.0/24 via 192.168.3.1 dev $ifName".
	// Unfortunately that won't work as the GW will be outside the interface's subnet.
	return nil, nil, nil
}

// Our solution is to configure the interface with 192.168.3.5/24, then delete the
// "192.168.3.0/24 dev $ifName" route that was automatically added. Then we add
// "192.168.3.1/32 dev $ifName" and "192.168.3.0/24 via 192.168.3.1 dev $ifName".
// In other words we force all traffic to ARP via the gateway except for GW itself.

// All addresses apply to the container veth interface

// Delete the route that was automatically added

func setupHostVeth(vethName string, result *current.Result) error {
	_ = "STUB: not implemented"
	// hostVeth moved namespaces and may have a new ifindex
	return nil
}

// dst happens to be the same as IP/net of host veth

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Invoke ipam del if err to avoid ip leak

// Convert whatever the IPAM result was into the current Result type

// Only override the DNS settings in the previous result if any DNS fields
// were provided to the ptp plugin. This allows, for example, IPAM plugins
// to specify the DNS settings instead of the ptp plugin.

func dnsConfSet(dnsConf types.DNS) bool { _ = "STUB: not implemented"; return false }

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// There is a netns so try to clean up. Delete can be called multiple times
// so don't return an error if the device is already removed.
// If the device isn't there then don't try to clean up IP masq either.

//  if NetNs is passed down by the Cloud Orchestration Engine, or if it called multiple times
// so don't return an error if the device is already removed.
// https://github.com/kubernetes/kubernetes/issues/43014#issuecomment-287164444

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:    cmdAdd,
		Check:  cmdCheck,
		Del:    cmdDel,
		Status: cmdStatus,
		/* FIXME GC */
	}, version.All, bv.BuildString("ptp"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Convert whatever the IPAM result was into the current Result type

// Find interfaces for name whe know, that of host-device inside container

// The namespace must be the same as what was configured

//
// Check prevResults for ips, routes and dns against values found in the container

// Check interface against values found in the container

func validateCniContainerInterface(intf current.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }
