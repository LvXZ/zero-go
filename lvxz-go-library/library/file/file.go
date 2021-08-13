package file

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

// PathExists 判断文件或者文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DownloadFile 下载文件
func DownloadFile(filePath string, url string) error {
	//获取data文件
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//创建file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	//写入file
	_, err = io.Copy(out, resp.Body)
	return err
}

// ReadFile 读取文件内容
func ReadFile(filePath string) (content string, err error) {
	//判断文件是否存在
	exist, err := PathExists(filePath)
	if !exist {
		return
	}

	//打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	//读取文件内容
	reader := bufio.NewReader(file)
	for {
		//读到一个换行就结束
		con, err2 := reader.ReadString('\n')

		//io.EOF 表示文件的末尾
		if err2 == io.EOF {
			break
		}

		//拼接内容
		content = content + con
	}
	//fmt.Print(content)
	//fmt.Println("文件读取结束...")
	return
}

// RemoveFile 删除指定文件
func RemoveFile(filePath string) (err error) {
	//判断文件是否存在
	exist, err := PathExists(filePath)
	if !exist {
		return
	}

	//删除文件
	return os.Remove(filePath)
}

// RemoveFiles 删除指定文件夹下的所有文件(不包括文件夹)
func RemoveFiles(folderPath string) (err error) {
	//判断文件夹是否存在
	exist, err := PathExists(folderPath)
	if !exist {
		return
	}

	//删除文件夹
	err = os.RemoveAll(folderPath)

	//删除指定文件夹及其所有文件后创建文件夹
	exist, err = PathExists(folderPath)
	if !exist {
		//判断不存在创建文件夹
		err = os.MkdirAll(folderPath, os.ModePerm)
	}
	return
}

// CreateFolder 创建文件夹
func CreateFolder(folderPath string) {
	exist, _ := PathExists(folderPath)
	if !exist {
		//判断不存在创建文件夹
		os.MkdirAll(folderPath, os.ModePerm)
	}
}

// RemoveFolder 删除指定文件夹及其所有文件
func RemoveFolder(folderPath string) (err error) {
	//判断文件夹是否存在
	exist, err := PathExists(folderPath)
	if !exist {
		return
	}

	//删除文件夹
	return os.RemoveAll(folderPath)
}
