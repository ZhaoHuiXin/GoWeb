package main

import(
	"net/http" // go 的http内核
	"log"
	"time"
	"os"
	"os/signal"
)

func main() {
	server := &http.Server{
		Addr:":4001",
		WriteTimeout:2*time.Second,
	}

	quit := make(chan os.Signal) //主动停止服务器
	signal.Notify(quit, os.Interrupt) // 注册通知事件（比如接收到signal就发送对象到channel

	go func(){
		<-quit
		if err := server.Close();err!=nil{
			log.Println("CLOSE SERVER",err)
		}
	}()

	mux := http.NewServeMux() // mux是实现了Handler接口的一个变量
	mux.Handle("/",&myHandler{}) // Handle第一参数和HandlerFunc无区别，第二个参数为handler类型
	mux.HandleFunc("/bye",sayBye1)
	server.Handler = mux // server.Handler类型也是Handler接口

	log.Println("starting server..v4")
	err:=server.ListenAndServe()
	if err!=nil{
		// 人为主动想要关闭服务器
		if err == http.ErrServerClosed{
			log.Println("server closed under request")
		}else{
		log.Println("server closed unexpeccted")
		}
		log.Fatal("server exit....")

	}
	log.Fatal(server.ListenAndServe()) // 开始监听时发生错误才调用;这是推出程序的写法，习惯性用例
}
type myHandler struct{}

func (_ *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello!"+r.URL.String()))
}

func sayBye1(w http.ResponseWriter, r *http.Request){
	time.Sleep(3 * time.Second)
	w.Write([]byte("bye bye, this is the version 3"))
}



