package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	opt := redis.Options{
		Addr:    os.Args[1],
		Network: "tcp",
	}
	c := redis.NewClient(&opt)
	valueSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	byteValue := [6000]byte{}
	for i := 0; i < 100_000; i++ {
		status := c.Set(strconv.Itoa(i), byteValue[:valueSize], time.Hour)
		if status.Err() != nil {
			log.Fatal(status.Err())
		}
	}

}
