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

// EthernetSegmentGroupIdentity represents the Identity of the object
var EthernetSegmentGroupIdentity = bambou.Identity {
    Name:     "ethernetsegmentgroup",
    Category: "ethernetsegmentgroups",
}

// EthernetSegmentGroupsList represents a list of EthernetSegmentGroups
type EthernetSegmentGroupsList []*EthernetSegmentGroup

// EthernetSegmentGroupsAncestor is the interface that an ancestor of a EthernetSegmentGroup must implement.
// An Ancestor is defined as an entity that has EthernetSegmentGroup as a descendant.
// An Ancestor can get a list of its child EthernetSegmentGroups, but not necessarily create one.
type EthernetSegmentGroupsAncestor interface {
    EthernetSegmentGroups(*bambou.FetchingInfo) (EthernetSegmentGroupsList, *bambou.Error)
}

// EthernetSegmentGroupsParent is the interface that a parent of a EthernetSegmentGroup must implement.
// A Parent is defined as an entity that has EthernetSegmentGroup as a child.
// A Parent is an Ancestor which can create a EthernetSegmentGroup.
type EthernetSegmentGroupsParent interface {
    EthernetSegmentGroupsAncestor
    CreateEthernetSegmentGroup(*EthernetSegmentGroup) (*bambou.Error)
}

// EthernetSegmentGroup represents the model of a ethernetsegmentgroup
type EthernetSegmentGroup struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    VLANRange string `json:"VLANRange,omitempty"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Virtual bool `json:"virtual"`
    PortType string `json:"portType,omitempty"`
    EthernetSegmentID string `json:"ethernetSegmentID,omitempty"`
    
}

// NewEthernetSegmentGroup returns a new *EthernetSegmentGroup
func NewEthernetSegmentGroup() *EthernetSegmentGroup {

    return &EthernetSegmentGroup{
        Virtual: false,
        }
}

// Identity returns the Identity of the object.
func (o *EthernetSegmentGroup) Identity() bambou.Identity {

    return EthernetSegmentGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EthernetSegmentGroup) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EthernetSegmentGroup) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the EthernetSegmentGroup from the server
func (o *EthernetSegmentGroup) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EthernetSegmentGroup into the server
func (o *EthernetSegmentGroup) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EthernetSegmentGroup from the server
func (o *EthernetSegmentGroup) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Permissions retrieves the list of child Permissions of the EthernetSegmentGroup
func (o *EthernetSegmentGroup) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

    var list PermissionsList
    err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
    return list, err
}



// CreatePermission creates a new child Permission under the EthernetSegmentGroup
func (o *EthernetSegmentGroup) CreatePermission(child *Permission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// VLANs retrieves the list of child VLANs of the EthernetSegmentGroup
func (o *EthernetSegmentGroup) VLANs(info *bambou.FetchingInfo) (VLANsList, *bambou.Error) {

    var list VLANsList
    err := bambou.CurrentSession().FetchChildren(o, VLANIdentity, &list, info)
    return list, err
}



// CreateVLAN creates a new child VLAN under the EthernetSegmentGroup
func (o *EthernetSegmentGroup) CreateVLAN(child *VLAN) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Alarms retrieves the list of child Alarms of the EthernetSegmentGroup
func (o *EthernetSegmentGroup) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

    var list AlarmsList
    err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
    return list, err
}




// EnterprisePermissions retrieves the list of child EnterprisePermissions of the EthernetSegmentGroup
func (o *EthernetSegmentGroup) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

    var list EnterprisePermissionsList
    err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
    return list, err
}



// CreateEnterprisePermission creates a new child EnterprisePermission under the EthernetSegmentGroup
func (o *EthernetSegmentGroup) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

