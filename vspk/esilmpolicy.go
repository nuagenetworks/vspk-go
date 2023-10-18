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

// EsIlmPolicyIdentity represents the Identity of the object
var EsIlmPolicyIdentity = bambou.Identity{
	Name:     "esilmpolicy",
	Category: "esilmpolicies",
}

// EsIlmPoliciesList represents a list of EsIlmPolicies
type EsIlmPoliciesList []*EsIlmPolicy

// EsIlmPoliciesAncestor is the interface that an ancestor of a EsIlmPolicy must implement.
// An Ancestor is defined as an entity that has EsIlmPolicy as a descendant.
// An Ancestor can get a list of its child EsIlmPolicies, but not necessarily create one.
type EsIlmPoliciesAncestor interface {
	EsIlmPolicies(*bambou.FetchingInfo) (EsIlmPoliciesList, *bambou.Error)
}

// EsIlmPoliciesParent is the interface that a parent of a EsIlmPolicy must implement.
// A Parent is defined as an entity that has EsIlmPolicy as a child.
// A Parent is an Ancestor which can create a EsIlmPolicy.
type EsIlmPoliciesParent interface {
	EsIlmPoliciesAncestor
	CreateEsIlmPolicy(*EsIlmPolicy) *bambou.Error
}

// EsIlmPolicy represents the model of a esilmpolicy
type EsIlmPolicy struct {
	ID                       string        `json:"ID,omitempty"`
	ParentID                 string        `json:"parentID,omitempty"`
	ParentType               string        `json:"parentType,omitempty"`
	Owner                    string        `json:"owner,omitempty"`
	Name                     string        `json:"name,omitempty"`
	WarmPhaseEnabled         bool          `json:"warmPhaseEnabled"`
	WarmTimer                int           `json:"warmTimer,omitempty"`
	DeletePhaseEnabled       bool          `json:"deletePhaseEnabled"`
	DeleteTimer              int           `json:"deleteTimer,omitempty"`
	Description              string        `json:"description,omitempty"`
	EmbeddedMetadata         []interface{} `json:"embeddedMetadata,omitempty"`
	IndexFreeze              bool          `json:"indexFreeze"`
	IndexReadOnly            bool          `json:"indexReadOnly"`
	EntityScope              string        `json:"entityScope,omitempty"`
	ColdPhaseEnabled         bool          `json:"coldPhaseEnabled"`
	ColdTimer                int           `json:"coldTimer,omitempty"`
	RolloverMaxAge           int           `json:"rolloverMaxAge,omitempty"`
	RolloverMaxDocs          int           `json:"rolloverMaxDocs,omitempty"`
	RolloverMaxSize          int           `json:"rolloverMaxSize,omitempty"`
	ForceMergeEnabled        bool          `json:"forceMergeEnabled"`
	ForceMergeMaxNumSegments int           `json:"forceMergeMaxNumSegments,omitempty"`
	EsIlmPolicyType          string        `json:"esIlmPolicyType,omitempty"`
	ExternalID               string        `json:"externalID,omitempty"`
}

// NewEsIlmPolicy returns a new *EsIlmPolicy
func NewEsIlmPolicy() *EsIlmPolicy {

	return &EsIlmPolicy{
		WarmPhaseEnabled:         false,
		WarmTimer:                0,
		DeletePhaseEnabled:       false,
		DeleteTimer:              168,
		IndexFreeze:              false,
		IndexReadOnly:            false,
		ColdPhaseEnabled:         false,
		ColdTimer:                0,
		RolloverMaxAge:           0,
		RolloverMaxDocs:          0,
		RolloverMaxSize:          150,
		ForceMergeEnabled:        false,
		ForceMergeMaxNumSegments: 1,
	}
}

// Identity returns the Identity of the object.
func (o *EsIlmPolicy) Identity() bambou.Identity {

	return EsIlmPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EsIlmPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EsIlmPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EsIlmPolicy from the server
func (o *EsIlmPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EsIlmPolicy into the server
func (o *EsIlmPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EsIlmPolicy from the server
func (o *EsIlmPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the EsIlmPolicy
func (o *EsIlmPolicy) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the EsIlmPolicy
func (o *EsIlmPolicy) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the EsIlmPolicy
func (o *EsIlmPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EsIlmPolicy
func (o *EsIlmPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EsIlmPolicy
func (o *EsIlmPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EsIlmPolicy
func (o *EsIlmPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
