package main

import (
	"net/http"
	"time"
	"log"
	"io"
)
func main() {
	server := http.Server{
		Addr:":8080",
		Handler: &myHandler3{},
		ReadTimeout:5 * time.Second,
	}

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

type myHandler3 struct{}

func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"URL: "+r.URL.String())
}





