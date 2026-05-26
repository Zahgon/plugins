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
	"encoding/json"
	"net"

	"github.com/Microsoft/hcsshim/hcn"
	"github.com/containernetworking/cni/pkg/types"
)

// NetConf is the CNI spec
type NetConf struct {
	types.NetConf
	// ApiVersion specifies the policies type of HNS or HCN, select one of [1, 2].
	// HNS is the v1 API, which is the default version and applies to dockershim.
	// HCN is the v2 API, which can leverage HostComputeNamespace and use in containerd.
	ApiVersion int `json:"apiVersion,omitempty"`
	// Policies specifies the policy list for HNSEndpoint or HostComputeEndpoint.
	Policies []Policy `json:"policies,omitempty"`
	// RuntimeConfig represents the options to be passed in by the runtime.
	RuntimeConfig RuntimeConfig `json:"runtimeConfig"`
	// LoopbackDSR specifies whether to support loopback direct server return.
	LoopbackDSR bool `json:"loopbackDSR,omitempty"`
}

type RuntimeDNS struct {
	Nameservers []string `json:"servers,omitempty"`
	Search      []string `json:"searches,omitempty"`
}

type PortMapEntry struct {
	HostPort      int    `json:"hostPort"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
	HostIP        string `json:"hostIP,omitempty"`
}

// constants of the supported Windows Socket protocol,
// ref to https://docs.microsoft.com/en-us/dotnet/api/system.net.sockets.protocoltype.
var protocolEnums = map[string]uint32{
	"icmpv4": 1,
	"igmp":   2,
	"tcp":    6,
	"udp":    17,
	"icmpv6": 58,
}

func (p *PortMapEntry) GetProtocolEnum() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

type RuntimeConfig struct {
	DNS      RuntimeDNS     `json:"dns"`
	PortMaps []PortMapEntry `json:"portMappings,omitempty"`
}

type Policy struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

// GetHNSEndpointPolicies converts the configuration policies to HNSEndpoint policies.
func (n *NetConf) GetHNSEndpointPolicies() []json.RawMessage { _ = "STUB: not implemented"; return nil }

// GetHostComputeEndpointPolicies converts the configuration policies to HostComputeEndpoint policies.
func (n *NetConf) GetHostComputeEndpointPolicies() []hcn.EndpointPolicy {
	_ = "STUB: not implemented"
	return nil
}

// GetDNS returns the DNS values if they are there use that else use netconf supplied DNS.
func (n *NetConf) GetDNS() types.DNS { _ = "STUB: not implemented"; return *new(types.DNS) }

// ApplyLoopbackDSRPolicy configures the given IP to support loopback DSR.
func (n *NetConf) ApplyLoopbackDSRPolicy(ip *net.IP) { _ = "STUB: not implemented"; return }

// find OutBoundNAT policy

// filter OutBoundNAT policy

// parse destination address list

// skip if Destinations/DestinationList field is not found

// return if found the given address

// or add a new OutBoundNAT if not found

// ApplyOutboundNatPolicy applies the sNAT policy in HNS/HCN and configures the given CIDR as an exception.
func (n *NetConf) ApplyOutboundNatPolicy(exceptionCIDR string) { _ = "STUB: not implemented"; return }

// find OutBoundNAT policy

// filter OutBoundNAT policy

// parse exception CIDR list

// skip if Exceptions/ExceptionList field is not found

// return if found the given CIDR

// or add a new OutBoundNAT if not found

// ApplyDefaultPAPolicy applies an endpoint PA policy in HNS/HCN.
func (n *NetConf) ApplyDefaultPAPolicy(address string) { _ = "STUB: not implemented"; return }

// find ProviderAddress policy

// filter ProviderAddress policy

// parse provider address

// skip if ProviderAddress/PA field is not found

// return if found the given address

// or add a new ProviderAddress if not found

// ApplyPortMappingPolicy applies the host/container port mapping policies in HNS/HCN.
func (n *NetConf) ApplyPortMappingPolicy(portMappings []PortMapEntry) {
	_ = "STUB: not implemented"
	return
}

// skip the invalid protocol mapping

// bprintf is similar to fmt.Sprintf and returns a byte array as result.
func bprintf(format string, a ...interface{}) []byte { _ = "STUB: not implemented"; return nil }
