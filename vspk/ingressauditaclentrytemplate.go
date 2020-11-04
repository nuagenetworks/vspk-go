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

// IngressAuditACLEntryTemplateIdentity represents the Identity of the object
var IngressAuditACLEntryTemplateIdentity = bambou.Identity{
	Name:     "ingressauditaclentrytemplate",
	Category: "ingressauditaclentrytemplates",
}

// IngressAuditACLEntryTemplatesList represents a list of IngressAuditACLEntryTemplates
type IngressAuditACLEntryTemplatesList []*IngressAuditACLEntryTemplate

// IngressAuditACLEntryTemplatesAncestor is the interface that an ancestor of a IngressAuditACLEntryTemplate must implement.
// An Ancestor is defined as an entity that has IngressAuditACLEntryTemplate as a descendant.
// An Ancestor can get a list of its child IngressAuditACLEntryTemplates, but not necessarily create one.
type IngressAuditACLEntryTemplatesAncestor interface {
	IngressAuditACLEntryTemplates(*bambou.FetchingInfo) (IngressAuditACLEntryTemplatesList, *bambou.Error)
}

// IngressAuditACLEntryTemplatesParent is the interface that a parent of a IngressAuditACLEntryTemplate must implement.
// A Parent is defined as an entity that has IngressAuditACLEntryTemplate as a child.
// A Parent is an Ancestor which can create a IngressAuditACLEntryTemplate.
type IngressAuditACLEntryTemplatesParent interface {
	IngressAuditACLEntryTemplatesAncestor
	CreateIngressAuditACLEntryTemplate(*IngressAuditACLEntryTemplate) *bambou.Error
}

// IngressAuditACLEntryTemplate represents the model of a ingressauditaclentrytemplate
type IngressAuditACLEntryTemplate struct {
	ID                                 string        `json:"ID,omitempty"`
	ParentID                           string        `json:"parentID,omitempty"`
	ParentType                         string        `json:"parentType,omitempty"`
	Owner                              string        `json:"owner,omitempty"`
	ACLTemplateName                    string        `json:"ACLTemplateName,omitempty"`
	ICMPCode                           string        `json:"ICMPCode,omitempty"`
	ICMPType                           string        `json:"ICMPType,omitempty"`
	IPv6AddressOverride                string        `json:"IPv6AddressOverride,omitempty"`
	DSCP                               string        `json:"DSCP,omitempty"`
	LastUpdatedBy                      string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                    string        `json:"lastUpdatedDate,omitempty"`
	Action                             string        `json:"action,omitempty"`
	AddressOverride                    string        `json:"addressOverride,omitempty"`
	WebFilterID                        string        `json:"webFilterID,omitempty"`
	WebFilterStatsLoggingEnabled       bool          `json:"webFilterStatsLoggingEnabled"`
	WebFilterType                      string        `json:"webFilterType,omitempty"`
	Description                        string        `json:"description,omitempty"`
	DestinationPort                    string        `json:"destinationPort,omitempty"`
	NetworkEntityType                  string        `json:"networkEntityType,omitempty"`
	NetworkID                          string        `json:"networkID,omitempty"`
	NetworkType                        string        `json:"networkType,omitempty"`
	MirrorDestinationGroupID           string        `json:"mirrorDestinationGroupID,omitempty"`
	MirrorDestinationID                string        `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled                 bool          `json:"flowLoggingEnabled"`
	EmbeddedMetadata                   []interface{} `json:"embeddedMetadata,omitempty"`
	EnterpriseName                     string        `json:"enterpriseName,omitempty"`
	EntityScope                        string        `json:"entityScope,omitempty"`
	LocationEntityType                 string        `json:"locationEntityType,omitempty"`
	LocationID                         string        `json:"locationID,omitempty"`
	LocationType                       string        `json:"locationType,omitempty"`
	PolicyState                        string        `json:"policyState,omitempty"`
	DomainName                         string        `json:"domainName,omitempty"`
	SourcePort                         string        `json:"sourcePort,omitempty"`
	CreationDate                       string        `json:"creationDate,omitempty"`
	Priority                           int           `json:"priority,omitempty"`
	Protocol                           string        `json:"protocol,omitempty"`
	AssociatedL7ApplicationSignatureID string        `json:"associatedL7ApplicationSignatureID,omitempty"`
	AssociatedLiveEntityID             string        `json:"associatedLiveEntityID,omitempty"`
	AssociatedLiveTemplateID           string        `json:"associatedLiveTemplateID,omitempty"`
	AssociatedTrafficType              string        `json:"associatedTrafficType,omitempty"`
	AssociatedTrafficTypeID            string        `json:"associatedTrafficTypeID,omitempty"`
	AssociatedVirtualFirewallRuleID    string        `json:"associatedVirtualFirewallRuleID,omitempty"`
	Stateful                           bool          `json:"stateful"`
	StatsID                            string        `json:"statsID,omitempty"`
	StatsLoggingEnabled                bool          `json:"statsLoggingEnabled"`
	EtherType                          string        `json:"etherType,omitempty"`
	OverlayMirrorDestinationID         string        `json:"overlayMirrorDestinationID,omitempty"`
	Owner                              string        `json:"owner,omitempty"`
	ExternalID                         string        `json:"externalID,omitempty"`
}

// NewIngressAuditACLEntryTemplate returns a new *IngressAuditACLEntryTemplate
func NewIngressAuditACLEntryTemplate() *IngressAuditACLEntryTemplate {

	return &IngressAuditACLEntryTemplate{
		Action:                       "TRANSPARENT",
		WebFilterStatsLoggingEnabled: false,
	}
}

// Identity returns the Identity of the object.
func (o *IngressAuditACLEntryTemplate) Identity() bambou.Identity {

	return IngressAuditACLEntryTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressAuditACLEntryTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressAuditACLEntryTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressAuditACLEntryTemplate from the server
func (o *IngressAuditACLEntryTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressAuditACLEntryTemplate into the server
func (o *IngressAuditACLEntryTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressAuditACLEntryTemplate from the server
func (o *IngressAuditACLEntryTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the IngressAuditACLEntryTemplate
func (o *IngressAuditACLEntryTemplate) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}
