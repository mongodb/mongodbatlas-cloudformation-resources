// Code generated by mockery. DO NOT EDIT.

package mocksvc

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// TeamUsersAPI is an autogenerated mock type for the TeamUsersAPI type
type TeamUsersAPI struct {
	mock.Mock
}

type TeamUsersAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *TeamUsersAPI) EXPECT() *TeamUsersAPI_Expecter {
	return &TeamUsersAPI_Expecter{mock: &_m.Mock}
}

// AddTeamUser provides a mock function with given fields: ctx, orgID, teamID, addUserToTeam
func (_m *TeamUsersAPI) AddTeamUser(ctx context.Context, orgID string, teamID string, addUserToTeam *[]admin.AddUserToTeam) (*admin.PaginatedApiAppUser, *http.Response, error) {
	ret := _m.Called(ctx, orgID, teamID, addUserToTeam)

	if len(ret) == 0 {
		panic("no return value specified for AddTeamUser")
	}

	var r0 *admin.PaginatedApiAppUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *[]admin.AddUserToTeam) (*admin.PaginatedApiAppUser, *http.Response, error)); ok {
		return rf(ctx, orgID, teamID, addUserToTeam)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *[]admin.AddUserToTeam) *admin.PaginatedApiAppUser); ok {
		r0 = rf(ctx, orgID, teamID, addUserToTeam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedApiAppUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *[]admin.AddUserToTeam) *http.Response); ok {
		r1 = rf(ctx, orgID, teamID, addUserToTeam)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *[]admin.AddUserToTeam) error); ok {
		r2 = rf(ctx, orgID, teamID, addUserToTeam)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TeamUsersAPI_AddTeamUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTeamUser'
type TeamUsersAPI_AddTeamUser_Call struct {
	*mock.Call
}

// AddTeamUser is a helper method to define mock.On call
//   - ctx context.Context
//   - orgID string
//   - teamID string
//   - addUserToTeam *[]admin.AddUserToTeam
func (_e *TeamUsersAPI_Expecter) AddTeamUser(ctx interface{}, orgID interface{}, teamID interface{}, addUserToTeam interface{}) *TeamUsersAPI_AddTeamUser_Call {
	return &TeamUsersAPI_AddTeamUser_Call{Call: _e.mock.On("AddTeamUser", ctx, orgID, teamID, addUserToTeam)}
}

func (_c *TeamUsersAPI_AddTeamUser_Call) Run(run func(ctx context.Context, orgID string, teamID string, addUserToTeam *[]admin.AddUserToTeam)) *TeamUsersAPI_AddTeamUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*[]admin.AddUserToTeam))
	})
	return _c
}

func (_c *TeamUsersAPI_AddTeamUser_Call) Return(_a0 *admin.PaginatedApiAppUser, _a1 *http.Response, _a2 error) *TeamUsersAPI_AddTeamUser_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *TeamUsersAPI_AddTeamUser_Call) RunAndReturn(run func(context.Context, string, string, *[]admin.AddUserToTeam) (*admin.PaginatedApiAppUser, *http.Response, error)) *TeamUsersAPI_AddTeamUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: ctx, userName
func (_m *TeamUsersAPI) GetUserByUsername(ctx context.Context, userName string) (*admin.CloudAppUser, *http.Response, error) {
	ret := _m.Called(ctx, userName)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *admin.CloudAppUser
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*admin.CloudAppUser, *http.Response, error)); ok {
		return rf(ctx, userName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *admin.CloudAppUser); ok {
		r0 = rf(ctx, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CloudAppUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, userName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, userName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TeamUsersAPI_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type TeamUsersAPI_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - userName string
func (_e *TeamUsersAPI_Expecter) GetUserByUsername(ctx interface{}, userName interface{}) *TeamUsersAPI_GetUserByUsername_Call {
	return &TeamUsersAPI_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", ctx, userName)}
}

func (_c *TeamUsersAPI_GetUserByUsername_Call) Run(run func(ctx context.Context, userName string)) *TeamUsersAPI_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TeamUsersAPI_GetUserByUsername_Call) Return(_a0 *admin.CloudAppUser, _a1 *http.Response, _a2 error) *TeamUsersAPI_GetUserByUsername_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *TeamUsersAPI_GetUserByUsername_Call) RunAndReturn(run func(context.Context, string) (*admin.CloudAppUser, *http.Response, error)) *TeamUsersAPI_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTeamUser provides a mock function with given fields: ctx, orgID, teamID, userID
func (_m *TeamUsersAPI) RemoveTeamUser(ctx context.Context, orgID string, teamID string, userID string) (*http.Response, error) {
	ret := _m.Called(ctx, orgID, teamID, userID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTeamUser")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*http.Response, error)); ok {
		return rf(ctx, orgID, teamID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *http.Response); ok {
		r0 = rf(ctx, orgID, teamID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, orgID, teamID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamUsersAPI_RemoveTeamUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTeamUser'
type TeamUsersAPI_RemoveTeamUser_Call struct {
	*mock.Call
}

// RemoveTeamUser is a helper method to define mock.On call
//   - ctx context.Context
//   - orgID string
//   - teamID string
//   - userID string
func (_e *TeamUsersAPI_Expecter) RemoveTeamUser(ctx interface{}, orgID interface{}, teamID interface{}, userID interface{}) *TeamUsersAPI_RemoveTeamUser_Call {
	return &TeamUsersAPI_RemoveTeamUser_Call{Call: _e.mock.On("RemoveTeamUser", ctx, orgID, teamID, userID)}
}

func (_c *TeamUsersAPI_RemoveTeamUser_Call) Run(run func(ctx context.Context, orgID string, teamID string, userID string)) *TeamUsersAPI_RemoveTeamUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *TeamUsersAPI_RemoveTeamUser_Call) Return(_a0 *http.Response, _a1 error) *TeamUsersAPI_RemoveTeamUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TeamUsersAPI_RemoveTeamUser_Call) RunAndReturn(run func(context.Context, string, string, string) (*http.Response, error)) *TeamUsersAPI_RemoveTeamUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewTeamUsersAPI creates a new instance of TeamUsersAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTeamUsersAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *TeamUsersAPI {
	mock := &TeamUsersAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
