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

type NetConf struct {
	types.NetConf
	Master     string `json:"master"`
	VlanID     int    `json:"vlanId"`
	MTU        int    `json:"mtu,omitempty"`
	LinkContNs bool   `json:"linkInContainer,omitempty"`
}

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

func loadConf(args *skel.CmdArgs) (*NetConf, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// check existing and MTU of master interface

func getMTUByName(ifName string, namespace string, inContainer bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func createVlan(conf *NetConf, ifName string, netns ns.NetNS) (*current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// due to kernel bug we have to create with tmpname or it might
// collide with the name on the host and error out

// Re-fetch interface to get all properties/attributes

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Invoke ipam del if err to avoid ip leak

// Convert whatever the IPAM result was into the current Result type

// All addresses belong to the vlan interface

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

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
	}, version.All, bv.BuildString("vlan"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Convert whatever the IPAM result was into the current Result type

// Find interfaces for name whe know, that of host-device inside container

// The namespace must be the same as what was configured

//
// Check prevResults for ips, routes and dns against values found in the container

// Check interface against values found in the container

func validateCniContainerInterface(intf current.Interface, vlanID int, mtu int) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO This works when unit testing via cnitool; fails with ./test.sh
// if masterIndex != vlan.Attrs().ParentIndex {
//   return fmt.Errorf("Container vlan Master %d does not match expected value: %d", vlan.Attrs().ParentIndex, masterIndex)
//}

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// TODO: Check if master interface exists.
