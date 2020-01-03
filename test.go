package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	resp,err:=http.Get(imgUrl)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	out,err:=os.Create("lalaalala.jpg")
	if err!=nil{
		panic(out)
	}
	defer out.Close()

	_,err=io.Copy(out,resp.Body)
	if err!=nil{
		panic(err)
	}

}