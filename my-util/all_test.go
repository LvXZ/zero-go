package my_util

import (
	"fmt"
	"testing"
	"time"
)

// @Author: lvxiaozheng
// @Date: 2021/7/6 15:52
// @Description:

func TestGetPinYin(t *testing.T) {
	GetPinYin()
}

func TestSkipList(t *testing.T) {
	InitSkipList()
}

func TestSnowflake(t *testing.T) {
	// 生成节点实例
	node, err := NewWorker(2)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second * 5)
		fmt.Println(node.GetId())
	}
}
