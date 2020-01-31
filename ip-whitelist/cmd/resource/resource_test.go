package resource

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/rs/xid"
)

const (
	publicKeyEnv  = "ATLAS_PUBLIC_KEY"
	privateKeyEnv = "ATLAS_PRIVATE_KEY"
	projectIDEnv  = "ATLAS_PROJECT_ID"
)

var (
	publicKey  = os.Getenv(publicKeyEnv)
	privateKey = os.Getenv(privateKeyEnv)
	projectID  = os.Getenv(projectIDEnv)
)

func new() *Model {
	return &Model{
		ProjectId: encoding.NewString(projectID),
		Whitelist: []WhitelistDefinition{
			WhitelistDefinition{
				Comment:   encoding.NewString("Testing range"),
				IpAddress: encoding.NewString("192.168.0.1"),
				ProjectId: encoding.NewString(projectID),
			},
		},
		ApiKeys: ApiKeyDefinition{
			PublicKey:  encoding.NewString(publicKey),
			PrivateKey: encoding.NewString(privateKey),
		},
		TotalCount: encoding.NewInt(1),
	}
}

func tearDown(model *Model) error {
	client, err := util.CreateMongoDBClient(publicKey, privateKey)
	if err != nil {
		return err
	}

	err = deleteEntries(model, client)
	if err != nil {
		return err
	}
	return nil
}

func setUp(model *Model) (*Model, error) {
	client, err := util.CreateMongoDBClient(*model.ApiKeys.PublicKey.Value(), *model.ApiKeys.PrivateKey.Value())
	projectID := *model.ProjectId.Value()
	request := getProjectIPWhitelistRequest(model)

	_, _, err = client.ProjectIPWhitelist.Create(context.Background(), projectID, request)
	if err != nil {
		return nil, err
	}

	guid := xid.New()

	model.Id = encoding.NewString(guid.String())
	return model, nil
}

func TestCreate(t *testing.T) {
	var testModel = new()

	type args struct {
		req          handler.Request
		prevModel    *Model
		currentModel *Model
	}
	tests := []struct {
		name    string
		args    args
		want    handler.ProgressEvent
		wantErr bool
	}{
		{"Simple CREATE", args{handler.Request{}, nil, testModel}, handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create Complete",
			ResourceModel:   testModel,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.req, tt.args.prevModel, tt.args.currentModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
			if err = tearDown(tt.args.currentModel); err != nil {
				t.Errorf("tearDown() failed error = %v", err.Error())
			}
		})
	}
}

func TestRead(t *testing.T) {
	var testModel = new()
	type args struct {
		req          handler.Request
		prevModel    *Model
		currentModel *Model
	}
	tests := []struct {
		name    string
		args    args
		want    handler.ProgressEvent
		wantErr bool
	}{
		{"Simple READ", args{handler.Request{}, nil, testModel}, handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Read Complete",
			ResourceModel:   testModel,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := setUp(tt.args.currentModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("setUp() error = %v", err)
				return
			}
			got, err := Read(tt.args.req, tt.args.prevModel, m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
			if err = tearDown(tt.args.currentModel); err != nil {
				t.Errorf("tearDown() failed error = %v", err.Error())
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	var testModel = new()
	type args struct {
		req          handler.Request
		prevModel    *Model
		currentModel *Model
	}
	tests := []struct {
		name    string
		args    args
		want    handler.ProgressEvent
		wantErr bool
	}{
		{"Simple UPDATE", args{handler.Request{}, nil, testModel}, handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Update Complete",
			ResourceModel:   testModel,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := setUp(tt.args.currentModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("setUp() error = %v", err)
				return
			}
			got, err := Update(tt.args.req, tt.args.prevModel, m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	var testModel = new()
	type args struct {
		req          handler.Request
		prevModel    *Model
		currentModel *Model
	}
	tests := []struct {
		name    string
		args    args
		want    handler.ProgressEvent
		wantErr bool
	}{
		{"Simple DELETE", args{handler.Request{}, nil, testModel}, handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Delete Complete",
			ResourceModel:   testModel,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := setUp(tt.args.currentModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("setUp() error = %v", err)
				return
			}
			got, err := Delete(tt.args.req, tt.args.prevModel, m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
				if err = tearDown(tt.args.currentModel); err != nil {
					t.Errorf("tearDown() failed error = %v", err.Error())
				}
			}
		})
	}
}
