package main

import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"kyc_urls/config"
	"kyc_urls/helpers"
	"kyc_urls/models"
	"net/http"
	"testing"
)

//testing GetLinks route
func Test_GetLinks(t *testing.T) {
	config.Connect()
	db := config.Db
	req, w := helpers.SetGetBooksRouter(db)
	defer db.Close()

	a := assert.New(t)
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Link{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.Link{}
	a.Equal(expected, actual) //comparing actual and expected results
	helpers.ClearTable()      //clearing the database to avoid duplicates
}

//Testing CreateLink Route
func Test_CreateBook_OK(t *testing.T) {
	a := assert.New(t)
	config.Connect()
	db := config.Db
	link := models.Link{
		Name:        "test",
		Url:         "test",
		Category:    "test",
		Environment: "test",
	}

	reqBody, err := json.Marshal(link)
	if err != nil {
		a.Error(err)
	}

	req, w, err := helpers.SetCreateBookRouter(db, bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Link{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := link
	a.Equal(expected, actual) //comparing actual and expected results
	helpers.ClearTable()      //clearing database to avoid duplicates
}

//Testing GetLinkByName
func Test_GetLink_OK(t *testing.T) {
	a := assert.New(t)
	config.Connect()
	db := config.Db

	link, err := helpers.InsertTestLink(db) //inserting test data
	if err != nil {
		a.Error(err)
	}

	req, w := helpers.SetGetLinkRouter(db) //getting response
	defer db.Close()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Link{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := link
	expected.Model = gorm.Model{}
	a.Equal(expected, actual) //comparing actual and expected results
	helpers.ClearTable()      //clearing database to avoid duplicates
}

//Testing GetLinkByCategoryRoute

func Test_GetLinkByCat_OK(t *testing.T) {
	a := assert.New(t)
	config.Connect()
	db := config.Db

	link, err := helpers.InsertTestLink(db) //inserting test data
	if err != nil {
		a.Error(err)
	}

	req, w := helpers.SetGetLinkByCatRouter(db) //getting response
	defer db.Close()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Link{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := link
	expected.Model = gorm.Model{}
	a.Equal(expected, actual) //comparing actual and expected results
	helpers.ClearTable()      //clearing database to avoid duplicates
}

//Testing GetLinkByEnvironment Route

func Test_GetLinkByEnv_OK(t *testing.T) {
	a := assert.New(t)
	config.Connect()
	db := config.Db

	link, err := helpers.InsertTestLink(db) //inserting test data
	if err != nil {
		a.Error(err)
	}

	req, w := helpers.SetGetLinkByEnvRouter(db) //getting response
	defer db.Close()

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Link{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := link
	expected.Model = gorm.Model{}
	a.Equal(expected, actual) //comparing actual and expected results
	helpers.ClearTable()      //clearing database to avoid duplicates
}
