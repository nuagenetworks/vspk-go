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

// NetconfGlobalConfigurationIdentity represents the Identity of the object
var NetconfGlobalConfigurationIdentity = bambou.Identity{
	Name:     "netconfglobalconfiguration",
	Category: "netconfglobalconfigurations",
}

// NetconfGlobalConfigurationsList represents a list of NetconfGlobalConfigurations
type NetconfGlobalConfigurationsList []*NetconfGlobalConfiguration

// NetconfGlobalConfigurationsAncestor is the interface that an ancestor of a NetconfGlobalConfiguration must implement.
// An Ancestor is defined as an entity that has NetconfGlobalConfiguration as a descendant.
// An Ancestor can get a list of its child NetconfGlobalConfigurations, but not necessarily create one.
type NetconfGlobalConfigurationsAncestor interface {
	NetconfGlobalConfigurations(*bambou.FetchingInfo) (NetconfGlobalConfigurationsList, *bambou.Error)
}

// NetconfGlobalConfigurationsParent is the interface that a parent of a NetconfGlobalConfiguration must implement.
// A Parent is defined as an entity that has NetconfGlobalConfiguration as a child.
// A Parent is an Ancestor which can create a NetconfGlobalConfiguration.
type NetconfGlobalConfigurationsParent interface {
	NetconfGlobalConfigurationsAncestor
	CreateNetconfGlobalConfiguration(*NetconfGlobalConfiguration) *bambou.Error
}

// NetconfGlobalConfiguration represents the model of a netconfglobalconfiguration
type NetconfGlobalConfiguration struct {
	ID                string        `json:"ID,omitempty"`
	ParentID          string        `json:"parentID,omitempty"`
	ParentType        string        `json:"parentType,omitempty"`
	Owner             string        `json:"owner,omitempty"`
	Name              string        `json:"name,omitempty"`
	LastUpdatedBy     string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate   string        `json:"lastUpdatedDate,omitempty"`
	Description       string        `json:"description,omitempty"`
	NetconfGatewayIDs []interface{} `json:"netconfGatewayIDs,omitempty"`
	EmbeddedMetadata  []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope       string        `json:"entityScope,omitempty"`
	ConfigDefinition  string        `json:"configDefinition,omitempty"`
	CreationDate      string        `json:"creationDate,omitempty"`
	Owner             string        `json:"owner,omitempty"`
	ExternalID        string        `json:"externalID,omitempty"`
}

// NewNetconfGlobalConfiguration returns a new *NetconfGlobalConfiguration
func NewNetconfGlobalConfiguration() *NetconfGlobalConfiguration {

	return &NetconfGlobalConfiguration{}
}

// Identity returns the Identity of the object.
func (o *NetconfGlobalConfiguration) Identity() bambou.Identity {

	return NetconfGlobalConfigurationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetconfGlobalConfiguration) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetconfGlobalConfiguration) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetconfGlobalConfiguration from the server
func (o *NetconfGlobalConfiguration) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetconfGlobalConfiguration into the server
func (o *NetconfGlobalConfiguration) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetconfGlobalConfiguration from the server
func (o *NetconfGlobalConfiguration) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// DeploymentFailures retrieves the list of child DeploymentFailures of the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) DeploymentFailures(info *bambou.FetchingInfo) (DeploymentFailuresList, *bambou.Error) {

	var list DeploymentFailuresList
	err := bambou.CurrentSession().FetchChildren(o, DeploymentFailureIdentity, &list, info)
	return list, err
}

// Permissions retrieves the list of child Permissions of the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NetconfGlobalConfiguration
func (o *NetconfGlobalConfiguration) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
