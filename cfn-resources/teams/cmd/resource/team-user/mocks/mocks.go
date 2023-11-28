// Code generated by MockGen. DO NOT EDIT.
// Source: team_user_service.go
//
// Generated by this command:
//
//	mockgen -source=team_user_service.go -destination=./mocks/mocks.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
	gomock "go.uber.org/mock/gomock"
)

// MockTeamUsersAPI is a mock of TeamUsersAPI interface.
type MockTeamUsersAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTeamUsersAPIMockRecorder
}

// MockTeamUsersAPIMockRecorder is the mock recorder for MockTeamUsersAPI.
type MockTeamUsersAPIMockRecorder struct {
	mock *MockTeamUsersAPI
}

// NewMockTeamUsersAPI creates a new mock instance.
func NewMockTeamUsersAPI(ctrl *gomock.Controller) *MockTeamUsersAPI {
	mock := &MockTeamUsersAPI{ctrl: ctrl}
	mock.recorder = &MockTeamUsersAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamUsersAPI) EXPECT() *MockTeamUsersAPIMockRecorder {
	return m.recorder
}

// AddTeamUser mocks base method.
func (m *MockTeamUsersAPI) AddTeamUser(ctx context.Context, orgID, teamID string, addUserToTeam *[]admin.AddUserToTeam) (*admin.PaginatedApiAppUser, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTeamUser", ctx, orgID, teamID, addUserToTeam)
	ret0, _ := ret[0].(*admin.PaginatedApiAppUser)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddTeamUser indicates an expected call of AddTeamUser.
func (mr *MockTeamUsersAPIMockRecorder) AddTeamUser(ctx, orgID, teamID, addUserToTeam any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTeamUser", reflect.TypeOf((*MockTeamUsersAPI)(nil).AddTeamUser), ctx, orgID, teamID, addUserToTeam)
}

// GetUserByUsername mocks base method.
func (m *MockTeamUsersAPI) GetUserByUsername(ctx context.Context, userName string) (*admin.CloudAppUser, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, userName)
	ret0, _ := ret[0].(*admin.CloudAppUser)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockTeamUsersAPIMockRecorder) GetUserByUsername(ctx, userName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockTeamUsersAPI)(nil).GetUserByUsername), ctx, userName)
}

// RemoveTeamUser mocks base method.
func (m *MockTeamUsersAPI) RemoveTeamUser(ctx context.Context, orgID, teamID, userID string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTeamUser", ctx, orgID, teamID, userID)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTeamUser indicates an expected call of RemoveTeamUser.
func (mr *MockTeamUsersAPIMockRecorder) RemoveTeamUser(ctx, orgID, teamID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTeamUser", reflect.TypeOf((*MockTeamUsersAPI)(nil).RemoveTeamUser), ctx, orgID, teamID, userID)
}
