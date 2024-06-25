package e2e_test

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

type GetSingleBookSuite struct{
	suite.Suite
}

func TestGetSingleBookSuite(t *testing.T){
	suite.Run(t, new(GetSingleBookSuite))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesNotExist(){
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/123456789")
	body, _ := ioutil.RealAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.JSONEq(t, `{}`, string(body))
}