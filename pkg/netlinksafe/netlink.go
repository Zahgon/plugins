// Package netlinksafe wraps vishvandanda/netlink functions that may return EINTR.
//
// A Handle instantiated using [NewHandle] or [NewHandleAt] can be used in place
// of a netlink.Handle, it's a wrapper that replaces methods that need to be
// wrapped. Functions that use the package handle need to be called as "netlinksafe.X"
// instead of "netlink.X".
//
// The wrapped functions currently return EINTR when NLM_F_DUMP_INTR flagged
// in a netlink response, meaning something changed during the dump so results
// may be incomplete or inconsistent.
//
// At present, the possibly incomplete/inconsistent results are not returned
// by netlink functions along with the EINTR. So, it's not possible to do
// anything but retry. After maxAttempts the EINTR will be returned to the
// caller.
package netlinksafe

import (
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
	"github.com/vishvananda/netns"
)

// Arbitrary limit on max attempts at netlink calls if they are repeatedly interrupted.
const maxAttempts = 5

type Handle struct {
	*netlink.Handle
}

func NewHandle(nlFamilies ...int) (Handle, error) {
	_ = "STUB: not implemented"
	return *new(Handle), nil
}

func NewHandleAt(ns netns.NsHandle, nlFamilies ...int) (Handle, error) {
	_ = "STUB: not implemented"
	return *new(Handle), nil
}

func (h Handle) Close() { _ = "STUB: not implemented"; return }

func retryOnIntr(f func() error) { _ = "STUB: not implemented"; return }

func discardErrDumpInterrupted(err error) error { _ = "STUB: not implemented"; return nil }

// The netlink function has returned possibly-inconsistent data along with the
// error. Discard the error and return the data. This restores the behaviour of
// the netlink package prior to v1.2.1, in which NLM_F_DUMP_INTR was ignored in
// the netlink response.

// AddrList calls netlink.AddrList, retrying if necessary.
func AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// LinkByName calls h.Handle.LinkByName, retrying if necessary. The netlink function
// doesn't normally ask the kernel for a dump of links. But, on an old kernel, it
// will do as a fallback and that dump may get inconsistent results.
func (h Handle) LinkByName(name string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

//nolint:forbidigo

// LinkByName calls netlink.LinkByName, retrying if necessary. The netlink
// function doesn't normally ask the kernel for a dump of links. But, on an old
// kernel, it will do as a fallback and that dump may get inconsistent results.
func LinkByName(name string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

//nolint:forbidigo

// LinkList calls h.Handle.LinkList, retrying if necessary.
func (h Handle) LinkList() ([]netlink.Link, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:forbidigo

// LinkList calls netlink.Handle.LinkList, retrying if necessary.
func LinkList() ([]netlink.Link, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:forbidigo

// RouteList calls h.Handle.RouteList, retrying if necessary.
func (h Handle) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// RouteList calls netlink.RouteList, retrying if necessary.
func RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// BridgeVlanList calls netlink.BridgeVlanList, retrying if necessary.
func BridgeVlanList() (map[int32][]*nl.BridgeVlanInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// RouteListFiltered calls h.Handle.RouteListFiltered, retrying if necessary.
func (h Handle) RouteListFiltered(family int, filter *netlink.Route, filterMask uint64) ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// RouteListFiltered calls netlink.RouteListFiltered, retrying if necessary.
func RouteListFiltered(family int, filter *netlink.Route, filterMask uint64) ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// QdiscList calls netlink.QdiscList, retrying if necessary.
func QdiscList(link netlink.Link) ([]netlink.Qdisc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// QdiscList calls h.Handle.QdiscList, retrying if necessary.
func (h *Handle) QdiscList(link netlink.Link) ([]netlink.Qdisc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// LinkGetProtinfo calls netlink.LinkGetProtinfo, retrying if necessary.
func LinkGetProtinfo(link netlink.Link) (netlink.Protinfo, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Protinfo), nil
}

//nolint:forbidigo

// LinkGetProtinfo calls h.Handle.LinkGetProtinfo, retrying if necessary.
func (h *Handle) LinkGetProtinfo(link netlink.Link) (netlink.Protinfo, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Protinfo), nil
}

//nolint:forbidigo

// RuleListFiltered calls netlink.RuleListFiltered, retrying if necessary.
func RuleListFiltered(family int, filter *netlink.Rule, filterMask uint64) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// RuleListFiltered calls h.Handle.RuleListFiltered, retrying if necessary.
func (h *Handle) RuleListFiltered(family int, filter *netlink.Rule, filterMask uint64) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// FilterList calls netlink.FilterList, retrying if necessary.
func FilterList(link netlink.Link, parent uint32) ([]netlink.Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// FilterList calls h.Handle.FilterList, retrying if necessary.
func (h *Handle) FilterList(link netlink.Link, parent uint32) ([]netlink.Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// RuleList calls netlink.RuleList, retrying if necessary.
func RuleList(family int) ([]netlink.Rule, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:forbidigo

// RuleList calls h.Handle.RuleList, retrying if necessary.
func (h *Handle) RuleList(family int) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forbidigo

// ConntrackDeleteFilters calls netlink.ConntrackDeleteFilters, retrying if necessary.
func ConntrackDeleteFilters(table netlink.ConntrackTableType, family netlink.InetFamily, filters ...netlink.CustomConntrackFilter) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:forbidigo

// ConntrackDeleteFilters calls h.Handle.ConntrackDeleteFilters, retrying if necessary.
func (h *Handle) ConntrackDeleteFilters(table netlink.ConntrackTableType, family netlink.InetFamily, filters ...netlink.CustomConntrackFilter) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:forbidigo
