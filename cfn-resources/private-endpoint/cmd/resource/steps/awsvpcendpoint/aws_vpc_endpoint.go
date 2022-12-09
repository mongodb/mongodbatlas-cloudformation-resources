package awsvpcendpoint

import (
	"fmt"
	log2 "log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func newEc2Client(region string, req handler.Request) *ec2.EC2 {
	return ec2.New(req.Session, aws.NewConfig().WithRegion(region))
}

type AwsPrivateEndpointInput struct {
	VpcID               string
	SubnetID            string
	InterfaceEndpointID *string
}

type AwsPrivateEndpointOutput struct {
	VpcID               string
	SubnetID            string
	InterfaceEndpointID string
}

func Create(req handler.Request, endpointServiceName string, region string, privateEndpointInputs []AwsPrivateEndpointInput) ([]AwsPrivateEndpointOutput, *handler.ProgressEvent) {
	svc := newEc2Client(region, req)

	vcpType := "Interface"

	subnetIds := make([]AwsPrivateEndpointOutput, len(privateEndpointInputs))

	log2.Print("Entered Point 6")
	for i, pe := range privateEndpointInputs {
		log2.Print("Entered Point 7")
		connection := ec2.CreateVpcEndpointInput{
			VpcId:           &pe.VpcID,
			ServiceName:     &endpointServiceName,
			VpcEndpointType: &vcpType,
			SubnetIds:       []*string{&pe.SubnetID},
		}
		log2.Print("Entered Point 8")
		vpcE, err := svc.CreateVpcEndpoint(&connection)
		if err != nil {
			fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating vcp Endpoint: %s", err.Error()),
				cloudformation.HandlerErrorCodeGeneralServiceException)
			return nil, &fpe
		}

		log2.Print("Entered Point 9")
		subnetIds[i] = AwsPrivateEndpointOutput{
			VpcID:               pe.VpcID,
			SubnetID:            pe.SubnetID,
			InterfaceEndpointID: *vpcE.VpcEndpoint.VpcEndpointId,
		}
	}

	log2.Print("Entered Point 10")
	return subnetIds, nil
}

func Delete(req handler.Request, interfaceEndpoints []string, region string) *handler.ProgressEvent {
	log2.Print("Entered Point 13.1")
	svc := newEc2Client(region, req)

	vpcEndpointIds := make([]*string, 0)

	for i := range interfaceEndpoints {
		vpcEndpointIds = append(vpcEndpointIds, &interfaceEndpoints[i])
	}
	log2.Print("Entered Point 13.2")
	connection := ec2.DeleteVpcEndpointsInput{
		DryRun:         nil,
		VpcEndpointIds: vpcEndpointIds,
	}
	log2.Print("Entered Point 13.3")
	_, err := svc.DeleteVpcEndpoints(&connection)
	log2.Print("Entered Point 13.4")
	if err != nil {
		fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error deleting vcp Endpoint: %s", err.Error()),
			cloudformation.HandlerErrorCodeGeneralServiceException)
		return &fpe
	}

	log2.Print("Entered Point 13.5")

	return nil
}

func (i AwsPrivateEndpointInput) ToString() string {
	return fmt.Sprintf("%s%s", i.VpcID, i.SubnetID)
}
