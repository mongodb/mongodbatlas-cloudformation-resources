// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"log"
	"strings"

	"github.com/rs/xid"
)

// Resource id's are used to generate
/// deployment secrets, which contain all the
// apikey and properties for a given resource.
// also used as ID, since can encode resource and it's parent ids.
// you should be able to use the Atlas api to GET the particular
// instance of this resource.

type ResourceIdentifier struct {
	Parent       *ResourceIdentifier
	ServiceName  string
	DeploymentID string
	ResourceType string // TODO - make enum for this?
	ResourceID   string
}

// Note string version is "+" delimited string of the fields, in proper heirachry
func (t ResourceIdentifier) String() string {
	fields := []string{}
	if t.DeploymentID != "" {
		fields = append(fields, "mongodb", t.DeploymentID)
	}
	fields = append(fields, t.ResourceType, t.ResourceID)
	if t.Parent != nil {
		fields = append(fields, t.Parent.String())
	}
	return strings.Join(fields, "+")
}

// ParseResourceIdentifier : Given a resource identifier and a kind, "Project", "Cluster", "DBUser", etc...
func ParseResourceIdentifier(resourceID string) (*ResourceIdentifier, error) {
	var r ResourceIdentifier
	parts := strings.Split(resourceID, "+")

	if len(parts) < 4 {
		return &r, fmt.Errorf("unable to parse input to resource identifier:%s", resourceID)
	}
	r.ServiceName = parts[0]
	r.DeploymentID = parts[1]
	r.ResourceType = parts[2]
	r.ResourceID = parts[3]
	if len(parts) < 5 { // so no parent
		log.Printf("ParseResourceIdentifier: r:%+v", r)
		return &r, nil
	}
	// handle parent id(s)
	if parts[4] == "mongodb" { // it's a service name
		// so parse recursive
		ps := strings.Join(parts[4:], "+")
		p, err := ParseResourceIdentifier(ps)
		if err != nil {
			return &r, fmt.Errorf("error parsing resource id:%+v", err)
		}
		r.Parent = p
	} else { // simple parent embedded id
		r.Parent = &ResourceIdentifier{
			ResourceType: parts[4],
			ResourceID:   parts[5],
		}
	}
	log.Printf("ParseResourceIdentifier: r:%+v", r)
	return &r, nil
}

func NewResourceIdentifier(resourceType, resourceID string, parent *ResourceIdentifier) *ResourceIdentifier {
	deployID := xid.New()
	log.Printf("NewResourceIdentifier new deployID:%s", deployID.String())
	r := ResourceIdentifier{
		DeploymentID: deployID.String(),
		ResourceType: resourceType,
		ResourceID:   resourceID,
	}
	if parent != nil {
		r.Parent = parent
	}
	return &r
}
