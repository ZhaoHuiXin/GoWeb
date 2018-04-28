package main

import (
	"encoding/json"
	"io"
	"strings"
	"log"
	"fmt"
)

type User struct {
UserName json.RawMessage `json:"username"`
Password string `json:"password"`

Email string
Phone int64
}

var jsonString string = `{
    "username": "phpgo@163.com",
    "password": "123"
}`

func Decode(r io.Reader) (u *User, err error) {
u = new(User)
if err = json.NewDecoder(r).Decode(u); err != nil {
return
}

var email string
if err = json.Unmarshal(u.UserName, &email); err == nil {
u.Email = email
return
}

var phone int64
if err = json.Unmarshal(u.UserName, &phone); err == nil {
u.Phone = phone
}

return
}

func main() {
user, err := Decode(strings.NewReader(jsonString))
if err != nil {
log.Fatal(err)
}
fmt.Printf(string(user.UserName))
//fmt.Println(user)
}