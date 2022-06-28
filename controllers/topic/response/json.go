package response

import (
	"fgd/core/topic"
	"time"
)

type Topic struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
	Description  string    `json:"description"`
	Rules        string    `json:"rules"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain topic.Domain) Topic {
	return Topic{
		ID:           domain.ID,
		Name:         domain.Name,
		ProfileImage: domain.ProfileImage,
		Description:  domain.Description,
		Rules:        domain.Rules,
	}
}
