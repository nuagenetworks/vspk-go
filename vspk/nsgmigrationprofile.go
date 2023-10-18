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

// NSGMigrationProfileIdentity represents the Identity of the object
var NSGMigrationProfileIdentity = bambou.Identity{
	Name:     "nsgmigrationprofile",
	Category: "nsgmigrationprofiles",
}

// NSGMigrationProfilesList represents a list of NSGMigrationProfiles
type NSGMigrationProfilesList []*NSGMigrationProfile

// NSGMigrationProfilesAncestor is the interface that an ancestor of a NSGMigrationProfile must implement.
// An Ancestor is defined as an entity that has NSGMigrationProfile as a descendant.
// An Ancestor can get a list of its child NSGMigrationProfiles, but not necessarily create one.
type NSGMigrationProfilesAncestor interface {
	NSGMigrationProfiles(*bambou.FetchingInfo) (NSGMigrationProfilesList, *bambou.Error)
}

// NSGMigrationProfilesParent is the interface that a parent of a NSGMigrationProfile must implement.
// A Parent is defined as an entity that has NSGMigrationProfile as a child.
// A Parent is an Ancestor which can create a NSGMigrationProfile.
type NSGMigrationProfilesParent interface {
	NSGMigrationProfilesAncestor
	CreateNSGMigrationProfile(*NSGMigrationProfile) *bambou.Error
}

// NSGMigrationProfile represents the model of a nsgmigrationprofile
type NSGMigrationProfile struct {
	ID                   string `json:"ID,omitempty"`
	ParentID             string `json:"parentID,omitempty"`
	ParentType           string `json:"parentType,omitempty"`
	Owner                string `json:"owner,omitempty"`
	Name                 string `json:"name,omitempty"`
	LastUpdatedBy        string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate      string `json:"lastUpdatedDate,omitempty"`
	Description          string `json:"description,omitempty"`
	DestinationProxyFQDN string `json:"destinationProxyFQDN,omitempty"`
	EntityScope          string `json:"entityScope,omitempty"`
	CreationDate         string `json:"creationDate,omitempty"`
	Owner                string `json:"owner,omitempty"`
	ExternalID           string `json:"externalID,omitempty"`
}

// NewNSGMigrationProfile returns a new *NSGMigrationProfile
func NewNSGMigrationProfile() *NSGMigrationProfile {

	return &NSGMigrationProfile{}
}

// Identity returns the Identity of the object.
func (o *NSGMigrationProfile) Identity() bambou.Identity {

	return NSGMigrationProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGMigrationProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGMigrationProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGMigrationProfile from the server
func (o *NSGMigrationProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGMigrationProfile into the server
func (o *NSGMigrationProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGMigrationProfile from the server
func (o *NSGMigrationProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the NSGMigrationProfile
func (o *NSGMigrationProfile) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the NSGMigrationProfile
func (o *NSGMigrationProfile) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
