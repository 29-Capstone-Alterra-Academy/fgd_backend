package reply

import "time"

type Domain struct {
	ID           int
	Author       DomainAuthor
	Image        *string
	Content      string
	LikedCount   int
	UnlikedCount int
	ReplyCount   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type DomainAuthor struct {
	ID           int
	Username     string
	ProfileImage *string
}

type Usecase interface {
	CreateReplyThread(data *Domain, userId, threadId int) (Domain, error)
	CreateReplyReply(data *Domain, userId, replyId int) (Domain, error)
	EditReply(data *Domain, userId, replyId int) (Domain, error)
	DeleteReply(userId, replyId int) error
	Like(userId, replyId int) error
	UndoLike(userId, replyId int) error
	Unlike(userId, replyId int) error
	UndoUnlike(userId, replyId int) error
}

type Repository interface {
	CreateReplyThread(data *Domain, userId, threadId int) (Domain, error)
	CreateReplyReply(data *Domain, userId, replyId int) (Domain, error)
	EditReply(data *Domain, userId, replyId int) (Domain, error)
	DeleteReply(userId, replyId int) error
	Like(userId, replyId int) error
	UndoLike(userId, replyId int) error
	Unlike(userId, replyId int) error
	UndoUnlike(userId, replyId int) error
}
