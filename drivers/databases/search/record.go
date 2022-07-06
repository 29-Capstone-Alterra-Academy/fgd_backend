package search

import (
	"fgd/core/search"
	"fgd/drivers/databases/user"
)

type SearchHistory struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   user.User
	Query  string
}

func (r *SearchHistory) toDomain() search.Domain {
	return search.Domain{
		UserID: r.UserID,
		Query:  r.Query,
	}
}

func fromDomain(searchDomain search.Domain) *SearchHistory {
	return &SearchHistory{
		UserID: searchDomain.UserID,
		Query:  searchDomain.Query,
	}
}
