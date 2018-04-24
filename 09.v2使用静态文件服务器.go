package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"fmt"
)

func main() {
	// 自定义mux控制路由的访问
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler4{})
	mux.HandleFunc("/hello",sayHello4)
	// os的方法，获取当前路径
	wd, err := os.Getwd()
	if err!= nil{
		log.Fatal(err)
	}
	// 如果没有static目录，http去除前缀，留下handler(http.FileServer)
	// http.FileServer()要求传进绝对路径
	// http.Dir(wd)会根据当前路径获取绝对路径
	// 静态路径中有index.html默认渲染
	lu := http.FileServer(http.Dir(wd))
	fmt.Println(lu)
	mux.Handle("/static/",
				http.StripPrefix("/static/",lu))

	err = http.ListenAndServe(":4000", mux)
	if err != nil{
		log.Fatal(err)
	}
}

type myHandler4 struct{}

func (*myHandler4)ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"URL "+r.URL.String())
}

func sayHello4(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello version 2.")
}

