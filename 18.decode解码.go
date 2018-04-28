package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var jsonString string = `{
    "username": "phpgo@163.com",
    "password": "123"
}`

func Decode(r io.Reader) (u *User, err error) {
	u = new(User)
	err = json.NewDecoder(r).Decode(u)
	if err != nil {
		return
	}
	return
}

func main() {
	user, err := Decode(strings.NewReader(jsonString))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", *user)
	fmt.Printf(user.UserName)
	fmt.Println()
	fmt.Printf(user.Password)
	fmt.Println()

}