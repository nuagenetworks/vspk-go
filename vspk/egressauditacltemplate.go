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

// EgressAuditACLTemplateIdentity represents the Identity of the object
var EgressAuditACLTemplateIdentity = bambou.Identity {
    Name:     "egressauditacltemplate",
    Category: "egressauditacltemplates",
}

// EgressAuditACLTemplatesList represents a list of EgressAuditACLTemplates
type EgressAuditACLTemplatesList []*EgressAuditACLTemplate

// EgressAuditACLTemplatesAncestor is the interface that an ancestor of a EgressAuditACLTemplate must implement.
// An Ancestor is defined as an entity that has EgressAuditACLTemplate as a descendant.
// An Ancestor can get a list of its child EgressAuditACLTemplates, but not necessarily create one.
type EgressAuditACLTemplatesAncestor interface {
    EgressAuditACLTemplates(*bambou.FetchingInfo) (EgressAuditACLTemplatesList, *bambou.Error)
}

// EgressAuditACLTemplatesParent is the interface that a parent of a EgressAuditACLTemplate must implement.
// A Parent is defined as an entity that has EgressAuditACLTemplate as a child.
// A Parent is an Ancestor which can create a EgressAuditACLTemplate.
type EgressAuditACLTemplatesParent interface {
    EgressAuditACLTemplatesAncestor
    CreateEgressAuditACLTemplate(*EgressAuditACLTemplate) (*bambou.Error)
}

// EgressAuditACLTemplate represents the model of a egressauditacltemplate
type EgressAuditACLTemplate struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
    LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`
    Active bool `json:"active"`
    DefaultAllowIP bool `json:"defaultAllowIP"`
    DefaultAllowNonIP bool `json:"defaultAllowNonIP"`
    DefaultInstallACLImplicitRules bool `json:"defaultInstallACLImplicitRules"`
    Description string `json:"description,omitempty"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    PolicyState string `json:"policyState,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    Priority int `json:"priority,omitempty"`
    PriorityType string `json:"priorityType,omitempty"`
    AssociatedLiveEntityID string `json:"associatedLiveEntityID,omitempty"`
    AssociatedVirtualFirewallPolicyID string `json:"associatedVirtualFirewallPolicyID,omitempty"`
    AutoGeneratePriority bool `json:"autoGeneratePriority"`
    Owner string `json:"owner,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    
}

// NewEgressAuditACLTemplate returns a new *EgressAuditACLTemplate
func NewEgressAuditACLTemplate() *EgressAuditACLTemplate {

    return &EgressAuditACLTemplate{
        }
}

// Identity returns the Identity of the object.
func (o *EgressAuditACLTemplate) Identity() bambou.Identity {

    return EgressAuditACLTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressAuditACLTemplate) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressAuditACLTemplate) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the EgressAuditACLTemplate from the server
func (o *EgressAuditACLTemplate) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressAuditACLTemplate into the server
func (o *EgressAuditACLTemplate) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressAuditACLTemplate from the server
func (o *EgressAuditACLTemplate) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Metadatas retrieves the list of child Metadatas of the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// EgressAuditACLEntryTemplates retrieves the list of child EgressAuditACLEntryTemplates of the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) EgressAuditACLEntryTemplates(info *bambou.FetchingInfo) (EgressAuditACLEntryTemplatesList, *bambou.Error) {

    var list EgressAuditACLEntryTemplatesList
    err := bambou.CurrentSession().FetchChildren(o, EgressAuditACLEntryTemplateIdentity, &list, info)
    return list, err
}



// CreateEgressAuditACLEntryTemplate creates a new child EgressAuditACLEntryTemplate under the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) CreateEgressAuditACLEntryTemplate(child *EgressAuditACLEntryTemplate) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressAuditACLTemplate
func (o *EgressAuditACLTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

