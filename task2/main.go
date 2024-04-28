package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB:0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	key := "goList"
	err = redisClient.RPush(key, "item1", "item2", "item3").Err()
	if err != nil {
		fmt.Println("Error pushing data into list:", err)
		return
	}
	fmt.Println("Data pushed into Redis list")

	hashKey := "myhash"
	err = redisClient.HSet(hashKey, hashKey, map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}).Err()
	if err != nil {
		fmt.Println("Error putting data into Redis hash map:", err)
		return
	}
	fmt.Println("Data put into Redis hash map successfully")

	hashData, err := redisClient.HGetAll(hashKey).Result()
	if err != nil {
		fmt.Println("Error reading hash map data from Redis:", err)
		return
	}
	fmt.Println("Data from Redis hash map:")
	for field, value := range hashData {
		fmt.Printf("%s: %s\n", field, value)
	}

	poppedItem, err := redisClient.LPop(key).Result()
	if err != nil {
		fmt.Println("Error popping data from list:", err)
		return
	}
	fmt.Println("Popped item from Redis list:", poppedItem)
}
