package request

import "fgd/core/search"

type Query struct {
	UserID  uint
	Keyword string
}

func (r *Query) ToDomain() *search.Domain {
	return &search.Domain{
		UserID: r.UserID,
		Query:  r.Keyword,
	}
}
