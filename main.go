package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	InitRedis()
	sub1 := r.Subscribe(context.Background(), "chanel").Channel()
	sub2 := r.Subscribe(context.Background(), "chanel").Channel()
	go func() {
		for {
			select {
			case msg := <-sub2:
				fmt.Println(msg)
			}
		}
	}()
	go func() {
		for {
			select {
			case msg := <-sub1:
				fmt.Println(msg)
			}
		}
	}()
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			r.Publish(context.Background(), "chanel", fmt.Sprintf("msg %d", i))
			i++
		}

	}()
	select {}
}

var r *redis.Client

func InitRedis() {
	r = redis.NewClient(&redis.Options{
		Addr: "192.168.69.65:6379",
	})
	err := r.Set(context.Background(), "k", "v1", 0).Err()
	if err != nil {
		fmt.Println("redis init failed")
		fmt.Println(err)
		return
	}
	fmt.Println("redis init success")
}
