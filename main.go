package main

import (
	"fmt"
	"net/http"
)
import "upload_cha/handler"

func main(){
	http.HandleFunc("/",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	http.HandleFunc("/file/download",handler.DownloadHandler)
	http.HandleFunc("/file/update",handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)
	http.HandleFunc("/show",handler.FillShowAll)
	err:=http.ListenAndServe(":8080",nil)

	if err!=nil{
		fmt.Print("failed to start server,err:%s",err.Error())
	}
}