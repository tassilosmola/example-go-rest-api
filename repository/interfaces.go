package repository

// Repository represents the basic interface for implemented repositories
type Repository interface {
	Get(id string) Article
	GetAll() []Article
	Delete(id string)
	Add(article Article)
}
