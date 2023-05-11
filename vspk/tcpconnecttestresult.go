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

// TCPConnectTestResultIdentity represents the Identity of the object
var TCPConnectTestResultIdentity = bambou.Identity {
    Name:     "None",
    Category: "None",
}

// TCPConnectTestResultsList represents a list of TCPConnectTestResults
type TCPConnectTestResultsList []*TCPConnectTestResult

// TCPConnectTestResultsAncestor is the interface that an ancestor of a TCPConnectTestResult must implement.
// An Ancestor is defined as an entity that has TCPConnectTestResult as a descendant.
// An Ancestor can get a list of its child TCPConnectTestResults, but not necessarily create one.
type TCPConnectTestResultsAncestor interface {
    TCPConnectTestResults(*bambou.FetchingInfo) (TCPConnectTestResultsList, *bambou.Error)
}

// TCPConnectTestResultsParent is the interface that a parent of a TCPConnectTestResult must implement.
// A Parent is defined as an entity that has TCPConnectTestResult as a child.
// A Parent is an Ancestor which can create a TCPConnectTestResult.
type TCPConnectTestResultsParent interface {
    TCPConnectTestResultsAncestor
    CreateTCPConnectTestResult(*TCPConnectTestResult) (*bambou.Error)
}

// TCPConnectTestResult represents the model of a None
type TCPConnectTestResult struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    FailedAttempts int `json:"failedAttempts,omitempty"`
    FailedPercent float64 `json:"failedPercent,omitempty"`
    MaximumRoundTripTime float64 `json:"maximumRoundTripTime,omitempty"`
    MinimumRoundTripTime float64 `json:"minimumRoundTripTime,omitempty"`
    ConnectionAttempts int `json:"connectionAttempts,omitempty"`
    SuccessfulConnections int `json:"successfulConnections,omitempty"`
    AverageRoundTripTime float64 `json:"averageRoundTripTime,omitempty"`
    
}

// NewTCPConnectTestResult returns a new *TCPConnectTestResult
func NewTCPConnectTestResult() *TCPConnectTestResult {

    return &TCPConnectTestResult{
        }
}

// Identity returns the Identity of the object.
func (o *TCPConnectTestResult) Identity() bambou.Identity {

    return TCPConnectTestResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TCPConnectTestResult) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TCPConnectTestResult) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the TCPConnectTestResult from the server
func (o *TCPConnectTestResult) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the TCPConnectTestResult into the server
func (o *TCPConnectTestResult) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the TCPConnectTestResult from the server
func (o *TCPConnectTestResult) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}

