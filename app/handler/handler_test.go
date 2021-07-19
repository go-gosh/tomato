package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type _testResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func (r *_testResponse) Unmarshal(v *bytes.Buffer) error {
	return json.Unmarshal(v.Bytes(), r)
}

type _handlerTestSuite struct {
	suite.Suite
	svc     *service.Service
	engine  *gin.Engine
	handler *Service
}

func (s *_handlerTestSuite) SetupSuite() {
	db, err := ent.Open("sqlite3", ":memory:?_fk=1")
	s.Require().NoError(err)
	db = db.Debug()
	s.Require().NoError(db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false)))
	s.svc = service.New(db)
	s.engine = gin.Default()
	s.handler = New(s.svc)
	s.handler.RegisterRoute(s.engine)
}

func (s _handlerTestSuite) Test_NormalCase() {
	s.testUserNoWorkingOnTomato()
	s.testStartTomato()
	s.testUserHasWorkingOnTomato()
	s.testCloseTomato()
	s.testUserNoWorkingOnTomato()
}

func (s _handlerTestSuite) serveApi(w *httptest.ResponseRecorder, req *http.Request) _testResponse {
	s.engine.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	resp := _testResponse{}
	s.NoError(resp.Unmarshal(w.Body))
	s.EqualValues(200, resp.Code)
	return resp
}

func (s _handlerTestSuite) testUserNoWorkingOnTomato() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/working-tomato", nil)
	s.engine.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal(`{"code":200,"data":null,"message":"not found tomato"}`, w.Body.String())
}

func (s _handlerTestSuite) testUserHasWorkingOnTomato() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/working-tomato", nil)
	resp := s.serveApi(w, req)
	s.EqualValues(200, resp.Code)
	s.NotEmpty(resp.Data["id"])
}

func (s _handlerTestSuite) testStartTomato() {
	w := httptest.NewRecorder()
	var b bytes.Buffer
	b.WriteString(`{"duration":60,"color":"red"}`)
	req, _ := http.NewRequest("POST", "/api/v1/tomato", &b)
	resp := s.serveApi(w, req)
	s.EqualValues(200, resp.Code)
	s.NotEmpty(resp.Data["id"])
	s.T().Logf("%+v", resp)
}

func (s _handlerTestSuite) testCloseTomato() {
	w := httptest.NewRecorder()
	var b bytes.Buffer
	b.WriteString(`{}`)
	req, _ := http.NewRequest("POST", "/api/v1/closing-tomato", &b)
	resp := s.serveApi(w, req)
	s.EqualValues(200, resp.Code)
	s.Equal("success", resp.Message)
	s.T().Logf("%+v", resp)
}

func TestService(t *testing.T) {
	suite.Run(t, &_handlerTestSuite{})
}
