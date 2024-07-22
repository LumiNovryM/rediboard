package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	// "github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Service Running...")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       15,
	})

	ctx := context.TODO()
	client.Set(ctx, "language", "Go", 0)
	language := client.Get(ctx, "language")
	year := client.Get(ctx, "year")

	fmt.Println(language.Val()) // "Go"
	fmt.Println(year.Val())     // ""
}
