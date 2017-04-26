// Automatically generated by MockGen. DO NOT EDIT!
// Source: common/interfaces.go

package common

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
)

// Mock of PersistenceAdapter interface
type MockPersistenceAdapter struct {
	ctrl     *gomock.Controller
	recorder *_MockPersistenceAdapterRecorder
}

// Recorder for MockPersistenceAdapter (not exported)
type _MockPersistenceAdapterRecorder struct {
	mock *MockPersistenceAdapter
}

func NewMockPersistenceAdapter(ctrl *gomock.Controller) *MockPersistenceAdapter {
	mock := &MockPersistenceAdapter{ctrl: ctrl}
	mock.recorder = &_MockPersistenceAdapterRecorder{mock}
	return mock
}

func (_m *MockPersistenceAdapter) EXPECT() *_MockPersistenceAdapterRecorder {
	return _m.recorder
}

func (_m *MockPersistenceAdapter) Create(msg Message) (string, error) {
	ret := _m.ctrl.Call(_m, "Create", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPersistenceAdapterRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockPersistenceAdapter) Update(id string, msg Message) error {
	ret := _m.ctrl.Call(_m, "Update", id, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPersistenceAdapterRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1)
}

func (_m *MockPersistenceAdapter) Delete(id string) error {
	ret := _m.ctrl.Call(_m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPersistenceAdapterRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

// Mock of Creator interface
type MockCreator struct {
	ctrl     *gomock.Controller
	recorder *_MockCreatorRecorder
}

// Recorder for MockCreator (not exported)
type _MockCreatorRecorder struct {
	mock *MockCreator
}

func NewMockCreator(ctrl *gomock.Controller) *MockCreator {
	mock := &MockCreator{ctrl: ctrl}
	mock.recorder = &_MockCreatorRecorder{mock}
	return mock
}

func (_m *MockCreator) EXPECT() *_MockCreatorRecorder {
	return _m.recorder
}

func (_m *MockCreator) Create(msg Message) (string, error) {
	ret := _m.ctrl.Call(_m, "Create", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCreatorRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

// Mock of Updater interface
type MockUpdater struct {
	ctrl     *gomock.Controller
	recorder *_MockUpdaterRecorder
}

// Recorder for MockUpdater (not exported)
type _MockUpdaterRecorder struct {
	mock *MockUpdater
}

func NewMockUpdater(ctrl *gomock.Controller) *MockUpdater {
	mock := &MockUpdater{ctrl: ctrl}
	mock.recorder = &_MockUpdaterRecorder{mock}
	return mock
}

func (_m *MockUpdater) EXPECT() *_MockUpdaterRecorder {
	return _m.recorder
}

func (_m *MockUpdater) Update(id string, msg Message) error {
	ret := _m.ctrl.Call(_m, "Update", id, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUpdaterRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1)
}

// Mock of Deleter interface
type MockDeleter struct {
	ctrl     *gomock.Controller
	recorder *_MockDeleterRecorder
}

// Recorder for MockDeleter (not exported)
type _MockDeleterRecorder struct {
	mock *MockDeleter
}

func NewMockDeleter(ctrl *gomock.Controller) *MockDeleter {
	mock := &MockDeleter{ctrl: ctrl}
	mock.recorder = &_MockDeleterRecorder{mock}
	return mock
}

func (_m *MockDeleter) EXPECT() *_MockDeleterRecorder {
	return _m.recorder
}

func (_m *MockDeleter) Delete(id string) error {
	ret := _m.ctrl.Call(_m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDeleterRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

// Mock of Searcher interface
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *_MockSearcherRecorder
}

// Recorder for MockSearcher (not exported)
type _MockSearcherRecorder struct {
	mock *MockSearcher
}

func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &_MockSearcherRecorder{mock}
	return mock
}

func (_m *MockSearcher) EXPECT() *_MockSearcherRecorder {
	return _m.recorder
}

func (_m *MockSearcher) SearchByUser(user string) ([]Message, error) {
	ret := _m.ctrl.Call(_m, "SearchByUser", user)
	ret0, _ := ret[0].([]Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSearcherRecorder) SearchByUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SearchByUser", arg0)
}

func (_m *MockSearcher) SearchByOrganization(organization string) ([]Message, error) {
	ret := _m.ctrl.Call(_m, "SearchByOrganization", organization)
	ret0, _ := ret[0].([]Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSearcherRecorder) SearchByOrganization(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SearchByOrganization", arg0)
}

func (_m *MockSearcher) SearchByProject(organization string) ([]Message, error) {
	ret := _m.ctrl.Call(_m, "SearchByProject", organization)
	ret0, _ := ret[0].([]Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSearcherRecorder) SearchByProject(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SearchByProject", arg0)
}

// Mock of Getter interface
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *_MockGetterRecorder
}

// Recorder for MockGetter (not exported)
type _MockGetterRecorder struct {
	mock *MockGetter
}

func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &_MockGetterRecorder{mock}
	return mock
}

func (_m *MockGetter) EXPECT() *_MockGetterRecorder {
	return _m.recorder
}

func (_m *MockGetter) GetById(id string) (Message, error) {
	ret := _m.ctrl.Call(_m, "GetById", id)
	ret0, _ := ret[0].(Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockGetterRecorder) GetById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetById", arg0)
}

func (_m *MockGetter) Get(beginning int, amount int) ([]Message, error) {
	ret := _m.ctrl.Call(_m, "Get", beginning, amount)
	ret0, _ := ret[0].([]Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockGetterRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

// Mock of MuxVarsGetter interface
type MockMuxVarsGetter struct {
	ctrl     *gomock.Controller
	recorder *_MockMuxVarsGetterRecorder
}

// Recorder for MockMuxVarsGetter (not exported)
type _MockMuxVarsGetterRecorder struct {
	mock *MockMuxVarsGetter
}

func NewMockMuxVarsGetter(ctrl *gomock.Controller) *MockMuxVarsGetter {
	mock := &MockMuxVarsGetter{ctrl: ctrl}
	mock.recorder = &_MockMuxVarsGetterRecorder{mock}
	return mock
}

func (_m *MockMuxVarsGetter) EXPECT() *_MockMuxVarsGetterRecorder {
	return _m.recorder
}

func (_m *MockMuxVarsGetter) Vars(r *http.Request) map[string]string {
	ret := _m.ctrl.Call(_m, "Vars", r)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

func (_mr *_MockMuxVarsGetterRecorder) Vars(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Vars", arg0)
}
