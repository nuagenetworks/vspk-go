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

// IDPProfileIdentity represents the Identity of the object
var IDPProfileIdentity = bambou.Identity {
    Name:     "idpprofile",
    Category: "idpprofiles",
}

// IDPProfilesList represents a list of IDPProfiles
type IDPProfilesList []*IDPProfile

// IDPProfilesAncestor is the interface that an ancestor of a IDPProfile must implement.
// An Ancestor is defined as an entity that has IDPProfile as a descendant.
// An Ancestor can get a list of its child IDPProfiles, but not necessarily create one.
type IDPProfilesAncestor interface {
    IDPProfiles(*bambou.FetchingInfo) (IDPProfilesList, *bambou.Error)
}

// IDPProfilesParent is the interface that a parent of a IDPProfile must implement.
// A Parent is defined as an entity that has IDPProfile as a child.
// A Parent is an Ancestor which can create a IDPProfile.
type IDPProfilesParent interface {
    IDPProfilesAncestor
    CreateIDPProfile(*IDPProfile) (*bambou.Error)
}

// IDPProfile represents the model of a idpprofile
type IDPProfile struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    ProtectAgainstInsertionEvasion bool `json:"protectAgainstInsertionEvasion"`
    AssociatedEnterpriseID string `json:"associatedEnterpriseID,omitempty"`
    
}

// NewIDPProfile returns a new *IDPProfile
func NewIDPProfile() *IDPProfile {

    return &IDPProfile{
        ProtectAgainstInsertionEvasion: true,
        }
}

// Identity returns the Identity of the object.
func (o *IDPProfile) Identity() bambou.Identity {

    return IDPProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IDPProfile) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IDPProfile) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the IDPProfile from the server
func (o *IDPProfile) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IDPProfile into the server
func (o *IDPProfile) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IDPProfile from the server
func (o *IDPProfile) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}


// IDPProfileActions retrieves the list of child IDPProfileActions of the IDPProfile
func (o *IDPProfile) IDPProfileActions(info *bambou.FetchingInfo) (IDPProfileActionsList, *bambou.Error) {

    var list IDPProfileActionsList
    err := bambou.CurrentSession().FetchChildren(o, IDPProfileActionIdentity, &list, info)
    return list, err
}



// CreateIDPProfileAction creates a new child IDPProfileAction under the IDPProfile
func (o *IDPProfile) CreateIDPProfileAction(child *IDPProfileAction) *bambou.Error {

    return bambou.CurrentSession().CreateChild(o, child)
}

