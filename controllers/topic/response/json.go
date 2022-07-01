package response

import (
	"fgd/core/topic"
	"time"
)

type Topic struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ProfileImage     string    `json:"profile_image"`
	Description      string    `json:"description"`
	Rules            string    `json:"rules,omitempty"`
	ActivityCount    int       `json:"activity_count,omitempty"`
	ContributorCount int       `json:"contributor_count,omitempty"`
	ModeratorCount   int       `json:"moderator_count,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

func FromDomain(domain topic.Domain) Topic {
	return Topic{
		ID:               domain.ID,
		Name:             domain.Name,
		ProfileImage:     domain.ProfileImage,
		Description:      domain.Description,
		Rules:            domain.Rules,
		ActivityCount:    domain.ActivityCount,
		ContributorCount: domain.ContributorCount,
		ModeratorCount:   domain.ModeratorCount,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
