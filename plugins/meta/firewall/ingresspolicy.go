// Copyright 2022 CNI authors
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

// This is a sample chained plugin that supports multiple CNI versions. It
// parses prevResult according to the cniVersion
package main

import (
	"github.com/coreos/go-iptables/iptables"

	types100 "github.com/containernetworking/cni/pkg/types/100"
)

func setupIngressPolicy(conf *FirewallNetConf, prevResult *types100.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// NOP

func setupIngressPolicyBridgeIsolation(conf *FirewallNetConf, prevResult *types100.Result, isolated bool) error {
	_ = "STUB: not implemented"
	return nil
}

func teardownIngressPolicy(conf *FirewallNetConf) error { _ = "STUB: not implemented"; return nil }

// NOP

// NOP
//
// We can't be sure whether conf.bridgeName is still in use by other containers.
// So we do not remove the iptable rules that are created per bridge.

const (
	filterTableName  = "filter"  // built-in
	forwardChainName = "FORWARD" // built-in
)

// setupIsolationChains executes the following iptables commands for isolating networks:
// ```
// iptables -N CNI-ISOLATION-STAGE-1
// iptables -N CNI-ISOLATION-STAGE-2
// # NOTE: "-j CNI-ISOLATION-STAGE-1" needs to be before "CNI-FORWARD" chain. So we use -I here.
// iptables -I FORWARD -j CNI-ISOLATION-STAGE-1
// iptables -A CNI-ISOLATION-STAGE-1 -i ${bridgeName} ! -o ${bridgeName} -j CNI-ISOLATION-STAGE-2
// [isolated=true] iptables -A CNI-ISOLATION-STAGE-1 -i ${bridgeName} -o ${bridgeName} -j DROP
// iptables -A CNI-ISOLATION-STAGE-1 -j RETURN
// iptables -A CNI-ISOLATION-STAGE-2 -o ${bridgeName} -j DROP
// iptables -A CNI-ISOLATION-STAGE-2 -j RETURN
// ```
func setupIsolationChains(ipt *iptables.IPTables, bridgeName string, isolated bool) error {
	_ = "STUB: not implemented"

	// Future version may support custom chain names
	return nil
}

// Commands:
// ```
// iptables -N CNI-ISOLATION-STAGE-1
// iptables -N CNI-ISOLATION-STAGE-2
// ```

// Commands:
// ```
// iptables -I FORWARD -j CNI-ISOLATION-STAGE-1
// ```

//  NOTE: "-j CNI-ISOLATION-STAGE-1" needs to be before "CNI-FORWARD" created by CNI firewall plugin.
// So we specify prepend = true .

// Commands:
// ```
// iptables -A CNI-ISOLATION-STAGE-1 -i ${bridgeName} ! -o ${bridgeName} -j CNI-ISOLATION-STAGE-2
// [isolate=true] iptables -A CNI-ISOLATION-STAGE-1 -i ${bridgeName} -o ${bridgeName} -j DROP
// iptables -A CNI-ISOLATION-STAGE-1 -j RETURN
// ```

// prepend = true because this needs to be before "-j RETURN"

// Commands:
// ```
// iptables -A CNI-ISOLATION-STAGE-2 -o ${bridgeName} -j DROP
// iptables -A CNI-ISOLATION-STAGE-2 -j RETURN
// ```

// prepend = true because this needs to be before "-j RETURN"

func isolationStage1BridgeRule(bridgeName, stage2Chain string) []string {
	_ = "STUB: not implemented"
	return nil
}

func isolationStage1BridgeDropRule(bridgeName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func isolationStage2BridgeRule(bridgeName string) []string { _ = "STUB: not implemented"; return nil }

func withPolicyComment(policy string) func([]string) []string {
	_ = "STUB: not implemented"
	return nil
}

func withDefaultComment(rule []string) []string { _ = "STUB: not implemented"; return nil }

func withComment(rule []string, comment string) []string { _ = "STUB: not implemented"; return nil }
