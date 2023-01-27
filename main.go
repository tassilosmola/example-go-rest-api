package main

import (
	"example-go-api/api"
	"example-go-api/controller"
	"example-go-api/repository"
	"example-go-api/service"
)

var (
	inMemoryRepository repository.Repository  = repository.NewInMemoryRepository()
	articleService     service.ArticleService = service.NewArticleService(inMemoryRepository)
	httpRouter         api.Router             = api.NewMuxRouter()
	articleController                         = controller.NewArticleController(httpRouter, articleService)
)

func main() {

}
