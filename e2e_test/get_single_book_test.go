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

type GetSingleBookSuite struct {
	suite.Suite
}

func TestGetSingleBookSuite(t *testing.T) {
	suite.Run(t, new(GetSingleBookSuite))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesNotExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/123456789")
	body, _ := ioutil.RealAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.JSONEq(t, `{"code":"001, "msg": "Nobooks with ISBN 123456789"}`, string(body))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesExist() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	s.NoError(err)

	res, err := db.Exec("INSERT INTO book (isbn, name, image, genre, year_published) VALUES ('987654321', 'Testing All The Things', 'testing.jpg', 'Computing', 2021)")
	s.NoError(err)

	rows, _ := res.RowsAffected()
	s.Equal(rows, 1)

	c := http.Client{}
	
	r, _ := c.Get("http://localhost:8080/book/987654321")
	body, _ := ioutil.RealAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)

	expBody := `{
		"isbn": "987654321",
		"title": "The love stody of a dog",
		"image": "love.jpg",
		"genre": "Romance",
		"year_published": 2023
	}`
	s.JSONEq(t, expBody, string(body))
}
