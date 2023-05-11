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

// EthernetSegmentGWGroupIdentity represents the Identity of the object
var EthernetSegmentGWGroupIdentity = bambou.Identity {
    Name:     "ethernetsegmentgwgroup",
    Category: "ethernetsegmentgwgroups",
}

// EthernetSegmentGWGroupsList represents a list of EthernetSegmentGWGroups
type EthernetSegmentGWGroupsList []*EthernetSegmentGWGroup

// EthernetSegmentGWGroupsAncestor is the interface that an ancestor of a EthernetSegmentGWGroup must implement.
// An Ancestor is defined as an entity that has EthernetSegmentGWGroup as a descendant.
// An Ancestor can get a list of its child EthernetSegmentGWGroups, but not necessarily create one.
type EthernetSegmentGWGroupsAncestor interface {
    EthernetSegmentGWGroups(*bambou.FetchingInfo) (EthernetSegmentGWGroupsList, *bambou.Error)
}

// EthernetSegmentGWGroupsParent is the interface that a parent of a EthernetSegmentGWGroup must implement.
// A Parent is defined as an entity that has EthernetSegmentGWGroup as a child.
// A Parent is an Ancestor which can create a EthernetSegmentGWGroup.
type EthernetSegmentGWGroupsParent interface {
    EthernetSegmentGWGroupsAncestor
    CreateEthernetSegmentGWGroup(*EthernetSegmentGWGroup) (*bambou.Error)
}

// EthernetSegmentGWGroup represents the model of a ethernetsegmentgwgroup
type EthernetSegmentGWGroup struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    Personality string `json:"personality,omitempty"`
    Description string `json:"description,omitempty"`
    AssocGatewaysNames []interface{} `json:"assocGatewaysNames,omitempty"`
    AssocGatewaysStatus []interface{} `json:"assocGatewaysStatus,omitempty"`
    AssociatedGatewayIDs []interface{} `json:"associatedGatewayIDs,omitempty"`
    
}

// NewEthernetSegmentGWGroup returns a new *EthernetSegmentGWGroup
func NewEthernetSegmentGWGroup() *EthernetSegmentGWGroup {

    return &EthernetSegmentGWGroup{
        Personality: "NETCONF_7X50",
        AssocGatewaysStatus: false,
        }
}

// Identity returns the Identity of the object.
func (o *EthernetSegmentGWGroup) Identity() bambou.Identity {

    return EthernetSegmentGWGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EthernetSegmentGWGroup) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EthernetSegmentGWGroup) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the EthernetSegmentGWGroup from the server
func (o *EthernetSegmentGWGroup) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EthernetSegmentGWGroup into the server
func (o *EthernetSegmentGWGroup) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EthernetSegmentGWGroup from the server
func (o *EthernetSegmentGWGroup) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// L2Domains retrieves the list of child L2Domains of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) L2Domains(info *bambou.FetchingInfo) (L2DomainsList, *bambou.Error) {

    var list L2DomainsList
    err := bambou.CurrentSession().FetchChildren(o, L2DomainIdentity, &list, info)
    return list, err
}




// MACFilterProfiles retrieves the list of child MACFilterProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) MACFilterProfiles(info *bambou.FetchingInfo) (MACFilterProfilesList, *bambou.Error) {

    var list MACFilterProfilesList
    err := bambou.CurrentSession().FetchChildren(o, MACFilterProfileIdentity, &list, info)
    return list, err
}




// SAPEgressQoSProfiles retrieves the list of child SAPEgressQoSProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) SAPEgressQoSProfiles(info *bambou.FetchingInfo) (SAPEgressQoSProfilesList, *bambou.Error) {

    var list SAPEgressQoSProfilesList
    err := bambou.CurrentSession().FetchChildren(o, SAPEgressQoSProfileIdentity, &list, info)
    return list, err
}




// SAPIngressQoSProfiles retrieves the list of child SAPIngressQoSProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) SAPIngressQoSProfiles(info *bambou.FetchingInfo) (SAPIngressQoSProfilesList, *bambou.Error) {

    var list SAPIngressQoSProfilesList
    err := bambou.CurrentSession().FetchChildren(o, SAPIngressQoSProfileIdentity, &list, info)
    return list, err
}




// DeploymentFailures retrieves the list of child DeploymentFailures of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) DeploymentFailures(info *bambou.FetchingInfo) (DeploymentFailuresList, *bambou.Error) {

    var list DeploymentFailuresList
    err := bambou.CurrentSession().FetchChildren(o, DeploymentFailureIdentity, &list, info)
    return list, err
}




// Permissions retrieves the list of child Permissions of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// EgressProfiles retrieves the list of child EgressProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) EgressProfiles(info *bambou.FetchingInfo) (EgressProfilesList, *bambou.Error) {

    var list EgressProfilesList
    err := bambou.CurrentSession().FetchChildren(o, EgressProfileIdentity, &list, info)
    return list, err
}



// CreateEgressProfile creates a new child EgressProfile under the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) CreateEgressProfile(child *EgressProfile) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Alarms retrieves the list of child Alarms of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

    var list AlarmsList
    err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
    return list, err
}




// IngressProfiles retrieves the list of child IngressProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) IngressProfiles(info *bambou.FetchingInfo) (IngressProfilesList, *bambou.Error) {

    var list IngressProfilesList
    err := bambou.CurrentSession().FetchChildren(o, IngressProfileIdentity, &list, info)
    return list, err
}



// CreateIngressProfile creates a new child IngressProfile under the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) CreateIngressProfile(child *IngressProfile) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// EnterprisePermissions retrieves the list of child EnterprisePermissions of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

    var list EnterprisePermissionsList
    err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
    return list, err
}



// CreateEnterprisePermission creates a new child EnterprisePermission under the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Jobs retrieves the list of child Jobs of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

    var list JobsList
    err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
    return list, err
}



// CreateJob creates a new child Job under the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) CreateJob(child *Job) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Domains retrieves the list of child Domains of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) Domains(info *bambou.FetchingInfo) (DomainsList, *bambou.Error) {

    var list DomainsList
    err := bambou.CurrentSession().FetchChildren(o, DomainIdentity, &list, info)
    return list, err
}




// IPFilterProfiles retrieves the list of child IPFilterProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) IPFilterProfiles(info *bambou.FetchingInfo) (IPFilterProfilesList, *bambou.Error) {

    var list IPFilterProfilesList
    err := bambou.CurrentSession().FetchChildren(o, IPFilterProfileIdentity, &list, info)
    return list, err
}




// IPv6FilterProfiles retrieves the list of child IPv6FilterProfiles of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) IPv6FilterProfiles(info *bambou.FetchingInfo) (IPv6FilterProfilesList, *bambou.Error) {

    var list IPv6FilterProfilesList
    err := bambou.CurrentSession().FetchChildren(o, IPv6FilterProfileIdentity, &list, info)
    return list, err
}




// EthernetSegmentGroups retrieves the list of child EthernetSegmentGroups of the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) EthernetSegmentGroups(info *bambou.FetchingInfo) (EthernetSegmentGroupsList, *bambou.Error) {

    var list EthernetSegmentGroupsList
    err := bambou.CurrentSession().FetchChildren(o, EthernetSegmentGroupIdentity, &list, info)
    return list, err
}



// CreateEthernetSegmentGroup creates a new child EthernetSegmentGroup under the EthernetSegmentGWGroup
func (o *EthernetSegmentGWGroup) CreateEthernetSegmentGroup(child *EthernetSegmentGroup) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

