package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//json post
func HttpPost(postUrl string, headers map[string]string, postJson interface{}) (string, error) {
	client := &http.Client{}
	//转换成postBody
	bytesData, err := json.Marshal(postJson)
	if err != nil {
		return "", err
	}
	postBody := bytes.NewReader(bytesData)
	//post请求
	req, err := http.NewRequest("POST", postUrl, postBody)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//返回内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
