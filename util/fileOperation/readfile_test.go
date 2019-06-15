package fileOperation

import (
	"fmt"
	"testing"
)

const (
	FileName = "test"
	FileType = "txt"
)

// 整个文件读取
// 按字节读取，将整个文件读取到缓冲区buffer
func TestReadWholeFile(t *testing.T) {
	ReadWholeFile(FileName, FileType)
}

// 按规定的长度读取
func TestReadFileBySpecialLength(t *testing.T) {
	ReadFileBySpecialLength(FileName, FileType, 10)
}

// 逐行读取
func TestReadFileByRow(t *testing.T) {
	ReadFileByRow(FileName, FileType)
}

// ioUtil读取文件
func TestReadFileByIoUtil(t *testing.T) {
	ReadFileByIoUtil(FileName, FileType)
}

// 读取文件夹; -- filename实际上是文件夹名字
func TestReadFileDir(t *testing.T) {
	ReadFileDir(FileName)
}

// 查看文件是否存在
func TestFileExist(t *testing.T) {
	fmt.Println(FileExist(FileName))
}
