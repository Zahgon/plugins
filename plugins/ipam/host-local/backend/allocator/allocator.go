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

package allocator

import (
	"net"

	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/plugins/plugins/ipam/host-local/backend"
)

type IPAllocator struct {
	rangeset *RangeSet
	store    backend.Store
	rangeID  string // Used for tracking last reserved ip
}

func NewIPAllocator(s *RangeSet, store backend.Store, id int) *IPAllocator {
	_ = "STUB: not implemented"
	return nil
}

// Get allocates an IP
func (a *IPAllocator) Get(id string, ifname string, requestedIP net.IP) (*current.IPConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// try to get allocated IPs for this given id, if exists, just return error
// because duplicate allocation is not allowed in SPEC
// https://github.com/containernetworking/cni/blob/master/SPEC.md

// check whether the existing IP belong to this range set

// Release clears all IPs allocated for the container with given ID
func (a *IPAllocator) Release(id string, ifname string) error {
	_ = "STUB: not implemented"
	return nil
}

type RangeIter struct {
	rangeset *RangeSet

	// The current range id
	rangeIdx int

	// Our current position
	cur net.IP

	// The IP where we started iterating; if we hit this again, we're done.
	startIP net.IP
}

// GetIter encapsulates the strategy for this allocator.
// We use a round-robin strategy, attempting to evenly use the whole set.
// More specifically, a crash-looping container will not see the same IP until
// the entire range has been run through.
// We may wish to consider avoiding recently-released IPs in the future.
func (a *IPAllocator) GetIter() (*RangeIter, error) { _ = "STUB: not implemented"; return nil, nil }

// Round-robin by trying to allocate from the last reserved IP + 1

// We might get a last reserved IP that is wrong if the range indexes changed.
// This is not critical, we just lose round-robin this one time.

// Find the range in the set with this IP

// We advance the cursor on every Next(), so the first call
// to next() will return lastReservedIP + 1

// Next returns the next IP, its mask, and its gateway. Returns nil
// if the iterator has been exhausted
func (i *RangeIter) Next() (*net.IPNet, net.IP) {
	_ = "STUB: not implemented"
	return nil, *new(net.IP)
}

// If this is the first time iterating and we're not starting in the middle
// of the range, then start at rangeStart, which is inclusive

// If we've reached the end of this range, we need to advance the range
// RangeEnd is inclusive as well

// IF we've looped back to where we started, give up
