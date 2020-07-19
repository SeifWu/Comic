package utils

import (
	"log"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// FileRename 文件重命名
func FileRename(fileName string) string {
	filenameWithSuffix := path.Base(fileName) //获取文件名带后缀
	log.Println(filenameWithSuffix)
	fileSuffix := path.Ext(filenameWithSuffix) //获取文件后缀
	unixStr := strconv.FormatInt(time.Now().Unix(), 10)

	return filepath.Base(unixStr + EncodeMd5(fileName) + fileSuffix)
}
