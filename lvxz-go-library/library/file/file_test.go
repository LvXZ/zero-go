package file

import (
	"fmt"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	err := RemoveFile("/usr/local/robot/20210126_144006_20326_724jc_8d93282c-5f55-427b-8383-dfb459c51d97")
	if err != nil {
		fmt.Println("删除文件异常:" + err.Error())
	}
}

func TestRemoveFiles(t *testing.T) {
	err := RemoveFiles("/usr/local/robot")
	if err != nil {
		fmt.Println("删除文件夹下所有文件（不包括文件夹）异常:" + err.Error())
	}
}

func TestRemoveFolder(t *testing.T) {
	err := RemoveFolder("/usr/local/robot")
	if err != nil {
		fmt.Println("删除文件夹及其所有文件异常:" + err.Error())
	}
}

func TestReadFile(t *testing.T) {
	content, err := ReadFile("/usr/local/robot/20210126_124038_18997_724jc_58f67819-b33c-4a3e-ab21-c87db3f52fa3")
	if err != nil {
		fmt.Println("读取文件内容异常:" + err.Error())
	}
	fmt.Println("读取文件内容:" + content)
}
