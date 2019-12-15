package result

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"path/filepath"
	"strings"
)

type FindFile struct {
	ostype   string //= os.Getenv("GOOS") // 获取系统类型
	suffix   string
	basePath string
	listfile []*string //获取文件列表
}

/*
*
**/
func NewFindfile(basePath string, suffix string) (fH *FindFile) {
	fH = &FindFile{}
	fH.ostype = os.Getenv("GOOS") // 获取系统类型
	fH.basePath = basePath
	fH.suffix = suffix
	return fH
}

/***

 */
func (fH *FindFile) listfunc(path string, f os.FileInfo, err error) error {
	var strRet string
	strRet, _ = os.Getwd()
	//ostype := os.Getenv("GOOS") // windows, linux

	if fH.ostype == "windows" {
		strRet += "\\"
	} else if fH.ostype == "linux" {
		strRet += "/"
	}

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	strRet += path //+ "\r\n"

	//用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
	ok := strings.HasSuffix(strRet, fH.suffix)
	if ok {
		fH.listfile = append(fH.listfile, &strRet) //将目录push到listfile []string中
	}
	//fmt.Println(ostype) // print ostype
	fmt.Println(strRet) //list the file

	return nil
}

/***

 */
func (fH *FindFile) FindFileList() (listfile []*string, err error) {
	//var strRet string
	err = filepath.Walk(fH.basePath, fH.listfunc) //

	if err != nil {
		logs.Error("GetFileList filepath.Walk() returned %v\n", err)
		return nil, err
	}

	return fH.listfile, nil
}

/***

 */
func (fH *FindFile) ListFileFunc(p []*string) {
	for index, value := range p {
		logs.Info("Index = %d, Value =  %v", index, *value)
	}
}
