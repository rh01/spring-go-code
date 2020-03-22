package utils

import (
	"log"
	"math/rand"
	"os"
)

//检测路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

//生成定锚随机数
func GetRandInt(anchor int) int {

	rand.Seed(int64(anchor))
	// 获取指定范围内的随机数
	return rand.Intn(100000)
}

//获取文件句柄
func GetFile(filePath string) (*os.File, error) {
	var f *os.File
	var err error

	if PathExists(filePath) {
		f, err = os.OpenFile(filePath, os.O_APPEND, 0666)
	} else {
		f, err = os.Create(filePath) //创建文件
	}
	if err != nil {
		log.Println(err)
		log.Println("open file err")
	}

	return f, err
}
