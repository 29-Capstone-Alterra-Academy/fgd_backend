package search

type Domain struct {
	UserID uint
	Query  string
}

type Usecase interface {
	GetSearchHistory(userId uint, keyword string, limit int) ([]Domain, error)
	StoreSearchKeyword(userId uint, data *Domain) error
}

type Repository interface {
	GetLastSearchHistory(userId uint, limit int) ([]Domain, error)
	QuerySearchHistory(userId uint, keyword string, limit int) ([]Domain, error)
	StoreSearchKeyword(userId uint, data *Domain) error
}
