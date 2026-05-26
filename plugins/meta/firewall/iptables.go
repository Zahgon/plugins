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
	"net"

	"github.com/coreos/go-iptables/iptables"

	current "github.com/containernetworking/cni/pkg/types/100"
)

func getPrivChainRules(ip string) [][]string { _ = "STUB: not implemented"; return nil }

func generateFilterRule(privChainName string) []string { _ = "STUB: not implemented"; return nil }

func generateAdminRule(adminChainName string) []string { _ = "STUB: not implemented"; return nil }

func cleanupRules(ipt *iptables.IPTables, privChainName string, rules [][]string) {
	_ = "STUB: not implemented"
	return
}

func ensureFirstChainRule(ipt *iptables.IPTables, chain string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ib *iptablesBackend) setupChains(ipt *iptables.IPTables) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure our private chains exist

// Ensure our filter rule exists in the forward chain

// Ensure our admin override chain rule exists in our private chain

func protoForIP(ip net.IPNet) iptables.Protocol {
	_ = "STUB: not implemented"
	return *new(iptables.Protocol)
}

func (ib *iptablesBackend) addRules(_ *FirewallNetConf, result *current.Result, ipt *iptables.IPTables, proto iptables.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up on any errors

func (ib *iptablesBackend) delRules(_ *FirewallNetConf, result *current.Result, ipt *iptables.IPTables, proto iptables.Protocol) {
	_ = "STUB: not implemented"
	return
}

func (ib *iptablesBackend) checkRules(_ *FirewallNetConf, result *current.Result, ipt *iptables.IPTables, proto iptables.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure our private chains exist

// Ensure our filter rule exists in the forward chain

// Ensure our admin override chain rule exists in our private chain

// ensure rules for this IP address exist

// Ensure our rule exists in our private chain

func findProtos(conf *FirewallNetConf) []iptables.Protocol { _ = "STUB: not implemented"; return nil }

// If PrevResult is given, scan all IP addresses to figure out
// which IP versions to use

type iptablesBackend struct {
	protos         map[iptables.Protocol]*iptables.IPTables
	privChainName  string
	adminChainName string
}

// iptablesBackend implements the FirewallBackend interface
var _ FirewallBackend = &iptablesBackend{}

func newIptablesBackend(conf *FirewallNetConf) (FirewallBackend, error) {
	_ = "STUB: not implemented"
	return *new(FirewallBackend), nil
}

func (ib *iptablesBackend) Add(conf *FirewallNetConf, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

func (ib *iptablesBackend) Del(conf *FirewallNetConf, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

func (ib *iptablesBackend) Check(conf *FirewallNetConf, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}
