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
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

// VRFNetConf represents the vrf configuration.
type VRFNetConf struct {
	types.NetConf

	// VRFName is the name of the vrf to add the interface to.
	VRFName string `json:"vrfname"`
	// Table is the optional name of the routing table set for the vrf
	Table uint32 `json:"table"`
}

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.VersionsStartingFrom("0.3.1"), bv.BuildString("vrf"))
}

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// If the user set a tableid and the vrf is already in the namespace
// we check if the tableid is the same one already assigned to the vrf.

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Meaning, we are deleting the last interface assigned to the VRF

//  if NetNs is passed down by the Cloud Orchestration Engine, or if it called multiple times
// so don't return an error if the device is already removed.
// https://github.com/kubernetes/kubernetes/issues/43014#issuecomment-287164444

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Ensure we have previous result.

func parseConf(data []byte) (*VRFNetConf, *current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// return early if there was no previous result, which is allowed for DEL calls

// Parse previous result.
