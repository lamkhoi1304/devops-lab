package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func rClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return client
}

func ping(client *redis.Client) error {
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	log.Println("Connect redis success")

	return nil
}

func main() {
	client := rClient()
	for {
		err := ping(client)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
