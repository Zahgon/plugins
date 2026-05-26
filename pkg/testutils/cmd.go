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

package testutils

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
)

func envCleanup() { _ = "STUB: not implemented"; return }

func CmdAdd(cniNetns, cniContainerID, cniIfname string, conf []byte, f func() error) (types.Result, []byte, error) {
	_ = "STUB: not implemented"
	return *new(types.Result), nil, nil
}

// Redirect stdout to capture plugin result

// Return errors after restoring stdout so Ginkgo will correctly
// emit verbose error information on stdout

// Plugin must return result in same version as specified in netconf

func CmdAddWithArgs(args *skel.CmdArgs, f func() error) (types.Result, []byte, error) {
	_ = "STUB: not implemented"
	return *new(types.Result), nil, nil
}

func CmdCheck(cniNetns, cniContainerID, cniIfname string, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func CmdCheckWithArgs(args *skel.CmdArgs, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func CmdDel(cniNetns, cniContainerID, cniIfname string, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func CmdDelWithArgs(args *skel.CmdArgs, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func CmdStatus(f func() error) error { _ = "STUB: not implemented"; return nil }
