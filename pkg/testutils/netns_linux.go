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

package testutils

import (
	"github.com/containernetworking/plugins/pkg/ns"
)

func getNsRunDir() string { _ = "STUB: not implemented"; return "" }

/// If XDG_RUNTIME_DIR is set, check if the current user owns /var/run.  If
// the owner is different, we are most likely running in a user namespace.
// In that case use $XDG_RUNTIME_DIR/netns as runtime dir.

// Creates a new persistent (bind-mounted) network namespace and returns an object
// representing that namespace, without switching to it.
func NewNS() (ns.NetNS, error) { _ = "STUB: not implemented"; return *new(ns.NetNS), nil }

// Create the directory for mounting network namespaces
// This needs to be a shared mountpoint in case it is mounted in to
// other namespaces (containers)

// Remount the namespace directory shared. This will fail if it is not
// already a mountpoint, so bind-mount it on to itself to "upgrade" it
// to a mountpoint.

// Recursively remount /var/run/netns on itself. The recursive flag is
// so that any existing netns bindmounts are carried over.

// Now we can make it shared

// create an empty file at the mount point

// Ensure the mount point is cleaned up on errors; if the namespace
// was successfully mounted this will have no effect because the file
// is in-use

// do namespace work in a dedicated goroutine, so that we can safely
// Lock/Unlock OSThread without upsetting the lock/unlock state of
// the caller of this function

// Don't unlock. By not unlocking, golang will kill the OS thread when the
// goroutine is done (for go1.10+)

// create a new netns on the current thread

// Put this thread back to the orig ns, since it might get reused (pre go1.10)

// bind mount the netns from the current thread (from /proc) onto the
// mount point. This causes the namespace to persist, even when there
// are no threads in the ns.

// UnmountNS unmounts the NS held by the netns object
func UnmountNS(ns ns.NetNS) error {
	_ = "STUB: not implemented"

	// Only unmount if it's been bind-mounted (don't touch namespaces in /proc...)
	return nil
}

// getCurrentThreadNetNSPath copied from pkg/ns
func getCurrentThreadNetNSPath() string {
	_ = "STUB: not implemented"
	// /proc/self/ns/net returns the namespace of the main thread, not
	// of whatever thread this goroutine is running on.  Make sure we
	// use the thread's net namespace since the thread is switching around
	return ""
}
