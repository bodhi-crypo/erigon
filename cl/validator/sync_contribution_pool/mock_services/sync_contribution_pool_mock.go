// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erigontech/erigon/cl/validator/sync_contribution_pool (interfaces: SyncContributionPool)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=mock_services/sync_contribution_pool_mock.go -package=mock_services . SyncContributionPool
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	common "github.com/erigontech/erigon-lib/common"
	cltypes "github.com/erigontech/erigon/cl/cltypes"
	state "github.com/erigontech/erigon/cl/phase1/core/state"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncContributionPool is a mock of SyncContributionPool interface.
type MockSyncContributionPool struct {
	ctrl     *gomock.Controller
	recorder *MockSyncContributionPoolMockRecorder
	isgomock struct{}
}

// MockSyncContributionPoolMockRecorder is the mock recorder for MockSyncContributionPool.
type MockSyncContributionPoolMockRecorder struct {
	mock *MockSyncContributionPool
}

// NewMockSyncContributionPool creates a new mock instance.
func NewMockSyncContributionPool(ctrl *gomock.Controller) *MockSyncContributionPool {
	mock := &MockSyncContributionPool{ctrl: ctrl}
	mock.recorder = &MockSyncContributionPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncContributionPool) EXPECT() *MockSyncContributionPoolMockRecorder {
	return m.recorder
}

// AddSyncCommitteeMessage mocks base method.
func (m *MockSyncContributionPool) AddSyncCommitteeMessage(headState *state.CachingBeaconState, subCommitee uint64, message *cltypes.SyncCommitteeMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSyncCommitteeMessage", headState, subCommitee, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSyncCommitteeMessage indicates an expected call of AddSyncCommitteeMessage.
func (mr *MockSyncContributionPoolMockRecorder) AddSyncCommitteeMessage(headState, subCommitee, message any) *MockSyncContributionPoolAddSyncCommitteeMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSyncCommitteeMessage", reflect.TypeOf((*MockSyncContributionPool)(nil).AddSyncCommitteeMessage), headState, subCommitee, message)
	return &MockSyncContributionPoolAddSyncCommitteeMessageCall{Call: call}
}

// MockSyncContributionPoolAddSyncCommitteeMessageCall wrap *gomock.Call
type MockSyncContributionPoolAddSyncCommitteeMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncContributionPoolAddSyncCommitteeMessageCall) Return(arg0 error) *MockSyncContributionPoolAddSyncCommitteeMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncContributionPoolAddSyncCommitteeMessageCall) Do(f func(*state.CachingBeaconState, uint64, *cltypes.SyncCommitteeMessage) error) *MockSyncContributionPoolAddSyncCommitteeMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncContributionPoolAddSyncCommitteeMessageCall) DoAndReturn(f func(*state.CachingBeaconState, uint64, *cltypes.SyncCommitteeMessage) error) *MockSyncContributionPoolAddSyncCommitteeMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AddSyncContribution mocks base method.
func (m *MockSyncContributionPool) AddSyncContribution(headState *state.CachingBeaconState, contribution *cltypes.Contribution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSyncContribution", headState, contribution)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSyncContribution indicates an expected call of AddSyncContribution.
func (mr *MockSyncContributionPoolMockRecorder) AddSyncContribution(headState, contribution any) *MockSyncContributionPoolAddSyncContributionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSyncContribution", reflect.TypeOf((*MockSyncContributionPool)(nil).AddSyncContribution), headState, contribution)
	return &MockSyncContributionPoolAddSyncContributionCall{Call: call}
}

// MockSyncContributionPoolAddSyncContributionCall wrap *gomock.Call
type MockSyncContributionPoolAddSyncContributionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncContributionPoolAddSyncContributionCall) Return(arg0 error) *MockSyncContributionPoolAddSyncContributionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncContributionPoolAddSyncContributionCall) Do(f func(*state.CachingBeaconState, *cltypes.Contribution) error) *MockSyncContributionPoolAddSyncContributionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncContributionPoolAddSyncContributionCall) DoAndReturn(f func(*state.CachingBeaconState, *cltypes.Contribution) error) *MockSyncContributionPoolAddSyncContributionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSyncAggregate mocks base method.
func (m *MockSyncContributionPool) GetSyncAggregate(slot uint64, beaconBlockRoot common.Hash) (*cltypes.SyncAggregate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncAggregate", slot, beaconBlockRoot)
	ret0, _ := ret[0].(*cltypes.SyncAggregate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncAggregate indicates an expected call of GetSyncAggregate.
func (mr *MockSyncContributionPoolMockRecorder) GetSyncAggregate(slot, beaconBlockRoot any) *MockSyncContributionPoolGetSyncAggregateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncAggregate", reflect.TypeOf((*MockSyncContributionPool)(nil).GetSyncAggregate), slot, beaconBlockRoot)
	return &MockSyncContributionPoolGetSyncAggregateCall{Call: call}
}

// MockSyncContributionPoolGetSyncAggregateCall wrap *gomock.Call
type MockSyncContributionPoolGetSyncAggregateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncContributionPoolGetSyncAggregateCall) Return(arg0 *cltypes.SyncAggregate, arg1 error) *MockSyncContributionPoolGetSyncAggregateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncContributionPoolGetSyncAggregateCall) Do(f func(uint64, common.Hash) (*cltypes.SyncAggregate, error)) *MockSyncContributionPoolGetSyncAggregateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncContributionPoolGetSyncAggregateCall) DoAndReturn(f func(uint64, common.Hash) (*cltypes.SyncAggregate, error)) *MockSyncContributionPoolGetSyncAggregateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSyncContribution mocks base method.
func (m *MockSyncContributionPool) GetSyncContribution(slot, subcommitteeIndex uint64, beaconBlockRoot common.Hash) *cltypes.Contribution {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncContribution", slot, subcommitteeIndex, beaconBlockRoot)
	ret0, _ := ret[0].(*cltypes.Contribution)
	return ret0
}

// GetSyncContribution indicates an expected call of GetSyncContribution.
func (mr *MockSyncContributionPoolMockRecorder) GetSyncContribution(slot, subcommitteeIndex, beaconBlockRoot any) *MockSyncContributionPoolGetSyncContributionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncContribution", reflect.TypeOf((*MockSyncContributionPool)(nil).GetSyncContribution), slot, subcommitteeIndex, beaconBlockRoot)
	return &MockSyncContributionPoolGetSyncContributionCall{Call: call}
}

// MockSyncContributionPoolGetSyncContributionCall wrap *gomock.Call
type MockSyncContributionPoolGetSyncContributionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSyncContributionPoolGetSyncContributionCall) Return(arg0 *cltypes.Contribution) *MockSyncContributionPoolGetSyncContributionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSyncContributionPoolGetSyncContributionCall) Do(f func(uint64, uint64, common.Hash) *cltypes.Contribution) *MockSyncContributionPoolGetSyncContributionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSyncContributionPoolGetSyncContributionCall) DoAndReturn(f func(uint64, uint64, common.Hash) *cltypes.Contribution) *MockSyncContributionPoolGetSyncContributionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}