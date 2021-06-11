package main

import (
	"github.com/robfig/cron"
	"robots/internal/service"
)

// @Author: lvxiaozheng
// @Date: 2021/6/11 11:49
// @Description:

func main() {

	c := cron.New()

	c.AddFunc("01 01 08 * * ?", func() {
		s := new(service.Service)
		s.Run()
	})

	// 启动计划任务
	c.Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}
