package response

import (
	"fgd/core/thread"
	"time"
)

type Thread struct {
	ID           int          `json:"id"`
	Author       ThreadAuthor `json:"author"`
	Topic        ThreadTopic  `json:"topic"`
	Title        string       `json:"title"`
	Content      *string      `json:"content"`
	Image1       *string      `json:"image_1"`
	Image2       *string      `json:"image_2"`
	Image3       *string      `json:"image_3"`
	Image4       *string      `json:"image_4"`
	Image5       *string      `json:"image_5"`
	LikedCount   int          `json:"liked_count"`
	UnlikedCount int          `json:"unliked_count"`
	ReplyCount   int          `json:"reply_count"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at"`
}

type ThreadTopic struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	ProfileImage *string `json:"profile_image"`
}

type ThreadAuthor struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func FromDomain(data *thread.Domain) Thread {
	return Thread{
		ID: data.ID,
		Author: ThreadAuthor{
			ID:        data.Author.ID,
			Username:  data.Author.Username,
			DeletedAt: data.Author.DeletedAt,
		},
		Topic: ThreadTopic{
			ID:           data.Topic.ID,
			Name:         data.Topic.Name,
			ProfileImage: data.Topic.ProfileImage,
		},
		Title:        data.Title,
		Content:      data.Content,
		Image1:       data.Image1,
		Image2:       data.Image2,
		Image3:       data.Image3,
		Image4:       data.Image4,
		Image5:       data.Image5,
		LikedCount:   data.LikeCount,
		UnlikedCount: data.UnlikeCount,
		ReplyCount:   data.ReplyCount,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt,
	}
}

func FromDomains(data *[]thread.Domain) []Thread {
	threads := []Thread{}

	for _, domain := range *data {
		threads = append(threads, Thread{
			ID: domain.ID,
			Author: ThreadAuthor{
				ID:        domain.Author.ID,
				Username:  domain.Author.Username,
				DeletedAt: domain.Author.DeletedAt,
			},
			Topic: ThreadTopic{
				ID:           domain.Topic.ID,
				Name:         domain.Topic.Name,
				ProfileImage: domain.Topic.ProfileImage,
			},
			Title:        domain.Title,
			Content:      domain.Content,
			Image1:       domain.Image1,
			Image2:       domain.Image2,
			Image3:       domain.Image3,
			Image4:       domain.Image4,
			Image5:       domain.Image5,
			LikedCount:   domain.LikeCount,
			UnlikedCount: domain.UnlikeCount,
			ReplyCount:   domain.ReplyCount,
			CreatedAt:    domain.CreatedAt,
			UpdatedAt:    domain.UpdatedAt,
			DeletedAt:    domain.DeletedAt,
		})
	}

	return threads
}
