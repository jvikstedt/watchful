package handler_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jvikstedt/watchful/pkg/handler"
	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/sqlite"
)

var testHandler http.Handler

func TestMain(m *testing.M) {
	logs := &bytes.Buffer{}
	logger := log.New(logs, "", log.LstdFlags)

	storage, err := sqlite.New(logger, "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	modelService := model.New(logger, storage)

	testHandler = handler.New(logger, modelService, nil)

	retCode := m.Run()
	os.Exit(retCode)
}

func makeRequest(t *testing.T, method string, path string, body string) (int, *bytes.Buffer) {
	b := bytes.NewBufferString(body)
	req, err := http.NewRequest(method, "/api/v1"+path, b)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	testHandler.ServeHTTP(rr, req)

	return rr.Result().StatusCode, rr.Body
}
