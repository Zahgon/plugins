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

package utils

import (
	"github.com/coreos/go-iptables/iptables"
)

const statusChainExists = 1

// EnsureChain idempotently creates the iptables chain. It does not
// return an error if the chain already exists.
func EnsureChain(ipt *iptables.IPTables, table, chain string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteRule idempotently delete the iptables rule in the specified table/chain.
// It does not return an error if the referring chain doesn't exist
func DeleteRule(ipt *iptables.IPTables, table, chain string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// swallow here, the chain was already deleted

// swallow here, invalid command line parameter because the referring rule is missing

// DeleteChain idempotently deletes the specified table/chain.
// It does not return an errors if the chain does not exist
func DeleteChain(ipt *iptables.IPTables, table, chain string) error {
	_ = "STUB: not implemented"
	return nil
}

// swallow here, the chain was already deleted

// ClearChain idempotently clear the iptables rules in the specified table/chain.
// If the chain does not exist, a new one will be created
func ClearChain(ipt *iptables.IPTables, table, chain string) error {
	_ = "STUB: not implemented"
	return nil
}

// swallow here, the chain was already deleted

// InsertUnique will add a rule to a chain if it does not already exist.
// By default the rule is appended, unless prepend is true.
func InsertUnique(ipt *iptables.IPTables, table, chain string, prepend bool, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}
