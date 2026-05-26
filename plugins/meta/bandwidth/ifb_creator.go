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

const latencyInMillis = 25

func CreateIfb(ifbDeviceName string, mtu int) error {
	_ = "STUB: not implemented"
	// do not set TxQLen > 0 nor TxQLen == -1 until issues have been fixed with numrxqueues / numtxqueues across interfaces
	// which needs to get set on IFB devices via upstream library: see hint https://github.com/containernetworking/plugins/pull/1097
	return nil
}

func TeardownIfb(deviceName string) error { _ = "STUB: not implemented"; return nil }

func CreateIngressQdisc(rateInBits, burstInBits uint64, hostDeviceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateEgressQdisc(rateInBits, burstInBits uint64, hostDeviceName string, ifbDeviceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// add qdisc ingress on host device

// ffff:

// add filter on host device to mirror traffic to ifb device

// throttle traffic on ifb device

func createTBF(rateInBits, burstInBits uint64, linkIndex int) error {
	_ = "STUB: not implemented"
	// Equivalent to
	// tc qdisc add dev link root tbf
	//		rate netConf.BandwidthLimits.Rate
	//		burst netConf.BandwidthLimits.Burst
	return nil
}

func time2Tick(time uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func buffer(rate uint64, burst uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func limit(rate uint64, latency float64, buffer uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func latencyInUsec(latencyInMillis float64) float64 { _ = "STUB: not implemented"; return 0 }
