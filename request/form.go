package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

//表单请求
func HttpPostForm(postUrl string, forms map[string]interface{}) (string, error) {
	urlValues := make(url.Values)
	for k, v := range forms {
		urlValues.Add(k, v.(string))
	}
	resp, err := http.PostForm(postUrl, urlValues)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
