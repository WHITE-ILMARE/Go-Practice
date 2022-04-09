package main

import "github.com/garyburd/redigo/redis"

var (
	Pool *redis.Pool
)

func init() {
	redisHost := "http://localhost:6379"
	Pool = redis.NewPool()
}
