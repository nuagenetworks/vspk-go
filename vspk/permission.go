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

// PermissionIdentity represents the Identity of the object
var PermissionIdentity = bambou.Identity{
	Name:     "permission",
	Category: "permissions",
}

// PermissionsList represents a list of Permissions
type PermissionsList []*Permission

// PermissionsAncestor is the interface that an ancestor of a Permission must implement.
// An Ancestor is defined as an entity that has Permission as a descendant.
// An Ancestor can get a list of its child Permissions, but not necessarily create one.
type PermissionsAncestor interface {
	Permissions(*bambou.FetchingInfo) (PermissionsList, *bambou.Error)
}

// PermissionsParent is the interface that a parent of a Permission must implement.
// A Parent is defined as an entity that has Permission as a child.
// A Parent is an Ancestor which can create a Permission.
type PermissionsParent interface {
	PermissionsAncestor
	CreatePermission(*Permission) *bambou.Error
}

// Permission represents the model of a permission
type Permission struct {
	ID                             string        `json:"ID,omitempty"`
	ParentID                       string        `json:"parentID,omitempty"`
	ParentType                     string        `json:"parentType,omitempty"`
	Owner                          string        `json:"owner,omitempty"`
	Name                           string        `json:"name,omitempty"`
	LastUpdatedBy                  string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                string        `json:"lastUpdatedDate,omitempty"`
	PermittedAction                string        `json:"permittedAction,omitempty"`
	PermittedEnterpriseDescription string        `json:"permittedEnterpriseDescription,omitempty"`
	PermittedEnterpriseID          string        `json:"permittedEnterpriseID,omitempty"`
	PermittedEnterpriseName        string        `json:"permittedEnterpriseName,omitempty"`
	PermittedEntityID              string        `json:"permittedEntityID,omitempty"`
	PermittedEntityName            string        `json:"permittedEntityName,omitempty"`
	PermittedEntityType            string        `json:"permittedEntityType,omitempty"`
	EmbeddedMetadata               []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope                    string        `json:"entityScope,omitempty"`
	CreationDate                   string        `json:"creationDate,omitempty"`
	AssociatedGroupDescription     string        `json:"associatedGroupDescription,omitempty"`
	AssociatedGroupID              string        `json:"associatedGroupID,omitempty"`
	AssociatedGroupName            string        `json:"associatedGroupName,omitempty"`
	AssociatedRoleDescription      string        `json:"associatedRoleDescription,omitempty"`
	AssociatedRoleID               string        `json:"associatedRoleID,omitempty"`
	AssociatedRoleName             string        `json:"associatedRoleName,omitempty"`
	Owner                          string        `json:"owner,omitempty"`
	ExternalID                     string        `json:"externalID,omitempty"`
}

// NewPermission returns a new *Permission
func NewPermission() *Permission {

	return &Permission{}
}

// Identity returns the Identity of the object.
func (o *Permission) Identity() bambou.Identity {

	return PermissionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Permission) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Permission) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Permission from the server
func (o *Permission) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Permission into the server
func (o *Permission) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Permission from the server
func (o *Permission) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the Permission
func (o *Permission) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the Permission
func (o *Permission) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the Permission
func (o *Permission) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Permission
func (o *Permission) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Permission
func (o *Permission) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Permission
func (o *Permission) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the Permission
func (o *Permission) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
