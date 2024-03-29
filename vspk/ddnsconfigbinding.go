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

// DdnsconfigbindingIdentity represents the Identity of the object
var DdnsconfigbindingIdentity = bambou.Identity{
	Name:     "ddnsconfigbinding",
	Category: "ddnsconfigbindings",
}

// DdnsconfigbindingsList represents a list of Ddnsconfigbindings
type DdnsconfigbindingsList []*Ddnsconfigbinding

// DdnsconfigbindingsAncestor is the interface that an ancestor of a Ddnsconfigbinding must implement.
// An Ancestor is defined as an entity that has Ddnsconfigbinding as a descendant.
// An Ancestor can get a list of its child Ddnsconfigbindings, but not necessarily create one.
type DdnsconfigbindingsAncestor interface {
	Ddnsconfigbindings(*bambou.FetchingInfo) (DdnsconfigbindingsList, *bambou.Error)
}

// DdnsconfigbindingsParent is the interface that a parent of a Ddnsconfigbinding must implement.
// A Parent is defined as an entity that has Ddnsconfigbinding as a child.
// A Parent is an Ancestor which can create a Ddnsconfigbinding.
type DdnsconfigbindingsParent interface {
	DdnsconfigbindingsAncestor
	CreateDdnsconfigbinding(*Ddnsconfigbinding) *bambou.Error
}

// Ddnsconfigbinding represents the model of a ddnsconfigbinding
type Ddnsconfigbinding struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	UplinkName         string `json:"uplinkName,omitempty"`
	UplinkPriority     int    `json:"uplinkPriority,omitempty"`
	AssocDDNSConfigId  string `json:"assocDDNSConfigId,omitempty"`
	AssociatedUplinkID string `json:"associatedUplinkID,omitempty"`
}

// NewDdnsconfigbinding returns a new *Ddnsconfigbinding
func NewDdnsconfigbinding() *Ddnsconfigbinding {

	return &Ddnsconfigbinding{
		UplinkPriority: 10,
	}
}

// Identity returns the Identity of the object.
func (o *Ddnsconfigbinding) Identity() bambou.Identity {

	return DdnsconfigbindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Ddnsconfigbinding) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Ddnsconfigbinding) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Ddnsconfigbinding from the server
func (o *Ddnsconfigbinding) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Ddnsconfigbinding into the server
func (o *Ddnsconfigbinding) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Ddnsconfigbinding from the server
func (o *Ddnsconfigbinding) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
