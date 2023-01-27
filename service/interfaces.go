package service

import "example-go-api/repository"

type ArticleService interface {
	AddArticle(article repository.Article) error
	DeleteArticle(id string) error
	GetArticle(id string) (repository.Article, error)
	GetAllArticles() ([]repository.Article, error)
}
