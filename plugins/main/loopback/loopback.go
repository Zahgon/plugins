// Copyright 2016 CNI authors
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
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

func parseNetConf(bytes []byte) (*types.NetConf, error) { _ = "STUB: not implemented"; return nil, nil }

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// ignore config, this only works for loopback

// not tested

// not tested

// not tested

// sanity check that this is a loopback address

// not tested

// sanity check that this is a loopback address

// not tested

// If loopback has previous result which passes from previous CNI plugin,
// loopback should pass it transparently

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// ignore config, this only works for loopback

// not tested

// not tested

//  if NetNs is passed down by the Cloud Orchestration Engine, or if it called multiple times
// so don't return an error if the device is already removed.
// https://github.com/kubernetes/kubernetes/issues/43014#issuecomment-287164444

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.All, bv.BuildString("loopback"))
}

func cmdCheck(args *skel.CmdArgs) error {
	_ = "STUB: not implemented"
	// ignore config, this only works for loopback
	return nil
}
