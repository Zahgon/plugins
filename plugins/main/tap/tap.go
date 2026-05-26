// Copyright 2022 CNI authors
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
	MultiQueue     bool      `json:"multiQueue"`
	MTU            int       `json:"mtu"`
	Mac            string    `json:"mac,omitempty"`
	Owner          *uint32   `json:"owner,omitempty"`
	Group          *uint32   `json:"group,omitempty"`
	SelinuxContext string    `json:"selinuxContext,omitempty"`
	Bridge         string    `json:"bridge,omitempty"`
	Args           *struct{} `json:"args,omitempty"`
	RuntimeConfig  struct {
		Mac string `json:"mac,omitempty"`
	} `json:"runtimeConfig,omitempty"`
}

// MacEnvArgs represents CNI_ARG
type MacEnvArgs struct {
	types.CommonArgs
	MAC types.UnmarshallableString `json:"mac,omitempty"`
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

// We want to share the parent process std{in|out|err} - fds 0 through 2.
// Since the FDs are inherited on fork / exec, we close on exec all others.
func closeFileDescriptorsOnExec() { _ = "STUB: not implemented"; return }

// Due to issues with the vishvananda/netlink library (fix pending) it is not possible to create an ownerless/groupless
// tap device. Until the issue is fixed, the workaround for creating a tap device with no owner/group is to use the iptool
func createTapWithIptool(tmpName string, mtu int, multiqueue bool, mac string, owner *uint32, group *uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func createLinkWithNetlink(tmpName string, mtu int, nsFd int, multiqueue bool, mac string, owner *uint32, group *uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func createLink(tmpName string, conf *NetConf, netns ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

func createTap(conf *NetConf, ifName string, netns ns.NetNS) (*current.Interface, error) {
	_ = "STUB: not implemented"
	return nil,

		// due to kernel bug we have to create with tmpName or it might
		// collide with the name on the host and error out
		nil
}

// Re-fetch link to get all properties/attributes

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Delete link if err to avoid link leak in this ns

// Assume L2 interface only

// run the IPAM plugin and get back the config to apply

// Invoke ipam del if err to avoid ip leak

// Convert whatever the IPAM result was into the current Result type

// All addresses apply to the container tap interface

// For L2 just change interface status to up

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// There is a netns so try to clean up. Delete can be called multiple times
// so don't return an error if the device is already removed.

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
	}, version.All, bv.BuildString("tap"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Parse previous result.

// Find interfaces for names whe know, tap device name inside container

// The namespace must be the same as what was configured

// Check prevResults for ips, routes and dns against values found in the container

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }
