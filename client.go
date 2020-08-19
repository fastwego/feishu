// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/*
飞书 api 服务器地址
*/
var FeishuServerUrl = "https://open.feishu.cn"
var FeishuServerUrl2 = "https://www.feishu.cn"

/*
HttpClient 用于向 接口发送请求
*/
type Client struct {
	Ctx *App
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(url string, header http.Header) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header = header

	return client.HTTPDo(req)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(url string, payload io.Reader, header http.Header) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return
	}
	req.Header = header

	return client.HTTPDo(req)
}

//HTTPDelete DELETE 请求
func (client *Client) HTTPDelete(url string, header http.Header) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}
	req.Header = header

	return client.HTTPDo(req)
}

//HTTPPatch PATCH 请求
func (client *Client) HTTPPatch(url string, payload io.Reader, header http.Header) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		return
	}
	req.Header = header

	return client.HTTPDo(req)
}

//HTTPut PUT 请求
func (client *Client) HTTPPut(url string, payload io.Reader, header http.Header) (resp []byte, err error) {

	req, err := http.NewRequest(http.MethodPut, url, payload)
	if err != nil {
		return
	}
	req.Header = header

	return client.HTTPDo(req)
}

//HTTPDo 执行 请求
func (client *Client) HTTPDo(req *http.Request) (resp []byte, err error) {
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("%s %s Headers %v", req.Method, req.URL.String(), req.Header)
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

/*
筛查 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应码 code 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if bytes.HasPrefix(resp, []byte(`{`)) { // 只针对 json 数据
		errorResponse := struct {
			Code int64  `json:"code"`
			Msg  string `json:"msg"`
		}{}
		err = json.Unmarshal(resp, &errorResponse)
		if err != nil {
			return
		}

		if errorResponse.Code != 0 {
			err = errors.New(string(resp))
			return
		}
	}

	return
}
