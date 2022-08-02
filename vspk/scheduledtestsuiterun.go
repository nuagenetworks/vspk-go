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

// ScheduledtestsuiterunIdentity represents the Identity of the object
var ScheduledtestsuiterunIdentity = bambou.Identity{
	Name:     "scheduledtestsuiterun",
	Category: "scheduledtestsuiteruns",
}

// ScheduledtestsuiterunsList represents a list of Scheduledtestsuiteruns
type ScheduledtestsuiterunsList []*Scheduledtestsuiterun

// ScheduledtestsuiterunsAncestor is the interface that an ancestor of a Scheduledtestsuiterun must implement.
// An Ancestor is defined as an entity that has Scheduledtestsuiterun as a descendant.
// An Ancestor can get a list of its child Scheduledtestsuiteruns, but not necessarily create one.
type ScheduledtestsuiterunsAncestor interface {
	Scheduledtestsuiteruns(*bambou.FetchingInfo) (ScheduledtestsuiterunsList, *bambou.Error)
}

// ScheduledtestsuiterunsParent is the interface that a parent of a Scheduledtestsuiterun must implement.
// A Parent is defined as an entity that has Scheduledtestsuiterun as a child.
// A Parent is an Ancestor which can create a Scheduledtestsuiterun.
type ScheduledtestsuiterunsParent interface {
	ScheduledtestsuiterunsAncestor
	CreateScheduledtestsuiterun(*Scheduledtestsuiterun) *bambou.Error
}

// Scheduledtestsuiterun represents the model of a scheduledtestsuiterun
type Scheduledtestsuiterun struct {
	ID                               string        `json:"ID,omitempty"`
	ParentID                         string        `json:"parentID,omitempty"`
	ParentType                       string        `json:"parentType,omitempty"`
	Owner                            string        `json:"owner,omitempty"`
	VPortName                        string        `json:"VPortName,omitempty"`
	NSGatewayName                    string        `json:"NSGatewayName,omitempty"`
	MacAddress                       string        `json:"macAddress,omitempty"`
	LastUpdatedBy                    string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                  string        `json:"lastUpdatedDate,omitempty"`
	DatapathID                       string        `json:"datapathID,omitempty"`
	SecondaryDatapathID              string        `json:"secondaryDatapathID,omitempty"`
	SecondaryNSGatewayName           string        `json:"secondaryNSGatewayName,omitempty"`
	SecondarySystemID                string        `json:"secondarySystemID,omitempty"`
	Destination                      string        `json:"destination,omitempty"`
	VlanID                           int           `json:"vlanID,omitempty"`
	EmbeddedMetadata                 []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope                      string        `json:"entityScope,omitempty"`
	DomainName                       string        `json:"domainName,omitempty"`
	ZoneName                         string        `json:"zoneName,omitempty"`
	SourceIP                         string        `json:"sourceIP,omitempty"`
	OperationStatus                  string        `json:"operationStatus,omitempty"`
	VportPortName                    string        `json:"vportPortName,omitempty"`
	VportVlanID                      int           `json:"vportVlanID,omitempty"`
	CreationDate                     string        `json:"creationDate,omitempty"`
	AssociatedScheduledTestSuiteID   string        `json:"associatedScheduledTestSuiteID,omitempty"`
	AssociatedScheduledTestSuiteName string        `json:"associatedScheduledTestSuiteName,omitempty"`
	SubnetName                       string        `json:"subnetName,omitempty"`
	Owner                            string        `json:"owner,omitempty"`
	ExternalID                       string        `json:"externalID,omitempty"`
	SystemID                         string        `json:"systemID,omitempty"`
}

// NewScheduledtestsuiterun returns a new *Scheduledtestsuiterun
func NewScheduledtestsuiterun() *Scheduledtestsuiterun {

	return &Scheduledtestsuiterun{}
}

// Identity returns the Identity of the object.
func (o *Scheduledtestsuiterun) Identity() bambou.Identity {

	return ScheduledtestsuiterunIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Scheduledtestsuiterun) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Scheduledtestsuiterun) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Scheduledtestsuiterun from the server
func (o *Scheduledtestsuiterun) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Scheduledtestsuiterun into the server
func (o *Scheduledtestsuiterun) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Scheduledtestsuiterun from the server
func (o *Scheduledtestsuiterun) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// TestRuns retrieves the list of child TestRuns of the Scheduledtestsuiterun
func (o *Scheduledtestsuiterun) TestRuns(info *bambou.FetchingInfo) (TestRunsList, *bambou.Error) {

	var list TestRunsList
	err := bambou.CurrentSession().FetchChildren(o, TestRunIdentity, &list, info)
	return list, err
}

// CreateTestRun creates a new child TestRun under the Scheduledtestsuiterun
func (o *Scheduledtestsuiterun) CreateTestRun(child *TestRun) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the Scheduledtestsuiterun
func (o *Scheduledtestsuiterun) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Scheduledtestsuiterun
func (o *Scheduledtestsuiterun) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Scheduledtestsuiterun
func (o *Scheduledtestsuiterun) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Scheduledtestsuiterun
func (o *Scheduledtestsuiterun) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
