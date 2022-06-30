package response

import (
	"fgd/core/reply"
	"time"
)

type Reply struct {
	ID           int         `json:"id"`
	Author       ReplyAuthor `json:"author"`
	Content      string      `json:"content"`
	Image        string      `json:"image,omitempty"`
	LikedCount   int         `json:"liked_count"`
	UnlikedCount int         `json:"unliked_count"`
	ReplyCount   int         `json:"reply_count"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type ReplyAuthor struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image,omitempty"`
}

func FromDomain(data *reply.Domain) Reply {
	return Reply{
		ID: data.ID,
		Author: ReplyAuthor{
			ID:           data.Author.ID,
			Username:     data.Author.Username,
			ProfileImage: data.Author.ProfileImage,
		},
		LikedCount:   data.LikedCount,
		UnlikedCount: data.UnlikedCount,
		ReplyCount:   data.ReplyCount,
		Content:      data.Content,
		Image:        *data.Image,
		UpdatedAt:    data.UpdatedAt,
	}
}