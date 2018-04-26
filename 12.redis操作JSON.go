package main

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
)

func main() {
	c, err := redis.Dial("tcp","localhost:6379")
	if err != nil{
		fmt.Println("connect to redis failed!",err)
		return
	}
	defer c.Close()
	// convert map to json
	key := "jsontest"
	imap := map[string]string{"username":"dongge","password":"666"}
	value,_ := json.Marshal(imap)
	// save json to redis
	n, err := c.Do("setnx",key,value)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(111,n)
	if n == int64(1){
		fmt.Println("success!")
	}
	// get json from redis
	var imapGet map[string]string
	valueGet, err:= redis.Bytes(c.Do("get",key))
	if err!=nil{
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil{
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}