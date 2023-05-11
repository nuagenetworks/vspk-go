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

// InfrastructureEVDFProfileIdentity represents the Identity of the object
var InfrastructureEVDFProfileIdentity = bambou.Identity {
    Name:     "infrastructureevdfprofile",
    Category: "infrastructureevdfprofiles",
}

// InfrastructureEVDFProfilesList represents a list of InfrastructureEVDFProfiles
type InfrastructureEVDFProfilesList []*InfrastructureEVDFProfile

// InfrastructureEVDFProfilesAncestor is the interface that an ancestor of a InfrastructureEVDFProfile must implement.
// An Ancestor is defined as an entity that has InfrastructureEVDFProfile as a descendant.
// An Ancestor can get a list of its child InfrastructureEVDFProfiles, but not necessarily create one.
type InfrastructureEVDFProfilesAncestor interface {
    InfrastructureEVDFProfiles(*bambou.FetchingInfo) (InfrastructureEVDFProfilesList, *bambou.Error)
}

// InfrastructureEVDFProfilesParent is the interface that a parent of a InfrastructureEVDFProfile must implement.
// A Parent is defined as an entity that has InfrastructureEVDFProfile as a child.
// A Parent is an Ancestor which can create a InfrastructureEVDFProfile.
type InfrastructureEVDFProfilesParent interface {
    InfrastructureEVDFProfilesAncestor
    CreateInfrastructureEVDFProfile(*InfrastructureEVDFProfile) (*bambou.Error)
}

// InfrastructureEVDFProfile represents the model of a infrastructureevdfprofile
type InfrastructureEVDFProfile struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    NTPServerKey string `json:"NTPServerKey,omitempty"`
    NTPServerKeyID int `json:"NTPServerKeyID,omitempty"`
    Name string `json:"name,omitempty"`
    LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
    LastUpdatedDate string `json:"lastUpdatedDate,omitempty"`
    ActiveController string `json:"activeController,omitempty"`
    ServiceIPv4Subnet string `json:"serviceIPv4Subnet,omitempty"`
    Description string `json:"description,omitempty"`
    EmbeddedMetadata []interface{} `json:"embeddedMetadata,omitempty"`
    EnterpriseID string `json:"enterpriseID,omitempty"`
    EntityScope string `json:"entityScope,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    ProxyDNSName string `json:"proxyDNSName,omitempty"`
    UseTwoFactor bool `json:"useTwoFactor"`
    StandbyController string `json:"standbyController,omitempty"`
    NuagePlatform string `json:"nuagePlatform,omitempty"`
    Owner string `json:"owner,omitempty"`
    ExternalID string `json:"externalID,omitempty"`
    
}

// NewInfrastructureEVDFProfile returns a new *InfrastructureEVDFProfile
func NewInfrastructureEVDFProfile() *InfrastructureEVDFProfile {

    return &InfrastructureEVDFProfile{
        ServiceIPv4Subnet: "0.0.0.0/8",
        UseTwoFactor: true,
        NuagePlatform: "KVM",
        }
}

// Identity returns the Identity of the object.
func (o *InfrastructureEVDFProfile) Identity() bambou.Identity {

    return InfrastructureEVDFProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructureEVDFProfile) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructureEVDFProfile) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the InfrastructureEVDFProfile from the server
func (o *InfrastructureEVDFProfile) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the InfrastructureEVDFProfile into the server
func (o *InfrastructureEVDFProfile) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the InfrastructureEVDFProfile from the server
func (o *InfrastructureEVDFProfile) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Metadatas retrieves the list of child Metadatas of the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

    var list MetadatasList
    err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
    return list, err
}



// CreateMetadata creates a new child Metadata under the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) CreateMetadata(child *Metadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// GlobalMetadatas retrieves the list of child GlobalMetadatas of the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

    var list GlobalMetadatasList
    err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
    return list, err
}



// CreateGlobalMetadata creates a new child GlobalMetadata under the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

