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

package aws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/awsconfig"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func newEc2Client(region string, req handler.Request) *ec2.Client {
	cfg := awsconfig.FromHandlerRequest(&req)
	cfg.Region = convertToAWSRegion(region)
	return ec2.NewFromConfig(cfg)
}

type PrivateEndpointInput struct {
	InterfaceEndpointID *string
	VpcID               string
	SubnetIDs           []string
}

type PrivateEndpointOutput struct {
	VpcID               string
	InterfaceEndpointID string
	Region              string
	SubnetIDs           []string
}

func convertToAWSRegion(region string) string {
	return strings.ReplaceAll(strings.ToLower(region), "_", "-")
}

func CreatePrivateEndpoint(req handler.Request, endpointServiceName string, region string, privateEndpointInputs []PrivateEndpointInput) ([]PrivateEndpointOutput, *handler.ProgressEvent) {
	svc := newEc2Client(convertToAWSRegion(region), req)

	subnetIDs := make([]PrivateEndpointOutput, len(privateEndpointInputs))

	for i, pe := range privateEndpointInputs {
		connection := &ec2.CreateVpcEndpointInput{
			VpcId:           aws.String(pe.VpcID),
			ServiceName:     aws.String(endpointServiceName),
			VpcEndpointType: ec2types.VpcEndpointTypeInterface,
			SubnetIds:       pe.SubnetIDs,
		}

		vpcE, err := svc.CreateVpcEndpoint(context.Background(), connection)
		if err != nil {
			fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating vcp Endpoint: %s", err.Error()),
				string(types.HandlerErrorCodeGeneralServiceException))
			return nil, &fpe
		}

		subnetIDs[i] = PrivateEndpointOutput{
			VpcID:               pe.VpcID,
			SubnetIDs:           pe.SubnetIDs,
			InterfaceEndpointID: *vpcE.VpcEndpoint.VpcEndpointId,
		}
	}

	return subnetIDs, nil
}

func DeletePrivateEndpoint(req handler.Request, interfaceEndpoints []string, region string) *handler.ProgressEvent {
	svc := newEc2Client(convertToAWSRegion(region), req)

	connection := &ec2.DeleteVpcEndpointsInput{
		VpcEndpointIds: interfaceEndpoints,
	}

	_, err := svc.DeleteVpcEndpoints(context.Background(), connection)

	if err != nil {
		fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error deleting vcp Endpoint: %s", err.Error()),
			string(types.HandlerErrorCodeGeneralServiceException))
		return &fpe
	}

	return nil
}

func (i PrivateEndpointInput) ToString() string {
	return fmt.Sprintf("%s%s", i.VpcID, i.SubnetIDs)
}
