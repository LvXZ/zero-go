package my_util

import (
	"fmt"
	"io/ioutil"
)

// @Author: lvxiaozheng
// @Date: 2021/2/18 10:53
// @Description: 读取文件树

//文件目录树形结构节点
type dirTreeNode struct {
	name  string
	child []dirTreeNode
}

//递归遍历文件目录
func getDirTree(pathName string) (dirTreeNode, error) {
	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		fmt.Printf("Read dir '%s' failed: %v \n", pathName, err)
	}
	var tree, childNode dirTreeNode
	tree.name = pathName
	var name, fullName string
	for _, fileDir := range rd {
		name = fileDir.Name()
		fullName = pathName + "/" + name
		if fileDir.IsDir() {
			childNode, err = getDirTree(fullName)
			if err != nil {
				fmt.Printf("Read dir '%s' failed: %v \n", fullName, err)
			}
		} else {
			childNode.name = name
			childNode.child = nil
		}
		tree.child = append(tree.child, childNode)
	}
	return tree, nil
}

//递归打印文件目录
func printDirTree(tree dirTreeNode, prefix string) {
	fmt.Println(prefix + tree.name)
	if len(tree.child) > 0 {
		prefix += "----"
		for _, childNode := range tree.child {
			printDirTree(childNode, prefix)
		}
	}
}
