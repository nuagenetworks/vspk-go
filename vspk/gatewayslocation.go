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

// GatewaysLocationIdentity represents the Identity of the object
var GatewaysLocationIdentity = bambou.Identity{
	Name:     "gatewayslocation",
	Category: "gatewayslocations",
}

// GatewaysLocationsList represents a list of GatewaysLocations
type GatewaysLocationsList []*GatewaysLocation

// GatewaysLocationsAncestor is the interface that an ancestor of a GatewaysLocation must implement.
// An Ancestor is defined as an entity that has GatewaysLocation as a descendant.
// An Ancestor can get a list of its child GatewaysLocations, but not necessarily create one.
type GatewaysLocationsAncestor interface {
	GatewaysLocations(*bambou.FetchingInfo) (GatewaysLocationsList, *bambou.Error)
}

// GatewaysLocationsParent is the interface that a parent of a GatewaysLocation must implement.
// A Parent is defined as an entity that has GatewaysLocation as a child.
// A Parent is an Ancestor which can create a GatewaysLocation.
type GatewaysLocationsParent interface {
	GatewaysLocationsAncestor
	CreateGatewaysLocation(*GatewaysLocation) *bambou.Error
}

// GatewaysLocation represents the model of a gatewayslocation
type GatewaysLocation struct {
	ID                   string        `json:"ID,omitempty"`
	ParentID             string        `json:"parentID,omitempty"`
	ParentType           string        `json:"parentType,omitempty"`
	Owner                string        `json:"owner,omitempty"`
	LastUpdatedBy        string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate      string        `json:"lastUpdatedDate,omitempty"`
	Latitude             float64       `json:"latitude,omitempty"`
	Address              string        `json:"address,omitempty"`
	IgnoreGeocode        bool          `json:"ignoreGeocode"`
	TimeZoneID           string        `json:"timeZoneID,omitempty"`
	EmbeddedMetadata     []interface{} `json:"embeddedMetadata,omitempty"`
	EntityScope          string        `json:"entityScope,omitempty"`
	Locality             string        `json:"locality,omitempty"`
	Longitude            float64       `json:"longitude,omitempty"`
	Country              string        `json:"country,omitempty"`
	CreationDate         string        `json:"creationDate,omitempty"`
	AssociatedEntityName string        `json:"associatedEntityName,omitempty"`
	AssociatedEntityType string        `json:"associatedEntityType,omitempty"`
	State                string        `json:"state,omitempty"`
	Owner                string        `json:"owner,omitempty"`
	ExternalID           string        `json:"externalID,omitempty"`
}

// NewGatewaysLocation returns a new *GatewaysLocation
func NewGatewaysLocation() *GatewaysLocation {

	return &GatewaysLocation{
		TimeZoneID: "UTC",
	}
}

// Identity returns the Identity of the object.
func (o *GatewaysLocation) Identity() bambou.Identity {

	return GatewaysLocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewaysLocation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewaysLocation) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewaysLocation from the server
func (o *GatewaysLocation) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewaysLocation into the server
func (o *GatewaysLocation) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewaysLocation from the server
func (o *GatewaysLocation) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GatewaysLocation
func (o *GatewaysLocation) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GatewaysLocation
func (o *GatewaysLocation) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GatewaysLocation
func (o *GatewaysLocation) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GatewaysLocation
func (o *GatewaysLocation) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
