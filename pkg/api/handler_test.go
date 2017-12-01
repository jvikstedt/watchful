package api_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/jvikstedt/watchful"
	"github.com/jvikstedt/watchful/pkg/api"
	"github.com/jvikstedt/watchful/pkg/exec/builtin"
	"github.com/jvikstedt/watchful/pkg/model"
	"github.com/jvikstedt/watchful/pkg/schedule"
)

var testHandler http.Handler
var db *sqlx.DB

func TestMain(m *testing.M) {
	logs := &bytes.Buffer{}
	logger := log.New(logs, "", log.LstdFlags)

	var err error
	db, err = model.NewDB("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := model.EnsureTables(db); err != nil {
		log.Fatal(err)
	}

	executorMock := &executorMock{}
	schedulerMock := &schedule.MockScheduler{}
	testHandler = api.New(logger, db, executorMock, schedulerMock)

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

type executorMock struct {
}

func (s *executorMock) AddJob(job *model.Job, isTestRun bool) string {
	return "test"
}

func (s *executorMock) Executables() map[string]watchful.Executable {
	return map[string]watchful.Executable{
		"http":  builtin.HTTP{},
		"equal": builtin.Equal{},
	}
}
