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

package testing

import (
	"net"

	"github.com/containernetworking/plugins/plugins/ipam/host-local/backend"
)

type FakeStore struct {
	ipMap          map[string]string
	lastReservedIP map[string]net.IP
}

// FakeStore implements the Store interface
var _ backend.Store = &FakeStore{}

func NewFakeStore(ipmap map[string]string, lastIPs map[string]net.IP) *FakeStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *FakeStore) Lock() error { _ = "STUB: not implemented"; return nil }

func (s *FakeStore) Unlock() error { _ = "STUB: not implemented"; return nil }

func (s *FakeStore) Close() error { _ = "STUB: not implemented"; return nil }

func (s *FakeStore) Reserve(id string, _ string, ip net.IP, rangeID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *FakeStore) LastReservedIP(rangeID string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func (s *FakeStore) ReleaseByID(id string, _ string) error { _ = "STUB: not implemented"; return nil }

func (s *FakeStore) GetByID(id string, _ string) []net.IP { _ = "STUB: not implemented"; return nil }

func (s *FakeStore) SetIPMap(m map[string]string) { _ = "STUB: not implemented"; return }
