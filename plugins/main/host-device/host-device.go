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

	"github.com/vishvananda/netlink"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

var (
	sysBusPCI       = "/sys/bus/pci/devices"
	sysBusAuxiliary = "/sys/bus/auxiliary/devices"
)

// Array of different linux drivers bound to network device needed for DPDK
var userspaceDrivers = []string{"vfio-pci", "uio_pci_generic", "igb_uio"}

// NetConf for host-device config, look the README to learn how to use those parameters
type NetConf struct {
	types.NetConf
	Device        string `json:"device"` // Device-Name, something like eth0 or can0 etc.
	HWAddr        string `json:"hwaddr"` // MAC Address of target network interface
	DPDKMode      bool
	KernelPath    string `json:"kernelpath"` // Kernelpath of the device
	PCIAddr       string `json:"pciBusID"`   // PCI Address of target network device
	RuntimeConfig struct {
		DeviceID string `json:"deviceID,omitempty"`
	} `json:"runtimeConfig,omitempty"`

	// for internal use
	auxDevice string `json:"-"` // Auxiliary device name as appears on Auxiliary bus (/sys/bus/auxiliary)
}

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

// handleDeviceID updates netconf fields with DeviceID runtime config
func handleDeviceID(netconf *NetConf) error { _ = "STUB: not implemented"; return nil }

// Check if deviceID is a PCI device

// Check if deviceID is an Auxiliary device

func loadConf(bytes []byte) (*NetConf, error) { _ = "STUB: not implemented"; return nil, nil }

// Override device with the standardized DeviceID if provided in Runtime Config.

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Override the device name with the name in the container namespace

// Set the MAC address of the interface

// run the IPAM plugin and get back the config to apply

// Invoke ipam del if err to avoid ip leak

// Convert whatever the IPAM result was into the current Result type

// All addresses apply to the container interface (move from host)

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func moveLinkIn(hostDev netlink.Link, containerNs ns.NetNS, containerIfName string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// With recent kernels we could do all changes in a single netlink call,
// but on failure the device is left in a partially modified state.
// Doing changes one by one allow us to (try to) rollback to the initial state.

// Create a temporary namespace to rename (and modify) the device in.
// We were previously using a temporary name, but rapid rename leads to
// race condition with udev and NetworkManager.

// Restore original up state in case of error
// This must be done in the hostNS as moving
// device between namespaces sets the link down

// lookup the device again (index might have changed)

// Move the host device into tempNS

// In a container in container scenario, hostNS is not the initial net namespace,
// but host / container naming is easier to follow.

// lookup the device in tempNS (index might have changed)

// detroying a non empty tempNS would move physical devices back to the initial net namespace,
// not the namespace of the "parent" process, and virtual devices would be destroyed,
// so we need to actively move the device back to hostNS on error

// Rename the device to the wanted name

// Restore the original device name in case of error

// Save host device name into the container device's alias property

// Remove the alias on error

// Move the device to the containerNS

// Lookup the device again on error, the index might have changed

// Move the interface back to tempNS on error

// Bring the device up
// This must be done in the containerNS

func moveLinkOut(containerNs ns.NetNS, containerIfName string) error {
	_ = "STUB: not implemented"
	// Create a temporary namespace to rename (and modify) the device in.
	// We were previously using a temporary name, but multiple rapid renames
	// leads to race condition with udev and NetworkManager.
	return nil
}

// Restore original up state in case of error
// This must be done in the containerNS as moving
// device between namespaces sets the link down

// lookup the device again (index might have changed)

// Lookup the device in the containerNS

// Verify we have the original name

// Move the device to the tempNS

// Lookup the device in tempNS (index might have changed)

// Move the device back to containerNS on error

// Rename container device to hostDevName

// Rename the device back to containerIfName on error

// Unset device's alias property

// Set back the device alias to hostDevName on error

// Finally move the device to the hostNS

// As we don't know the previous state, leave the link down

func hasDpdkDriver(pciaddr string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func printLink(dev netlink.Link, cniVersion string, containerNs ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

func linkFromPath(path string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// grab the first net device

func getLink(devname, hwaddr, kernelpath, pciaddr string, auxDev string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:    cmdAdd,
		Check:  cmdCheck,
		Del:    cmdDel,
		Status: cmdStatus,
		/* FIXME GC */
	}, version.All, bv.BuildString("host-device"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Parse previous result.

// Find interfaces for name we know, that of host-device inside container

// The namespace must be the same as what was configured

//
// Check prevResults for ips, routes and dns against values found in the container

// Check interface against values found in the container

//

func validateCniContainerInterface(intf current.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// TODO: Check if host device exists.
