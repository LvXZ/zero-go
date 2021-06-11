package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"my-notify/internal/model"
	"my-notify/internal/utils"
	"net/http"
	"os"
	"strings"
)

// @Author: lvxiaozheng
// @Date: 2021/6/11 11:56
// @Description:

// Service 基类
type Service struct {
}

func (s *Service) Run() {

	httpResponse := getHttpInfo()
	if httpResponse == nil {
		return
	}

	var content []string
	if len(httpResponse.Data) == 0 {
		fmt.Println(httpResponse.Message)
		content = append(content, httpResponse.Message)
	} else {
		for _, data := range httpResponse.Data {
			if data.V != "" {
				content = append(content, strings.Replace(data.V, "自营店", "", -1))
			}
		}
	}

	sendText("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=dcf9d9b5-eb3b-4d2a-86d8-d9a02e14c6da", strings.Join(content, " "), nil)
	return
}

//获取信息
func getHttpInfo() (httpResponse *model.HttpResponse) {

	//直接请求信息
	content, err := utils.GetPostHttp()
	if err != nil || content == "" {
		return
	}

	httpResponse = &model.HttpResponse{}
	err = json.Unmarshal([]byte(content), httpResponse)
	if err != nil {
		return
	}
	return
}

// MarkdownMsg 消息
type MarkdownMsg struct {
	Msgtype  string            `json:"msgtype"`
	Markdown map[string]string `json:"markdown"`
}

// TextMsg 文本消息
type TextMsg struct {
	Msgtype string  `json:"msgtype"`
	Text    Content `json:"text"`
}

// Content
type Content struct {
	Content             string   `json:"content"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

// send: 发送文本数据格式
func sendText(url string, content string, attention []string) {
	msg := new(TextMsg)
	msg.Msgtype = "text"
	cnt := new(Content)
	cnt.Content = content
	if len(attention) > 0 {
		cnt.MentionedMobileList = attention
	}
	msg.Text = *cnt
	c := new(bytes.Buffer)
	json.NewEncoder(c).Encode(msg)
	res, _ := http.Post(url, "application/json; charset=utf-8", c)
	io.Copy(os.Stdout, res.Body)
	fmt.Println("发送成功")
}
