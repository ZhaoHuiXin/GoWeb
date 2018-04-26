package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"database/sql"
	"regexp"
	_ "github.com/go-sql-driver/mysql"
)
// 连接mysql 返回DB对象
//func Connectdb() (res interface{}) {
//	db, err := sql.Open("mysql", "zhx:666666@tcp" +
//		"(localhost:3306)/test?charset=utf8")
//		if err != nil{
//			fmt.Println("connect to mysql failed!", err)
//			return
//		}
//		return db
//	}
func main() {

	path :="/home/zhx/access.log"
	Read(path)

}

func Read (path string){
	// connect to db
	db, err := sql.Open("mysql", "zhx:666666@tcp" +
		"(localhost:3306)/test?charset=utf8")
	if err != nil{
		fmt.Println("connect to mysql failed!", err)
		return
	}
	stmt, err := db.Prepare("insert into wftest(url) values(?)")

	// regexp format
	r := regexp.MustCompile("http://(.*)/ca")
	// open file
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	// read file
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		data := r.FindStringSubmatch(string(a))
		if len(data) >= 2{
			fmt.Println(data[1])
			stmt.Exec(data[1])
		}
	}
}
