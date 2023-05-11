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

// BandwidthTestResultIdentity represents the Identity of the object
var BandwidthTestResultIdentity = bambou.Identity {
    Name:     "None",
    Category: "None",
}

// BandwidthTestResultsList represents a list of BandwidthTestResults
type BandwidthTestResultsList []*BandwidthTestResult

// BandwidthTestResultsAncestor is the interface that an ancestor of a BandwidthTestResult must implement.
// An Ancestor is defined as an entity that has BandwidthTestResult as a descendant.
// An Ancestor can get a list of its child BandwidthTestResults, but not necessarily create one.
type BandwidthTestResultsAncestor interface {
    BandwidthTestResults(*bambou.FetchingInfo) (BandwidthTestResultsList, *bambou.Error)
}

// BandwidthTestResultsParent is the interface that a parent of a BandwidthTestResult must implement.
// A Parent is defined as an entity that has BandwidthTestResult as a child.
// A Parent is an Ancestor which can create a BandwidthTestResult.
type BandwidthTestResultsParent interface {
    BandwidthTestResultsAncestor
    CreateBandwidthTestResult(*BandwidthTestResult) (*bambou.Error)
}

// BandwidthTestResult represents the model of a None
type BandwidthTestResult struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Bandwidth float64 `json:"bandwidth,omitempty"`
    
}

// NewBandwidthTestResult returns a new *BandwidthTestResult
func NewBandwidthTestResult() *BandwidthTestResult {

    return &BandwidthTestResult{
        }
}

// Identity returns the Identity of the object.
func (o *BandwidthTestResult) Identity() bambou.Identity {

    return BandwidthTestResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BandwidthTestResult) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BandwidthTestResult) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the BandwidthTestResult from the server
func (o *BandwidthTestResult) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BandwidthTestResult into the server
func (o *BandwidthTestResult) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BandwidthTestResult from the server
func (o *BandwidthTestResult) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}

