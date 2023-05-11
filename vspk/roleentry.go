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

// RoleentryIdentity represents the Identity of the object
var RoleentryIdentity = bambou.Identity {
    Name:     "roleentry",
    Category: "roleentries",
}

// RoleentriesList represents a list of Roleentries
type RoleentriesList []*Roleentry

// RoleentriesAncestor is the interface that an ancestor of a Roleentry must implement.
// An Ancestor is defined as an entity that has Roleentry as a descendant.
// An Ancestor can get a list of its child Roleentries, but not necessarily create one.
type RoleentriesAncestor interface {
    Roleentries(*bambou.FetchingInfo) (RoleentriesList, *bambou.Error)
}

// RoleentriesParent is the interface that a parent of a Roleentry must implement.
// A Parent is defined as an entity that has Roleentry as a child.
// A Parent is an Ancestor which can create a Roleentry.
type RoleentriesParent interface {
    RoleentriesAncestor
    CreateRoleentry(*Roleentry) (*bambou.Error)
}

// Roleentry represents the model of a roleentry
type Roleentry struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EndPointRestName string `json:"endPointRestName,omitempty"`
    EndPointType string `json:"endPointType,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    RoleAccessTypeList []interface{} `json:"roleAccessTypeList,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    
}

// NewRoleentry returns a new *Roleentry
func NewRoleentry() *Roleentry {

    return &Roleentry{
        }
}

// Identity returns the Identity of the object.
func (o *Roleentry) Identity() bambou.Identity {

    return RoleentryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Roleentry) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Roleentry) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the Roleentry from the server
func (o *Roleentry) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Roleentry into the server
func (o *Roleentry) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Roleentry from the server
func (o *Roleentry) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the Roleentry
func (o *Roleentry) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the Roleentry
func (o *Roleentry) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Metadatas retrieves the list of child Metadatas of the Roleentry
func (o *Roleentry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the Roleentry
func (o *Roleentry) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Roleentry
func (o *Roleentry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the Roleentry
func (o *Roleentry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

