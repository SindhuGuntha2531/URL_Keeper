package helpers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"kyc_urls/config"
	"kyc_urls/controllers"
	"kyc_urls/models"
	"net/http"
	"net/http/httptest"
)

//function to clear the database
func ClearTable() {
	db := config.Db
	db.Exec("DELETE FROM links")
	db.Exec("ALTER SEQUENCE id RESTART WITH 1")
}

func SetGetBooksRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New() //creating a new router
	//api := &APIEnv{DB: db}
	r.GET("/links", controllers.GetLinks) //http GET request
	req, err := http.NewRequest(http.MethodGet, "/links", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json") //setting up the header

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w
}

func SetCreateBookRouter(db *gorm.DB,
	body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()                           //creating a new router
	r.POST("/links", controllers.CreateLink) //http POST request
	req, err := http.NewRequest(http.MethodPost, "/links", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json") //setting up the header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func SetGetLinkRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New() //creating a new router

	r.GET("/link/:id", controllers.GetLinkByName) //http GET request for a specific name
	URL := models.Link{
		Name:        "test",
		Url:         "test",
		Category:    "test",
		Environment: "test",
	}
	req, err := http.NewRequest(http.MethodGet, "/link/"+URL.Name, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json") //setting up the header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w
}

func InsertTestLink(db *gorm.DB) (models.Link, error) {

	b := models.Link{
		Name:        "test",
		Url:         "test",
		Category:    "test",
		Environment: "test",
	}

	//creating a new test link
	if err := db.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

func SetGetLinkByCatRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New() //creating a new router

	r.GET("/cat/:id", controllers.GetLinkByName) //http GET request for a specific name
	URL := models.Link{
		Name:        "test",
		Url:         "test",
		Category:    "test",
		Environment: "test",
	}
	req, err := http.NewRequest(http.MethodGet, "/cat/"+URL.Category, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json") //setting up the header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w
}

func SetGetLinkByEnvRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New() //creating a new router

	r.GET("/env/:id", controllers.GetLinkByName) //http GET request for a specific name
	URL := models.Link{
		Name:        "test",
		Url:         "test",
		Category:    "test",
		Environment: "test",
	}
	req, err := http.NewRequest(http.MethodGet, "/env/"+URL.Environment, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json") //setting up the header
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w
}
