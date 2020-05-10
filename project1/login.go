package main

import(
	"fmt"
	"text/template"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()//解析url传递的参数，对于POST则解析响应包的主体
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key",k)
		fmt.Println("value",strings.Join(v,""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)//获取请求的方法
	if r.Method == "GET"{
		t,_ := template.ParseFiles("src/project1/login.gtpl")
		log.Println(t.Execute(w,nil))
	}else{
		//请求的是登录数据，则执行登录的逻辑判断
		r.ParseForm()
		fmt.Println("username:",template.HTMLEscaper(r.Form["username"]))
		fmt.Println("password:",template.HTMLEscapeString(r.FormValue("password")))
		template.HTMLEscape(w,[]byte(r.Form.Get("username")))
	}
}

func pwn(w http.ResponseWriter,r *http.Request){
	t,err := template.New("foo").Parse(`{{define "T}}Hello,{{.}}!{{end}}`)
	err = t.ExecuteTemplate(out,"T","<script>alert('you have been pwned')</script>")
}

func main(){
	http.HandleFunc("/",sayHelloName)//设置主页路由
	http.HandleFunc("/login",login)//设置登录路由
	err := http.ListenAndServe(":9090",nil)//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}