package main

// @Author: lvxiaozheng
// @Date: 2021/2/4 10:52
// @Description: 

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "你好")
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	// /user?name=lvxz
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "不存在")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})


	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
