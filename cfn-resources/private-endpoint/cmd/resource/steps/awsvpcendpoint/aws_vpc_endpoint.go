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

package awsvpcendpoint

import (
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func newEc2Client(region string, req handler.Request) *ec2.EC2 {
	return ec2.New(req.Session, aws.NewConfig().WithRegion(region))
}

type AwsPrivateEndpointInput struct {
	VpcID               string
	SubnetIDs           []string
	InterfaceEndpointID *string
}

type AwsPrivateEndpointOutput struct {
	VpcID               string
	SubnetIDs           []string
	InterfaceEndpointID string
}

func convertToAWSRegion(region string) string {
	return strings.ReplaceAll(strings.ToLower(region), "_", "-")
}

func Create(req handler.Request, endpointServiceName string, region string, privateEndpointInputs []AwsPrivateEndpointInput) ([]AwsPrivateEndpointOutput, *handler.ProgressEvent) {
	svc := newEc2Client(convertToAWSRegion(region), req)

	vcpType := "Interface"

	subnetIds := make([]AwsPrivateEndpointOutput, len(privateEndpointInputs))

	for i, pe := range privateEndpointInputs {
		subnetIdsIn := make([]*string, len(pe.SubnetIDs))

		for i := range pe.SubnetIDs {
			subnetIdsIn[i] = &(pe.SubnetIDs[i])
		}

		connection := ec2.CreateVpcEndpointInput{
			VpcId:           &pe.VpcID,
			ServiceName:     &endpointServiceName,
			VpcEndpointType: &vcpType,
			SubnetIds:       subnetIdsIn,
		}

		vpcE, err := svc.CreateVpcEndpoint(&connection)
		if err != nil {
			fpe := progressevent.GetFailedEventByCode(fmt.Sprintf("Error creating vcp Endpoint: %s", err.Error()),
				cloudformation.HandlerErrorCodeGeneralServiceException)
			return nil, &fpe
		}

		subnetIds[i] = AwsPrivateEndpointOutput{
			VpcID:               pe.VpcID,
			SubnetIDs:           pe.SubnetIDs,
			InterfaceEndpointID: *vpcE.VpcEndpoint.VpcEndpointId,
		}
	}

	return subnetIds, nil
}

func Delete(req handler.Request, interfaceEndpoints []string, region string) *handler.ProgressEvent {
	svc := newEc2Client(convertToAWSRegion(region), req)

	vpcEndpointIds := make([]*string, 0)

	for i := range interfaceEndpoints {
		vpcEndpointIds = append(vpcEndpointIds, &interfaceEndpoints[i])
	}

	connection := ec2.DeleteVpcEndpointsInput{
		DryRun:         nil,
		VpcEndpointIds: vpcEndpointIds,
	}

	_, err := svc.DeleteVpcEndpoints(&connection)

	if err != nil {
		fpe := progressevent.GetFailedEventByCode(fmt.Sprintf("Error deleting vcp Endpoint: %s", err.Error()),
			cloudformation.HandlerErrorCodeGeneralServiceException)
		return &fpe
	}

	return nil
}

func (i AwsPrivateEndpointInput) ToString() string {
	return fmt.Sprintf("%s%s", i.VpcID, i.SubnetIDs)
}
