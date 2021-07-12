package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type _handlerTestSuite struct {
	suite.Suite
	svc    *Service
	engine *gin.Engine
}

func (s *_handlerTestSuite) SetupSuite() {
	db, err := ent.Open("sqlite3", ":memory:?_fk=1")
	s.Require().NoError(err)
	db = db.Debug()
	s.Require().NoError(db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false)))
	s.svc = New(db)
	s.engine = gin.Default()
	s.svc.RegisterRoute(s.engine)
}

func (s _handlerTestSuite) Test_WorkingOnTomato() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/working-tomato", nil)
	s.engine.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	s.Equal(`{"code":200,"data":null,"message":"not found tomato"}`, w.Body.String())
}

func TestService(t *testing.T) {
	suite.Run(t, &_handlerTestSuite{})
}
