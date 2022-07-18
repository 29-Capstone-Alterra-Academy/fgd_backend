package drivers

import (
	authDomain "fgd/core/auth"
	moderatorDomain "fgd/core/moderator"
	replyDomain "fgd/core/reply"
	reportDomain "fgd/core/report"
	searchDomain "fgd/core/search"
	threadDomain "fgd/core/thread"
	topicDomain "fgd/core/topic"
	userDomain "fgd/core/user"
	verifyDomain "fgd/core/verify"
	"fgd/drivers/databases/auth"
	"fgd/drivers/databases/moderator"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/report"
	"fgd/drivers/databases/search"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"fgd/drivers/databases/verify"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

func NewAuthRepository(c *redis.Client) authDomain.Repository {
	return auth.InitCacheAuthRepository(c)
}

func NewModeratorRepository(c *gorm.DB) moderatorDomain.Repository {
	return moderator.InitPersistenceModeratorRepository(c)
}

func NewReportRepository(c *gorm.DB) reportDomain.Repository {
	return report.InitPersistenceReportRepository(c)
}

func NewReplyRepository(c *gorm.DB) replyDomain.Repository {
	return reply.InitPersistenceReplyRepository(c)
}

func NewSearchRepository(c *gorm.DB) searchDomain.Repository {
	return search.InitPersistenceSearchRepository(c)
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
