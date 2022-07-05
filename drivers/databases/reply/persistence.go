package reply

import (
	"fgd/core/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceReplyRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceReplyRepository) CreateReplyReply(data *reply.Domain, userId, replyId int) (reply.Domain, error) {
	parentReply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	fetchParentResult := rp.Conn.Take(&parentReply)
	if fetchParentResult.Error != nil {
		return reply.Domain{}, fetchParentResult.Error
	}

	author := user.User{Model: gorm.Model{ID: uint(userId)}}
	reply := fromDomain(data)

	reply.Author = author
	reply.Parent = &parentReply
	reply.Thread = parentReply.Thread

	res := rp.Conn.Create(&reply)
	return reply.toDomain(), res.Error
}

func (rp *persistenceReplyRepository) CreateReplyThread(data *reply.Domain, userId, threadId int) (reply.Domain, error) {
	thread := thread.Thread{Model: gorm.Model{ID: uint(threadId)}}
	author := user.User{Model: gorm.Model{ID: uint(userId)}}
	reply := fromDomain(data)

	reply.Author = author
	reply.Thread = thread

	res := rp.Conn.Create(&reply)
	return reply.toDomain(), res.Error
}

func (rp *persistenceReplyRepository) DeleteReply(userId, replyId int) error {
	res := rp.Conn.Where("author_id = ?", userId).Delete(&Reply{}, replyId)
	return res.Error
}

func (rp *persistenceReplyRepository) EditReply(data *reply.Domain, userId, replyId int) (reply.Domain, error) {
	updatedReply := fromDomain(data)

	existingReply := Reply{}
	fetchResult := rp.Conn.Where("author_id = ?", userId).Take(&existingReply, replyId)
	if fetchResult.Error != nil {
		return reply.Domain{}, fetchResult.Error
	}

	existingReply.Content = updatedReply.Content
	existingReply.Image = &updatedReply.Content

	res := rp.Conn.Where("author_id = ?", userId).Save(&existingReply)
	return existingReply.toDomain(), res.Error
}

func (rp *persistenceReplyRepository) Like(userId, replyId int) error {
	undoErr := rp.UndoUnlike(userId, replyId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	return rp.Conn.Model(&reply).Association("LikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceReplyRepository) UndoLike(userId, replyId int) error {
	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	return rp.Conn.Model(&reply).Association("LikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceReplyRepository) UndoUnlike(userId, replyId int) error {
	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	return rp.Conn.Model(&reply).Association("UnlikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceReplyRepository) Unlike(userId, replyId int) error {
	undoErr := rp.UndoLike(userId, replyId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	return rp.Conn.Model(&reply).Association("UnlikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func InitPersistenceReplyRepository(c *gorm.DB) reply.Repository {
	return &persistenceReplyRepository{
		Conn: c,
	}
}
