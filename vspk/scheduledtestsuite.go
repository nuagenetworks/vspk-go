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

// ScheduledTestSuiteIdentity represents the Identity of the object
var ScheduledTestSuiteIdentity = bambou.Identity {
    Name:     "scheduledtestsuite",
    Category: "scheduledtestsuites",
}

// ScheduledTestSuitesList represents a list of ScheduledTestSuites
type ScheduledTestSuitesList []*ScheduledTestSuite

// ScheduledTestSuitesAncestor is the interface that an ancestor of a ScheduledTestSuite must implement.
// An Ancestor is defined as an entity that has ScheduledTestSuite as a descendant.
// An Ancestor can get a list of its child ScheduledTestSuites, but not necessarily create one.
type ScheduledTestSuitesAncestor interface {
    ScheduledTestSuites(*bambou.FetchingInfo) (ScheduledTestSuitesList, *bambou.Error)
}

// ScheduledTestSuitesParent is the interface that a parent of a ScheduledTestSuite must implement.
// A Parent is defined as an entity that has ScheduledTestSuite as a child.
// A Parent is an Ancestor which can create a ScheduledTestSuite.
type ScheduledTestSuitesParent interface {
    ScheduledTestSuitesAncestor
    CreateScheduledTestSuite(*ScheduledTestSuite) (*bambou.Error)
}

// ScheduledTestSuite represents the model of a scheduledtestsuite
type ScheduledTestSuite struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
    LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`
    ScheduleInterval int `json:"scheduleInterval,omitempty"`
    ScheduleIntervalUnits string `json:"scheduleIntervalUnits,omitempty"`
    Description string `json:"description,omitempty"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EndDateTime float64 `json:"endDateTime,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    StartDateTime float64 `json:"startDateTime,omitempty"`
    Owner string `json:"owner,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    
}

// NewScheduledTestSuite returns a new *ScheduledTestSuite
func NewScheduledTestSuite() *ScheduledTestSuite {

    return &ScheduledTestSuite{
        ScheduleInterval: 1,
        ScheduleIntervalUnits: "DAYS",
        }
}

// Identity returns the Identity of the object.
func (o *ScheduledTestSuite) Identity() bambou.Identity {

    return ScheduledTestSuiteIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ScheduledTestSuite) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ScheduledTestSuite) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the ScheduledTestSuite from the server
func (o *ScheduledTestSuite) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ScheduledTestSuite into the server
func (o *ScheduledTestSuite) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ScheduledTestSuite from the server
func (o *ScheduledTestSuite) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Scheduledtestsuiteruns retrieves the list of child Scheduledtestsuiteruns of the ScheduledTestSuite
func (o *ScheduledTestSuite) Scheduledtestsuiteruns(info *bambou.FetchingInfo) (ScheduledtestsuiterunsList, *bambou.Error) {

    var list ScheduledtestsuiterunsList
    err := bambou.CurrentSession().FetchChildren(o, ScheduledtestsuiterunIdentity, &list, info)
    return list, err
}




// Tests retrieves the list of child Tests of the ScheduledTestSuite
func (o *ScheduledTestSuite) Tests(info *bambou.FetchingInfo) (TestsList, *bambou.Error) {

    var list TestsList
    err := bambou.CurrentSession().FetchChildren(o, TestIdentity, &list, info)
    return list, err
}



// CreateTest creates a new child Test under the ScheduledTestSuite
func (o *ScheduledTestSuite) CreateTest(child *Test) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Metadatas retrieves the list of child Metadatas of the ScheduledTestSuite
func (o *ScheduledTestSuite) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the ScheduledTestSuite
func (o *ScheduledTestSuite) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ScheduledTestSuite
func (o *ScheduledTestSuite) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the ScheduledTestSuite
func (o *ScheduledTestSuite) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

