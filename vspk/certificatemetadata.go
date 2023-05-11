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

// CertificateMetadataIdentity represents the Identity of the object
var CertificateMetadataIdentity = bambou.Identity {
    Name:     "None",
    Category: "None",
}

// CertificateMetadatasList represents a list of CertificateMetadatas
type CertificateMetadatasList []*CertificateMetadata

// CertificateMetadatasAncestor is the interface that an ancestor of a CertificateMetadata must implement.
// An Ancestor is defined as an entity that has CertificateMetadata as a descendant.
// An Ancestor can get a list of its child CertificateMetadatas, but not necessarily create one.
type CertificateMetadatasAncestor interface {
    CertificateMetadatas(*bambou.FetchingInfo) (CertificateMetadatasList, *bambou.Error)
}

// CertificateMetadatasParent is the interface that a parent of a CertificateMetadata must implement.
// A Parent is defined as an entity that has CertificateMetadata as a child.
// A Parent is an Ancestor which can create a CertificateMetadata.
type CertificateMetadatasParent interface {
    CertificateMetadatasAncestor
    CreateCertificateMetadata(*CertificateMetadata) (*bambou.Error)
}

// CertificateMetadata represents the model of a None
type CertificateMetadata struct {
    ID         string `json:"ID,omitempty"`
    ParentID   string `json:"parentID,omitempty"`
    ParentType string `json:"parentType,omitempty"`
    Owner      string `json:"owner,omitempty"`
    SHA1Fingerprint string `json:"SHA1Fingerprint,omitempty"`
    Serial string `json:"serial,omitempty"`
    NotAfter float64 `json:"notAfter,omitempty"`
    NotBefore float64 `json:"notBefore,omitempty"`
    Issuer string `json:"issuer,omitempty"`
    Subject string `json:"subject,omitempty"`
    
}

// NewCertificateMetadata returns a new *CertificateMetadata
func NewCertificateMetadata() *CertificateMetadata {

    return &CertificateMetadata{
        }
}

// Identity returns the Identity of the object.
func (o *CertificateMetadata) Identity() bambou.Identity {

    return CertificateMetadataIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CertificateMetadata) Identifier() string {

    return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CertificateMetadata) SetIdentifier(ID string) {

    o.ID = ID
}

// Fetch retrieves the CertificateMetadata from the server
func (o *CertificateMetadata) Fetch() *bambou.Error {

    return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CertificateMetadata into the server
func (o *CertificateMetadata) Save() *bambou.Error {

    return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CertificateMetadata from the server
func (o *CertificateMetadata) Delete() *bambou.Error {

    return bambou.CurrentSession().DeleteEntity(o)
}

