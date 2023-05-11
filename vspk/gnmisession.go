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

// GNMISessionIdentity represents the Identity of the object
var GNMISessionIdentity = bambou.Identity {
    Name:     "gnmisession",
    Category: "gnmisessions",
}

// GNMISessionsList represents a list of GNMISessions
type GNMISessionsList []*GNMISession

// GNMISessionsAncestor is the interface that an ancestor of a GNMISession must implement.
// An Ancestor is defined as an entity that has GNMISession as a descendant.
// An Ancestor can get a list of its child GNMISessions, but not necessarily create one.
type GNMISessionsAncestor interface {
    GNMISessions(*bambou.FetchingInfo) (GNMISessionsList, *bambou.Error)
}

// GNMISessionsParent is the interface that a parent of a GNMISession must implement.
// A Parent is defined as an entity that has GNMISession as a child.
// A Parent is an Ancestor which can create a GNMISession.
type GNMISessionsParent interface {
    GNMISessionsAncestor
    CreateGNMISession(*GNMISession) (*bambou.Error)
}

// GNMISession represents the model of a gnmisession
type GNMISession struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
    LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`
    GatewayModel string `json:"gatewayModel,omitempty"`
    GatewayVendor string `json:"gatewayVendor,omitempty"`
    GatewayVersion string `json:"gatewayVersion,omitempty"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    AssociatedGatewayID string `json:"associatedGatewayID,omitempty"`
    AssociatedGatewayName string `json:"associatedGatewayName,omitempty"`
    Status string `json:"status,omitempty"`
    SubscriptionError string `json:"subscriptionError,omitempty"`
    SubscriptionFailureReason string `json:"subscriptionFailureReason,omitempty"`
    SubscriptionFailureRetryCount int `json:"subscriptionFailureRetryCount,omitempty"`
    SubscriptionState string `json:"subscriptionState,omitempty"`
    Owner string `json:"owner,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    
}

// NewGNMISession returns a new *GNMISession
func NewGNMISession() *GNMISession {

    return &GNMISession{
        }
}

// Identity returns the Identity of the object.
func (o *GNMISession) Identity() bambou.Identity {

    return GNMISessionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GNMISession) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GNMISession) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the GNMISession from the server
func (o *GNMISession) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GNMISession into the server
func (o *GNMISession) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GNMISession from the server
func (o *GNMISession) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Metadatas retrieves the list of child Metadatas of the GNMISession
func (o *GNMISession) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the GNMISession
func (o *GNMISession) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Alarms retrieves the list of child Alarms of the GNMISession
func (o *GNMISession) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

    var list AlarmsList
    err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
    return list, err
}




// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GNMISession
func (o *GNMISession) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the GNMISession
func (o *GNMISession) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

