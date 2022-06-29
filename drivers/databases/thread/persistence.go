package thread

import (
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
	fetchTopic := rp.Conn.Take(&topic, topicId)
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

	return *newThread.toDomain(), nil
}

func (rp *persistenceThreadRepository) DeleteThread(userId int, threadId int) error {
	res := rp.Conn.Where("author_id = ?", userId).Delete(&Thread{}, threadId)
	return res.Error
}

func (rp *persistenceThreadRepository) GetThreadByAuthorID(userId, limit, offset int) ([]thread.Domain, error) {
	threads := []Thread{}
	fetchResult := rp.Conn.Limit(limit).Offset(offset).Where("author_id = ?", userId).Find(&threads)
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
		rp.Conn.Table("reply").Where("thread_id", threadDomain.ID).Count(&replyCount)
		threadDomain.ReplyCount = int(replyCount)

		threadDomains = append(threadDomains, *threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) GetThreadByTopicID(topicId, limit, offset int) ([]thread.Domain, error) {
	threads := []Thread{}
	fetchResult := rp.Conn.Limit(limit).Offset(offset).Where("topic_id = ?", topicId).Find(&threads)
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
		rp.Conn.Table("reply").Where("thread_id", threadDomain.ID).Count(&replyCount)
		threadDomain.ReplyCount = int(replyCount)

		threadDomains = append(threadDomains, *threadDomain)
	}

	return threadDomains, nil
}

func (rp *persistenceThreadRepository) Like(userId int, threadId int) error {
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
	thread := Thread{Model: gorm.Model{ID: uint(threadId)}}
	return rp.Conn.Model(&thread).Association("UnlikedBy").Append(&user.User{Model: gorm.Model{ID: uint(userId)}})
}

func (rp *persistenceThreadRepository) UpdateThread(data *thread.Domain, threadId, userId int) (thread.Domain, error) {
	existingThread := Thread{}
	fetchResult := rp.Conn.Where("author_id = ?", userId).Take(&existingThread, threadId)
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

	return *existingThread.toDomain(), nil
}

func InitPersistenceThreadRepository(c *gorm.DB) thread.Repository {
	return &persistenceThreadRepository{
		Conn: c,
	}
}
