package response

import (
	"fgd/core/moderator"
	"time"
)

type PromotionRequest struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Topic     Topic     `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username"`
	ProfileImage *string `json:"profile_image"`
}

type Topic struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FromDomain(moderatorDomain moderator.Domain) PromotionRequest {
	return PromotionRequest{
		ID: moderatorDomain.ID,
		User: User{
			ID:           moderatorDomain.UserID,
			Username:     moderatorDomain.Username,
			ProfileImage: moderatorDomain.UserProfileImage,
		},
		Topic: Topic{
			ID:   moderatorDomain.TopicID,
			Name: moderatorDomain.TopicName,
		},
		CreatedAt: moderatorDomain.CreatedAt,
	}
}

func FromDomains(moderatorDomains *[]moderator.Domain) []PromotionRequest {
	reqs := []PromotionRequest{}

	for _, domains := range *moderatorDomains {
		reqs = append(reqs, PromotionRequest{
			ID: domains.ID,
			User: User{
				ID:           domains.UserID,
				Username:     domains.Username,
				ProfileImage: domains.UserProfileImage,
			},
			Topic: Topic{
				ID:   domains.TopicID,
				Name: domains.TopicName,
			},
			CreatedAt: domains.CreatedAt,
		})
	}

	return reqs
}
