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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

const ServerUrl = `https://open.feishu.cn`

var apiUniqMap = map[string]bool{}

func run() {

	file, err := ioutil.ReadFile("./data/docs2.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	matched := regexp.MustCompile(`url="(/document/\S+)"`).FindAllStringSubmatch(string(file), -1)

	for _, match := range matched {
		//fmt.Println(match)

		docUrl := match[1]

		split := strings.Split(docUrl, "/")
		docId := split[3]

		detailDocUrl := fmt.Sprintf("https://open.feishu.cn/opendoc/page/node/detail?id=%s&tab=1", docId)

		resp, err := http.Get(detailDocUrl)
		//fmt.Println(detailDocUrl, "---------------")
		body, err := ioutil.ReadAll(resp.Body)

		//fmt.Println(string(body))

		var data = struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
			Data struct {
				Detail struct {
					Content     string `json:"content"`
					Format      int    `json:"format"`
					HasNewLabel bool   `json:"has_new_label"`
					Hidden      int    `json:"hidden"`
					ID          string `json:"id"`
					Name        string `json:"name"`
					NodeType    int    `json:"node_type"`
					ParentID    string `json:"parent_id"`
					Seq         int    `json:"seq"`
					Status      int    `json:"status"`
					SubName     string `json:"sub_name"`
					Tab         int    `json:"tab"`
				} `json:"detail"`
			} `json:"data"`
		}{}

		err = json.Unmarshal(body, &data)
		if err != nil {
			continue
		}

		line := strings.ReplaceAll(data.Data.Detail.Content, "<br>", "")
		line = strings.Trim(line, "\n\r")

		m2 := regexp.MustCompile(`(?s)#\s+\S+\s(.+).+\*\*请求方式.+(GET|POST|DELETE|PATCH).+请求地址.+(https://open\.feishu\.cn/\S+|https://www\.feishu\.cn/\S+)`).FindStringSubmatch(line)

		//fmt.Println(line, "\n", m2)

		if len(m2) != 4 {
			continue
		}

		apiurl := m2[3]
		_DESCRIPTION_ := m2[1]
		method := m2[2]

		_REQUEST_ := method + " " + apiurl

		parse, _ := url.Parse(apiurl)
		if _, ok := apiUniqMap[parse.Path]; !ok {
			apiUniqMap[parse.Path] = true
		} else {
			continue
		}

		_NAME_ := data.Data.Detail.Name

		_SEE_ := ServerUrl + docUrl
		_FUNCNAME_ := strcase.ToCamel(path.Base(parse.Path))

		//GetParams: []Param{
		//	{Name: `appid`, appType: `string`},
		//	{Name: `redirect_uri`, appType: `string`},
		//},

		_GET_PARAMS_ := ""
		fields := strings.Fields(_REQUEST_)
		parse, _ = url.Parse(fields[1])
		uniqParams := map[string]bool{}
		for param_name, _ := range parse.Query() {
			if uniqParams[param_name] {
				continue
			}
			_GET_PARAMS_ += "\t\t{Name: `" + param_name + "`, Type: `string`},\n"

			uniqParams[param_name] = true
		}
		if _GET_PARAMS_ != "" {
			_GET_PARAMS_ = `GetParams: []Param{
` + _GET_PARAMS_ + "\t},"
		}

		_DESCRIPTION_ = strings.ReplaceAll(_DESCRIPTION_, "\"", "'")
		_DESCRIPTION_ = strings.ReplaceAll(_DESCRIPTION_, "`", " ")

		tpl := strings.ReplaceAll(itemTpl, "_NAME_", _NAME_)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", "`"+_DESCRIPTION_+"`")
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", _REQUEST_)
		tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
		tpl = strings.ReplaceAll(tpl, "_FUNCNAME_", _FUNCNAME_)
		tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)

		fmt.Println(tpl)

	}
}

var itemTpl = `{
	Name:        "_NAME_",
	Description: _DESCRIPTION_,
	Request:     "_REQUEST_",
	See:         "_SEE_",
	FuncName:    "_FUNCNAME_",
	_GET_PARAMS_
},`
