// Copyright 2020 CNI authors
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
)

// findVRF finds a VRF link with the provided name.
func findVRF(name string) (*netlink.Vrf, error) { _ = "STUB: not implemented"; return nil, nil }

// createVRF creates a new VRF and sets it up.
func createVRF(name string, tableID uint32) (*netlink.Vrf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assignedInterfaces returns the list of interfaces associated to the given vrf.
func assignedInterfaces(vrf *netlink.Vrf) ([]netlink.Link, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addInterface adds the given interface to the VRF
func addInterface(vrf *netlink.Vrf, intf string) error { _ = "STUB: not implemented"; return nil }

// Global IPV6 addresses are not maintained unless
// sysctl -w net.ipv6.conf.all.keep_addr_on_down=1 is called
// so we save it, and restore it back.

// Save all routes that are not local and connected, before setting master,
// because otherwise those routes will be deleted after interface is moved.

// Exclude local and connected routes

// Filter based on link index and scope

// Used to identify which global IPV6 addresses are missing

// Since keeping the ipv6 address depends on net.ipv6.conf.all.keep_addr_on_down ,
// we check if the new interface does not have them and in case we restore them.

// Not found, re-adding it

// Waits for global IPV6 addresses to be added by the kernel.

// Exponential backoff - 10ms, 20m, 40ms, 80ms, 160ms, 320ms, 640ms, 1280ms
// Approx 2,5 seconds total

// Apply all saved routes for the interface that was moved to the VRF

// Modify original table to vrf one,

// equivalent of 'ip route replace <address> table <int>'.

func findFreeRoutingTableID(links []netlink.Link) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func resetMaster(interfaceName string) error { _ = "STUB: not implemented"; return nil }

// getGlobalAddresses returns the global addresses of the given interface
func getGlobalAddresses(link netlink.Link, family int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
