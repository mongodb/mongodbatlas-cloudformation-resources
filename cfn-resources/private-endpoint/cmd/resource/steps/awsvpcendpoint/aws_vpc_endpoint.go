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

	log2.Print("Entered Point 6")
	for i, pe := range privateEndpointInputs {
		log2.Print("Entered Point 7")
		connection := ec2.CreateVpcEndpointInput{
			VpcId:           &pe.VpcId,
			ServiceName:     &endpointServiceName,
			VpcEndpointType: &vcpType,
			SubnetIds:       []*string{&pe.SubnetId},
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
			VpcId:               pe.VpcId,
			SubnetId:            pe.SubnetId,
			InterfaceEndpointId: *vpcE.VpcEndpoint.VpcEndpointId,
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

func Update(req handler.Request, endpointServiceName, region string, previousEndpointInput []AwsPrivateEndpointInput, currentEndpointInput []AwsPrivateEndpointInput) (*UpdateOutPut, *handler.ProgressEvent) {
	log2.Print("Entered Point 4")
	toAdd := sliceDifference(currentEndpointInput, previousEndpointInput)
	toDelete := sliceDifference(previousEndpointInput, currentEndpointInput)

	log2.Printf("To add : %v", len(toAdd))
	log2.Printf("To delete : %v", len(toDelete))

	if len(toAdd) == 0 && len(toDelete) == 0 {
		progressEvent := handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "no private endpoints to delete or create to update",
			ResourceModel:   currentEndpointInput,
		}
		return nil, &progressEvent
	}

	log2.Print("Entered Point 5")
	updateOutput := UpdateOutPut{}

	if len(toAdd) > 0 {
		log2.Printf("ADDING %v private endpoints", len(toAdd))
		addProgressEvent := &handler.ProgressEvent{}
		outPut, addProgressEvent := Create(req, endpointServiceName, region, toAdd)
		if addProgressEvent != nil && addProgressEvent.OperationStatus == handler.Failed {
			return nil, addProgressEvent
		}
		log2.Print("Entered Point 5 - 2")
		updateOutput.ToAdd = outPut
		log2.Print("Entered Point 5 - 3")
	}

	if len(toDelete) > 0 {
		log2.Printf("Removing %v private endpoints", len(toDelete))
		deleteInput := make([]string, 0, len(toDelete))
		updateToDeleteOutput := make([]AwsPrivateEndpointOutput, 0, len(toDelete))
		log2.Print("Entered Point 11")
		for i := range toDelete {
			log2.Printf("To Delete %v", toDelete[i].InterfaceEndpointId)
			deleteInput = append(deleteInput, *toDelete[i].InterfaceEndpointId)
			updateToDelete := AwsPrivateEndpointOutput{
				VpcId:               toDelete[i].VpcId,
				SubnetId:            toDelete[i].SubnetId,
				InterfaceEndpointId: *toDelete[i].InterfaceEndpointId,
			}
			log2.Print("Entered Point 12")
			updateToDeleteOutput = append(updateToDeleteOutput, updateToDelete)
		}
		log2.Print("Entered Point 13")
		addProgressEvent := Delete(req, deleteInput, region)
		if addProgressEvent != nil && addProgressEvent.OperationStatus == handler.Failed {
			return nil, addProgressEvent
		}
		updateOutput.ToDelete = updateToDeleteOutput
	}
	log2.Print("Entered Point 14")
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
