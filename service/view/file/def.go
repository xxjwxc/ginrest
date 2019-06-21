package file

//
type FileInfo struct {
	File_name string `json:"file_name"` //文件地址
}

//
type DeleteParam struct {
	File_name string `json:"file_name"` //文件地址
	Business  string `json:"business"`  //文件所在业务
}
