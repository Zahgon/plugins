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
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.All, bv.BuildString("host-local"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Look to see if there is at least one IP address allocated to the container
// in the data dir, irrespective of what that address actually is

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Keep the allocators we used, so we can release all IPs if an error
// occurs after we start allocating

// Store all requested IPs in a map, so we can easily remove ones we use
// and error if some remain
// net.IP cannot be a key

// Check to see if there are any custom IPs requested in this range.

// Deallocate all already allocated IPs

// If an IP was requested that wasn't fulfilled, fail

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Loop through all ranges, releasing all IPs, even if an error occurs
