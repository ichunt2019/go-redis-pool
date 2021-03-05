package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
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

		Password:"xxx",
		ReadonlyPassword:"xxxx",
		Options:&redis.Options{
			DialTimeout:10*time.Second,//连接超时
			MinIdleConns:10,//空闲链接数
			ReadTimeout:10*time.Second,
			WriteTimeout: 30 * time.Second,
		},

		// optional
		AutoEjectHost:      true,//是否弹出故障主机
		ServerFailureLimit: 3,//达到失败次数时弹出
		ServerRetryTimeout: 5 * time.Second,//在“ServerRetryTimeout”之后重试弹出的主机`
		MinServerNum:       1,//保留min服务器 针对从服务器
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
