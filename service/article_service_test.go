package service_test

import (
	"example-go-api/repository"
	"example-go-api/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type serviceSuite struct {
	suite.Suite
	service service.ArticleService
	repo    repository.MockRepository
}

func (s *serviceSuite) SetupTest() {
	s.repo = repository.MockRepository{}

	s.service = service.NewArticleService(&s.repo)
}

func (s *serviceSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(serviceSuite))
}

func (s *serviceSuite) TestNewService() {
	service := service.NewArticleService(&s.repo)
	s.NotNil(service)
}

func (s *serviceSuite) TestAddArticle_ShouldSucceed() {
	art := repository.Article{ID: "1", Title: "TestTitle", Desc: "Desc", Content: "Content"}
	s.repo.On("Get", art.ID).Return(repository.Article{})
	s.repo.On("Add", art).Return(nil)

	err := s.service.AddArticle(art)

	s.NoError(err)
}

func (s *serviceSuite) TestAddExistingArticle_ShouldFail() {
	art := repository.Article{ID: "1", Title: "TestTitle", Desc: "Desc", Content: "Content"}
	s.repo.On("Get", art.ID).Return(repository.Article{ID: "1"})

	err := s.service.AddArticle(art)

	s.Error(err)
}

func (s *serviceSuite) TestDeleteExistingArticle_ShouldSucceed() {
	art := repository.Article{ID: "1", Title: "TestTitle", Desc: "Desc", Content: "Content"}
	s.repo.On("Get", art.ID).Return(repository.Article{ID: "1"})
	s.repo.On("Delete", art.ID).Return(nil)

	err := s.service.DeleteArticle(art.ID)

	s.NoError(err)
}

func (s *serviceSuite) TestDeleteNonExistingArticle_ShouldFail() {
	art := repository.Article{ID: "1", Title: "TestTitle", Desc: "Desc", Content: "Content"}
	s.repo.On("Get", art.ID).Return(repository.Article{})

	err := s.service.DeleteArticle(art.ID)

	s.Error(err)
}

func (s *serviceSuite) TestGetExistingArticle_ShouldSucceed() {
	art := repository.Article{ID: "1", Title: "TestTitle", Desc: "Desc", Content: "Content"}
	s.repo.On("Get", art.ID).Return(art)

	exp, err := s.service.GetArticle(art.ID)

	s.NoError(err)
	assert.Equal(s.T(), art.ID, exp.ID)
}

func (s *serviceSuite) TestGetNonExistingArticle_ShouldFail() {
	art := repository.Article{ID: "1", Title: "TestTitle", Desc: "Desc", Content: "Content"}
	s.repo.On("Get", art.ID).Return(repository.Article{})

	exp, err := s.service.GetArticle(art.ID)

	assert.NotEqual(s.T(), art.ID, exp.ID)
	s.Error(err)
}
