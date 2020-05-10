package main

import(
	"net/http"
	"log"
	."project2"
	."project2/client"	
)

func main(){
	http.HandleFunc("/",SayHelloName)//设置主页路由
	http.HandleFunc("/login",Login)//设置登录路由
	http.HandleFunc("/upload", Upload)//设置上传文件路由
	err := http.ListenAndServe(":9090",nil)//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
	uploadtest()
}

func uploadtest() {
	target_url := "http://localhost:9090/upload"
	filename := "./uploadedWengan Chen 陈文干-副本.jpg"
	PostFile(filename, target_url)
}