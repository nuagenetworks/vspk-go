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

// StatisticsprofileIdentity represents the Identity of the object
var StatisticsprofileIdentity = bambou.Identity {
    Name:     "statisticsprofile",
    Category: "statisticsprofiles",
}

// StatisticsprofilesList represents a list of Statisticsprofiles
type StatisticsprofilesList []*Statisticsprofile

// StatisticsprofilesAncestor is the interface that an ancestor of a Statisticsprofile must implement.
// An Ancestor is defined as an entity that has Statisticsprofile as a descendant.
// An Ancestor can get a list of its child Statisticsprofiles, but not necessarily create one.
type StatisticsprofilesAncestor interface {
    Statisticsprofiles(*bambou.FetchingInfo) (StatisticsprofilesList, *bambou.Error)
}

// StatisticsprofilesParent is the interface that a parent of a Statisticsprofile must implement.
// A Parent is defined as an entity that has Statisticsprofile as a child.
// A Parent is an Ancestor which can create a Statisticsprofile.
type StatisticsprofilesParent interface {
    StatisticsprofilesAncestor
    CreateStatisticsprofile(*Statisticsprofile) (*bambou.Error)
}

// Statisticsprofile represents the model of a statisticsprofile
type Statisticsprofile struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    CloneFrom string `json:"cloneFrom,omitempty"`
    FlowStatsAggregationEnabled bool `json:"flowStatsAggregationEnabled"`
    
}

// NewStatisticsprofile returns a new *Statisticsprofile
func NewStatisticsprofile() *Statisticsprofile {

    return &Statisticsprofile{
        FlowStatsAggregationEnabled: false,
        }
}

// Identity returns the Identity of the object.
func (o *Statisticsprofile) Identity() bambou.Identity {

    return StatisticsprofileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Statisticsprofile) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Statisticsprofile) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the Statisticsprofile from the server
func (o *Statisticsprofile) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Statisticsprofile into the server
func (o *Statisticsprofile) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Statisticsprofile from the server
func (o *Statisticsprofile) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// Flowstatisticsaggregationrules retrieves the list of child Flowstatisticsaggregationrules of the Statisticsprofile
func (o *Statisticsprofile) Flowstatisticsaggregationrules(info *bambou.FetchingInfo) (FlowstatisticsaggregationrulesList, *bambou.Error) {

    var list FlowstatisticsaggregationrulesList
    err := bambou.CurrentSession().FetchChildren(o, FlowstatisticsaggregationruleIdentity, &list, info)
    return list, err
}



// CreateFlowstatisticsaggregationrule creates a new child Flowstatisticsaggregationrule under the Statisticsprofile
func (o *Statisticsprofile) CreateFlowstatisticsaggregationrule(child *Flowstatisticsaggregationrule) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}


// Enterprises retrieves the list of child Enterprises of the Statisticsprofile
func (o *Statisticsprofile) Enterprises(info *bambou.FetchingInfo) (EnterprisesList, *bambou.Error) {

    var list EnterprisesList
    err := bambou.CurrentSession().FetchChildren(o, EnterpriseIdentity, &list, info)
    return list, err
}



// AssignEnterprises assigns the list of Enterprises to the Statisticsprofile
func (o *Statisticsprofile) AssignEnterprises(children EnterprisesList) *bambou.Error {

    list := []bambou.Identifiable{}
    for _, c := range children {
        list = append(list, c)
    }

    return bambou.CurrentSession().AssignChildren(o, list, EnterpriseIdentity)
}


// NSGateways retrieves the list of child NSGateways of the Statisticsprofile
func (o *Statisticsprofile) NSGateways(info *bambou.FetchingInfo) (NSGatewaysList, *bambou.Error) {

    var list NSGatewaysList
    err := bambou.CurrentSession().FetchChildren(o, NSGatewayIdentity, &list, info)
    return list, err
}



// AssignNSGateways assigns the list of NSGateways to the Statisticsprofile
func (o *Statisticsprofile) AssignNSGateways(children NSGatewaysList) *bambou.Error {

    list := []bambou.Identifiable{}
    for _, c := range children {
        list = append(list, c)
    }

    return bambou.CurrentSession().AssignChildren(o, list, NSGatewayIdentity)
}

