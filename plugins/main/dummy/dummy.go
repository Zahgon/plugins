// Copyright 2022 Arista Networks
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
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

func parseNetConf(bytes []byte) (*types.NetConf, error) { _ = "STUB: not implemented"; return nil, nil }

func createDummy(ifName string, netns ns.NetNS) (*current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Re-fetch interface to get all properties/attributes

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Delete link if err to avoid link leak in this ns

// defer ipam deletion to avoid ip leak

// convert IPAMResult to current Result type

// all addresses apply to the container dummy interface

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
	}, version.All, bv.BuildString("dummy"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Convert whatever the IPAM result was into the current Result type

// Find interfaces for name whe know, that of dummy device inside container

// The namespace must be the same as what was configured

//
// Check prevResults for ips, routes and dns against values found in the container

// Check interface against values found in the container

func validateCniContainerInterface(intf current.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }
