package main

import(
	"net/http" // go 的http内核
	"log"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello, this is version 1!"))

	})
	http.HandleFunc("/bye",sayBye)
	log.Println("starting server..")
	log.Fatal(http.ListenAndServe(":4000",nil)) // 开始监听时发生错误才调用;这是推出程序的写法，习惯性用例
}

func sayBye(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("byebye, this is the version 1"))
}