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