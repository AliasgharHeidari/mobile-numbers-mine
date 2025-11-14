package dataonredis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/config"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedisClient() {
	if RedisClient != nil {
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:        config.AppConfig.Redis.Address,
		Password:    config.AppConfig.Redis.Password,
		DB:          config.AppConfig.Redis.DB,
		DialTimeout: time.Duration(config.AppConfig.Redis.Timeout) * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Panicf("Failed to connect to redis client, error: %+v", err)
	}
}

func SaveUserToRedis(user model.User) error {
	key := user.ID
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return RedisClient.Set(context.Background(), string(key), data, 0).Err()
}

func LoadUserFromRedis(userID int) (*model.User, error) {
	val, err := RedisClient.Get(context.Background(), string(userID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // User not found
		}
		return nil, err
	}
	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsersFromRedis() ([]model.User, error) {
	var users []model.User

	keys, err := RedisClient.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		val, err := RedisClient.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		var user model.User
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func DeleteUserFromRedis(userID int) error {
	return RedisClient.Del(context.Background(), string(userID)).Err()
}
