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

// L7applicationsignatureIdentity represents the Identity of the object
var L7applicationsignatureIdentity = bambou.Identity{
	Name:     "l7applicationsignature",
	Category: "l7applicationsignatures",
}

// L7applicationsignaturesList represents a list of L7applicationsignatures
type L7applicationsignaturesList []*L7applicationsignature

// L7applicationsignaturesAncestor is the interface that an ancestor of a L7applicationsignature must implement.
// An Ancestor is defined as an entity that has L7applicationsignature as a descendant.
// An Ancestor can get a list of its child L7applicationsignatures, but not necessarily create one.
type L7applicationsignaturesAncestor interface {
	L7applicationsignatures(*bambou.FetchingInfo) (L7applicationsignaturesList, *bambou.Error)
}

// L7applicationsignaturesParent is the interface that a parent of a L7applicationsignature must implement.
// A Parent is defined as an entity that has L7applicationsignature as a child.
// A Parent is an Ancestor which can create a L7applicationsignature.
type L7applicationsignaturesParent interface {
	L7applicationsignaturesAncestor
	CreateL7applicationsignature(*L7applicationsignature) *bambou.Error
}

// L7applicationsignature represents the model of a l7applicationsignature
type L7applicationsignature struct {
	ID                string        `json:"ID,omitempty"`
	ParentID          string        `json:"parentID,omitempty"`
	ParentType        string        `json:"parentType,omitempty"`
	Owner             string        `json:"owner,omitempty"`
	Name              string        `json:"name,omitempty"`
	LastUpdatedBy     string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate   string        `json:"lastUpdatedDate,omitempty"`
	Category          string        `json:"category,omitempty"`
	Readonly          bool          `json:"readonly"`
	Reference         string        `json:"reference,omitempty"`
	Deprecated        bool          `json:"deprecated"`
	DeprecatedVersion string        `json:"deprecatedVersion,omitempty"`
	Description       string        `json:"description,omitempty"`
	DictionaryVersion int           `json:"dictionaryVersion,omitempty"`
	SignatureIndex    int           `json:"signatureIndex,omitempty"`
	SignatureVersion  string        `json:"signatureVersion,omitempty"`
	Risk              int           `json:"risk,omitempty"`
	PluginName        string        `json:"pluginName,omitempty"`
	EmbeddedMetadata  []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope       string        `json:"entityScope,omitempty"`
	SoftwareFlags     string        `json:"softwareFlags,omitempty"`
	CreationDate      string        `json:"creationDate,omitempty"`
	Productivity      int           `json:"productivity,omitempty"`
	Guidstring        string        `json:"guidstring,omitempty"`
	Owner             string        `json:"owner,omitempty"`
	ExternalID        string        `json:"externalID,omitempty"`
}

// NewL7applicationsignature returns a new *L7applicationsignature
func NewL7applicationsignature() *L7applicationsignature {

	return &L7applicationsignature{
		Readonly:   false,
		Deprecated: false,
	}
}

// Identity returns the Identity of the object.
func (o *L7applicationsignature) Identity() bambou.Identity {

	return L7applicationsignatureIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *L7applicationsignature) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *L7applicationsignature) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the L7applicationsignature from the server
func (o *L7applicationsignature) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the L7applicationsignature into the server
func (o *L7applicationsignature) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the L7applicationsignature from the server
func (o *L7applicationsignature) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the L7applicationsignature
func (o *L7applicationsignature) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the L7applicationsignature
func (o *L7applicationsignature) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the L7applicationsignature
func (o *L7applicationsignature) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the L7applicationsignature
func (o *L7applicationsignature) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the L7applicationsignature
func (o *L7applicationsignature) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the L7applicationsignature
func (o *L7applicationsignature) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Applications retrieves the list of child Applications of the L7applicationsignature
func (o *L7applicationsignature) Applications(info *bambou.FetchingInfo) (ApplicationsList, *bambou.Error) {

	var list ApplicationsList
	err := bambou.CurrentSession().FetchChildren(o, ApplicationIdentity, &list, info)
	return list, err
}
