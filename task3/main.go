package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	log.Printf("Connected to Redis: %s", pong)


	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	for {
		data, err := rdb.LPop(ctx, "data").Result()
		if err != nil && err != redis.Nil {
			log.Printf("Error reading from Redis: %v", err)
			continue
		}
		if data != "" {
			_, err := file.WriteString(data + "\n")
			if err != nil {
				log.Printf("Error writing to file: %v", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
