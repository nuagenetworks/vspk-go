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

// PolicyStatementIdentity represents the Identity of the object
var PolicyStatementIdentity = bambou.Identity{
	Name:     "policystatement",
	Category: "policystatements",
}

// PolicyStatementsList represents a list of PolicyStatements
type PolicyStatementsList []*PolicyStatement

// PolicyStatementsAncestor is the interface that an ancestor of a PolicyStatement must implement.
// An Ancestor is defined as an entity that has PolicyStatement as a descendant.
// An Ancestor can get a list of its child PolicyStatements, but not necessarily create one.
type PolicyStatementsAncestor interface {
	PolicyStatements(*bambou.FetchingInfo) (PolicyStatementsList, *bambou.Error)
}

// PolicyStatementsParent is the interface that a parent of a PolicyStatement must implement.
// A Parent is defined as an entity that has PolicyStatement as a child.
// A Parent is an Ancestor which can create a PolicyStatement.
type PolicyStatementsParent interface {
	PolicyStatementsAncestor
	CreatePolicyStatement(*PolicyStatement) *bambou.Error
}

// PolicyStatement represents the model of a policystatement
type PolicyStatement struct {
	ID               string        `json:"ID,omitempty"`
	ParentID         string        `json:"parentID,omitempty"`
	ParentType       string        `json:"parentType,omitempty"`
	Owner            string        `json:"owner,omitempty"`
	Name             string        `json:"name,omitempty"`
	LastUpdatedBy    string        `json:"lastUpdatedBy,omitempty"`
	Description      string        `json:"description,omitempty"`
	EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope      string        `json:"entityScope,omitempty"`
	ExternalID       string        `json:"externalID,omitempty"`
}

// NewPolicyStatement returns a new *PolicyStatement
func NewPolicyStatement() *PolicyStatement {

	return &PolicyStatement{}
}

// Identity returns the Identity of the object.
func (o *PolicyStatement) Identity() bambou.Identity {

	return PolicyStatementIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyStatement) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyStatement) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyStatement from the server
func (o *PolicyStatement) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyStatement into the server
func (o *PolicyStatement) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyStatement from the server
func (o *PolicyStatement) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the PolicyStatement
func (o *PolicyStatement) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the PolicyStatement
func (o *PolicyStatement) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the PolicyStatement
func (o *PolicyStatement) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyStatement
func (o *PolicyStatement) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyStatement
func (o *PolicyStatement) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyStatement
func (o *PolicyStatement) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PolicyEntries retrieves the list of child PolicyEntries of the PolicyStatement
func (o *PolicyStatement) PolicyEntries(info *bambou.FetchingInfo) (PolicyEntriesList, *bambou.Error) {

	var list PolicyEntriesList
	err := bambou.CurrentSession().FetchChildren(o, PolicyEntryIdentity, &list, info)
	return list, err
}

// CreatePolicyEntry creates a new child PolicyEntry under the PolicyStatement
func (o *PolicyStatement) CreatePolicyEntry(child *PolicyEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
