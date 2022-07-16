package thread

import (
	"errors"
	"fgd/core/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceThreadRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceThreadRepository) CreateThread(data *thread.Domain, userId, topicId int) (thread.Domain, error) {
	tx := rp.Conn.Begin()

	topic := topic.Topic{}
	fetchTopicErr := tx.Take(&topic, topicId).Error
	if fetchTopicErr != nil {
		return thread.Domain{}, fetchTopicErr
	}

	author := user.User{}
	fetchAuthorErr := tx.Take(&author, userId).Error
	if fetchAuthorErr != nil {
		return thread.Domain{}, fetchAuthorErr
	}
	newThread := fromDomain(*data)
	newThread.Topic = topic
	newThread.Author = author

	err := tx.Create(&newThread).Error
	if err != nil {
		tx.Rollback()
		return thread.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return thread.Domain{}, err
	}

	return *newThread.toDomain(), nil
}

func (rp *persistenceThreadRepository) DeleteThread(userId int, threadId int) error {
	return rp.Conn.Where("author_id = ?", userId).Delete(&Thread{}, threadId).Error
}

func (rp *persistenceThreadRepository) GetThreadByID(threadId int) (thread.Domain, error) {
	thread := Thread{}

	res := rp.Conn.Preload("Author").Preload("Topic").Take(&thread, threadId)

	domain := thread.toDomain()

	var likeCount int64
	var unlikeCount int64
	var replyCount int64

	rp.Conn.Table("liked_thread").Where("thread_id", domain.ID).Count(&likeCount)
	domain.LikeCount = int(likeCount)
	rp.Conn.Table("unliked_thread").Where("thread_id", domain.ID).Count(&unlikeCount)
	domain.UnlikeCount = int(unlikeCount)
	rp.Conn.Table("replies").Where("thread_id", domain.ID).Count(&replyCount)
	domain.ReplyCount = int(replyCount)

	return *domain, res.Error
}

func (rp *persistenceThreadRepository) GetThreadByAuthorID(userId, limit, offset int) ([]thread.Domain, error) {
	threads := []Thread{}
	fetchResult := rp.Conn.Preload("Author").Preload("Topic").Limit(limit).Offset(offset).Where("author_id = ?", userId).Find(&threads)
	if fetchResult.Error != nil {
		return []thread.Domain{}, fetchResult.Error
	}

	threadDomains := []thread.Domain{}
	for _, thread := range threads {
		threadDomain := thread.toDomain()
		var likeCount int64
		var unlikeCount int64
		var replyCount int64

		rp.Conn.Table("liked_thread").Where("thread_id", threadDomain.ID).Count(&likeCount)
		threadDomain.LikeCount = int(likeCount)
		rp.Conn.Table("unliked_thread").Where("thread_id", threadDomain.ID).Count(&unlikeCount)
		threadDomain.UnlikeCount = int(unlikeCount)
		rp.Conn.Table("replies").Where("thread_id", threadDomain.ID).Count(&replyCount)
		threadDomain.ReplyCount = int(replyCount)

		threadDomains = append(threadDomains, *threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) GetThreadByTopicID(topicId, limit, offset int) ([]thread.Domain, error) {
	threads := []Thread{}
	fetchResult := rp.Conn.Unscoped().Preload("Author").Preload("Topic").Limit(limit).Offset(offset).Where("topic_id = ?", topicId).Find(&threads)
	if fetchResult.Error != nil {
		return []thread.Domain{}, fetchResult.Error
	}

	threadDomains := []thread.Domain{}
	for _, thread := range threads {
		threadDomain := thread.toDomain()
		var likeCount int64
		var unlikeCount int64
		var replyCount int64

		rp.Conn.Table("liked_thread").Where("thread_id", threadDomain.ID).Count(&likeCount)
		threadDomain.LikeCount = int(likeCount)
		rp.Conn.Table("unliked_thread").Where("thread_id", threadDomain.ID).Count(&unlikeCount)
		threadDomain.UnlikeCount = int(unlikeCount)
		rp.Conn.Table("replies").Where("thread_id", threadDomain.ID).Count(&replyCount)
		threadDomain.ReplyCount = int(replyCount)

		threadDomains = append(threadDomains, *threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) GetThreadByKeyword(keyword string, limit, offset int) ([]thread.Domain, error) {
	threads := []Thread{}
	fetchResult := rp.Conn.Unscoped().Preload("Author").Preload("Topic").Limit(limit).Offset(offset).Where("UPPER(title) LIKE UPPER(?)", "%"+keyword+"%").Find(&threads)
	if fetchResult.Error != nil {
		return []thread.Domain{}, fetchResult.Error
	}

	threadDomains := []thread.Domain{}
	for _, thread := range threads {
		threadDomain := thread.toDomain()
		var likeCount int64
		var unlikeCount int64
		var replyCount int64

		rp.Conn.Table("liked_thread").Where("thread_id", threadDomain.ID).Count(&likeCount)
		threadDomain.LikeCount = int(likeCount)
		rp.Conn.Table("unliked_thread").Where("thread_id", threadDomain.ID).Count(&unlikeCount)
		threadDomain.UnlikeCount = int(unlikeCount)
		rp.Conn.Table("replies").Where("thread_id", threadDomain.ID).Count(&replyCount)
		threadDomain.ReplyCount = int(replyCount)

		threadDomains = append(threadDomains, *threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) Like(userId int, threadId int) error {
	undoErr := rp.UndoUnlike(userId, threadId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}

	tx := rp.Conn.Begin()

	err := tx.Model(&thread).Association("LikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceThreadRepository) UndoLike(userId int, threadId int) error {
	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}

	tx := rp.Conn.Begin()

	err := tx.Model(&thread).Association("LikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceThreadRepository) UndoUnlike(userId, threadId int) error {
	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}

	tx := rp.Conn.Begin()

	err := tx.Model(&thread).Association("UnlikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceThreadRepository) Unlike(userId int, threadId int) error {
	undoErr := rp.UndoLike(userId, threadId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}

	tx := rp.Conn.Begin()

	err := tx.Model(&thread).Association("UnlikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceThreadRepository) UpdateThread(data *thread.Domain, threadId, userId int) (thread.Domain, error) {
	tx := rp.Conn.Begin()

	existingThread := Thread{}
	fetchResultErr := tx.Preload("Author").Preload("Topic").Where("author_id = ?", userId).Take(&existingThread, threadId).Error
	if fetchResultErr != nil {
		return thread.Domain{}, fetchResultErr
	}
	updatedThread := fromDomain(*data)

	if updatedThread.Content != nil {
		existingThread.Content = updatedThread.Content
	}
	existingThread.Title = updatedThread.Title
	if updatedThread.Image1 != nil {
		existingThread.Image1 = updatedThread.Image1
	}
	if updatedThread.Image2 != nil {
		existingThread.Image2 = updatedThread.Image2
	}
	if updatedThread.Image3 != nil {
		existingThread.Image3 = updatedThread.Image3
	}
	if updatedThread.Image4 != nil {
		existingThread.Image4 = updatedThread.Image4
	}
	if updatedThread.Image5 != nil {
		existingThread.Image5 = updatedThread.Image5
	}

	err := tx.Save(&existingThread).Error
	if err != nil {
		return thread.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return thread.Domain{}, err
	}

	domain := existingThread.toDomain()

	var likeCount int64
	var unlikeCount int64
	var replyCount int64

	rp.Conn.Table("liked_thread").Where("thread_id", domain.ID).Count(&likeCount)
	domain.LikeCount = int(likeCount)
	rp.Conn.Table("unliked_thread").Where("thread_id", domain.ID).Count(&unlikeCount)
	domain.UnlikeCount = int(unlikeCount)
	rp.Conn.Table("replies").Where("thread_id", domain.ID).Count(&replyCount)
	domain.ReplyCount = int(replyCount)

	return *domain, nil
}

func InitPersistenceThreadRepository(c *gorm.DB) thread.Repository {
	return &persistenceThreadRepository{
		Conn: c,
	}
}
