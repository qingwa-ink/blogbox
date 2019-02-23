package tools

import (
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPGet 简单的Get网络请求
func HTTPGet(url string) (string, error) {

	httpClient := &http.Client{
		Timeout: 20 * time.Second,
	}
	res, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
