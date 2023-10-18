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

// VMIPReservationIdentity represents the Identity of the object
var VMIPReservationIdentity = bambou.Identity{
	Name:     "vmipreservation",
	Category: "vmipreservations",
}

// VMIPReservationsList represents a list of VMIPReservations
type VMIPReservationsList []*VMIPReservation

// VMIPReservationsAncestor is the interface that an ancestor of a VMIPReservation must implement.
// An Ancestor is defined as an entity that has VMIPReservation as a descendant.
// An Ancestor can get a list of its child VMIPReservations, but not necessarily create one.
type VMIPReservationsAncestor interface {
	VMIPReservations(*bambou.FetchingInfo) (VMIPReservationsList, *bambou.Error)
}

// VMIPReservationsParent is the interface that a parent of a VMIPReservation must implement.
// A Parent is defined as an entity that has VMIPReservation as a child.
// A Parent is an Ancestor which can create a VMIPReservation.
type VMIPReservationsParent interface {
	VMIPReservationsAncestor
	CreateVMIPReservation(*VMIPReservation) *bambou.Error
}

// VMIPReservation represents the model of a vmipreservation
type VMIPReservation struct {
	ID                  string        `json:"ID,omitempty"`
	ParentID            string        `json:"parentID,omitempty"`
	ParentType          string        `json:"parentType,omitempty"`
	Owner               string        `json:"owner,omitempty"`
	IPType              string        `json:"IPType,omitempty"`
	IPV4Address         string        `json:"IPV4Address,omitempty"`
	IPV6Address         string        `json:"IPV6Address,omitempty"`
	IPV6AllocationPools []interface{} `json:"IPV6AllocationPools,omitempty"`
	LastUpdatedBy       string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate     string        `json:"lastUpdatedDate,omitempty"`
	AllocationPools     []interface{} `json:"allocationPools,omitempty"`
	EmbeddedMetadata    []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope         string        `json:"entityScope,omitempty"`
	CreationDate        string        `json:"creationDate,omitempty"`
	State               string        `json:"state,omitempty"`
	Owner               string        `json:"owner,omitempty"`
	ExternalID          string        `json:"externalID,omitempty"`
}

// NewVMIPReservation returns a new *VMIPReservation
func NewVMIPReservation() *VMIPReservation {

	return &VMIPReservation{
		State: "UNASSIGNED",
	}
}

// Identity returns the Identity of the object.
func (o *VMIPReservation) Identity() bambou.Identity {

	return VMIPReservationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VMIPReservation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VMIPReservation) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VMIPReservation from the server
func (o *VMIPReservation) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VMIPReservation into the server
func (o *VMIPReservation) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VMIPReservation from the server
func (o *VMIPReservation) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the VMIPReservation
func (o *VMIPReservation) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the VMIPReservation
func (o *VMIPReservation) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the VMIPReservation
func (o *VMIPReservation) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VMIPReservation
func (o *VMIPReservation) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VMIPReservation
func (o *VMIPReservation) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VMIPReservation
func (o *VMIPReservation) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
