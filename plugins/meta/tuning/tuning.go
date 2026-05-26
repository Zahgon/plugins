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

// This is a "meta-plugin". It reads in its own netconf, it does not create
// any network interface but just changes the network sysctl.

package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

const (
	defaultDataDir       = "/run/cni/tuning"
	defaultAllowlistDir  = "/etc/cni/tuning/"
	defaultAllowlistFile = "allowlist.conf"
)

// TuningConf represents the network tuning configuration.
type TuningConf struct {
	types.NetConf
	DataDir  string            `json:"dataDir,omitempty"`
	SysCtl   map[string]string `json:"sysctl"`
	Mac      string            `json:"mac,omitempty"`
	Promisc  bool              `json:"promisc,omitempty"`
	Mtu      int               `json:"mtu,omitempty"`
	TxQLen   *int              `json:"txQLen,omitempty"`
	Allmulti *bool             `json:"allmulti,omitempty"`

	RuntimeConfig struct {
		Mac string `json:"mac,omitempty"`
	} `json:"runtimeConfig,omitempty"`
	Args *struct {
		A *IPAMArgs `json:"cni"`
	} `json:"args"`
}

type IPAMArgs struct {
	SysCtl   *map[string]string `json:"sysctl"`
	Mac      *string            `json:"mac,omitempty"`
	Promisc  *bool              `json:"promisc,omitempty"`
	Mtu      *int               `json:"mtu,omitempty"`
	Allmulti *bool              `json:"allmulti,omitempty"`
	TxQLen   *int               `json:"txQLen,omitempty"`
}

// configToRestore will contain interface attributes that should be restored on cmdDel
type configToRestore struct {
	Mac      string `json:"mac,omitempty"`
	Promisc  *bool  `json:"promisc,omitempty"`
	Mtu      int    `json:"mtu,omitempty"`
	Allmulti *bool  `json:"allmulti,omitempty"`
	TxQLen   *int   `json:"txQLen,omitempty"`
}

// MacEnvArgs represents CNI_ARG
type MacEnvArgs struct {
	types.CommonArgs
	MAC types.UnmarshallableString `json:"mac,omitempty"`
}

func parseConf(data []byte, envArgs string) (*TuningConf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse custom Mac from both env args

// Parse custom Mac from RuntimeConfig

// Get args

func changeMacAddr(ifName string, newMacAddr string) error { _ = "STUB: not implemented"; return nil }

func updateResultsMacAddr(config *TuningConf, ifName string, newMacAddr string) {
	_ = "STUB: not implemented"
	// Parse previous result.
	return
}

func changePromisc(ifName string, val bool) error { _ = "STUB: not implemented"; return nil }

func changeMtu(ifName string, mtu int) error { _ = "STUB: not implemented"; return nil }

func changeAllmulti(ifName string, val bool) error { _ = "STUB: not implemented"; return nil }

func changeTxQLen(ifName string, txQLen int) error { _ = "STUB: not implemented"; return nil }

func createBackup(ifName, containerID, backupPath string, tuningConf *TuningConf) error {
	_ = "STUB: not implemented"
	return nil
}

func restoreBackup(ifName, containerID, backupPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// No backup file - nothing to revert

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Parse previous result.

// The directory /proc/sys/net is per network namespace. Enter in the
// network namespace before writing on it.

// cmdDel will restore NIC attributes to the original ones when called
func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// MAC address, MTU, promiscuous and all-multicast mode settings will be restored

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
		/* FIXME GC */
		/* FIXME Status */
	}, version.All, bv.BuildString("tuning"))
}

func cmdCheck(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Parse previous result.

// Check each configured value vs what's currently in the container

// Validate the sysctls in the tuning config are on the sysctl allowlist file.
// Note that if the allowlist file is missing no validation takes place.
func validateSysctlConf(tuningConf *TuningConf) error { _ = "STUB: not implemented"; return nil }

// Validate the allowList contains the given sysctl
func contains(sysctl string, allowList []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Read the systctl allowlist from file. Return info if the file is present and the read allowList if it is
func readAllowlist() (bool, []string, error) { _ = "STUB: not implemented"; return false, nil, nil }

type sysctlKey string

type sysctlCheck struct {
	SysCtl map[sysctlKey]string `json:"sysctl"`
}

var sysctlDuplicatesMap = map[sysctlKey]interface{}{}

func (d *sysctlKey) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

func validateSysctlConflictingKeys(data []byte) error { _ = "STUB: not implemented"; return nil }

func validateArgs(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func getSysctlFilename(key, ifName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If the key contains `IFNAME` - substitute it with args.IfName
// to allow setting sysctls on a particular interface, on which
// other operations (like mac/mtu setting) are performed

// Refuse to modify sysctl parameters that don't belong
// to the network subsystem.
