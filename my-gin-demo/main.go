package main

// @Author: lvxiaozheng
// @Date: 2021/2/4 10:52
// @Description: 在Go的路上无法自拔

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello "+time.Now().Format("2006-01-01 15:04:05"))
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

	//表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	//文件上传
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})

	//上传图片
	r.POST("/upload/image", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			c.String(http.StatusNotAcceptable, "只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "/Users/lvxiaozheng/Documents/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})

	//上传多个文件
	// 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload/files", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有文件
		files := form.File["files"]
		// 遍历所有文件
		for _, file := range files {
			// 逐个存储
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})

	//HTML返回渲染
	r.LoadHTMLGlob("template/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "content": "123456"})
	})

	r.GET("/info", func(c *gin.Context) {
		c.HTML(http.StatusOK, "info.html", gin.H{"title": "我是测试", "address": "www.bilibili.com"})
	})

	//重定向
	r.GET("/redirect/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.bilibili.com")
	})

	// 1.异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		var content string
		content = "123456"

		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "content": content})
	})

	// 2.同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "content": time.Now().Format("2006-01-01 15:04:05")})
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
