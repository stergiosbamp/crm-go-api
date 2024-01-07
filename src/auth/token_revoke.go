package auth

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const BlacklistNamespace string = "blacklist:"
const BlacklistExpirationTime time.Duration = time.Minute * TOKEN_EXP_MINS // Redis TTL cleanup for blacklisted tokens, must be greater than token expiration time

// TokenRevoker is an interface that handles revoking tokens for server-side loggout
type TokenRevoker interface {
	RevokeToken(token string) error
	IsTokenRevoked(token string) bool
}

type RedisTokenRevoker struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewRedisTokenRevoker(ctx context.Context) *RedisTokenRevoker {

	cl := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisTokenRevoker{
		redisClient: cl,
		ctx:         ctx,
	}
}

func (revoker *RedisTokenRevoker) RevokeToken(token string) error {

	key := BlacklistNamespace + token

	return revoker.redisClient.Set(revoker.ctx, key, nil, BlacklistExpirationTime).Err()
}

func (revoker *RedisTokenRevoker) IsTokenRevoked(token string) bool {
	key := BlacklistNamespace + token

	// If key exists, token is revoked
	_, err := revoker.redisClient.Get(revoker.ctx, key).Result()
	return err == nil
}
