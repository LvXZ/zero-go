package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// @Author: lvxiaozheng
// @Date: 2021/6/11 11:50
// @Description:

func GetPostHttp() (string, error) {

	url := "https://reserve.moutai.com.cn/api/rsv-server/anon/consumer/getShops"
	method := "POST"

	payload := strings.NewReader(`custId=******`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Host", " reserve.moutai.com.cn")
	req.Header.Add("Accept", " application/json, text/plain, */*")
	req.Header.Add("Accept-Language", " zh-cn")
	req.Header.Add("Origin", " https://reserve.moutai.com.cn")
	req.Header.Add("Content-Length", " 13")
	req.Header.Add("User-Agent", " Mozilla/5.0 (Macintosh; Intel Mac OS X 10_16) AppleWebKit/605.1.15 (KHTML, like Gecko) MicroMessenger/6.8.0(0x16080000) MacWechat/3.0.2(0x13000210) Chrome/39.0.2171.95 Safari/537.36 NetType/WIFI WindowsWechat MicroMessenger/6.8.0(0x16080000) MacWechat/3.0.2(0x13000210) NetType/WIFI WindowsWechat")
	req.Header.Add("Referer", " https://reserve.moutai.com.cn/mconsumer/?a=1&token=023TS5000FOFRL1rwm300JEZcf1TS50m")
	req.Header.Add("Connection", " keep-alive")
	req.Header.Add("Content-Type", " application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Cookie", " lambo-sso-key_0_=023TS5000FOFRL1rwm300JEZcf1TS50m#8cJWkB3u66nJS03K5Qnx8J8C1+KZbLm+Iu5r0VPkKTI=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
