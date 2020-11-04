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

// VirtualUplinkIdentity represents the Identity of the object
var VirtualUplinkIdentity = bambou.Identity{
	Name:     "virtualuplink",
	Category: "virtualuplinks",
}

// VirtualUplinksList represents a list of VirtualUplinks
type VirtualUplinksList []*VirtualUplink

// VirtualUplinksAncestor is the interface that an ancestor of a VirtualUplink must implement.
// An Ancestor is defined as an entity that has VirtualUplink as a descendant.
// An Ancestor can get a list of its child VirtualUplinks, but not necessarily create one.
type VirtualUplinksAncestor interface {
	VirtualUplinks(*bambou.FetchingInfo) (VirtualUplinksList, *bambou.Error)
}

// VirtualUplinksParent is the interface that a parent of a VirtualUplink must implement.
// A Parent is defined as an entity that has VirtualUplink as a child.
// A Parent is an Ancestor which can create a VirtualUplink.
type VirtualUplinksParent interface {
	VirtualUplinksAncestor
	CreateVirtualUplink(*VirtualUplink) *bambou.Error
}

// VirtualUplink represents the model of a virtualuplink
type VirtualUplink struct {
	ID                                    string `json:"ID,omitempty"`
	ParentID                              string `json:"parentID,omitempty"`
	ParentType                            string `json:"parentType,omitempty"`
	Owner                                 string `json:"owner,omitempty"`
	PeerEndpoint                          string `json:"peerEndpoint,omitempty"`
	PeerGatewayID                         string `json:"peerGatewayID,omitempty"`
	PeerGatewayName                       string `json:"peerGatewayName,omitempty"`
	PeerGatewaySystemID                   string `json:"peerGatewaySystemID,omitempty"`
	PeerPortID                            string `json:"peerPortID,omitempty"`
	PeerUplinkID                          int    `json:"peerUplinkID,omitempty"`
	PeerVLANID                            string `json:"peerVLANID,omitempty"`
	ShuntEndpoint                         string `json:"shuntEndpoint,omitempty"`
	ShuntPortID                           string `json:"shuntPortID,omitempty"`
	ShuntVLANID                           string `json:"shuntVLANID,omitempty"`
	VirtualUplinkDatapathID               string `json:"virtualUplinkDatapathID,omitempty"`
	EnableNATProbes                       bool   `json:"enableNATProbes"`
	UnderlayID                            int    `json:"underlayID,omitempty"`
	UnderlayNAT                           bool   `json:"underlayNAT"`
	UnderlayName                          string `json:"underlayName,omitempty"`
	UnderlayRouting                       bool   `json:"underlayRouting"`
	Role                                  string `json:"role,omitempty"`
	RoleOrder                             int    `json:"roleOrder,omitempty"`
	TrafficThroughUBROnly                 bool   `json:"trafficThroughUBROnly"`
	AssociatedEgressQoSPolicyID           string `json:"associatedEgressQoSPolicyID,omitempty"`
	AssociatedIngressOverlayQoSPolicerID  string `json:"associatedIngressOverlayQoSPolicerID,omitempty"`
	AssociatedIngressQoSPolicyID          string `json:"associatedIngressQoSPolicyID,omitempty"`
	AssociatedIngressUnderlayQoSPolicerID string `json:"associatedIngressUnderlayQoSPolicerID,omitempty"`
	AssociatedUplinkConnectionID          string `json:"associatedUplinkConnectionID,omitempty"`
	AssociatedVSCProfileID                string `json:"associatedVSCProfileID,omitempty"`
	AuxMode                               string `json:"auxMode,omitempty"`
}

// NewVirtualUplink returns a new *VirtualUplink
func NewVirtualUplink() *VirtualUplink {

	return &VirtualUplink{
		EnableNATProbes:       false,
		UnderlayNAT:           true,
		UnderlayRouting:       true,
		Role:                  "PRIMARY",
		TrafficThroughUBROnly: false,
		AuxMode:               "NONE",
	}
}

// Identity returns the Identity of the object.
func (o *VirtualUplink) Identity() bambou.Identity {

	return VirtualUplinkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VirtualUplink) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VirtualUplink) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VirtualUplink from the server
func (o *VirtualUplink) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VirtualUplink into the server
func (o *VirtualUplink) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VirtualUplink from the server
func (o *VirtualUplink) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// IKEGatewayConnections retrieves the list of child IKEGatewayConnections of the VirtualUplink
func (o *VirtualUplink) IKEGatewayConnections(info *bambou.FetchingInfo) (IKEGatewayConnectionsList, *bambou.Error) {

	var list IKEGatewayConnectionsList
	err := bambou.CurrentSession().FetchChildren(o, IKEGatewayConnectionIdentity, &list, info)
	return list, err
}
