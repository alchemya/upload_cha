package meta


//FileMeta 文件元信息结构
type FileMeta struct{
	FileShal string
	FileName string
	FileSize int64
	Location string
	UploadAt string

}



var fileMetas map[string]FileMeta

func init(){
	fileMetas=make(map[string]FileMeta)
}


//updateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta){
	fileMetas[fmeta.FileShal]=fmeta
}



func GetFileMeta(fileSha1 string)FileMeta{
	return fileMetas[fileSha1]
}


//删除信息
func RemoveFileMeta(fileSha1 string){
	delete(fileMetas,fileSha1)

}



