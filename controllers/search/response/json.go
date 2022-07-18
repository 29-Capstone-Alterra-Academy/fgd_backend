package response

import (
	"fgd/core/search"
	"fgd/core/thread"
	"fgd/core/topic"
	"fgd/core/user"
	"time"
)

type SearchHistory struct {
	Entry string
}

type UserResult struct {
	ID             uint    `json:"id"`
	Username       string  `json:"username"`
	ProfileImage   *string `json:"profile_image"`
	FollowersCount int     `json:"followers_count"`
}

type TopicResult struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	ProfileImage *string `json:"profile_image"`
	ThreadCount  int     `json:"thread_count"`
}

type ThreadResult struct {
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
}

type ThreadTopic struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	ProfileImage *string `json:"profile_image"`
}

type ThreadAuthor struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func FromSearchDomains(data *[]search.Domain) []SearchHistory {
	entries := []SearchHistory{}

	for _, domain := range *data {
		entries = append(entries, SearchHistory{
			Entry: domain.Query,
		})
	}

	return entries
}

func FromUserDomains(data *[]user.Domain) []UserResult {
	users := []UserResult{}

	for _, domain := range *data {
		users = append(users, UserResult{
			ID:             uint(domain.ID),
			Username:       domain.Username,
			ProfileImage:   domain.ProfileImage,
			FollowersCount: domain.FollowersCount,
		})
	}

	return users
}

func FromTopicDomains(data *[]topic.Domain) []TopicResult {
	topics := []TopicResult{}

	for _, domain := range *data {
		topics = append(topics, TopicResult{
			ID:           uint(domain.ID),
			Name:         domain.Name,
			ProfileImage: domain.ProfileImage,
			ThreadCount:  domain.ActivityCount,
		})
	}

	return topics
}

func FromThreadDomains(data *[]thread.Domain) []ThreadResult {
	threads := []ThreadResult{}

	for _, domain := range *data {
		threads = append(threads, ThreadResult{
			ID: domain.ID,
			Author: ThreadAuthor{
				ID:       domain.Author.ID,
				Username: domain.Author.Username,
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
		})
	}

	return threads
}
