package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	// 自定义mux控制路由的访问
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello",sayHello2)

	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"URL "+r.URL.String())
}

func sayHello2(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello version 2.")
}

