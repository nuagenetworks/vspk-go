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

// IPFilterProfileIdentity represents the Identity of the object
var IPFilterProfileIdentity = bambou.Identity{
	Name:     "ipfilterprofile",
	Category: "ipfilterprofiles",
}

// IPFilterProfilesList represents a list of IPFilterProfiles
type IPFilterProfilesList []*IPFilterProfile

// IPFilterProfilesAncestor is the interface that an ancestor of a IPFilterProfile must implement.
// An Ancestor is defined as an entity that has IPFilterProfile as a descendant.
// An Ancestor can get a list of its child IPFilterProfiles, but not necessarily create one.
type IPFilterProfilesAncestor interface {
	IPFilterProfiles(*bambou.FetchingInfo) (IPFilterProfilesList, *bambou.Error)
}

// IPFilterProfilesParent is the interface that a parent of a IPFilterProfile must implement.
// A Parent is defined as an entity that has IPFilterProfile as a child.
// A Parent is an Ancestor which can create a IPFilterProfile.
type IPFilterProfilesParent interface {
	IPFilterProfilesAncestor
	CreateIPFilterProfile(*IPFilterProfile) *bambou.Error
}

// IPFilterProfile represents the model of a ipfilterprofile
type IPFilterProfile struct {
	ID                 string        `json:"ID,omitempty"`
	ParentID           string        `json:"parentID,omitempty"`
	ParentType         string        `json:"parentType,omitempty"`
	Owner              string        `json:"owner,omitempty"`
	Name               string        `json:"name,omitempty"`
	LastUpdatedBy      string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate    string        `json:"lastUpdatedDate,omitempty"`
	Description        string        `json:"description,omitempty"`
	BlobText           string        `json:"blobText,omitempty"`
	BlobText_SRL_23_10 string        `json:"blobText_SRL_23_10,omitempty"`
	BlobType           string        `json:"blobType,omitempty"`
	EmbeddedMetadata   []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope        string        `json:"entityScope,omitempty"`
	CreationDate       string        `json:"creationDate,omitempty"`
	AssocEntityType    string        `json:"assocEntityType,omitempty"`
	CustomerID         int           `json:"customerID,omitempty"`
	Owner              string        `json:"owner,omitempty"`
	ExternalID         string        `json:"externalID,omitempty"`
}

// NewIPFilterProfile returns a new *IPFilterProfile
func NewIPFilterProfile() *IPFilterProfile {

	return &IPFilterProfile{}
}

// Identity returns the Identity of the object.
func (o *IPFilterProfile) Identity() bambou.Identity {

	return IPFilterProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IPFilterProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IPFilterProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IPFilterProfile from the server
func (o *IPFilterProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IPFilterProfile into the server
func (o *IPFilterProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IPFilterProfile from the server
func (o *IPFilterProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the IPFilterProfile
func (o *IPFilterProfile) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the IPFilterProfile
func (o *IPFilterProfile) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the IPFilterProfile
func (o *IPFilterProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IPFilterProfile
func (o *IPFilterProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IPFilterProfile
func (o *IPFilterProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IPFilterProfile
func (o *IPFilterProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
