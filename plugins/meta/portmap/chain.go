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
	"github.com/coreos/go-iptables/iptables"
)

type chain struct {
	table       string
	name        string
	entryChains []string // the chains to add the entry rule

	entryRules [][]string // the rules that "point" to this chain
	rules      [][]string // the rules this chain contains

	prependEntry bool // whether or not the entry rules should be prepended
}

// setup idempotently creates the chain. It will not error if the chain exists.
func (c *chain) setup(ipt *iptables.IPTables) error { _ = "STUB: not implemented"; return nil }

// Add the rules to the chain

// Add the entry rules to the entry chains

// teardown idempotently deletes a chain. It will not error if the chain doesn't exist.
// It will first delete all references to this chain in the entryChains.
func (c *chain) teardown(ipt *iptables.IPTables) error {
	_ = "STUB: not implemented"
	// nothing to do if the custom chain doesn't exist to begin with
	return nil
}

// delete references created by setup()

// if chain deletion succeeds now, all references are gone

// find references the hard way

// Swallow error here - probably the chain doesn't exist.
// If we miss something the deletion will fail

// List results always include an -A CHAINNAME

// check the chain.
func (c *chain) check(ipt *iptables.IPTables) error { _ = "STUB: not implemented"; return nil }

func checkRule(ipt *iptables.IPTables, table, chain string, rule []string) bool {
	_ = "STUB: not implemented"
	return false
}
