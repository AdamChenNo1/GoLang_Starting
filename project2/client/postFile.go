package client

import(
	"fmt"
	"os"
	"bytes"
	"mime/multipart"
	"net/http"
	"io"
	"io/ioutil"
)

func PostFile(filename string,target_url string) error{
	bodyBuf := &bytes.Buffer{}
	bodyWritter := multipart.NewWriter(bodyBuf)

	fileWritter,err := bodyWritter.CreateFormFile("uploadfile",filename)
	if err != nil {
		fmt.Println("error writting to buffer")
		return err
	}

	//打开文件句柄
	fh,err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	_,err1 := io.Copy(fileWritter,fh)
	if err1 != nil {
		return err1
	}

	contentType := bodyWritter.FormDataContentType()
	bodyWritter.Close()

	resp,err := http.Post(target_url,contentType,bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
