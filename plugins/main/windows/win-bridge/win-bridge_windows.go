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

package main

import (
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/cni/pkg/version"

	"github.com/containernetworking/plugins/pkg/hns"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

type NetConf struct {
	hns.NetConf

	IPMasqNetwork string `json:"ipMasqNetwork,omitempty"`
}

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

func loadNetConf(bytes []byte) (*NetConf, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func processEndpointArgs(args *skel.CmdArgs, n *NetConf) (*hns.EndpointInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// it's not necessary to have have an IPAM in windows as HNS can provide IP/GW

// convert whatever the IPAM result was into the current result

// Calculate gateway for bridge network (needs to be x.2)

// configure sNAT exception

// add port mapping if any present

func cmdHnsAdd(args *skel.CmdArgs, n *NetConf) (*current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cmdHcnAdd(args *skel.CmdArgs, n *NetConf) (*current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func cmdCheck(_ *skel.CmdArgs) error {
	_ = "STUB: not implemented"
	// TODO: implement
	return nil
}

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:    cmdAdd,
		Check:  cmdCheck,
		Del:    cmdDel,
		Status: cmdStatus,
		/* FIXME GC */
	}, version.All, bv.BuildString("win-bridge"))
}

func cmdStatus(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }
