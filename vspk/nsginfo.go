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

import "github.com/FlorianOtel/go-bambou/bambou"

// NSGInfoIdentity represents the Identity of the object
var NSGInfoIdentity = bambou.Identity{
	Name:     "nsginfo",
	Category: "nsginfos",
}

// NSGInfosList represents a list of NSGInfos
type NSGInfosList []*NSGInfo

// NSGInfosAncestor is the interface of an ancestor of a NSGInfo must implement.
type NSGInfosAncestor interface {
	NSGInfos(*bambou.FetchingInfo) (NSGInfosList, *bambou.Error)
	CreateNSGInfos(*NSGInfo) *bambou.Error
}

// NSGInfo represents the model of a nsginfo
type NSGInfo struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	MACAddress            string `json:"MACAddress,omitempty"`
	SKU                   string `json:"SKU,omitempty"`
	TPMStatus             string `json:"TPMStatus,omitempty"`
	CPUType               string `json:"CPUType,omitempty"`
	NSGVersion            string `json:"NSGVersion,omitempty"`
	UUID                  string `json:"UUID,omitempty"`
	Family                string `json:"family,omitempty"`
	SerialNumber          string `json:"serialNumber,omitempty"`
	Libraries             string `json:"libraries,omitempty"`
	EntityScope           string `json:"entityScope,omitempty"`
	AssociatedNSGatewayID string `json:"associatedNSGatewayID,omitempty"`
	ExternalID            string `json:"externalID,omitempty"`
}

// NewNSGInfo returns a new *NSGInfo
func NewNSGInfo() *NSGInfo {

	return &NSGInfo{}
}

// Identity returns the Identity of the object.
func (o *NSGInfo) Identity() bambou.Identity {

	return NSGInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGInfo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGInfo) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGInfo from the server
func (o *NSGInfo) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGInfo into the server
func (o *NSGInfo) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGInfo from the server
func (o *NSGInfo) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
