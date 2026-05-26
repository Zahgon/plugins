// Copyright 2014 CNI authors
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
	"net"
	"runtime"

	"github.com/vishvananda/netlink"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/containernetworking/plugins/pkg/ns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

// For testcases to force an error after IPAM has been performed
var debugPostIPAMError error

const defaultBrName = "cni0"

type NetConf struct {
	types.NetConf
	BrName                    string       `json:"bridge"`
	IsGW                      bool         `json:"isGateway"`
	IsDefaultGW               bool         `json:"isDefaultGateway"`
	ForceAddress              bool         `json:"forceAddress"`
	IPMasq                    bool         `json:"ipMasq"`
	IPMasqBackend             *string      `json:"ipMasqBackend,omitempty"`
	MTU                       int          `json:"mtu"`
	HairpinMode               bool         `json:"hairpinMode"`
	PromiscMode               bool         `json:"promiscMode"`
	Vlan                      int          `json:"vlan"`
	VlanTrunk                 []*VlanTrunk `json:"vlanTrunk,omitempty"`
	PreserveDefaultVlan       bool         `json:"preserveDefaultVlan"`
	MacSpoofChk               bool         `json:"macspoofchk,omitempty"`
	EnableDad                 bool         `json:"enabledad,omitempty"`
	DisableContainerInterface bool         `json:"disableContainerInterface,omitempty"`
	PortIsolation             bool         `json:"portIsolation,omitempty"`

	Args struct {
		Cni BridgeArgs `json:"cni,omitempty"`
	} `json:"args,omitempty"`
	RuntimeConfig struct {
		Mac string `json:"mac,omitempty"`
	} `json:"runtimeConfig,omitempty"`

	mac   string
	vlans []int
}

type VlanTrunk struct {
	MinID *int `json:"minID,omitempty"`
	MaxID *int `json:"maxID,omitempty"`
	ID    *int `json:"id,omitempty"`
}

type BridgeArgs struct {
	Mac string `json:"mac,omitempty"`
}

// MacEnvArgs represents CNI_ARGS
type MacEnvArgs struct {
	types.CommonArgs
	MAC types.UnmarshallableString `json:"mac,omitempty"`
}

type gwInfo struct {
	gws               []net.IPNet
	family            int
	defaultRouteFound bool
}

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

func loadNetConf(bytes []byte, envArgs string) (*NetConf, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// fail to parsing

// This method is copied from https://github.com/k8snetworkplumbingwg/ovs-cni/blob/v0.27.2/pkg/plugin/plugin.go
func collectVlanTrunk(vlanTrunk []*VlanTrunk) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// single vid

// calcGateways processes the results from the IPAM plugin and does the
// following for each IP family:
//   - Calculates and compiles a list of gateway addresses
//   - Adds a default route if needed
func calcGateways(result *current.Result, n *NetConf) (*gwInfo, *gwInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Determine if this config is IPv4 or IPv6

// All IPs currently refer to the container interface

// If not provided, calculate the gateway address corresponding
// to the selected IP address

// Add a default route for this family using the current
// gateway address if necessary.

// Append this gateway address to the list of gateways

func ensureAddr(br netlink.Link, family int, ipn *net.IPNet, forceAddress bool) error {
	_ = "STUB: not implemented"
	return nil
}

// string comp is actually easiest for doing IPNet comps

// Multiple IPv6 addresses are allowed on the bridge if the
// corresponding subnets do not overlap. For IPv4 or for
// overlapping IPv6 subnets, reconfigure the IP address if
// forceAddress is true, otherwise throw an error.

// Set the bridge's MAC to itself. Otherwise, the bridge will take the
// lowest-numbered mac on the bridge, and will change as ifs churn

func deleteAddr(br netlink.Link, ipn *net.IPNet) error { _ = "STUB: not implemented"; return nil }

func bridgeByName(name string) (*netlink.Bridge, error) { _ = "STUB: not implemented"; return nil, nil }

func ensureBridge(brName string, mtu int, promiscMode, vlanFiltering bool) (*netlink.Bridge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Re-fetch link to read all attributes and if it already existed,
// ensure it's really a bridge with similar configuration

// we want to own the routes for this interface

func ensureVlanInterface(br *netlink.Bridge, vlanID int, preserveDefaultVlan bool) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

func setupVeth(
	netns ns.NetNS,
	br *netlink.Bridge,
	ifName string,
	mtu int,
	hairpinMode bool,
	vlanID int,
	vlans []int,
	preserveDefaultVlan bool,
	mac string,
	portIsolation bool,
) (*current.Interface, *current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// create the veth pair in the container and move host end into host netns

// need to lookup hostVeth again as its index has changed during ns move

// connect host veth end to the bridge

// set hairpin mode

// set isolation mode

// Backwards compatibility with users that did not specify a vlanID

// If no vlan is specified, we set the native vlan on the trunk equal to 1

func removeDefaultVlan(hostVeth netlink.Link) error { _ = "STUB: not implemented"; return nil }

func calcGatewayIP(ipn *net.IPNet) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func setupBridge(n *NetConf) (*netlink.Bridge, *current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// create bridge if necessary

func enableIPForward(family int) error { _ = "STUB: not implemented"; return nil }

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Assume L2 interface only

// run the IPAM plugin and get back the config to apply

// release IP in case of failure

// Convert whatever the IPAM result was into the current Result type

// Gather gateway information for each IP family

// Configure the container hardware address and IP address(es)

// Add the IP to the interface

// Set the IP address(es) on the bridge and enable forwarding

// If layer 2 we still need to set the container veth to up

// check bridge port state

// In certain circumstances, the host-side of the veth may change addrs

// Refetch the bridge since its MAC address may change when the first
// veth is added or after its IP address is set

// Return an error requested by testcases, if any

// Use incoming DNS settings if provided, otherwise use the
// settings that were already configured by the IPAM plugin

func dnsConfSet(dnsConf types.DNS) bool { _ = "STUB: not implemented"; return false }

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// There is a netns so try to clean up. Delete can be called multiple times
// so don't return an error if the device is already removed.
// If the device isn't there then don't try to clean up IP masq either.

//  if NetNs is passed down by the Cloud Orchestration Engine, or if it called multiple times
// so don't return an error if the device is already removed.
// https://github.com/kubernetes/kubernetes/issues/43014#issuecomment-287164444

// call ipam.ExecDel after clean up device in netns

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:    cmdAdd,
		Check:  cmdCheck,
		Del:    cmdDel,
		Status: cmdStatus,
		/* FIXME GC */
	}, version.All, bv.BuildString("bridge"))
}

type cniBridgeIf struct {
	Name        string
	ifIndex     int
	peerIndex   int
	masterIndex int
	found       bool
}

func validateInterface(intf current.Interface, expectInSb bool) (cniBridgeIf, netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(cniBridgeIf), *new(netlink.Link), nil
}

func validateCniBrInterface(intf current.Interface, n *NetConf) (cniBridgeIf, error) {
	_ = "STUB: not implemented"
	return *new(cniBridgeIf), nil
}

func validateCniVethInterface(intf *current.Interface, brIf cniBridgeIf, contIf cniBridgeIf) (cniBridgeIf, error) {
	_ = "STUB: not implemented"
	return *new(cniBridgeIf), nil
}

// just skip it, it's not what CNI created

func validateCniContainerInterface(intf current.Interface) (cniBridgeIf, error) {
	_ = "STUB: not implemented"
	return *new(cniBridgeIf), nil
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// run the IPAM plugin and get back the config to apply

// Parse previous result.

// Find interfaces for names whe know, CNI Bridge and container

// The namespace must be the same as what was configured

// Check interface against values found in the container

// Now look for veth that is peer with container interface.
// Anything else wasn't created by CNI, skip it

// Skip this result if name is the same as cni bridge
// It's either the cni bridge we dealt with above, or something with the
// same name in a different namespace.  We just skip since it's not ours

// same here for container name

// veth with container interface as peer and bridge as master found

// Check prevResults for ips, routes and dns against values found in the container

func uniqueID(containerID, cniIface string) string { _ = "STUB: not implemented"; return "" }

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }
