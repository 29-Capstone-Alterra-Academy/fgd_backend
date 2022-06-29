package response

import (
	"fgd/core/thread"
	"time"
)

type Thread struct {
	ID           int          `json:"id"`
	Author       ThreadAuthor `json:"author"`
	Title        string       `json:"title"`
	Content      string       `json:"content,omitempty"`
	Image1       string       `json:"image_1,omitempty"`
	Image2       string       `json:"image_2,omitempty"`
	Image3       string       `json:"image_3,omitempty"`
	Image4       string       `json:"image_4,omitempty"`
	Image5       string       `json:"image_5,omitempty"`
	LikedCount   int          `json:"liked_count"`
	UnlikedCount int          `json:"unliked_count"`
	ReplyCount   int          `json:"reply_count"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type ThreadAuthor struct {
	ID           int     `json:"id"`
	Username     string  `json:"username"`
	ProfileImage *string `json:"profile_image,omitempty"`
}

func FromDomain(data *thread.Domain) Thread {
	return Thread{
		ID: data.ID,
		Author: ThreadAuthor{
			ID:           data.Author.ID,
			Username:     data.Author.Username,
			ProfileImage: &data.Author.ProfileImage,
		},
		Title:        data.Title,
		Content:      data.Content,
		Image1:       *data.Image1,
		Image2:       *data.Image2,
		Image3:       *data.Image3,
		Image4:       *data.Image4,
		Image5:       *data.Image5,
		LikedCount:   data.LikeCount,
		UnlikedCount: data.UnlikeCount,
		ReplyCount:   data.ReplyCount,
		UpdatedAt:    data.UpdatedAt,
	}
}
