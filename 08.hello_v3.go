package main

import (
	"net/http"
	"time"
	"log"
	"io"
)

// 自定义路由
var mux map[string]func(http.ResponseWriter, *http.Request)
func main() {
	server := http.Server{
		Addr:":8080",
		Handler: &myHandler3{},
		ReadTimeout:5 * time.Second,
	}
	// 路由注册
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = sayHello3
	mux["/bye"] = sayBye3
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

type myHandler3 struct{}

func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// 判断，如果mux中有Request中的路径,找到的话ok为true
	// h为路径对应的函数func(http.ResponseWriter, *http.Request)
	// 直接将w，r转发到该函数即可，即调用该函数h（sayHello和sayBye）
	if h, ok:=mux[r.URL.String()]; ok{
		h(w, r)
		return
	}
	io.WriteString(w,"URL: "+r.URL.String())
}

func sayHello3(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello this is version 3")
}

func sayBye3(w http.ResponseWriter,r *http.Request){
	io.WriteString(w, "bye bye, this is version 3")
}





