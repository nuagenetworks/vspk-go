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

// IngressAuditACLTemplateIdentity represents the Identity of the object
var IngressAuditACLTemplateIdentity = bambou.Identity {
    Name:     "ingressauditacltemplate",
    Category: "ingressauditacltemplates",
}

// IngressAuditACLTemplatesList represents a list of IngressAuditACLTemplates
type IngressAuditACLTemplatesList []*IngressAuditACLTemplate

// IngressAuditACLTemplatesAncestor is the interface that an ancestor of a IngressAuditACLTemplate must implement.
// An Ancestor is defined as an entity that has IngressAuditACLTemplate as a descendant.
// An Ancestor can get a list of its child IngressAuditACLTemplates, but not necessarily create one.
type IngressAuditACLTemplatesAncestor interface {
    IngressAuditACLTemplates(*bambou.FetchingInfo) (IngressAuditACLTemplatesList, *bambou.Error)
}

// IngressAuditACLTemplatesParent is the interface that a parent of a IngressAuditACLTemplate must implement.
// A Parent is defined as an entity that has IngressAuditACLTemplate as a child.
// A Parent is an Ancestor which can create a IngressAuditACLTemplate.
type IngressAuditACLTemplatesParent interface {
    IngressAuditACLTemplatesAncestor
    CreateIngressAuditACLTemplate(*IngressAuditACLTemplate) (*bambou.Error)
}

// IngressAuditACLTemplate represents the model of a ingressauditacltemplate
type IngressAuditACLTemplate struct {
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
    Description string `json:"description,omitempty"`
    AllowAddressSpoof bool `json:"allowAddressSpoof"`
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

// NewIngressAuditACLTemplate returns a new *IngressAuditACLTemplate
func NewIngressAuditACLTemplate() *IngressAuditACLTemplate {

    return &IngressAuditACLTemplate{
        DefaultAllowNonIP: true,
        }
}

// Identity returns the Identity of the object.
func (o *IngressAuditACLTemplate) Identity() bambou.Identity {

    return IngressAuditACLTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressAuditACLTemplate) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressAuditACLTemplate) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the IngressAuditACLTemplate from the server
func (o *IngressAuditACLTemplate) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressAuditACLTemplate into the server
func (o *IngressAuditACLTemplate) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressAuditACLTemplate from the server
func (o *IngressAuditACLTemplate) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Metadatas retrieves the list of child Metadatas of the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// IngressAuditACLEntryTemplates retrieves the list of child IngressAuditACLEntryTemplates of the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) IngressAuditACLEntryTemplates(info *bambou.FetchingInfo) (IngressAuditACLEntryTemplatesList, *bambou.Error) {

    var list IngressAuditACLEntryTemplatesList
    err := bambou.CurrentSession().FetchChildren(o, IngressAuditACLEntryTemplateIdentity, &list, info)
    return list, err
}



// CreateIngressAuditACLEntryTemplate creates a new child IngressAuditACLEntryTemplate under the IngressAuditACLTemplate
func (o *IngressAuditACLTemplate) CreateIngressAuditACLEntryTemplate(child *IngressAuditACLEntryTemplate) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

