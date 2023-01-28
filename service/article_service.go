package service

import (
	"example-go-api/repository"
	"fmt"
)

type ArticleService interface {
	AddArticle(article repository.Article) error
	DeleteArticle(id string) error
	GetArticle(id string) (repository.Article, error)
	GetAllArticles() ([]repository.Article, error)
}

var (
	repo repository.Repository
)

type svc struct{}

// NewArticleService provides a ctor function for creating new ArticleService
func NewArticleService(r repository.Repository) ArticleService {
	repo = r
	return &svc{}
}

func (*svc) AddArticle(article repository.Article) error {
	art := repo.Get(article.ID)
	if art.ID != "" {
		return fmt.Errorf("Article with Id %v already exists in storage", article.ID)
	}

	repo.Add(article)

	return nil
}

func (*svc) DeleteArticle(id string) error {
	art := repo.Get(id)
	if art.ID != id {
		return fmt.Errorf("No article with id %v to delete", id)
	}

	repo.Delete(id)

	return nil
}

func (*svc) GetArticle(id string) (repository.Article, error) {
	art := repo.Get(id)
	if art.ID != id {
		return repository.Article{}, fmt.Errorf("No article with id %v in storage", id)
	}
	return art, nil
}

func (*svc) GetAllArticles() ([]repository.Article, error) {

	art := repo.GetAll()
	if len(art) == 0 {
		return nil, fmt.Errorf("No articles in storage")
	}
	return art, nil
}
