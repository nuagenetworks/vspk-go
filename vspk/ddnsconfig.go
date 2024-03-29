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

// DdnsconfigIdentity represents the Identity of the object
var DdnsconfigIdentity = bambou.Identity{
	Name:     "ddnsconfig",
	Category: "ddnsconfigs",
}

// DdnsconfigsList represents a list of Ddnsconfigs
type DdnsconfigsList []*Ddnsconfig

// DdnsconfigsAncestor is the interface that an ancestor of a Ddnsconfig must implement.
// An Ancestor is defined as an entity that has Ddnsconfig as a descendant.
// An Ancestor can get a list of its child Ddnsconfigs, but not necessarily create one.
type DdnsconfigsAncestor interface {
	Ddnsconfigs(*bambou.FetchingInfo) (DdnsconfigsList, *bambou.Error)
}

// DdnsconfigsParent is the interface that a parent of a Ddnsconfig must implement.
// A Parent is defined as an entity that has Ddnsconfig as a child.
// A Parent is an Ancestor which can create a Ddnsconfig.
type DdnsconfigsParent interface {
	DdnsconfigsAncestor
	CreateDdnsconfig(*Ddnsconfig) *bambou.Error
}

// Ddnsconfig represents the model of a ddnsconfig
type Ddnsconfig struct {
	ID               string `json:"ID,omitempty"`
	ParentID         string `json:"parentID,omitempty"`
	ParentType       string `json:"parentType,omitempty"`
	Owner            string `json:"owner,omitempty"`
	Password         string `json:"password,omitempty"`
	EnableDDNSConfig bool   `json:"enableDDNSConfig"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Hostname         string `json:"hostname,omitempty"`
	ProviderName     string `json:"providerName,omitempty"`
	Username         string `json:"username,omitempty"`
	AssocGatewayId   string `json:"assocGatewayId,omitempty"`
}

// NewDdnsconfig returns a new *Ddnsconfig
func NewDdnsconfig() *Ddnsconfig {

	return &Ddnsconfig{
		EnableDDNSConfig: true,
		ConnectionStatus: "UNKNOWN",
		ProviderName:     "DYN_DNS",
	}
}

// Identity returns the Identity of the object.
func (o *Ddnsconfig) Identity() bambou.Identity {

	return DdnsconfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Ddnsconfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Ddnsconfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Ddnsconfig from the server
func (o *Ddnsconfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Ddnsconfig into the server
func (o *Ddnsconfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Ddnsconfig from the server
func (o *Ddnsconfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Ddnsconfigbindings retrieves the list of child Ddnsconfigbindings of the Ddnsconfig
func (o *Ddnsconfig) Ddnsconfigbindings(info *bambou.FetchingInfo) (DdnsconfigbindingsList, *bambou.Error) {

	var list DdnsconfigbindingsList
	err := bambou.CurrentSession().FetchChildren(o, DdnsconfigbindingIdentity, &list, info)
	return list, err
}

// CreateDdnsconfigbinding creates a new child Ddnsconfigbinding under the Ddnsconfig
func (o *Ddnsconfig) CreateDdnsconfigbinding(child *Ddnsconfigbinding) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
