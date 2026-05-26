// Copyright 2021 CNI authors
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

package link

import (
	"github.com/networkplumbing/go-nft/nft"
	"github.com/networkplumbing/go-nft/nft/schema"
)

const (
	natTableName            = "nat"
	preRoutingBaseChainName = "PREROUTING"
)

type NftConfigurer interface {
	Apply(*nft.Config) (*nft.Config, error)
	Read(filterCommands ...string) (*nft.Config, error)
}

type SpoofChecker struct {
	iface      string
	macAddress string
	refID      string
	configurer NftConfigurer
	rulestore  *nft.Config
}

type defaultNftConfigurer struct{}

func (dnc defaultNftConfigurer) Apply(cfg *nft.Config) (*nft.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dnc defaultNftConfigurer) Read(filterCommands ...string) (*nft.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSpoofChecker(iface, macAddress, refID string) *SpoofChecker {
	_ = "STUB: not implemented"
	return nil
}

func NewSpoofCheckerWithConfigurer(iface, macAddress, refID string, configurer NftConfigurer) *SpoofChecker {
	_ = "STUB: not implemented"
	return nil
}

// Setup applies nftables configuration to restrict traffic
// from the provided interface. Only traffic with the mentioned mac address
// is allowed to pass, all others are blocked.
// The configuration follows the format libvirt and ebtables implemented, allowing
// extensions to the rules in the future.
// refID is used to label the rules with a unique comment, identifying the rule-set.
//
// In order to take advantage of the nftables configuration change atomicity, the
// following steps are taken to apply the configuration:
// - Declare the table and chains (they will be created in case not present).
// - Apply the rules, while first flushing the iface/mac specific regular chain rules.
// Two transactions are used because the flush succeeds only if the table/chain it targets
// exists. This avoids the need to query the existing state and acting upon it (a raceful pattern).
// Although two transactions are taken place, only the 2nd one where the rules
// are added has a real impact on the system.
func (sc *SpoofChecker) Setup() error { _ = "STUB: not implemented"; return nil }

func (sc *SpoofChecker) findPreroutingRule(ruleToFind *schema.Rule) ([]*schema.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Teardown removes the interface and mac-address specific chains and their rules.
// The table and base-chain are expected to survive while the base-chain rule that matches the
// interface is removed.
func (sc *SpoofChecker) Teardown() error { _ = "STUB: not implemented"; return nil }

// It is safer to exclude the statement matching, avoiding cases where a current statement includes
// additional default entries (e.g. counters).

// Drop the cache, it should contain deleted rule(s) now

func (sc *SpoofChecker) matchIfaceJumpToChainRule(chain, toChain string) *schema.Rule {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SpoofChecker) jumpToChainRule(chain, toChain string) *schema.Rule {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SpoofChecker) matchMacRule(chain string) *schema.Rule {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SpoofChecker) dropRule(chain string) *schema.Rule { _ = "STUB: not implemented"; return nil }

func (sc *SpoofChecker) baseChain() *schema.Chain { _ = "STUB: not implemented"; return nil }

func (sc *SpoofChecker) ifaceChain() *schema.Chain { _ = "STUB: not implemented"; return nil }

func (sc *SpoofChecker) macChain(ifaceChainName string) *schema.Chain {
	_ = "STUB: not implemented"
	return nil
}

func ruleComment(id string) string { _ = "STUB: not implemented"; return "" }

func listChainBridgeNatPrerouting() []string { _ = "STUB: not implemented"; return nil }
