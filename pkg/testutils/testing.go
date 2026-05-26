// Copyright 2016 CNI authors
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

// AllSpecVersions contains all CNI spec version numbers
var AllSpecVersions = [...]string{"0.1.0", "0.2.0", "0.3.0", "0.3.1", "0.4.0", "1.0.0", "1.1.0"}

// SpecVersionHasIPVersion returns true if the given CNI specification version
// includes the "version" field in the IP address elements
func SpecVersionHasIPVersion(ver string) bool { _ = "STUB: not implemented"; return false }

// SpecVersionHasCHECK returns true if the given CNI specification version
// supports the CHECK command
func SpecVersionHasCHECK(ver string) bool { _ = "STUB: not implemented"; return false }

// SpecVersionHasSTATUS returns true if the given CNI specification version
// supports the STATUS command
func SpecVersionHasSTATUS(ver string) bool { _ = "STUB: not implemented"; return false }

// SpecVersionHasChaining returns true if the given CNI specification version
// supports plugin chaining
func SpecVersionHasChaining(ver string) bool { _ = "STUB: not implemented"; return false }

// SpecVersionHasMultipleIPs returns true if the given CNI specification version
// supports more than one IP address of each family
func SpecVersionHasMultipleIPs(ver string) bool { _ = "STUB: not implemented"; return false }
