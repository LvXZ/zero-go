package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"robots/internal/model"
	"robots/internal/utils"
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

	if len(httpResponse.Data) == 0 {
		fmt.Println(httpResponse.Message)
		return
	}

	//data, _ := json.Marshal(httpResponse)
	//fmt.Println(string(data))

	var content string
	for _, data := range httpResponse.Data {
		content = content + data.V + "\n"
	}

	sendText("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=9573b0db-7504-4119-ad84-84e15c2db651", content, nil)
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
