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

package hns

import (
	"net"

	"github.com/Microsoft/hcsshim"
	"github.com/Microsoft/hcsshim/hcn"

	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
)

const (
	pauseContainerNetNS = "none"
)

type EndpointInfo struct {
	EndpointName string
	DNS          types.DNS
	NetworkName  string
	NetworkId    string
	Gateway      net.IP
	IpAddress    net.IP
	MacAddress   string
}

// GetSandboxContainerID returns the sandbox ID of this pod.
func GetSandboxContainerID(containerID string, netNs string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetIpString returns the given IP as a string.
func GetIpString(ip *net.IP) string { _ = "STUB: not implemented"; return "" }

// GetDefaultDestinationPrefix returns the default destination prefix according to the given IP type.
func GetDefaultDestinationPrefix(ip *net.IP) string { _ = "STUB: not implemented"; return "" }

// ConstructEndpointName constructs endpoint id which is used to identify an endpoint from HNS/HCN.
func ConstructEndpointName(containerID string, netNs string, networkName string) string {
	_ = "STUB: not implemented"
	return ""
}

// GenerateHnsEndpoint generates an HNSEndpoint with given info and config.
func GenerateHnsEndpoint(epInfo *EndpointInfo, n *NetConf) (*hcsshim.HNSEndpoint, error) {
	_ = "STUB: not implemented"
	// run the IPAM plugin and get back the config to apply
	return nil, nil
}

// remove endpoint if corrupted

// RemoveHnsEndpoint detaches the given name endpoint from container specified by containerID,
// or removes the given name endpoint completely.
func RemoveHnsEndpoint(epName string, netns string, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// for shared endpoint, detach it from the container

// for removing the endpoint completely, hot detach is used at first

type HnsEndpointMakerFunc func() (*hcsshim.HNSEndpoint, error)

// AddHnsEndpoint attaches an HNSEndpoint to a container specified by containerID.
func AddHnsEndpoint(epName string, expectedNetworkId string, containerID string, netns string, makeEndpoint HnsEndpointMakerFunc) (*hcsshim.HNSEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for shared endpoint, we expect that the endpoint already exists

// verify the existing endpoint is corrupted or not

// create endpoint if not found

// attach to container

// ConstructHnsResult constructs the CNI result for the HNSEndpoint.
func ConstructHnsResult(hnsNetwork *hcsshim.HNSNetwork, hnsEndpoint *hcsshim.HNSEndpoint) (*current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateHcnEndpoint generates a HostComputeEndpoint with given info and config.
func GenerateHcnEndpoint(epInfo *EndpointInfo, n *NetConf) (*hcn.HostComputeEndpoint, error) {
	_ = "STUB: not implemented"
	// run the IPAM plugin and get back the config to apply
	return nil, nil
}

// verify the existing endpoint is corrupted or not

// remove endpoint if corrupted

// RemoveHcnEndpoint removes the given name endpoint from namespace.
func RemoveHcnEndpoint(epName string) error { _ = "STUB: not implemented"; return nil }

type HcnEndpointMakerFunc func() (*hcn.HostComputeEndpoint, error)

// AddHcnEndpoint attaches a HostComputeEndpoint to the given namespace.
func AddHcnEndpoint(epName string, expectedNetworkId string, namespace string, makeEndpoint HcnEndpointMakerFunc) (*hcn.HostComputeEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verify the existing endpoint is corrupted or not

// create endpoint if not found

// add to namespace

// ConstructHcnResult constructs the CNI result for the HostComputeEndpoint.
func ConstructHcnResult(hcnNetwork *hcn.HostComputeNetwork, hcnEndpoint *hcn.HostComputeEndpoint) (*current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
