package reply

import (
	"errors"
	"fgd/core/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceReplyRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceReplyRepository) GetReplyByThreadID(threadId, limit, offset int) ([]reply.Domain, error) {
	replies := []Reply{}

	fetchErr := rp.Conn.Unscoped().Preload("Author").Where("thread_id = ?", threadId).Limit(limit).Offset(offset).Find(&replies).Error
	if fetchErr != nil {
		return []reply.Domain{}, fetchErr
	}

	domains := []reply.Domain{}
	for _, rep := range replies {
		replyDomain := rep.toDomain()

		var likeCount int64
		var unlikeCount int64
		var replyCount int64

		rp.Conn.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
		replyDomain.LikeCount = int(likeCount)
		rp.Conn.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
		replyDomain.UnlikeCount = int(unlikeCount)
		rp.Conn.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
		replyDomain.ReplyCount = int(replyCount)

		domains = append(domains, replyDomain)
	}

	return domains, nil
}

func (rp *persistenceReplyRepository) GetReplyByID(replyId int, limit int) (reply.Domain, error) {
	parentReply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	fetchErr := rp.Conn.Unscoped().Preload("Author").Take(&parentReply).Error
	if fetchErr != nil {
		return parentReply.toDomain(), fetchErr
	}

	parentDomain := parentReply.toDomain()
	childs := []Reply{}
	childsDomains := []reply.Domain{}
	fetchChildErr := rp.Conn.Preload("Author").Preload("Parent").Limit(limit).Where("parent_id = ?", parentReply.ID).Find(&childs).Error
	if fetchChildErr != nil {
		return parentDomain, fetchChildErr
	}

	for _, childReply := range childs {
		replyDomain := childReply.toDomain()

		var likeCount int64
		var unlikeCount int64
		var replyCount int64
		var childCount int64

		rp.Conn.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
		replyDomain.LikeCount = int(likeCount)
		rp.Conn.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
		replyDomain.UnlikeCount = int(unlikeCount)
		rp.Conn.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
		replyDomain.ReplyCount = int(replyCount)
		rp.Conn.Where("parent_id = ?", parentReply.ID).Count(&childCount)
		replyDomain.ChildCount = int(childCount)

		childsDomains = append(childsDomains, childReply.toDomain())
	}

	parentDomain.Child = &childsDomains

	return parentDomain, nil
}

func (rp *persistenceReplyRepository) GetReplyChilds(replyId, limit, offset int) ([]reply.Domain, error) {
	childs := []Reply{}
	domains := []reply.Domain{}
	fetchChildErr := rp.Conn.Preload("Author").Limit(limit).Offset(offset).Where("parent_id = ?", replyId).Find(&childs).Error
	if fetchChildErr != nil {
		return domains, fetchChildErr
	}

	for _, reply := range childs {
		replyDomain := reply.toDomain()

		var likeCount int64
		var unlikeCount int64
		var replyCount int64
		var childCount int64

		rp.Conn.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
		replyDomain.LikeCount = int(likeCount)
		rp.Conn.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
		replyDomain.UnlikeCount = int(unlikeCount)
		rp.Conn.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
		replyDomain.ReplyCount = int(replyCount)
		rp.Conn.Where("parent_id = ?", reply.ID).Count(&childCount)
		replyDomain.ChildCount = int(childCount)

		domains = append(domains, replyDomain)
	}

	return domains, nil
}

func (rp *persistenceReplyRepository) CreateReplyReply(data *reply.Domain, userId, replyId int) (reply.Domain, error) {
	parentReply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	author := user.User{Model: gorm.Model{ID: uint(userId)}}
	fetchParentResult := rp.Conn.Preload("Topic").Preload("Thread").Take(&parentReply)
	if fetchParentResult.Error != nil {
		return reply.Domain{}, fetchParentResult.Error
	}

	reply := fromDomain(data)

	authorErr := rp.Conn.Take(&author).Error
	if authorErr != nil {
		return reply.toDomain(), authorErr
	}

	reply.Author = author
	reply.Topic = parentReply.Topic
	reply.Parent = &parentReply
	reply.Thread = parentReply.Thread

	res := rp.Conn.Preload("Author").Create(&reply)
	return reply.toDomain(), res.Error
}

func (rp *persistenceReplyRepository) CreateReplyThread(data *reply.Domain, userId, threadId int) (reply.Domain, error) {
	thread := thread.Thread{Model: gorm.Model{ID: uint(threadId)}}
	author := user.User{Model: gorm.Model{ID: uint(userId)}}
	reply := fromDomain(data)

	fetchErr := rp.Conn.Preload("Topic").Take(&thread).Error
	if fetchErr != nil {
		return reply.toDomain(), fetchErr
	}

	authorErr := rp.Conn.Take(&author).Error
	if authorErr != nil {
		return reply.toDomain(), authorErr
	}

	reply.Author = author
	reply.Topic = thread.Topic
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
	if updatedReply.Image != nil {
		existingReply.Image = updatedReply.Image
	}

	err := rp.Conn.Where("author_id = ?", userId).Save(&existingReply).Error
	if err != nil {
		return existingReply.toDomain(), err
	}

	replyDomain := existingReply.toDomain()
	var likeCount int64
	var unlikeCount int64
	var replyCount int64
	var childCount int64

	rp.Conn.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
	replyDomain.LikeCount = int(likeCount)
	rp.Conn.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
	replyDomain.UnlikeCount = int(unlikeCount)
	rp.Conn.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
	replyDomain.ReplyCount = int(replyCount)
	rp.Conn.Where("parent_id = ?", existingReply.ID).Count(&childCount)
	replyDomain.ChildCount = int(childCount)

	return existingReply.toDomain(), err
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
