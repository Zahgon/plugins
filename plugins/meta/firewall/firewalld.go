// Copyright 2018 CNI authors
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
	"github.com/godbus/dbus/v5"

	current "github.com/containernetworking/cni/pkg/types/100"
)

const (
	dbusName               = "org.freedesktop.DBus"
	dbusPath               = "/org/freedesktop/DBus"
	dbusGetNameOwnerMethod = "GetNameOwner"

	firewalldName               = "org.fedoraproject.FirewallD1"
	firewalldPath               = "/org/fedoraproject/FirewallD1"
	firewalldZoneInterface      = "org.fedoraproject.FirewallD1.zone"
	firewalldAddSourceMethod    = "addSource"
	firewalldRemoveSourceMethod = "removeSource"
	firewalldQuerySourceMethod  = "querySource"

	errZoneAlreadySet = "ZONE_ALREADY_SET"
)

// Only used for testcases to override the D-Bus connection
var testConn *dbus.Conn

type fwdBackend struct {
	conn *dbus.Conn
}

// fwdBackend implements the FirewallBackend interface
var _ FirewallBackend = &fwdBackend{}

func getConn() (*dbus.Conn, error) { _ = "STUB: not implemented"; return nil, nil }

// isFirewalldRunning checks whether firewalld is running.
func isFirewalldRunning() bool { _ = "STUB: not implemented"; return false }

func newFirewalldBackend() (FirewallBackend, error) {
	_ = "STUB: not implemented"
	return *new(FirewallBackend), nil
}

func (fb *fwdBackend) Add(conf *FirewallNetConf, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// Add a firewalld rule which assigns the given source IP to the given zone

func (fb *fwdBackend) Del(conf *FirewallNetConf, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove firewalld rules which assigned the given source IP to the given zone

func (fb *fwdBackend) Check(conf *FirewallNetConf, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// Check for a firewalld rule for the given source IP to the given zone
