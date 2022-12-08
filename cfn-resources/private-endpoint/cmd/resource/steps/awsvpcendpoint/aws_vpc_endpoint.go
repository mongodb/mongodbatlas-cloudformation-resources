package awsvpcendpoint

import (
	"fmt"

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
	VpcId               string
	SubnetId            string
	InterfaceEndpointId *string
}

type AwsPrivateEndpointOutput struct {
	VpcId               string
	SubnetId            string
	InterfaceEndpointId string
}

func Create(req handler.Request, endpointServiceName string, region string, privateEndpointInputs []AwsPrivateEndpointInput) ([]AwsPrivateEndpointOutput, *handler.ProgressEvent) {
	svc := newEc2Client(region, req)

	vcpType := "Interface"

	subnetIds := make([]AwsPrivateEndpointOutput, len(privateEndpointInputs))

	for i, pe := range privateEndpointInputs {
		connection := ec2.CreateVpcEndpointInput{
			VpcId:           &pe.VpcId,
			ServiceName:     &endpointServiceName,
			VpcEndpointType: &vcpType,
			SubnetIds:       []*string{&pe.SubnetId},
		}

		vpcE, err := svc.CreateVpcEndpoint(&connection)
		if err != nil {
			fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating vcp Endpoint: %s", err.Error()),
				cloudformation.HandlerErrorCodeGeneralServiceException)
			return nil, &fpe
		}

		subnetIds[i] = AwsPrivateEndpointOutput{
			VpcId:               pe.VpcId,
			SubnetId:            pe.SubnetId,
			InterfaceEndpointId: *vpcE.VpcEndpoint.VpcEndpointId,
		}
	}

	return subnetIds, nil
}

func Delete(req handler.Request, interfaceEndpoints []string, region string) *handler.ProgressEvent {
	svc := newEc2Client(region, req)

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
		fpe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error deleting vcp Endpoint: %s", err.Error()),
			cloudformation.HandlerErrorCodeGeneralServiceException)
		return &fpe
	}

	return nil
}

func Update(req handler.Request, endpointServiceName, region string, previousEndpointInput []AwsPrivateEndpointInput, currentEndpointInput []AwsPrivateEndpointInput) (*UpdateOutPut, *handler.ProgressEvent) {
	toAdd := sliceDifference(currentEndpointInput, previousEndpointInput)
	toDelete := sliceDifference(previousEndpointInput, currentEndpointInput)

	if len(toAdd) == 0 && len(toDelete) == 0 {
		progressEvent := handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "no private endpoints to delete or create to update",
			ResourceModel:   currentEndpointInput,
		}
		return nil, &progressEvent
	}

	updateOutput := UpdateOutPut{}

	if len(toAdd) > 0 {
		addProgressEvent := &handler.ProgressEvent{}
		outPut, addProgressEvent := Create(req, endpointServiceName, region, toAdd)
		if addProgressEvent.OperationStatus == handler.Failed {
			return nil, addProgressEvent
		}
		updateOutput.ToAdd = outPut
	}

	if len(toDelete) > 0 {
		addProgressEvent := &handler.ProgressEvent{}
		deleteInput := make([]string, 0, len(toDelete))
		updateToDeleteOutput := make([]AwsPrivateEndpointOutput, 0, len(toDelete))

		for i := range toDelete {
			deleteInput = append(deleteInput, *toDelete[i].InterfaceEndpointId)
			updateToDelete := AwsPrivateEndpointOutput{
				VpcId:               toDelete[i].VpcId,
				SubnetId:            toDelete[i].SubnetId,
				InterfaceEndpointId: *toDelete[i].InterfaceEndpointId,
			}
			updateToDeleteOutput = append(updateToDeleteOutput, updateToDelete)
		}

		addProgressEvent = Delete(req, deleteInput, region)
		if addProgressEvent.OperationStatus == handler.Failed {
			return nil, addProgressEvent
		}
		updateOutput.ToDelete = updateToDeleteOutput
	}

	return &updateOutput, nil
}

type UpdateOutPut struct {
	ToAdd    []AwsPrivateEndpointOutput
	ToDelete []AwsPrivateEndpointOutput
}

func sliceDifference(current, previous []AwsPrivateEndpointInput) []AwsPrivateEndpointInput {
	pSlice := make(map[string]string, len(previous))
	for _, p := range previous {
		pSlice[p.ToString()] = ""
	}
	var diff []AwsPrivateEndpointInput
	for _, x := range current {
		if _, found := pSlice[x.ToString()]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func (i AwsPrivateEndpointInput) ToString() string {
	return fmt.Sprintf("%s%s", i.VpcId, i.SubnetId)
}
