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

// ThreatPreventionServerConnectionIdentity represents the Identity of the object
var ThreatPreventionServerConnectionIdentity = bambou.Identity{
	Name:     "threatpreventionserverconnection",
	Category: "threatpreventionserverconnections",
}

// ThreatPreventionServerConnectionsList represents a list of ThreatPreventionServerConnections
type ThreatPreventionServerConnectionsList []*ThreatPreventionServerConnection

// ThreatPreventionServerConnectionsAncestor is the interface that an ancestor of a ThreatPreventionServerConnection must implement.
// An Ancestor is defined as an entity that has ThreatPreventionServerConnection as a descendant.
// An Ancestor can get a list of its child ThreatPreventionServerConnections, but not necessarily create one.
type ThreatPreventionServerConnectionsAncestor interface {
	ThreatPreventionServerConnections(*bambou.FetchingInfo) (ThreatPreventionServerConnectionsList, *bambou.Error)
}

// ThreatPreventionServerConnectionsParent is the interface that a parent of a ThreatPreventionServerConnection must implement.
// A Parent is defined as an entity that has ThreatPreventionServerConnection as a child.
// A Parent is an Ancestor which can create a ThreatPreventionServerConnection.
type ThreatPreventionServerConnectionsParent interface {
	ThreatPreventionServerConnectionsAncestor
	CreateThreatPreventionServerConnection(*ThreatPreventionServerConnection) *bambou.Error
}

// ThreatPreventionServerConnection represents the model of a threatpreventionserverconnection
type ThreatPreventionServerConnection struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
	FQDN       string `json:"FQDN,omitempty"`
	VSDName    string `json:"VSDName,omitempty"`
	Status     string `json:"status,omitempty"`
}

// NewThreatPreventionServerConnection returns a new *ThreatPreventionServerConnection
func NewThreatPreventionServerConnection() *ThreatPreventionServerConnection {

	return &ThreatPreventionServerConnection{}
}

// Identity returns the Identity of the object.
func (o *ThreatPreventionServerConnection) Identity() bambou.Identity {

	return ThreatPreventionServerConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ThreatPreventionServerConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ThreatPreventionServerConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ThreatPreventionServerConnection from the server
func (o *ThreatPreventionServerConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ThreatPreventionServerConnection into the server
func (o *ThreatPreventionServerConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ThreatPreventionServerConnection from the server
func (o *ThreatPreventionServerConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
