package aws_vpc_endpoint

import (
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas/mongodbatlas"
)

func newEc2Client(region string, req handler.Request) *ec2.EC2 {
	return ec2.New(req.Session, aws.NewConfig().WithRegion(region))
}

func Create(req handler.Request, peCon mongodbatlas.PrivateEndpointConnection, region string, subnetId string, VpcId string) (*string, *handler.ProgressEvent) {
	svc := newEc2Client(region, req)

	vcpType := "Interface"

	connection := ec2.CreateVpcEndpointInput{
		VpcId:           &VpcId,
		ServiceName:     &peCon.EndpointServiceName,
		VpcEndpointType: &vcpType,
		SubnetIds:       []*string{&subnetId},
	}

	vpcE, err := svc.CreateVpcEndpoint(&connection)
	if err != nil {
		fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating vcp Endpoint: %s", err.Error()),
			cloudformation.HandlerErrorCodeGeneralServiceException)
		return nil, &fpe
	}

	return vpcE.VpcEndpoint.VpcEndpointId, nil
}

func Delete(req handler.Request, interfaceEndpoints []string, region string) (*ec2.DeleteVpcEndpointsOutput, *handler.ProgressEvent) {
	svc := newEc2Client(region, req)

	vpcEndpointIds := make([]*string, 0)

	for _, i := range interfaceEndpoints {
		vpcEndpointIds = append(vpcEndpointIds, &i)
	}

	connection := ec2.DeleteVpcEndpointsInput{
		DryRun:         nil,
		VpcEndpointIds: vpcEndpointIds,
	}

	vpcE, err := svc.DeleteVpcEndpoints(&connection)
	if err != nil {
		fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error deleting vcp Endpoint: %s", err.Error()),
			cloudformation.HandlerErrorCodeGeneralServiceException)
		return nil, &fpe
	}

	return vpcE, nil
}
