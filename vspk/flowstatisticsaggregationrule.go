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

// FlowstatisticsaggregationruleIdentity represents the Identity of the object
var FlowstatisticsaggregationruleIdentity = bambou.Identity {
    Name:     "flowstatisticsaggregationrule",
    Category: "flowstatisticsaggregationrules",
}

// FlowstatisticsaggregationrulesList represents a list of Flowstatisticsaggregationrules
type FlowstatisticsaggregationrulesList []*Flowstatisticsaggregationrule

// FlowstatisticsaggregationrulesAncestor is the interface that an ancestor of a Flowstatisticsaggregationrule must implement.
// An Ancestor is defined as an entity that has Flowstatisticsaggregationrule as a descendant.
// An Ancestor can get a list of its child Flowstatisticsaggregationrules, but not necessarily create one.
type FlowstatisticsaggregationrulesAncestor interface {
    Flowstatisticsaggregationrules(*bambou.FetchingInfo) (FlowstatisticsaggregationrulesList, *bambou.Error)
}

// FlowstatisticsaggregationrulesParent is the interface that a parent of a Flowstatisticsaggregationrule must implement.
// A Parent is defined as an entity that has Flowstatisticsaggregationrule as a child.
// A Parent is an Ancestor which can create a Flowstatisticsaggregationrule.
type FlowstatisticsaggregationrulesParent interface {
    FlowstatisticsaggregationrulesAncestor
    CreateFlowstatisticsaggregationrule(*Flowstatisticsaggregationrule) (*bambou.Error)
}

// Flowstatisticsaggregationrule represents the model of a flowstatisticsaggregationrule
type Flowstatisticsaggregationrule struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    MatchingCriteria string `json:"matchingCriteria,omitempty"`
    Description string `json:"description,omitempty"`
    AggregationCriteria string `json:"aggregationCriteria,omitempty"`
    AssociatedTrafficTypeID string `json:"associatedTrafficTypeID,omitempty"`
    
}

// NewFlowstatisticsaggregationrule returns a new *Flowstatisticsaggregationrule
func NewFlowstatisticsaggregationrule() *Flowstatisticsaggregationrule {

    return &Flowstatisticsaggregationrule{
        MatchingCriteria: "L4_SERVICE",
        AggregationCriteria: "FORWARD_AND_REVERSE_TRAFFIC_PORT_AGG",
        }
}

// Identity returns the Identity of the object.
func (o *Flowstatisticsaggregationrule) Identity() bambou.Identity {

    return FlowstatisticsaggregationruleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Flowstatisticsaggregationrule) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Flowstatisticsaggregationrule) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the Flowstatisticsaggregationrule from the server
func (o *Flowstatisticsaggregationrule) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Flowstatisticsaggregationrule into the server
func (o *Flowstatisticsaggregationrule) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Flowstatisticsaggregationrule from the server
func (o *Flowstatisticsaggregationrule) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}

