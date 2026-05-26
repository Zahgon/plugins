// Copyright 2015 CNI authors
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
	"errors"
	"net"
	"sync"
	"time"

	"github.com/containernetworking/cni/pkg/skel"
	current "github.com/containernetworking/cni/pkg/types/100"
)

var errNoMoreTries = errors.New("no more tries")

type DHCP struct {
	mux                 sync.Mutex
	leases              map[string]*DHCPLease
	hostNetnsPrefix     string
	clientTimeout       time.Duration
	clientResendMax     time.Duration
	clientResendTimeout time.Duration
	broadcast           bool
}

func newDHCP(clientTimeout, clientResendMax time.Duration, resendTimeout time.Duration) *DHCP {
	_ = "STUB: not implemented"
	return nil
}

// TODO: current client ID is too long. At least the container ID should not be used directly.
// A separate issue is necessary to ensure no breaking change is affecting other users.
func generateClientID(containerID string, netName string, ifName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defined in RFC 2132, length size can not be larger than 1 octet. So we truncate 254 to make everyone happy.

// Allocate acquires an IP from a DHCP server for a specified container.
// The acquired lease will be maintained until Release() is called.
func (d *DHCP) Allocate(args *skel.CmdArgs, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// If we already have an active lease for this clientID, do not create
// another one

// Release stops maintenance of the lease acquired in Allocate()
// and sends a release msg to the DHCP server.
func (d *DHCP) Release(args *skel.CmdArgs, _ *struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DHCP) getLease(clientID string) *DHCPLease { _ = "STUB: not implemented"; return nil }

// TODO(eyakubovich): hash it to avoid collisions

func (d *DHCP) setLease(clientID string, l *DHCPLease) { _ = "STUB: not implemented"; return }

// TODO(eyakubovich): hash it to avoid collisions

// func (d *DHCP) clearLease(contID, netName, ifName string) {
func (d *DHCP) clearLease(clientID string) { _ = "STUB: not implemented"; return }

// TODO(eyakubovich): hash it to avoid collisions

func getListener(socketPath string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func runDaemon(
	pidfilePath, hostPrefix, socketPath string,
	dhcpClientTimeout time.Duration, resendMax time.Duration, resendTimeout time.Duration,
	broadcast bool,
) error {
	_ = "STUB: not implemented"
	// since other goroutines (on separate threads) will change namespaces,
	// ensure the RPC server does not get scheduled onto those
	return nil
}

// Write the pidfile
