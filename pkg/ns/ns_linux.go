// Copyright 2015-2017 CNI authors
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

package ns

import (
	"os"

	"golang.org/x/sys/unix"
)

// Returns an object representing the current OS thread's network namespace
func GetCurrentNS() (NetNS, error) {
	_ = "STUB: not implemented"
	// Lock the thread in case other goroutine executes in it and changes its
	// network namespace after getCurrentThreadNetNSPath(), otherwise it might
	// return an unexpected network namespace.
	return *new(NetNS), nil
}

func getCurrentNSNoLock() (NetNS, error) { _ = "STUB: not implemented"; return *new(NetNS), nil }

func getCurrentThreadNetNSPath() string {
	_ = "STUB: not implemented"
	// /proc/self/ns/net returns the namespace of the main thread, not
	// of whatever thread this goroutine is running on.  Make sure we
	// use the thread's net namespace since the thread is switching around
	return ""
}

func (ns *netNS) Close() error { _ = "STUB: not implemented"; return nil }

func (ns *netNS) Set() error { _ = "STUB: not implemented"; return nil }

type NetNS interface {
	// Executes the passed closure in this object's network namespace,
	// attempting to restore the original namespace before returning.
	// However, since each OS thread can have a different network namespace,
	// and Go's thread scheduling is highly variable, callers cannot
	// guarantee any specific namespace is set unless operations that
	// require that namespace are wrapped with Do().  Also, no code called
	// from Do() should call runtime.UnlockOSThread(), or the risk
	// of executing code in an incorrect namespace will be greater.  See
	// https://github.com/golang/go/wiki/LockOSThread for further details.
	Do(toRun func(NetNS) error) error

	// Sets the current network namespace to this object's network namespace.
	// Note that since Go's thread scheduling is highly variable, callers
	// cannot guarantee the requested namespace will be the current namespace
	// after this function is called; to ensure this wrap operations that
	// require the namespace with Do() instead.
	Set() error

	// Returns the filesystem path representing this object's network namespace
	Path() string

	// Returns a file descriptor representing this object's network namespace
	Fd() uintptr

	// Cleans up this instance of the network namespace; if this instance
	// is the last user the namespace will be destroyed
	Close() error
}

type netNS struct {
	file   *os.File
	closed bool
}

// netNS implements the NetNS interface
var _ NetNS = &netNS{}

const (
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/magic.h
	NSFS_MAGIC   = unix.NSFS_MAGIC
	PROCFS_MAGIC = unix.PROC_SUPER_MAGIC
)

type NSPathNotExistErr struct{ msg string }

func (e NSPathNotExistErr) Error() string { _ = "STUB: not implemented"; return "" }

type NSPathNotNSErr struct{ msg string }

func (e NSPathNotNSErr) Error() string { _ = "STUB: not implemented"; return "" }

func IsNSorErr(nspath string) error { _ = "STUB: not implemented"; return nil }

// Returns an object representing the namespace referred to by @path
func GetNS(nspath string) (NetNS, error) { _ = "STUB: not implemented"; return *new(NetNS), nil }

// Returns a new empty NetNS.
// Calling Close() let the kernel garbage collect the network namespace.
func TempNetNS() (NetNS, error) { _ = "STUB: not implemented"; return *new(NetNS), nil }

// Create the new namespace in a new goroutine so that if we later fail
// to switch the namespace back to the original one, we can safely
// leave the thread locked to die without a risk of the current thread
// left lingering with incorrect namespace.

// save a handle to current network namespace

// create the temporary network namespace

// get a handle to the temporary network namespace

// Unlock the current thread only when we successfully switched back
// to the original namespace; otherwise leave the thread locked which
// will force the runtime to scrap the current thread, that is maybe
// not as optimal but at least always safe to do.

func (ns *netNS) Path() string { _ = "STUB: not implemented"; return "" }

func (ns *netNS) Fd() uintptr { _ = "STUB: not implemented"; return 0 }

func (ns *netNS) errorIfClosed() error { _ = "STUB: not implemented"; return nil }

func (ns *netNS) Do(toRun func(NetNS) error) error { _ = "STUB: not implemented"; return nil }

// switch to target namespace

// switch back

// Unlock the current thread only when we successfully switched back
// to the original namespace; otherwise leave the thread locked which
// will force the runtime to scrap the current thread, that is maybe
// not as optimal but at least always safe to do.

// save a handle to current network namespace

// Start the callback in a new green thread so that if we later fail
// to switch the namespace back to the original one, we can safely
// leave the thread locked to die without a risk of the current thread
// left lingering with incorrect namespace.

// WithNetNSPath executes the passed closure under the given network
// namespace, restoring the original namespace afterwards.
func WithNetNSPath(nspath string, toRun func(NetNS) error) error {
	_ = "STUB: not implemented"
	return nil
}
