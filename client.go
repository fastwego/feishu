// Copyright 2021 FastWeGo
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
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	contentTypeApplicationJson = "application/json"
)

var (
	ServerUrl  = "https://open.feishu.cn"
	ServerUrl2 = "https://www.feishu.cn"
	UserAgent  = "fastwego/feishu"
)

// NewClient
func NewClient() (client *Client) {
	return &Client{
		HttpClient: http.DefaultClient,
	}
}

/*
Client 用于向接口发送请求
*/
type Client struct {
	HttpClient *http.Client
}

// Do 执行 请求
func (client *Client) Do(req *http.Request, accessToken string) (resp []byte, err error) {

	// 默认 Header Content-Type
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentTypeApplicationJson)
	}

	// 添加 access_token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// 添加 User-Agent
	req.Header.Set("User-Agent", UserAgent)

	if Logger != nil {
		Logger.Printf("%s %s %v", req.Method, req.URL.String(), req.Header)
	}

	response, err := client.HttpClient.Do(req)
	if err != nil {
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	_ = response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("response.Status %s, response.Body %s", response.Status, resp)
		return
	}

	return
}
