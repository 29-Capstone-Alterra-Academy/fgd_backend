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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	replies := []Reply{}
	fetchErr := tx.Unscoped().Preload("Author").Where("thread_id = ?", threadId).Limit(limit).Offset(offset).Find(&replies).Error
	if fetchErr != nil {
		return []reply.Domain{}, fetchErr
	}

	domains := []reply.Domain{}
	for _, rep := range replies {
		replyDomain := rep.toDomain()

		var likeCount int64
		var unlikeCount int64
		var replyCount int64

		tx.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
		replyDomain.LikeCount = int(likeCount)
		tx.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
		replyDomain.UnlikeCount = int(unlikeCount)
		tx.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
		replyDomain.ReplyCount = int(replyCount)

		domains = append(domains, replyDomain)
	}

	return domains, nil
}

func (rp *persistenceReplyRepository) GetReplyByID(replyId int, limit int) (reply.Domain, error) {
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	parentReply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	fetchErr := tx.Unscoped().Preload("Author").Take(&parentReply).Error
	if fetchErr != nil {
		return parentReply.toDomain(), fetchErr
	}

	parentDomain := parentReply.toDomain()
	childs := []Reply{}
	childsDomains := []reply.Domain{}
	fetchChildErr := tx.Preload("Author").Preload("Parent").Limit(limit).Where("parent_id = ?", parentReply.ID).Find(&childs).Error
	if fetchChildErr != nil {
		return parentDomain, fetchChildErr
	}

	for _, childReply := range childs {
		replyDomain := childReply.toDomain()

		var likeCount int64
		var unlikeCount int64
		var replyCount int64
		var childCount int64

		tx.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
		replyDomain.LikeCount = int(likeCount)
		tx.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
		replyDomain.UnlikeCount = int(unlikeCount)
		tx.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
		replyDomain.ReplyCount = int(replyCount)
		tx.Where("parent_id = ?", parentReply.ID).Count(&childCount)
		replyDomain.ChildCount = int(childCount)

		childsDomains = append(childsDomains, childReply.toDomain())
	}

	parentDomain.Child = &childsDomains

	return parentDomain, nil
}

func (rp *persistenceReplyRepository) GetReplyChilds(replyId, limit, offset int) ([]reply.Domain, error) {
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	childs := []Reply{}
	domains := []reply.Domain{}
	fetchChildErr := tx.Preload("Author").Limit(limit).Offset(offset).Where("parent_id = ?", replyId).Find(&childs).Error
	if fetchChildErr != nil {
		return domains, fetchChildErr
	}

	for _, reply := range childs {
		replyDomain := reply.toDomain()

		var likeCount int64
		var unlikeCount int64
		var replyCount int64
		var childCount int64

		tx.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
		replyDomain.LikeCount = int(likeCount)
		tx.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
		replyDomain.UnlikeCount = int(unlikeCount)
		tx.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
		replyDomain.ReplyCount = int(replyCount)
		tx.Where("parent_id = ?", reply.ID).Count(&childCount)
		replyDomain.ChildCount = int(childCount)

		domains = append(domains, replyDomain)
	}

	return domains, nil
}

func (rp *persistenceReplyRepository) CreateReplyReply(data *reply.Domain, userId, replyId int) (reply.Domain, error) {
	tx := rp.Conn.Begin()

	parentReply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	author := user.User{Model: gorm.Model{ID: uint(userId)}}
	fetchParentErr := tx.Preload("Topic").Preload("Thread").Take(&parentReply).Error
	if fetchParentErr != nil {
		return reply.Domain{}, fetchParentErr
	}

	newReply := fromDomain(data)

	authorErr := tx.Take(&author).Error
	if authorErr != nil {
		return newReply.toDomain(), authorErr
	}

	newReply.Author = author
	newReply.Topic = parentReply.Topic
	newReply.Parent = &parentReply
	newReply.Thread = parentReply.Thread

	err := tx.Preload("Author").Create(&newReply).Error
	if err != nil {
		tx.Rollback()
		return reply.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return reply.Domain{}, err
	}

	return newReply.toDomain(), nil
}

func (rp *persistenceReplyRepository) CreateReplyThread(data *reply.Domain, userId, threadId int) (reply.Domain, error) {
	tx := rp.Conn.Begin()

	thread := thread.Thread{Model: gorm.Model{ID: uint(threadId)}}
	author := user.User{Model: gorm.Model{ID: uint(userId)}}
	newReply := fromDomain(data)

	fetchErr := tx.Preload("Topic").Take(&thread).Error
	if fetchErr != nil {
		return reply.Domain{}, fetchErr
	}

	authorErr := tx.Take(&author).Error
	if authorErr != nil {
		return reply.Domain{}, authorErr
	}

	newReply.Author = author
	newReply.Topic = thread.Topic
	newReply.Thread = thread

	err := tx.Create(&newReply).Error
	if err != nil {
		tx.Rollback()
		return reply.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return reply.Domain{}, err
	}

	return newReply.toDomain(), nil
}

func (rp *persistenceReplyRepository) DeleteReply(userId, replyId int) error {
	res := rp.Conn.Where("author_id = ?", userId).Delete(&Reply{}, replyId)
	return res.Error
}

func (rp *persistenceReplyRepository) EditReply(data *reply.Domain, userId, replyId int) (reply.Domain, error) {
	tx := rp.Conn.Begin()

	updatedReply := fromDomain(data)

	existingReply := Reply{}
	fetchResultErr := tx.Where("author_id = ?", userId).Take(&existingReply, replyId).Error
	if fetchResultErr != nil {
		return reply.Domain{}, fetchResultErr
	}

	existingReply.Content = updatedReply.Content
	if updatedReply.Image != nil {
		existingReply.Image = updatedReply.Image
	}

	err := tx.Where("author_id = ?", userId).Save(&existingReply).Error
	if err != nil {
		tx.Rollback()
		return reply.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return reply.Domain{}, err
	}

	replyDomain := existingReply.toDomain()

	var likeCount int64
	var unlikeCount int64
	var replyCount int64
	var childCount int64

	tx.Table("liked_reply").Where("reply_id", replyDomain.ID).Count(&likeCount)
	replyDomain.LikeCount = int(likeCount)
	tx.Table("unliked_reply").Where("reply_id", replyDomain.ID).Count(&unlikeCount)
	replyDomain.UnlikeCount = int(unlikeCount)
	tx.Model(&Reply{}).Where("parent_id", replyDomain.ID).Count(&replyCount)
	replyDomain.ReplyCount = int(replyCount)
	tx.Where("parent_id = ?", existingReply.ID).Count(&childCount)
	replyDomain.ChildCount = int(childCount)

	return existingReply.toDomain(), nil
}

func (rp *persistenceReplyRepository) Like(userId, replyId int) error {
	tx := rp.Conn.Begin()

	undoErr := rp.UndoUnlike(userId, replyId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	err := tx.Model(&reply).Association("LikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (rp *persistenceReplyRepository) UndoLike(userId, replyId int) error {
	tx := rp.Conn.Begin()

	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	err := tx.Model(&reply).Association("LikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (rp *persistenceReplyRepository) UndoUnlike(userId, replyId int) error {
	tx := rp.Conn.Begin()

	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	err := tx.Model(&reply).Association("UnlikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (rp *persistenceReplyRepository) Unlike(userId, replyId int) error {
	tx := rp.Conn.Begin()

	undoErr := rp.UndoLike(userId, replyId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	reply := Reply{Model: gorm.Model{ID: uint(replyId)}}
	err := tx.Model(&reply).Association("UnlikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func InitPersistenceReplyRepository(c *gorm.DB) reply.Repository {
	return &persistenceReplyRepository{
		Conn: c,
	}
}
