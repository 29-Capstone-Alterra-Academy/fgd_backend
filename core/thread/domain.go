package thread

import "time"

type Domain struct {
	ID          int
	Author      DomainAuthor
	Topic       DomainTopic
	Image1      *string
	Image2      *string
	Image3      *string
	Image4      *string
	Image5      *string
	Title       string
	Content     *string
	LikeCount   int
	UnlikeCount int
	ReplyCount  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type DomainAuthor struct {
	ID           int
	Username     string
	ProfileImage *string
}

type DomainTopic struct {
	ID   int
	Name string
}

type Usecase interface {
	GetThreadByAuthorID(userId, limit, offset int) ([]Domain, error)
	GetThreadByTopicID(topicId, limit, offset int) ([]Domain, error)
	CreateThread(data *Domain, userId, topicId int) (Domain, error)
	UpdateThread(data *Domain, userId, threadId int) (Domain, error)
	DeleteThread(userId, threadId int) error
	Like(userId, threadId int) error
	UndoLike(userId, threadId int) error
	Unlike(userId, threadId int) error
	UndoUnlike(userId, threadId int) error
}

type Repository interface {
	GetThreadByAuthorID(userId, limit, offset int) ([]Domain, error)
	GetThreadByTopicID(topicId, limit, offset int) ([]Domain, error)
	CreateThread(data *Domain, userId, topicId int) (Domain, error)
	UpdateThread(data *Domain, userId, threadId int) (Domain, error)
	DeleteThread(userId, threadId int) error
	Like(userId, threadId int) error
	UndoLike(userId, threadId int) error
	Unlike(userId, threadId int) error
	UndoUnlike(userId, threadId int) error
}
