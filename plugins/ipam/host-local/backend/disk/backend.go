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

package disk

import (
	"net"

	"github.com/containernetworking/plugins/plugins/ipam/host-local/backend"
)

const (
	lastIPFilePrefix = "last_reserved_ip."
	LineBreak        = "\r\n"
)

var defaultDataDir = "/var/lib/cni/networks"

// Store is a simple disk-backed store that creates one file per IP
// address in a given directory. The contents of the file are the container ID.
type Store struct {
	*FileLock
	dataDir string
}

// Store implements the Store interface
var _ backend.Store = &Store{}

func New(network, dataDir string) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Store) Reserve(id string, ifname string, ip net.IP, rangeID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// store the reserved ip in lastIPFile

// LastReservedIP returns the last reserved IP if exists
func (s *Store) LastReservedIP(rangeID string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func (s *Store) FindByKey(match string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Store) FindByID(id string, ifname string) bool { _ = "STUB: not implemented"; return false }

// Match anything created by this id

func (s *Store) ReleaseByKey(match string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// N.B. This function eats errors to be tolerant and
// release as much as possible
func (s *Store) ReleaseByID(id string, ifname string) error { _ = "STUB: not implemented"; return nil }

// For backwards compatibility, look for files written by a previous version

// GetByID returns the IPs which have been allocated to the specific ID
func (s *Store) GetByID(id string, ifname string) []net.IP { _ = "STUB: not implemented"; return nil }

// matchOld for backwards compatibility

// walk through all ips in this network to get the ones which belong to a specific ID

func GetEscapedPath(dataDir string, fname string) string { _ = "STUB: not implemented"; return "" }
