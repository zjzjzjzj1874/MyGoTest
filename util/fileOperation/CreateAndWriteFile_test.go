package fileOperation

import (
	"fmt"
	"testing"
)

// 创建新的文件
func TestCreateFile(t *testing.T) {
	CreateFile(FileName, FileType)
}

//写入文件
func TestWriteFile(t *testing.T) {
	fmt.Println(WriteFile(FileName, "This is test content"))
}
