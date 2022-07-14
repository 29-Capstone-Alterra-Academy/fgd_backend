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
	topic := topic.Topic{}
	fetchTopic := rp.Conn.Unscoped().Take(&topic, topicId)
	if fetchTopic.Error != nil {
		return thread.Domain{}, fetchTopic.Error
	}

	author := user.User{}
	fetchAuthor := rp.Conn.Take(&author, userId)
	if fetchAuthor.Error != nil {
		return thread.Domain{}, fetchAuthor.Error
	}
	newThread := fromDomain(*data)
	newThread.Topic = topic
	newThread.Author = author

	res := rp.Conn.Create(&newThread)
	if res.Error != nil {
		return thread.Domain{}, res.Error
	}

	return newThread.toDomain(), nil
}

func (rp *persistenceThreadRepository) DeleteThread(userId int, threadId int) error {
	res := rp.Conn.Where("author_id = ?", userId).Delete(&Thread{}, threadId)
	return res.Error
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

	return domain, res.Error
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

		threadDomains = append(threadDomains, threadDomain)
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

		threadDomains = append(threadDomains, threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) GetThreadByKeyword(keyword string, limit, offset int) ([]thread.Domain, error) {
	threads := []Thread{}
	fetchResult := rp.Conn.Unscoped().Preload("Author").Preload("Topic").Limit(limit).Offset(offset).Where("title LIKE ?", keyword+"%").Find(&threads)
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

		threadDomains = append(threadDomains, threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) Like(userId int, threadId int) error {
	undoErr := rp.UndoUnlike(userId, threadId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}
	return rp.Conn.Model(&thread).Association("LikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceThreadRepository) UndoLike(userId int, threadId int) error {
	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}
	return rp.Conn.Model(&thread).Association("LikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceThreadRepository) UndoUnlike(userId, threadId int) error {
	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}
	return rp.Conn.Model(&thread).Association("UnlikedBy").Delete(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceThreadRepository) Unlike(userId int, threadId int) error {
	undoErr := rp.UndoLike(userId, threadId)
	if undoErr != nil && !errors.Is(undoErr, gorm.ErrRecordNotFound) {
		return undoErr
	}

	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}
	return rp.Conn.Model(&thread).Association("UnlikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceThreadRepository) UpdateThread(data *thread.Domain, threadId, userId int) (thread.Domain, error) {
	existingThread := Thread{}
	fetchResult := rp.Conn.Preload("Author").Preload("Topic").Where("author_id = ?", userId).Take(&existingThread, threadId)
	if fetchResult.Error != nil {
		return thread.Domain{}, fetchResult.Error
	}
	updatedThread := fromDomain(*data)

	existingThread.Content = updatedThread.Content
	existingThread.Title = updatedThread.Title
	existingThread.Image1 = updatedThread.Image1
	existingThread.Image2 = updatedThread.Image2
	existingThread.Image3 = updatedThread.Image3
	existingThread.Image4 = updatedThread.Image4
	existingThread.Image5 = updatedThread.Image5

	res := rp.Conn.Save(&existingThread)
	if res.Error != nil {
		return thread.Domain{}, res.Error
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

	return domain, nil
}

func InitPersistenceThreadRepository(c *gorm.DB) thread.Repository {
	return &persistenceThreadRepository{
		Conn: c,
	}
}
