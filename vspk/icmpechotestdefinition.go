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

// ICMPEchoTestDefinitionIdentity represents the Identity of the object
var ICMPEchoTestDefinitionIdentity = bambou.Identity {
    Name:     "icmpechotestdefinition",
    Category: "icmpechotestdefinitions",
}

// ICMPEchoTestDefinitionsList represents a list of ICMPEchoTestDefinitions
type ICMPEchoTestDefinitionsList []*ICMPEchoTestDefinition

// ICMPEchoTestDefinitionsAncestor is the interface that an ancestor of a ICMPEchoTestDefinition must implement.
// An Ancestor is defined as an entity that has ICMPEchoTestDefinition as a descendant.
// An Ancestor can get a list of its child ICMPEchoTestDefinitions, but not necessarily create one.
type ICMPEchoTestDefinitionsAncestor interface {
    ICMPEchoTestDefinitions(*bambou.FetchingInfo) (ICMPEchoTestDefinitionsList, *bambou.Error)
}

// ICMPEchoTestDefinitionsParent is the interface that a parent of a ICMPEchoTestDefinition must implement.
// A Parent is defined as an entity that has ICMPEchoTestDefinition as a child.
// A Parent is an Ancestor which can create a ICMPEchoTestDefinition.
type ICMPEchoTestDefinitionsParent interface {
    ICMPEchoTestDefinitionsAncestor
    CreateICMPEchoTestDefinition(*ICMPEchoTestDefinition) (*bambou.Error)
}

// ICMPEchoTestDefinition represents the model of a icmpechotestdefinition
type ICMPEchoTestDefinition struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    PacketCount int `json:"packetCount,omitempty"`
    PacketInterval int `json:"packetInterval,omitempty"`
    PacketSize int `json:"packetSize,omitempty"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    ThresholdAverageRoundTripTime float64 `json:"thresholdAverageRoundTripTime,omitempty"`
    ThresholdPacketLoss float64 `json:"thresholdPacketLoss,omitempty"`
    Timeout int `json:"timeout,omitempty"`
    SlaMonitoring bool `json:"slaMonitoring"`
    DonotFragment bool `json:"donotFragment"`
    Tos int `json:"tos,omitempty"`
    
}

// NewICMPEchoTestDefinition returns a new *ICMPEchoTestDefinition
func NewICMPEchoTestDefinition() *ICMPEchoTestDefinition {

    return &ICMPEchoTestDefinition{
        PacketCount: 5,
        PacketInterval: 1000,
        PacketSize: 64,
        Timeout: 60,
        SlaMonitoring: false,
        DonotFragment: false,
        }
}

// Identity returns the Identity of the object.
func (o *ICMPEchoTestDefinition) Identity() bambou.Identity {

    return ICMPEchoTestDefinitionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ICMPEchoTestDefinition) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ICMPEchoTestDefinition) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the ICMPEchoTestDefinition from the server
func (o *ICMPEchoTestDefinition) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ICMPEchoTestDefinition into the server
func (o *ICMPEchoTestDefinition) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ICMPEchoTestDefinition from the server
func (o *ICMPEchoTestDefinition) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}

