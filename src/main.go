package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"tingfm/config"
	"tingfm/dao"
	"tingfm/routers"
)

func init() {
	///配置初始化
	config.Init()
	///数据库初始化
	dao.Init()
}

func main() {
	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v \n", err)
	}
}

func fileProcess() {
	path := "."
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		// 带扩展名的文件名
		fullFilename := f.Name()
		//扩展名
		fileExt := filepath.Ext(f.Name())

		fmt.Println(fileExt)
		if fileExt != ".mp3" && fileExt != ".aac" && fileExt != ".m4a" {
			continue
		}
		fmt.Println(fullFilename)
		// 不带扩展名的文件名
		filenameOnly := strings.TrimSuffix(fullFilename, fileExt)
		fmt.Println(filenameOnly)
		//将每个文件名后面加上1，扩展名不变
		err :=  os.Rename(path+"/"+f.Name(), path+"/"+fmt.Sprintf("%s%s%s", filenameOnly,"胜英传", fileExt))
		if err != nil {
			println(err.Error())
		}
	}
}
