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

// GNMIProfileIdentity represents the Identity of the object
var GNMIProfileIdentity = bambou.Identity{
	Name:     "gnmiprofile",
	Category: "gnmiprofiles",
}

// GNMIProfilesList represents a list of GNMIProfiles
type GNMIProfilesList []*GNMIProfile

// GNMIProfilesAncestor is the interface that an ancestor of a GNMIProfile must implement.
// An Ancestor is defined as an entity that has GNMIProfile as a descendant.
// An Ancestor can get a list of its child GNMIProfiles, but not necessarily create one.
type GNMIProfilesAncestor interface {
	GNMIProfiles(*bambou.FetchingInfo) (GNMIProfilesList, *bambou.Error)
}

// GNMIProfilesParent is the interface that a parent of a GNMIProfile must implement.
// A Parent is defined as an entity that has GNMIProfile as a child.
// A Parent is an Ancestor which can create a GNMIProfile.
type GNMIProfilesParent interface {
	GNMIProfilesAncestor
	CreateGNMIProfile(*GNMIProfile) *bambou.Error
}

// GNMIProfile represents the model of a gnmiprofile
type GNMIProfile struct {
	ID               string        `json:"ID,omitempty"`
	ParentID         string        `json:"parentID,omitempty"`
	ParentType       string        `json:"parentType,omitempty"`
	Owner            string        `json:"owner,omitempty"`
	Name             string        `json:"name,omitempty"`
	Password         string        `json:"password,omitempty"`
	LastUpdatedBy    string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate  string        `json:"lastUpdatedDate,omitempty"`
	Description      string        `json:"description,omitempty"`
	KeyAlias         string        `json:"keyAlias,omitempty"`
	EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope      string        `json:"entityScope,omitempty"`
	Port             int           `json:"port,omitempty"`
	CreationDate     string        `json:"creationDate,omitempty"`
	UserName         string        `json:"userName,omitempty"`
	AssocEntityType  string        `json:"assocEntityType,omitempty"`
	Owner            string        `json:"owner,omitempty"`
	ExternalID       string        `json:"externalID,omitempty"`
}

// NewGNMIProfile returns a new *GNMIProfile
func NewGNMIProfile() *GNMIProfile {

	return &GNMIProfile{
		Port: 57400,
	}
}

// Identity returns the Identity of the object.
func (o *GNMIProfile) Identity() bambou.Identity {

	return GNMIProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GNMIProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GNMIProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GNMIProfile from the server
func (o *GNMIProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GNMIProfile into the server
func (o *GNMIProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GNMIProfile from the server
func (o *GNMIProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GNMIProfile
func (o *GNMIProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GNMIProfile
func (o *GNMIProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GNMIProfile
func (o *GNMIProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GNMIProfile
func (o *GNMIProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
