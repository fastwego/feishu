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

/*
HttpClient 用于向 接口发送请求
*/
type Client struct {
	Ctx *App
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string, header http.Header) (resp []byte, err error) {
	url := FeishuServerUrl + uri

	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header = header

	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("GET %s Headers %v", req.URL.String(), req.Header)
	}
	response, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, header http.Header) (resp []byte, err error) {

	url := FeishuServerUrl + uri

	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return
	}
	req.Header = header

	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("POST %s Headers %v", req.URL.String(), req.Header)
	}
	response, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

func (client *Client) HTTPDelete(uri string, header http.Header) (resp []byte, err error) {
	url := FeishuServerUrl + uri

	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}
	req.Header = header

	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("DELETE %s Headers %v", req.URL.String(), req.Header)
	}
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

	return
}
