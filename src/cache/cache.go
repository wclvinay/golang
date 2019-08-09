package cache
import (
     "fmt"
     "github.com/go-redis/redis"
)
var Cacheobj *redis.Client
func init() {
	Cacheobj = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       13,  // use default DB
	})
	_,err := Cacheobj.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("redis connected")
	}
}