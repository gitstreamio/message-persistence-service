package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"message-persistence-service/common"
	"message-persistence-service/server/mocks"
	"net/http"
	"testing"
)

//go:generate $GOPATH/bin/mockgen -destination mocks/mock_http.go -package mocks net/http ResponseWriter

type testUpdateData struct {
	testName       string
	testOrga       string
	testProj       string
	testId         string
	testAuthor     string
	testBody       string
	elasticErr     error
	expectedStatus int
}

var updateTestCases = []testUpdateData{
	{"sunny day test", "someorga", "someproject", "someid", "ksj", "anmsb", nil, 200},
	{"empty author", "someorga", "someproject", "someid", "", "anmsb", nil, 400},
	{"empty body", "someorga", "someproject", "someid", "ksj", "", nil, 400},
	{"empty id", "someorga", "someproject", "", "ksj", "anmsb", nil, 400},
	{"elastic error", "someorga", "someproject", "someid", "ksj", "anmsb", errors.New("someerr"), 502},
}

// Data driven test for updates
func Test_handleUpdate(t *testing.T) {
	for _, test := range updateTestCases {
		t.Run(test.testName, func(t *testing.T) {
			handleUpdate_test_function(t, test)
		})
	}
}

func handleUpdate_test_function(t *testing.T, data testUpdateData) {
	//given

	testVars := make(map[string]string)
	testVars[organization] = data.testOrga
	testVars[project] = data.testProj
	testVars[id] = data.testId

	msg := &common.Message{Author: data.testAuthor, Body: data.testBody}

	jsonMsg, _ := json.Marshal(msg)

	ctrl := gomock.NewController(t)
	persisterMock := common.NewMockPersistenceAdapter(ctrl)
	varsGetterMock := common.NewMockMuxVarsGetter(ctrl)

	//Request and RespWriter for the test
	req := &http.Request{Method: "UPDATE", Body: ioutil.NopCloser(bytes.NewReader(jsonMsg))}
	rw := mocks.NewMockResponseWriter(ctrl)

	toTest := &writeHandler{persisterMock, varsGetterMock}

	//expect
	rw.EXPECT().WriteHeader(data.expectedStatus).Times(1)
	varsGetterMock.EXPECT().Vars(req).Times(1).Return(testVars)
	persisterMock.EXPECT().Update(data.testId, gomock.Any()).Return(data.elasticErr)

	//when
	toTest.ServeHTTP(rw, req)

}
