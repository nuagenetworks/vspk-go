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

// AutoDiscoveredGatewayIdentity represents the Identity of the object
var AutoDiscoveredGatewayIdentity = bambou.Identity {
    Name:     "autodiscoveredgateway",
    Category: "autodiscoveredgateways",
}

// AutoDiscoveredGatewaysList represents a list of AutoDiscoveredGateways
type AutoDiscoveredGatewaysList []*AutoDiscoveredGateway

// AutoDiscoveredGatewaysAncestor is the interface that an ancestor of a AutoDiscoveredGateway must implement.
// An Ancestor is defined as an entity that has AutoDiscoveredGateway as a descendant.
// An Ancestor can get a list of its child AutoDiscoveredGateways, but not necessarily create one.
type AutoDiscoveredGatewaysAncestor interface {
    AutoDiscoveredGateways(*bambou.FetchingInfo) (AutoDiscoveredGatewaysList, *bambou.Error)
}

// AutoDiscoveredGatewaysParent is the interface that a parent of a AutoDiscoveredGateway must implement.
// A Parent is defined as an entity that has AutoDiscoveredGateway as a child.
// A Parent is an Ancestor which can create a AutoDiscoveredGateway.
type AutoDiscoveredGatewaysParent interface {
    AutoDiscoveredGatewaysAncestor
    CreateAutoDiscoveredGateway(*AutoDiscoveredGateway) (*bambou.Error)
}

// AutoDiscoveredGateway represents the model of a autodiscoveredgateway
type AutoDiscoveredGateway struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
    LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`
    GatewayID string `json:"gatewayID,omitempty"`
    Peer string `json:"peer,omitempty"`
    Personality string `json:"personality,omitempty"`
    Description string `json:"description,omitempty"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    Controllers []interface{} `json:"controllers,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    Vtep string `json:"vtep,omitempty"`
    Owner string `json:"owner,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    SystemID string `json:"systemID,omitempty"`
    
}

// NewAutoDiscoveredGateway returns a new *AutoDiscoveredGateway
func NewAutoDiscoveredGateway() *AutoDiscoveredGateway {

    return &AutoDiscoveredGateway{
        }
}

// Identity returns the Identity of the object.
func (o *AutoDiscoveredGateway) Identity() bambou.Identity {

    return AutoDiscoveredGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutoDiscoveredGateway) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutoDiscoveredGateway) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the AutoDiscoveredGateway from the server
func (o *AutoDiscoveredGateway) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AutoDiscoveredGateway into the server
func (o *AutoDiscoveredGateway) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AutoDiscoveredGateway from the server
func (o *AutoDiscoveredGateway) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// WANServices retrieves the list of child WANServices of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) WANServices(info *bambou.FetchingInfo) (WANServicesList, *bambou.Error) {

    var list WANServicesList
    err := bambou.CurrentSession().FetchChildren(o, WANServiceIdentity, &list, info)
    return list, err
}




// Metadatas retrieves the list of child Metadatas of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// WirelessPorts retrieves the list of child WirelessPorts of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) WirelessPorts(info *bambou.FetchingInfo) (WirelessPortsList, *bambou.Error) {

    var list WirelessPortsList
    err := bambou.CurrentSession().FetchChildren(o, WirelessPortIdentity, &list, info)
    return list, err
}




// GlobalMetadatas retrieves the list of child GlobalMetadatas of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Ports retrieves the list of child Ports of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) Ports(info *bambou.FetchingInfo) (PortsList, *bambou.Error) {

    var list PortsList
    err := bambou.CurrentSession().FetchChildren(o, PortIdentity, &list, info)
    return list, err
}




// NSPorts retrieves the list of child NSPorts of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) NSPorts(info *bambou.FetchingInfo) (NSPortsList, *bambou.Error) {

    var list NSPortsList
    err := bambou.CurrentSession().FetchChildren(o, NSPortIdentity, &list, info)
    return list, err
}




// EventLogs retrieves the list of child EventLogs of the AutoDiscoveredGateway
func (o *AutoDiscoveredGateway) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

    var list EventLogsList
    err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
    return list, err
}



