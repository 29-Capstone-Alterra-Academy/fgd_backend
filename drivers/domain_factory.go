package drivers

import (
	authDomain "fgd/core/auth"
	replyDomain "fgd/core/reply"
	threadDomain "fgd/core/thread"
	topicDomain "fgd/core/topic"
	userDomain "fgd/core/user"
	verifyDomain "fgd/core/verify"
	"fgd/drivers/databases/auth"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"fgd/drivers/databases/verify"

	"gorm.io/gorm"
)

func NewAuthRepository(c *gorm.DB) authDomain.Repository {
	return auth.InitPersistenceAuthRepository(c)
}

func NewReplyRepository(c *gorm.DB) replyDomain.Repository {
	return reply.InitPersistenceReplyRepository(c)
}

func NewThreadRepository(c *gorm.DB) threadDomain.Repository {
	return thread.InitPersistenceThreadRepository(c)
}

func NewTopicRepository(c *gorm.DB) topicDomain.Repository {
	return topic.InitPersistenceTopicRepository(c)
}

func NewUserRepository(c *gorm.DB) userDomain.Repository {
	return user.InitPersistenceUserRepository(c)
}

func NewVerifyRepository(c *gorm.DB) verifyDomain.Repository {
	return verify.InitPersistenceVerifyRepository(c)
}