package file

import (
	"net/http"
	"path"
	"regexp"
	"strings"

	"github.com/xie1xiao1jun/ginrest/service/config"
	"github.com/xie1xiao1jun/public/message"
	"github.com/xie1xiao1jun/public/myhttp"
)

//
var O_file File

//
type File struct {
}

/*
	上传文件
*/
func (u *File) Upload(w http.ResponseWriter, r *http.Request) {
	//form提交
	if r.Method == "POST" {
		var paths []string
		result, tmp_paths := myhttp.UploadMoreFile(r, "")
		if !result {
			myhttp.WriteJson(w, message.GetErrorMsg(message.UnknownError))
			return
		}
		if len(tmp_paths) > 0 {
			//返回绝对路径
			for _, v := range tmp_paths {
				if !strings.Contains(config.GetDomainName(), "https") {
					paths = append(paths, config.GetDomainName()+":"+config.GetServerPort()+v)
				} else {
					paths = append(paths, config.GetDomainName()+v)
				}
			}
		}
		msg := message.GetSuccessMsg(message.NormalMessageId)
		msg.Data = paths
		myhttp.WriteJson(w, msg)
	} else {
		myhttp.WriteJson(w, message.MessageBody{State: false, Code: 405, Error: "method not allowed", Data: nil})
		return
	}
}

func removeStrArray(old, split, deletes string) (news string, ret bool) {
	//判断是否包含要删除的
	b1, _ := regexp.MatchString(deletes, old)
	if b1 {
		//判断是否包含分割号
		b2, _ := regexp.MatchString(split, old)
		if b2 {
			array := strings.Split(old, split)
			//移除元素
			for k, v := range array {
				if v == deletes {
					array = append(array[:k], array[k+1:]...)
				}

			}
			//拼接
			for i := 0; i < len(array); i++ {
				if i != len(array)-1 {
					news += array[i] + split
				} else {
					news += array[i]
				}

			}
			return news, true
		}
	}

	return "", false
}

//获取文件后缀
func getFileType(exp string) string {
	fileSuffix := path.Ext(exp) //获取文件后缀
	if len(fileSuffix) > 1 {
		return fileSuffix[1:]
	}
	return ""
}
