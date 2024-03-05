package resource

import (
	"errors"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/testutil/mocksvc"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

type stateTransitionTestCase struct {
	name                string
	respModel           *admin.ApiSearchDeploymentResponse
	respHttp            *http.Response
	respError           error
	targetState         string
	expectedEventStatus handler.Status
}

var prevModel = Model{
	Profile:     admin.PtrString(profile),
	ClusterName: admin.PtrString(clusterName),
	ProjectId:   admin.PtrString(dummyProjectID),
}

func TestStateTransitionProgressEvents(t *testing.T) {
	testCases := []stateTransitionTestCase{
		{
			name: "State in WORKING with target IDLE should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("UPDATING"),
			},
			respHttp: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.IdleState,
			expectedEventStatus: handler.InProgress,
		},
		{
			name: "State in IDLE with target IDLE should return success event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("IDLE"),
			},
			respHttp: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.IdleState,
			expectedEventStatus: handler.Success,
		},
		{
			name:      "400 response with target DELETED should return success event",
			respModel: nil,
			respHttp: &http.Response{
				StatusCode: 400,
			},
			respError:           errors.New(searchDeploymentDoesNotExistsError),
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.Success,
		},
		{
			name: "State in WORKING with target DELETED should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("UPDATING"),
			},
			respHttp: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.InProgress,
		},
		{
			name: "State in IDLE with target DELETED should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("IDLE"),
			},
			respHttp: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.InProgress,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := mocksvc.NewAtlasSearchApi(t)
			m.EXPECT().GetAtlasSearchDeployment(mock.Anything, mock.Anything, mock.Anything).Return(admin.GetAtlasSearchDeploymentApiRequest{ApiService: m}).Once()
			m.EXPECT().GetAtlasSearchDeploymentExecute(mock.Anything).Return(tc.respModel, tc.respHttp, tc.respError).Once()

			client := admin.APIClient{AtlasSearchApi: m}
			eventResult := handleStateTransition(client, &prevModel, tc.targetState)
			assert.Equal(t, tc.expectedEventStatus, eventResult.OperationStatus)
		})
	}
}