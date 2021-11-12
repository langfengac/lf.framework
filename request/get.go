package request

import (
	"io/ioutil"
	"net/http"
)

//k=v get
func HttpGet(getUrl string, headers map[string]string) (string, error) {
	client := &http.Client{}

	//get请求
	req, err := http.NewRequest("GET", getUrl, nil)
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
