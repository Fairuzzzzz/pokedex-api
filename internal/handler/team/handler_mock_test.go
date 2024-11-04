// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go
//
// Generated by this command:
//
//	mockgen -source=handler.go -destination=handler_mock_test.go -package=team
//

// Package team is a generated GoMock package.
package team

import (
	context "context"
	reflect "reflect"

	team "github.com/Fairuzzzzz/pokedex-api/internal/models/team"
	gomock "go.uber.org/mock/gomock"
)

// Mockservice is a mock of service interface.
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
	isgomock struct{}
}

// MockserviceMockRecorder is the mock recorder for Mockservice.
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance.
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// CreateTeam mocks base method.
func (m *Mockservice) CreateTeam(ctx context.Context, userID uint, request team.PokeTeamNameRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockserviceMockRecorder) CreateTeam(ctx, userID, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*Mockservice)(nil).CreateTeam), ctx, userID, request)
}

// DeleteTeam mocks base method.
func (m *Mockservice) DeleteTeam(ctx context.Context, request team.PokeTeamRequestByID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeam", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeam indicates an expected call of DeleteTeam.
func (mr *MockserviceMockRecorder) DeleteTeam(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeam", reflect.TypeOf((*Mockservice)(nil).DeleteTeam), ctx, request)
}

// GetTeam mocks base method.
func (m *Mockservice) GetTeam(ctx context.Context, request team.PokeTeamRequestByID) (*team.PokeTeam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeam", ctx, request)
	ret0, _ := ret[0].(*team.PokeTeam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeam indicates an expected call of GetTeam.
func (mr *MockserviceMockRecorder) GetTeam(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeam", reflect.TypeOf((*Mockservice)(nil).GetTeam), ctx, request)
}

// ListTeam mocks base method.
func (m *Mockservice) ListTeam(ctx context.Context, userID uint) ([]team.PokeTeam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeam", ctx, userID)
	ret0, _ := ret[0].([]team.PokeTeam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeam indicates an expected call of ListTeam.
func (mr *MockserviceMockRecorder) ListTeam(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeam", reflect.TypeOf((*Mockservice)(nil).ListTeam), ctx, userID)
}
