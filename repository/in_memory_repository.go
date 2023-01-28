package repository

// Repository represents the basic interface for implemented repositories
type Repository interface {
	Get(id string) Article
	GetAll() []Article
	Delete(id string)
	Add(article Article)
}

type repo struct{}

var (
	// Articles represents the in memory storage of the repository
	Articles []Article
)

// NewInMemoryRepository creates a new in memory repository
func NewInMemoryRepository() Repository {
	return &repo{}
}

func (*repo) Get(id string) Article {
	for _, article := range Articles {
		if article.ID == id {
			return article
		}
	}
	return Article{}
}

func (*repo) GetAll() []Article {
	return Articles
}

func (*repo) Delete(id string) {
	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func (*repo) Add(article Article) {
	Articles = append(Articles, article)
}
