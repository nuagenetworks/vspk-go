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

// ApplicationperformancemanagementbindingIdentity represents the Identity of the object
var ApplicationperformancemanagementbindingIdentity = bambou.Identity {
    Name:     "applicationperformancemanagementbinding",
    Category: "applicationperformancemanagementbindings",
}

// ApplicationperformancemanagementbindingsList represents a list of Applicationperformancemanagementbindings
type ApplicationperformancemanagementbindingsList []*Applicationperformancemanagementbinding

// ApplicationperformancemanagementbindingsAncestor is the interface that an ancestor of a Applicationperformancemanagementbinding must implement.
// An Ancestor is defined as an entity that has Applicationperformancemanagementbinding as a descendant.
// An Ancestor can get a list of its child Applicationperformancemanagementbindings, but not necessarily create one.
type ApplicationperformancemanagementbindingsAncestor interface {
    Applicationperformancemanagementbindings(*bambou.FetchingInfo) (ApplicationperformancemanagementbindingsList, *bambou.Error)
}

// ApplicationperformancemanagementbindingsParent is the interface that a parent of a Applicationperformancemanagementbinding must implement.
// A Parent is defined as an entity that has Applicationperformancemanagementbinding as a child.
// A Parent is an Ancestor which can create a Applicationperformancemanagementbinding.
type ApplicationperformancemanagementbindingsParent interface {
    ApplicationperformancemanagementbindingsAncestor
    CreateApplicationperformancemanagementbinding(*Applicationperformancemanagementbinding) (*bambou.Error)
}

// Applicationperformancemanagementbinding represents the model of a applicationperformancemanagementbinding
type Applicationperformancemanagementbinding struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
    LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`
    ReadOnly bool `json:"readOnly"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    Priority int `json:"priority,omitempty"`
    AssociatedApplicationPerformanceManagementID string `json:"associatedApplicationPerformanceManagementID,omitempty"`
    Owner string `json:"owner,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    
}

// NewApplicationperformancemanagementbinding returns a new *Applicationperformancemanagementbinding
func NewApplicationperformancemanagementbinding() *Applicationperformancemanagementbinding {

    return &Applicationperformancemanagementbinding{
        ReadOnly: false,
        }
}

// Identity returns the Identity of the object.
func (o *Applicationperformancemanagementbinding) Identity() bambou.Identity {

    return ApplicationperformancemanagementbindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Applicationperformancemanagementbinding) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Applicationperformancemanagementbinding) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the Applicationperformancemanagementbinding from the server
func (o *Applicationperformancemanagementbinding) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Applicationperformancemanagementbinding into the server
func (o *Applicationperformancemanagementbinding) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Applicationperformancemanagementbinding from the server
func (o *Applicationperformancemanagementbinding) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the Applicationperformancemanagementbinding
func (o *Applicationperformancemanagementbinding) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the Applicationperformancemanagementbinding
func (o *Applicationperformancemanagementbinding) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Metadatas retrieves the list of child Metadatas of the Applicationperformancemanagementbinding
func (o *Applicationperformancemanagementbinding) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the Applicationperformancemanagementbinding
func (o *Applicationperformancemanagementbinding) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Applicationperformancemanagementbinding
func (o *Applicationperformancemanagementbinding) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the Applicationperformancemanagementbinding
func (o *Applicationperformancemanagementbinding) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

