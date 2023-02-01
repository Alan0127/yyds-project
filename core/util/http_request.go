package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	const1 "yyds-pro/core/const"
)

var client = &http.Client{}

// http远程调用
func DoHttp(ch chan struct{}, method, url string, data interface{}) (body []byte, err error) {
	var (
		req *http.Request
		res *http.Response
	)
	if url == "" || len(url) == 0 {
		err = errors.New("url is empty, please check")
		return
	}
	bytesData, _ := json.Marshal(data)
	switch method {
	case const1.Post:
		req, err = http.NewRequest(const1.Post, url, bytes.NewReader(bytesData))
	case const1.Get:
		req, err = http.NewRequest(const1.Get, url, bytes.NewReader(bytesData))
	case const1.Delete:
		req, err = http.NewRequest(const1.Delete, url, bytes.NewReader(bytesData))
	case const1.Put:
		req, err = http.NewRequest(const1.Put, url, bytes.NewReader(bytesData))
	default:
		err = errors.New("method not allowed")
		return
	}
	if req != nil {
		res, err = client.Do(req)
	}
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		err = errors.New("http service error")
		return
	}
	body, _ = ioutil.ReadAll(res.Body)
	ch <- struct{}{}
	return
}
