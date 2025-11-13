package onmemory

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"sync"
	"time"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	"github.com/go-redis/redis/v8"
)

type UserData struct {
	Users           []model.User
	UserChangeMutex *sync.Mutex
}

var UsersRepo UserData

var (
	RedisClient *redis.Client
	redisKey    = "users_repo"
)

func initRedisClient() {
	if RedisClient != nil {
		return
	}
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return
	}
	pass := os.Getenv("REDIS_PASS")
	db := 0
	if dbStr := os.Getenv("REDIS_DB"); dbStr != "" {
		if v, err := strconv.Atoi(dbStr); err == nil {
			db = v
		}
	}


	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    pass,
		DB:          db,
		DialTimeout: 5 * time.Second,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := client.Ping(ctx).Result(); err != nil {
	
		return
	}
	RedisClient = client
}

func saveUsersToRedis() error {
	if RedisClient == nil {
		return nil
	}
	ctx := context.Background()
	data, err := json.Marshal(UsersRepo.Users)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, redisKey, data, 0).Err()
}

func loadUsersFromRedis() (bool, error) {
	if RedisClient == nil {
		return false, nil
	}
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil // Redis key does not exist
		}
		return false, err
	}
	var users []model.User
	if err := json.Unmarshal([]byte(val), &users); err != nil {
		return false, err
	}
	UsersRepo.Users = users

	if UsersRepo.UserChangeMutex == nil {
		UsersRepo.UserChangeMutex = &sync.Mutex{}
	}
	return true, nil
}

func init() {
	initRedisClient()
	if UsersRepo.UserChangeMutex == nil {
		UsersRepo.UserChangeMutex = &sync.Mutex{}
	}
}

func InitRedis() bool {
	initRedisClient()
	return RedisClient != nil
}

func SaveUsers(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if RedisClient == nil {
		return nil
	}
	data, err := json.Marshal(UsersRepo.Users)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, redisKey, data, 0).Err()
}

func LoadUsersFromRedis() (bool, error) {
	return loadUsersFromRedis()
}

func LoadInitUsers() {

	initRedisClient()

	if ok, err := loadUsersFromRedis(); err == nil && ok {
		return
	}

	UsersRepo = UserData{}
	UsersRepo.Users = []model.User{
		{
			ID:         1,
			Name:       "Ali",
			FamilyName: "Heidari",
			Age:        18,
			IsMarried:  false,
		},

		{
			ID:         2,
			Name:       "Amir",
			FamilyName: "Barkhordari",
			Age:        21,
			IsMarried:  false,
		},
	}

	UsersRepo.UserChangeMutex = &sync.Mutex{}
	_ = saveUsersToRedis()
}
