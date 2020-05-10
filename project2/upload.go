package project2

import(
	"fmt"
	"net/http"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"html/template"
	"log"
	"os"
)

//处理upload逻辑
func Upload(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)//获取请求的方法
	if r.Method == "GET"{
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))

		t,_ := template.ParseFiles("src/project2/views/upload.gtpl")
		log.Println(t.Execute(w,token))
	}else{
		r.ParseMultipartForm(32 << 20)
		file,handler,err := r.FormFile("uploadfile")
		if err != nil{
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f,err := os.OpenFile("src/project2/uploaded/" + handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil{
			fmt.Println(err)
			return
		}
		defer file.Close()
		io.Copy(f,file)
	}
}