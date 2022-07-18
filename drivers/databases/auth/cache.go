package auth

import (
	"context"
	"fgd/core/auth"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type cacheAuthRepository struct {
	Client *redis.Client
}

func (rp *cacheAuthRepository) FetchAuth(uuid string) error {
	return rp.Client.Get(context.Background(), uuid).Err()
}

func (rp *cacheAuthRepository) StoreAuth(userId int, accessUuid, refreshUuid string, accessExpiry, refreshExpiry time.Duration) error {
	storeAccessErr := rp.Client.Set(context.Background(), accessUuid, strconv.Itoa(userId), accessExpiry).Err()
	if storeAccessErr != nil {
		return storeAccessErr
	}

	storeRefreshErr := rp.Client.Set(context.Background(), refreshUuid, strconv.Itoa(userId), refreshExpiry).Err()
	if storeRefreshErr != nil {
		return storeRefreshErr
	}

	return nil
}

func (rp *cacheAuthRepository) DeleteAuth(uuid string) error {
	return rp.Client.Del(context.Background(), uuid).Err()
}

func InitCacheAuthRepository(c *redis.Client) auth.Repository {
	return &cacheAuthRepository{
		Client: c,
	}
}
