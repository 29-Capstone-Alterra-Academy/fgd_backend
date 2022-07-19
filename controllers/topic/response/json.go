package response

import (
	"fgd/core/topic"
	"time"
)

type Topic struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	ProfileImage     *string    `json:"profile_image"`
	Description      string     `json:"description"`
	Rules            *string    `json:"rules"`
	ActivityCount    int        `json:"activity_count"`
	ContributorCount int        `json:"contributor_count"`
	ModeratorCount   int        `json:"moderator_count"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
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
		DeletedAt:        domain.DeletedAt,
	}
}

func FromDomains(domains *[]topic.Domain) []Topic {
	topics := []Topic{}

	for _, domain := range *domains {
		topics = append(topics, Topic{
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
			DeletedAt:        domain.DeletedAt,
		})
	}

	return topics
}
