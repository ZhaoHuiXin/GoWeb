package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Account struct {
	Email string
	password string
	Money float64
}

func main() {
	account := Account{
		Email: "phpgo@163.com",
		password: "123456",
		Money: 100.5,
	}

	rs, err := json.Marshal(account)
	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
}