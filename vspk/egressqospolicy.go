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

// EgressQOSPolicyIdentity represents the Identity of the object
var EgressQOSPolicyIdentity = bambou.Identity{
	Name:     "egressqospolicy",
	Category: "egressqospolicies",
}

// EgressQOSPoliciesList represents a list of EgressQOSPolicies
type EgressQOSPoliciesList []*EgressQOSPolicy

// EgressQOSPoliciesAncestor is the interface that an ancestor of a EgressQOSPolicy must implement.
// An Ancestor is defined as an entity that has EgressQOSPolicy as a descendant.
// An Ancestor can get a list of its child EgressQOSPolicies, but not necessarily create one.
type EgressQOSPoliciesAncestor interface {
	EgressQOSPolicies(*bambou.FetchingInfo) (EgressQOSPoliciesList, *bambou.Error)
}

// EgressQOSPoliciesParent is the interface that a parent of a EgressQOSPolicy must implement.
// A Parent is defined as an entity that has EgressQOSPolicy as a child.
// A Parent is an Ancestor which can create a EgressQOSPolicy.
type EgressQOSPoliciesParent interface {
	EgressQOSPoliciesAncestor
	CreateEgressQOSPolicy(*EgressQOSPolicy) *bambou.Error
}

// EgressQOSPolicy represents the model of a egressqospolicy
type EgressQOSPolicy struct {
	ID                                      string        `json:"ID,omitempty"`
	ParentID                                string        `json:"parentID,omitempty"`
	ParentType                              string        `json:"parentType,omitempty"`
	Owner                                   string        `json:"owner,omitempty"`
	Name                                    string        `json:"name,omitempty"`
	ParentQueueAssociatedRateLimiterID      string        `json:"parentQueueAssociatedRateLimiterID,omitempty"`
	LastUpdatedBy                           string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                         string        `json:"lastUpdatedDate,omitempty"`
	DefaultServiceClass                     string        `json:"defaultServiceClass,omitempty"`
	Description                             string        `json:"description,omitempty"`
	NetworkCtrlQueueAssociatedRateLimiterID string        `json:"networkCtrlQueueAssociatedRateLimiterID,omitempty"`
	MgmtQueueAssociatedRateLimiterID        string        `json:"mgmtQueueAssociatedRateLimiterID,omitempty"`
	EmbeddedMetadata                        []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope                             string        `json:"entityScope,omitempty"`
	CreationDate                            string        `json:"creationDate,omitempty"`
	AssocEgressQosId                        string        `json:"assocEgressQosId,omitempty"`
	AssociatedCOSRemarkingPolicyTableID     string        `json:"associatedCOSRemarkingPolicyTableID,omitempty"`
	AssociatedDSCPRemarkingPolicyTableID    string        `json:"associatedDSCPRemarkingPolicyTableID,omitempty"`
	Queue1AssociatedRateLimiterID           string        `json:"queue1AssociatedRateLimiterID,omitempty"`
	Queue1ForwardingClasses                 []interface{} `json:"queue1ForwardingClasses,omitempty"`
	Queue2AssociatedRateLimiterID           string        `json:"queue2AssociatedRateLimiterID,omitempty"`
	Queue2ForwardingClasses                 []interface{} `json:"queue2ForwardingClasses,omitempty"`
	Queue3AssociatedRateLimiterID           string        `json:"queue3AssociatedRateLimiterID,omitempty"`
	Queue3ForwardingClasses                 []interface{} `json:"queue3ForwardingClasses,omitempty"`
	Queue4AssociatedRateLimiterID           string        `json:"queue4AssociatedRateLimiterID,omitempty"`
	Queue4ForwardingClasses                 []interface{} `json:"queue4ForwardingClasses,omitempty"`
	CustomSpqDepth                          int           `json:"customSpqDepth,omitempty"`
	Owner                                   string        `json:"owner,omitempty"`
	ExternalID                              string        `json:"externalID,omitempty"`
}

// NewEgressQOSPolicy returns a new *EgressQOSPolicy
func NewEgressQOSPolicy() *EgressQOSPolicy {

	return &EgressQOSPolicy{
		CustomSpqDepth: 0,
	}
}

// Identity returns the Identity of the object.
func (o *EgressQOSPolicy) Identity() bambou.Identity {

	return EgressQOSPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressQOSPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressQOSPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressQOSPolicy from the server
func (o *EgressQOSPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressQOSPolicy into the server
func (o *EgressQOSPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressQOSPolicy from the server
func (o *EgressQOSPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the EgressQOSPolicy
func (o *EgressQOSPolicy) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the EgressQOSPolicy
func (o *EgressQOSPolicy) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the EgressQOSPolicy
func (o *EgressQOSPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressQOSPolicy
func (o *EgressQOSPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressQOSPolicy
func (o *EgressQOSPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressQOSPolicy
func (o *EgressQOSPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
