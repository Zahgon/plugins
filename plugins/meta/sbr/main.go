// Copyright 2017 CNI authors
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

// This is the Source Based Routing plugin that sets up source based routing.
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

const firstTableID = 100

// PluginConf is the configuration document passed in.
type PluginConf struct {
	types.NetConf

	// This is the previous result, when called in the context of a chained
	// plugin. Because this plugin supports multiple versions, we'll have to
	// parse this in two passes. If your plugin is not chained, this can be
	// removed (though you may wish to error if a non-chainable plugin is
	// chained).
	RawPrevResult *map[string]interface{} `json:"prevResult"`
	PrevResult    *current.Result         `json:"-"`

	// Add plugin-specific flags here
	Table *int `json:"table,omitempty"`
	// Gateways allows specifying static/hardcoded gateway IP addresses
	// If set, these will be used instead of the gateway from prevResult
	// Supports dual-stack: provide one IPv4 and/or one IPv6 gateway
	// Note: Currently applies the same gateway to all IPs of the same family.
	// Per-subnet gateway mapping is not yet supported.
	Gateways []string `json:"gateways,omitempty"`
	// AddSourceHints adds source IP hints to subnet routes in the main table,
	// enabling destination-based routing when no explicit source IP is specified.
	// This allows both source-based and destination-based routing to work together.
	AddSourceHints bool `json:"addSourceHints,omitempty"`
}

// Wrapper that does a lock before and unlock after operations to serialise
// this plugin.
func withLockAndNetNS(nspath string, toRun func(_ ns.NetNS) error) error {
	_ = "STUB: not implemented"
	// We lock on the network namespace to ensure that no other instance
	// clashes with this one.
	return nil
}

// Cleaner to unlock even though about to exit

// parseConfig parses the supplied configuration (and prevResult) from stdin.
func parseConfig(stdin []byte) (*PluginConf, error) { _ = "STUB: not implemented"; return nil, nil }

// Parse previous result.

// End previous result parsing

// getIPCfgs finds the IPs on the supplied interface, returning as IPConfig structures
func getIPCfgs(iface string, prevResult *current.Result) ([]*current.IPConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// No IP addresses; that makes no sense. Pack it in.
}

// We do a single interface name, stored in args.IfName

// ips contains the IPConfig structures that were passed, filtered somewhat

// IPs have an interface that is an index into the interfaces array.
// We assume a match if this index is missing.

// Skip all IPs we know belong to an interface with the wrong name.

// cmdAdd is called for ADD requests
func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Get the list of relevant IPs.

// Do the actual work.

// Pass through the result for the next plugin

// getNextTableID picks the first free table id from a giveen candidate id
func getNextTableID(rules []netlink.Rule, routes []netlink.Route, candidateID int) int {
	_ = "STUB: not implemented"
	return 0
}

// doRoutes does all the work to set up routes and rules during an add.
func doRoutes(ipCfgs []*current.IPConfig, iface string, staticGateways []string, addSourceHints bool) error {
	_ = "STUB: not implemented"
	// Get a list of rules and routes ready.
	return nil
}

// Pick a table ID to use. We pick the first table ID from firstTableID
// on that has no existing rules mapping to it and no existing routes in
// it.

// Get all routes for the interface in the default routing table

// Parse static gateways if provided (supports dual-stack: one IPv4, one IPv6)

// We already have an IPv4 static gateway

// We already have an IPv6 static gateway

// Loop through setting up source based rules and default routes.

// Source must be restricted to a single IP, not a full subnet

// Determine which gateway to use: static gateway takes precedence
// Match gateway by IP family (IPv4 vs IPv6)

// Add a default route, since this may have been removed by previous
// plugin.

// Copy the previously added routes for the interface to the correct
// table; all the routes have been added to the interface anyway but
// in the wrong table, so instead of removing them we just move them
// to the table we want them in.

// (r.Src == nil && r.Gw == nil) is inferred as a generic route

// Reset the route flags since if it is dynamically created,
// adding it to the new table will fail with "invalid argument"

// We use route replace in case the route already exists, which
// is possible for the default gateway we added above.

// Use a different table for each ipCfg

// Handle routes in the default routing table

// Keep or re-add subnet routes in main table with source IP hints
// for destination-based routing when no explicit source IP is specified

// Find the subnet route for this IP

// Add route to main table with source IP hint

// Use RouteReplace to update if it exists

// Don't fail completely, just warn

// Delete non-subnet routes from main table (gateway routes, etc.)

// Skip subnet routes (scope link), only delete other routes

// Not deleting them while copying to accommodate for multiple ipCfgs from
// the same subnet. Else, (error for network is unreachable while adding gateway)

func doRoutesWithTable(ipCfgs []*current.IPConfig, table int) error {
	_ = "STUB: not implemented"
	return nil
}

// Source must be restricted to a single IP, not a full subnet

// cmdDel is called for DELETE requests
func cmdDel(args *skel.CmdArgs) error {
	_ = "STUB: not implemented"
	// We care a bit about config because it sets log level.
	return nil
}

// Tidy up the rules for the deleted interface
func tidyRules(iface string, table *int) error {
	_ = "STUB: not implemented"
	// We keep on going on rule deletion error, but return the last failure.
	return nil
}

// If interface is not found by any reason it's safe to ignore an error. Also, we don't need to raise an error
// during cmdDel call according to CNI spec:
// https://github.com/containernetworking/cni/blob/main/SPEC.md#del-remove-container-from-network-or-un-apply-modifications

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.All, bv.BuildString("sbr"))
}

func cmdCheck(_ *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }
