package project2

import(
	"fmt"
	//"text/template"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func SayHelloName(w http.ResponseWriter,r *http.Request){
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

func Login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)//获取请求的方法
	if r.Method == "GET"{
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))

		t,_ := template.ParseFiles("src/project2/views/login.gtpl")
		log.Println(t.Execute(w,token))
	}else{
		//请求的是登录数据，则执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != ""{
			//验证token的合法性
		}else{
			//不存在token报错
		}
		fmt.Println("username length:",len(r.Form["username"][0]))
		fmt.Println("username:",template.HTMLEscaper(r.Form["username"]))
		fmt.Println("password:",template.HTMLEscapeString(r.FormValue("password")))
		template.HTMLEscape(w,[]byte(r.Form.Get("username")))
	}
}
/*
func pwn(w http.ResponseWriter,r *http.Request){
	t,err := template.New("foo").Parse(`{{define "T}}Hello,{{.}}!{{end}}`)
	err = t.ExecuteTemplate(out,"T","<script>alert('you have been pwned')</script>")
}
*/
