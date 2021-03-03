package main

import (
	"fmt"
	"time"

	pool "github.com/ichunt2019/go-redis-pool"
)

func main() {
	pool, err := pool.NewHA(&pool.HAConfig{
		Master: "192.168.1.235:6379",
		Slaves: []string{
			"192.168.1.235:6379",
			"192.168.1.237:6379",
		},
		Password:"xxxxx",
		ReadonlyPassword:"xxxxxxx",
		

		// optional
		AutoEjectHost:      true,
		ServerFailureLimit: 3,
		ServerRetryTimeout: 5 * time.Second,
		MinServerNum:       1,
	})
	if err != nil {
		fmt.Println(err)
		// log the error
	}
	pool.Set("foo12345678", "bar44", 0)

	fmt.Println(pool.Get("foo12345678"))
	redismaster,_ :=pool.WithMaster()
	fmt.Println(redismaster.Get("foo12345678"));
}
