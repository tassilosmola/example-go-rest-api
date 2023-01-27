package controller

import (
	"encoding/json"
	"example-go-api/api"
	"example-go-api/repository"
	"example-go-api/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	port = ":10000"
)

// Controller struct
type Controller struct {
	router  api.Router
	service service.ArticleService
}

// Controller for Articles
func NewArticleController(router api.Router, service service.ArticleService) *Controller {
	c := &Controller{
		router:  router,
		service: service,
	}
	c.handleRequests()
	return c
}

func (c *Controller) handleRequests() {
	c.router.GET("/", c.homePage)
	c.router.GET("/articles", c.returnAllArticles)
	c.router.POST("/article", c.createNewArticle)
	c.router.DELETE("/article/{id}", c.deleteArticle)
	c.router.GET("/article/{id}", c.returnSingleArticle)
	c.router.SERVE(port)
}

// swagger:route GET / homePage
// Main entry point for the application
// responses:
// 	200: noContentResponse
func (c *Controller) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// swagger:route GET /articles articles returnAllArticles
// Return a list of available articles stored in the database
// responses:
// 	200: articlesResponse
// 	500: errorResponse
func (c *Controller) returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	ret, err := c.service.GetAllArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while fetching Articles \n" + err.Error()))
		return
	}
	json.NewEncoder(w).Encode(ret)
}

// swagger:route GET /article/{id} article returnSingleArticle
// Returns a single article with given id {id}
// responses:
// 	200: articleResponse
// 	500: errorResponse
func (c *Controller) returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/article/")

	ret, err := c.service.GetArticle(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while fetching Article \n" + err.Error()))
		return
	}

	json.NewEncoder(w).Encode(ret)

}

// swagger:route POST /article article createNewArticle
// Creates a new article with provided values
// responses:
// 	200: articleResponse
// 	500: errorValidation
func (c *Controller) createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article repository.Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	err := c.service.AddArticle(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while creating Article \n" + err.Error()))
		return
	}

	json.NewEncoder(w).Encode(article)
}

// swagger:route DELETE /article article deleteArticle
// Deletes a given article from the database
// responses:
// 	200: noContentResponse
// 	500: errorResponse
func (c *Controller) deleteArticle(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/article/")
	err := c.service.DeleteArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while deleting Article \n" + err.Error()))
	}
}
