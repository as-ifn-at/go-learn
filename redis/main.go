package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB: 0,
	})

	ping, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("ping: %v\n", ping)
	}

	err = client.Set(context.Background(),"key", "value", time.Second * 10).Err()
	if err != nil {
		fmt.Printf("failed to set value err: %v\n", err)
	}

	time.Sleep(3 * time.Second)

	val, err := client.Get(context.Background(), "key").Result()
	if err != nil {
		fmt.Printf("failed to get value err: %v\n", err)
	}
	fmt.Printf("val: %v\n", val)

}
