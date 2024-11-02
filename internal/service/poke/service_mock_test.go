// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock_test.go -package=poke
//

// Package poke is a generated GoMock package.
package poke

import (
	context "context"
	reflect "reflect"

	poke "github.com/Fairuzzzzz/pokedex-api/internal/repository/poke"
	gomock "go.uber.org/mock/gomock"
)

// MockpokemonOutbound is a mock of pokemonOutbound interface.
type MockpokemonOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockpokemonOutboundMockRecorder
	isgomock struct{}
}

// MockpokemonOutboundMockRecorder is the mock recorder for MockpokemonOutbound.
type MockpokemonOutboundMockRecorder struct {
	mock *MockpokemonOutbound
}

// NewMockpokemonOutbound creates a new mock instance.
func NewMockpokemonOutbound(ctrl *gomock.Controller) *MockpokemonOutbound {
	mock := &MockpokemonOutbound{ctrl: ctrl}
	mock.recorder = &MockpokemonOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpokemonOutbound) EXPECT() *MockpokemonOutboundMockRecorder {
	return m.recorder
}

// SearchPokemon mocks base method.
func (m *MockpokemonOutbound) SearchPokemon(ctx context.Context, name string) (*poke.Pokemon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPokemon", ctx, name)
	ret0, _ := ret[0].(*poke.Pokemon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPokemon indicates an expected call of SearchPokemon.
func (mr *MockpokemonOutboundMockRecorder) SearchPokemon(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPokemon", reflect.TypeOf((*MockpokemonOutbound)(nil).SearchPokemon), ctx, name)
}
