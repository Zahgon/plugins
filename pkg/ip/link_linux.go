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
	"errors"
	"net"

	"github.com/vishvananda/netlink"

	"github.com/containernetworking/plugins/pkg/ns"
)

var ErrLinkNotFound = errors.New("link not found")

// makeVethPair is called from within the container's network namespace
func makeVethPair(name, peer string, mtu int, mac string, hostNS ns.NetNS) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// Re-fetch the container link to get its creation-time parameters, e.g. index and mac

// try and clean up the link if possible.

func peerExists(name string) bool { _ = "STUB: not implemented"; return false }

func makeVeth(name, vethPeerName string, mtu int, mac string, hostNS ns.NetNS) (string, netlink.Link, error) {
	_ = "STUB: not implemented"
	return "", *new(netlink.Link), nil
}

// should really never be hit

// RandomVethName returns string "veth" with random prefix (hashed from entropy)
func RandomVethName() (string, error) { _ = "STUB: not implemented"; return "", nil }

// NetworkManager (recent versions) will ignore veth devices that start with "veth"

func RenameLink(curName, newName string) error { _ = "STUB: not implemented"; return nil }

func ifaceFromNetlinkLink(l netlink.Link) net.Interface {
	_ = "STUB: not implemented"
	return *new(net.Interface)
}

// SetupVethWithName sets up a pair of virtual ethernet devices.
// Call SetupVethWithName from inside the container netns.  It will create both veth
// devices and move the host-side veth into the provided hostNS namespace.
// hostVethName: If hostVethName is not specified, the host-side veth name will use a random string.
// On success, SetupVethWithName returns (hostVeth, containerVeth, nil)
func SetupVethWithName(contVethName, hostVethName string, mtu int, contVethMac string, hostNS ns.NetNS) (net.Interface, net.Interface, error) {
	_ = "STUB: not implemented"
	return *new(net.Interface), *new(net.Interface), nil
}

// we want to own the routes for this interface

// SetupVeth sets up a pair of virtual ethernet devices.
// Call SetupVeth from inside the container netns.  It will create both veth
// devices and move the host-side veth into the provided hostNS namespace.
// On success, SetupVeth returns (hostVeth, containerVeth, nil)
func SetupVeth(contVethName string, mtu int, contVethMac string, hostNS ns.NetNS) (net.Interface, net.Interface, error) {
	_ = "STUB: not implemented"
	return *new(net.Interface), *new(net.Interface), nil
}

// DelLinkByName removes an interface link.
func DelLinkByName(ifName string) error { _ = "STUB: not implemented"; return nil }

// DelLinkByNameAddr remove an interface and returns its addresses
func DelLinkByNameAddr(ifName string) ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVethPeerIfindex returns the veth link object, the peer ifindex of the
// veth, or an error. This peer ifindex will only be valid in the peer's
// network namespace.
func GetVethPeerIfindex(ifName string) (netlink.Link, int, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), 0, nil
}

// veth supports IFLA_LINK (what vishvananda/netlink calls ParentIndex)
// on 4.1 and higher kernels

// Fall back to ethtool for 4.0 and earlier kernels
