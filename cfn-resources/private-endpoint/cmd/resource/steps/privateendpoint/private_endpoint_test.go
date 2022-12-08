package privateendpoint

import (
	"fmt"
	"testing"
)

func TestDifference(t *testing.T) {
	/*Aca se borra uno, asi que deberia aparecer uno solo*/
	prevSlice := make([]AtlasPrivateEndpointInput, 2)
	currSlice := make([]AtlasPrivateEndpointInput, 1)

	prevSlice = append(prevSlice, AtlasPrivateEndpointInput{
		VpcId:               "a",
		SubnetId:            "b",
		InterfaceEndpointId: "dbbf",
	})

	prevSlice = append(prevSlice, AtlasPrivateEndpointInput{
		VpcId:               "h",
		SubnetId:            "f",
		InterfaceEndpointId: "wrwfdsfs",
	})

	currSlice = append(currSlice, AtlasPrivateEndpointInput{
		VpcId:               "a",
		SubnetId:            "b",
		InterfaceEndpointId: "dbbf",
	})

	currSlice = append(currSlice, AtlasPrivateEndpointInput{
		VpcId:               "f",
		SubnetId:            "x",
		InterfaceEndpointId: "asdgsg",
	})

	toDelete := sliceDifference(prevSlice, currSlice)

	t.Log(fmt.Sprintf("To delete amount %v", len(toDelete)))

	t.Log(fmt.Sprintf("To delete %s", toDelete[0].VpcId))

	toAdd := sliceDifference(currSlice, prevSlice)

	t.Log(fmt.Sprintf("To delete amount %v", len(toAdd)))

	t.Log(fmt.Sprintf("To delete %s", toAdd[0].VpcId))
}
