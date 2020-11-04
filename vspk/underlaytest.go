/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// UnderlayTestIdentity represents the Identity of the object
var UnderlayTestIdentity = bambou.Identity{
	Name:     "underlaytest",
	Category: "underlaytests",
}

// UnderlayTestsList represents a list of UnderlayTests
type UnderlayTestsList []*UnderlayTest

// UnderlayTestsAncestor is the interface that an ancestor of a UnderlayTest must implement.
// An Ancestor is defined as an entity that has UnderlayTest as a descendant.
// An Ancestor can get a list of its child UnderlayTests, but not necessarily create one.
type UnderlayTestsAncestor interface {
	UnderlayTests(*bambou.FetchingInfo) (UnderlayTestsList, *bambou.Error)
}

// UnderlayTestsParent is the interface that a parent of a UnderlayTest must implement.
// A Parent is defined as an entity that has UnderlayTest as a child.
// A Parent is an Ancestor which can create a UnderlayTest.
type UnderlayTestsParent interface {
	UnderlayTestsAncestor
	CreateUnderlayTest(*UnderlayTest) *bambou.Error
}

// UnderlayTest represents the model of a underlaytest
type UnderlayTest struct {
	ID                           string `json:"ID,omitempty"`
	ParentID                     string `json:"parentID,omitempty"`
	ParentType                   string `json:"parentType,omitempty"`
	Owner                        string `json:"owner,omitempty"`
	Name                         string `json:"name,omitempty"`
	TestResult                   string `json:"testResult,omitempty"`
	UnderlayTestServer           string `json:"underlayTestServer,omitempty"`
	UnderlayTestType             string `json:"underlayTestType,omitempty"`
	CreateOnly                   bool   `json:"createOnly"`
	AssociatedDataPathID         string `json:"associatedDataPathID,omitempty"`
	AssociatedNSGatewayID        string `json:"associatedNSGatewayID,omitempty"`
	AssociatedNSGatewayName      string `json:"associatedNSGatewayName,omitempty"`
	AssociatedSystemID           string `json:"associatedSystemID,omitempty"`
	AssociatedTestSuiteRunID     string `json:"associatedTestSuiteRunID,omitempty"`
	AssociatedUplinkConnectionID string `json:"associatedUplinkConnectionID,omitempty"`
	AssociatedUplinkInterface    string `json:"associatedUplinkInterface,omitempty"`
	StartDateTime                int    `json:"startDateTime,omitempty"`
	StopDateTime                 int    `json:"stopDateTime,omitempty"`
	RunBandwidthTest             bool   `json:"runBandwidthTest"`
	RunConnectivityTest          bool   `json:"runConnectivityTest"`
	RunMTUDiscoveryTest          bool   `json:"runMTUDiscoveryTest"`
	Duration                     int    `json:"duration,omitempty"`
}

// NewUnderlayTest returns a new *UnderlayTest
func NewUnderlayTest() *UnderlayTest {

	return &UnderlayTest{
		CreateOnly:          false,
		RunBandwidthTest:    true,
		RunConnectivityTest: true,
		RunMTUDiscoveryTest: true,
	}
}

// Identity returns the Identity of the object.
func (o *UnderlayTest) Identity() bambou.Identity {

	return UnderlayTestIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UnderlayTest) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UnderlayTest) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the UnderlayTest from the server
func (o *UnderlayTest) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the UnderlayTest into the server
func (o *UnderlayTest) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the UnderlayTest from the server
func (o *UnderlayTest) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
