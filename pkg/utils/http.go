package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
	GET 请求
*/
func HttpGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

/*
	POST 请求
*/

func HttpPost(url string, data interface{}) ([]byte, error) {
	jsonStr, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	req.Header.Add("content-type", "application/json;charset=utf-8")
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

//PostJSON post json 数据请求
func PostJSON(uri string, obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)

	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

func PostUrlEncoded(urls string, data map[string]string) ([]byte, error) {
	client := &http.Client{}

	//构建from表单的数据
	DataUrlVal := url.Values{}
	for key, val := range data {
		DataUrlVal.Add(key, val)
	}

	req, err := http.NewRequest("POST", urls, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		return nil, err
	}
	//设置请求头部
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	//提交请求
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	//判断状态吗
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", urls, resp.StatusCode)
	}

	//读取返回值
	return ioutil.ReadAll(resp.Body)

}
