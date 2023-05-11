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

// UDPProbeTestResultIdentity represents the Identity of the object
var UDPProbeTestResultIdentity = bambou.Identity {
    Name:     "None",
    Category: "None",
}

// UDPProbeTestResultsList represents a list of UDPProbeTestResults
type UDPProbeTestResultsList []*UDPProbeTestResult

// UDPProbeTestResultsAncestor is the interface that an ancestor of a UDPProbeTestResult must implement.
// An Ancestor is defined as an entity that has UDPProbeTestResult as a descendant.
// An Ancestor can get a list of its child UDPProbeTestResults, but not necessarily create one.
type UDPProbeTestResultsAncestor interface {
    UDPProbeTestResults(*bambou.FetchingInfo) (UDPProbeTestResultsList, *bambou.Error)
}

// UDPProbeTestResultsParent is the interface that a parent of a UDPProbeTestResult must implement.
// A Parent is defined as an entity that has UDPProbeTestResult as a child.
// A Parent is an Ancestor which can create a UDPProbeTestResult.
type UDPProbeTestResultsParent interface {
    UDPProbeTestResultsAncestor
    CreateUDPProbeTestResult(*UDPProbeTestResult) (*bambou.Error)
}

// UDPProbeTestResult represents the model of a None
type UDPProbeTestResult struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    PacketLossPercent float64 `json:"packetLossPercent,omitempty"`
    PacketsLost int `json:"packetsLost,omitempty"`
    PacketsReceived int `json:"packetsReceived,omitempty"`
    PacketsTransmitted int `json:"packetsTransmitted,omitempty"`
    
}

// NewUDPProbeTestResult returns a new *UDPProbeTestResult
func NewUDPProbeTestResult() *UDPProbeTestResult {

    return &UDPProbeTestResult{
        }
}

// Identity returns the Identity of the object.
func (o *UDPProbeTestResult) Identity() bambou.Identity {

    return UDPProbeTestResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UDPProbeTestResult) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UDPProbeTestResult) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the UDPProbeTestResult from the server
func (o *UDPProbeTestResult) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the UDPProbeTestResult into the server
func (o *UDPProbeTestResult) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the UDPProbeTestResult from the server
func (o *UDPProbeTestResult) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}

