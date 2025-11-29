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
	"context"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/awsconfig"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func newEc2Client(region string, req handler.Request) *ec2.Client {
	cfg := awsconfig.FromHandlerRequest(&req)
	if region != "" {
		cfg.Region = region
	}
	return ec2.NewFromConfig(cfg)
}

type AwsPrivateEndpointInput struct {
	InterfaceEndpointID *string
	VpcID               string
	SubnetIDs           []string
}

type AwsPrivateEndpointOutput struct {
	VpcID               string
	InterfaceEndpointID string
	SubnetIDs           []string
}

func convertToAWSRegion(region string) string {
	return strings.ReplaceAll(strings.ToLower(region), "_", "-")
}

func Create(req handler.Request, endpointServiceName string, region string, privateEndpointInputs []AwsPrivateEndpointInput) ([]AwsPrivateEndpointOutput, *handler.ProgressEvent) {
	svc := newEc2Client(convertToAWSRegion(region), req)

	subnetIDs := make([]AwsPrivateEndpointOutput, len(privateEndpointInputs))

	for i, pe := range privateEndpointInputs {
		connection := &ec2.CreateVpcEndpointInput{
			VpcId:           &pe.VpcID,
			ServiceName:     &endpointServiceName,
			VpcEndpointType: ec2types.VpcEndpointTypeInterface,
			SubnetIds:       pe.SubnetIDs,
		}

		vpcE, err := svc.CreateVpcEndpoint(context.Background(), connection)
		if err != nil {
			fpe := progressevent.GetFailedEventByCode(fmt.Sprintf("Error creating vcp Endpoint: %s", err.Error()),
				string(types.HandlerErrorCodeGeneralServiceException))
			return nil, &fpe
		}

		subnetIDs[i] = AwsPrivateEndpointOutput{
			VpcID:               pe.VpcID,
			SubnetIDs:           pe.SubnetIDs,
			InterfaceEndpointID: *vpcE.VpcEndpoint.VpcEndpointId,
		}
	}

	return subnetIDs, nil
}

func Delete(req handler.Request, interfaceEndpoints []string, region string) *handler.ProgressEvent {
	svc := newEc2Client(convertToAWSRegion(region), req)

	connection := &ec2.DeleteVpcEndpointsInput{
		VpcEndpointIds: interfaceEndpoints,
	}

	_, err := svc.DeleteVpcEndpoints(context.Background(), connection)
	if err != nil {
		fpe := progressevent.GetFailedEventByCode(fmt.Sprintf("Error deleting vcp Endpoint: %s", err.Error()),
			string(types.HandlerErrorCodeGeneralServiceException))
		return &fpe
	}

	return nil
}

func (i AwsPrivateEndpointInput) ToString() string {
	return fmt.Sprintf("%s%s", i.VpcID, i.SubnetIDs)
}
