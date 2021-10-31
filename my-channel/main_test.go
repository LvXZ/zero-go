package my_channel

import (
	"fmt"
	"testing"
)

// @Author: lvxiaozheng
// @Date: 2021/10/27 11:21
// @Description:

func TestChannel1(t *testing.T) {

	//定义一个通道(当前无缓冲通道)
	channel := make(chan int)

	// 启用goroutine从通道接收值
	go receiverChannel(channel)

	//发送值到通道
	channel <- 1000

	//close函数来关闭通道
	//close(channel)

}

func receiverChannel(channel chan int) {
	//接收通道中的值
	value := <-channel
	fmt.Println("接收成功:", value)
}

func TestString(t *testing.T) {

	ss := "我们"
	fmt.Println(len(ss)) //6个字节

}
