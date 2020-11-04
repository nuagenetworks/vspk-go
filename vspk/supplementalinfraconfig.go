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

// SupplementalInfraConfigIdentity represents the Identity of the object
var SupplementalInfraConfigIdentity = bambou.Identity{
	Name:     "supplementalinfraconfig",
	Category: "supplementalinfraconfigs",
}

// SupplementalInfraConfigsList represents a list of SupplementalInfraConfigs
type SupplementalInfraConfigsList []*SupplementalInfraConfig

// SupplementalInfraConfigsAncestor is the interface that an ancestor of a SupplementalInfraConfig must implement.
// An Ancestor is defined as an entity that has SupplementalInfraConfig as a descendant.
// An Ancestor can get a list of its child SupplementalInfraConfigs, but not necessarily create one.
type SupplementalInfraConfigsAncestor interface {
	SupplementalInfraConfigs(*bambou.FetchingInfo) (SupplementalInfraConfigsList, *bambou.Error)
}

// SupplementalInfraConfigsParent is the interface that a parent of a SupplementalInfraConfig must implement.
// A Parent is defined as an entity that has SupplementalInfraConfig as a child.
// A Parent is an Ancestor which can create a SupplementalInfraConfig.
type SupplementalInfraConfigsParent interface {
	SupplementalInfraConfigsAncestor
	CreateSupplementalInfraConfig(*SupplementalInfraConfig) *bambou.Error
}

// SupplementalInfraConfig represents the model of a supplementalinfraconfig
type SupplementalInfraConfig struct {
	ID                 string      `json:"ID,omitempty"`
	ParentID           string      `json:"parentID,omitempty"`
	ParentType         string      `json:"parentType,omitempty"`
	Owner              string      `json:"owner,omitempty"`
	SupplementalConfig interface{} `json:"supplementalConfig,omitempty"`
}

// NewSupplementalInfraConfig returns a new *SupplementalInfraConfig
func NewSupplementalInfraConfig() *SupplementalInfraConfig {

	return &SupplementalInfraConfig{}
}

// Identity returns the Identity of the object.
func (o *SupplementalInfraConfig) Identity() bambou.Identity {

	return SupplementalInfraConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SupplementalInfraConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SupplementalInfraConfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SupplementalInfraConfig from the server
func (o *SupplementalInfraConfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SupplementalInfraConfig into the server
func (o *SupplementalInfraConfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SupplementalInfraConfig from the server
func (o *SupplementalInfraConfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
