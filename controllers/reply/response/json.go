package response

import (
	"fgd/core/reply"
	"time"
)

type Reply struct {
	ID           int         `json:"id"`
	Author       ReplyAuthor `json:"author"`
	Content      string      `json:"content"`
	Image        *string     `json:"image"`
	LikedCount   int         `json:"liked_count"`
	UnlikedCount int         `json:"unliked_count"`
	ReplyCount   int         `json:"reply_count"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	DeletedAt    *time.Time  `json:"deleted_at"`
}

type ReplyAuthor struct {
	ID           int        `json:"id"`
	Username     string     `json:"username"`
	ProfileImage *string    `json:"profile_image"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func FromDomain(data *reply.Domain) Reply {
	return Reply{
		ID: data.ID,
		Author: ReplyAuthor{
			ID:           data.Author.ID,
			Username:     data.Author.Username,
			ProfileImage: data.Author.ProfileImage,
		},
		LikedCount:   data.LikeCount,
		UnlikedCount: data.UnlikeCount,
		ReplyCount:   data.ReplyCount,
		Content:      data.Content,
		Image:        data.Image,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt,
	}
}

func FromDomains(datas *[]reply.Domain) []Reply {
	replies := []Reply{}

	for _, data := range *datas {
		replies = append(replies, Reply{
			ID: data.ID,
			Author: ReplyAuthor{
				ID:           data.Author.ID,
				Username:     data.Author.Username,
				ProfileImage: data.Author.ProfileImage,
			},
			LikedCount:   data.LikeCount,
			UnlikedCount: data.UnlikeCount,
			ReplyCount:   data.ReplyCount,
			Content:      data.Content,
			Image:        data.Image,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
			DeletedAt:    data.DeletedAt,
		})
	}

	return replies
}
