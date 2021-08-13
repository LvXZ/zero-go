package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func CurlPostHttp(reqUrl string, params map[string]string) (resBody string, err error) {
	// 1.构建JSON参数
	bodyParam, err := json.Marshal(params)
	reqBody := bytes.NewReader(bodyParam)

	// 2.构建请求header
	req, _ := http.NewRequest(http.MethodPost, reqUrl, reqBody)
	req.Header.Set("Content-Type", "application/json")

	// 3.请求
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		err = errors.New("curl url fail，url=" + reqUrl + " statsCode=" + strconv.Itoa(res.StatusCode))
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	resBody = string(body[:])

	return
}

func CurlHttp(method string, reqUrl string, header map[string]string, params map[string]string) (resBody string, err error) {
	// 1.构建参数
	v := url.Values{}
	for key, value := range params {
		v.Add(key, value)
	}
	url := fmt.Sprintf("%s?%s", reqUrl, v.Encode())

	// 2.构建header
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// 3.请求
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		err = errors.New("curl url fail，url=" + url + " statsCode=" + strconv.Itoa(res.StatusCode))
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	resBody = string(body[:])
	return
}
