package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"upload_cha/meta"
	"upload_cha/util"
)


func UploadHandler(w http.ResponseWriter,request *http.Request){
	if request.Method=="GET"{
		data,err:=ioutil.ReadFile("./static/view/index.html")
		if err!=nil{
			io.WriteString(w,"internal server error")
			return
		}
		io.WriteString(w,string(data))
	}else if request.Method=="POST"{
		file,head,err:=request.FormFile("file")
		if err!=nil{
			fmt.Printf("failed to %s\n",err.Error())
		}
		defer file.Close()
		fileMeta:=meta.FileMeta{
			FileName:head.Filename,
			Location:"/Users/alchemy/Documents/fileChan/"+head.Filename,

			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile,err:=os.Create(fileMeta.Location)
		if err!=nil{
			fmt.Printf("Failed to creat,err is %s\n",err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize,err = io.Copy(newFile,file)
		if err!=nil{
			fmt.Printf("Failed to save data into file,err is %#v",err.Error())
			return
		}
		newFile.Seek(0,0)
		fileMeta.FileShal=util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w,request,"/file/upload/suc",http.StatusFound)

	}
}



func UploadSucHandler(w http.ResponseWriter,request *http.Request){
	io.WriteString(w,"upload finished")
}

//GetFileMetaHandler 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter,request *http.Request){
	request.ParseForm()
	filehash:=request.Form["filehash"][0]
	fMeta:=meta.GetFileMeta(filehash)
	data,err:=json.Marshal(fMeta)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}



func DownloadHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fsha1:=r.Form.Get("filehash")
	fm:=meta.GetFileMeta(fsha1)
	f,err:=os.Open(fm.Location)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data,err:=ioutil.ReadAll(f)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("Content-Disposition","attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

func FileMetaUpdateHandler(w http.ResponseWriter,request *http.Request){
	request.ParseForm()
	opType:=request.Form.Get("op")
	fileSha1:=request.Form.Get("filehash")
	newFileName:=request.Form.Get("filename")
	if opType!="0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	curFileMeta:=meta.GetFileMeta(fileSha1)
	curFileMeta.FileName=newFileName
	meta.UpdateFileMeta(curFileMeta)
	data,err:=json.Marshal(curFileMeta)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}


//删除索引
func FileDeleteHandler(w http.ResponseWriter,request *http.Request){
	request.ParseForm()
	fileSha1:=request.Form.Get("filehash")

	fMeta:=meta.GetFileMeta(fileSha1)
	err:=os.Remove(fMeta.Location)
	if err!=nil{
		io.WriteString(w,err.Error())
		return
	}

	meta.RemoveFileMeta(fileSha1)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w,"删除成功")

}

func FillShowAll(w http.ResponseWriter,request *http.Request){
	z:=meta.GetAllFiles()
	alchemy, _ :=json.Marshal(z)
	w.Write(alchemy)
}

