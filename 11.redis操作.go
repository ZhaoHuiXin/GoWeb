package main
import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)
func main() {
	// connect to redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	checkErr(err)

	defer  c.Close()
	// set but never expire
	//_,err = c.Do("set","s1","donge")
	// set and expire
	_,err = c.Do("set","s1","donge","ex","5")
	if err!= nil{
		fmt.Println("redis set failed", err)	}
	// check key is exist or not
	is_key_exists, err := redis.Bool(c.Do("exists","s1"))
	if err!=nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("exists or not: %v ! \n", is_key_exists)
	}
	// get s1
	username, err := redis.String(c.Do("get","s1"))
	if err != nil{
		fmt.Println("redis get failed", err)
	}else{
		fmt.Println(111111)
		fmt.Printf("get s1: %v \n", username)
	}
	// wait a moment
	time.Sleep(8 * time.Second)
	// check key is exist or not
	is_key_exists, err = redis.Bool(c.Do("exists","s1"))
	if err!=nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("exists or not: %v ! \n", is_key_exists)
	// get s1
	username, err = redis.String(c.Do("get","s1"))
	if err != nil{
		fmt.Println("redis get failed", err)
	} else {
		fmt.Printf("get mykey:%v \n", username)
	}
}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}