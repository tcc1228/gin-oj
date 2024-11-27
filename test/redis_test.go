package test

import (
	"context"
	"fmt"
	"gin-gorm-oj/models"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestRedisSet(t *testing.T) {
	rdb.Set("key", "mmc", time.Second*10)
}

func TestRedisGet(t *testing.T) {
	v, err := rdb.Get("key").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestRedisGetByModel(t *testing.T) {
	v, err := models.RDB.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
