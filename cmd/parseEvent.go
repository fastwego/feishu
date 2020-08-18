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
	"regexp"
	"strings"

	"github.com/fastwego/feishu/types/event_types"
	"github.com/iancoleman/strcase"
)

func parse() {

	file, err := ioutil.ReadFile("./data/docs3.html")
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

		//fmt.Println(data.Data.Detail.Content)

		line := strings.ReplaceAll(data.Data.Detail.Content, "<br>", "")
		line = strings.Trim(line, "\n\r")

		m2 := regexp.MustCompile(`(?sU)\x60\x60\x60json(.+)\x60\x60\x60`).FindAllStringSubmatch(line, -1)

		//fmt.Println(line, "\n", m2)

		for _, m := range m2 {
			if len(m) > 1 {
				data := m[1]
				noComment := regexp.MustCompile(`(?m)//.+$`).ReplaceAllString(data, "")
				//fmt.Println(noComment)

				e := event_types.BaseEvent{}
				eventType := ""
				err := json.Unmarshal([]byte(noComment), &e)
				if err != nil {

					submatch := regexp.MustCompile(`(?sU)"event".+"type":\s*"(.+)"`).FindStringSubmatch(noComment)
					if len(submatch) > 1 {
						eventType = submatch[1]
					} else {
						continue
					}

				} else {
					eventType = e.Event.Type
				}

				if eventType == "message" {
					msgType := ""
					submatch := regexp.MustCompile(`(?sU)"event".+"msg_type":\s*"(.+)"`).FindStringSubmatch(noComment)
					if len(submatch) > 1 {
						msgType = submatch[1]
					} else {
						continue
					}

					fmt.Println(`
const MsgType` + strcase.ToCamel(msgType) + ` = "` + msgType + `"
/*
` + data + `
*/

type EventMessage` + strcase.ToCamel(msgType) + ` struct {
` + noComment + `
}
`)
				} else {

					fmt.Println(`
const EventType` + strcase.ToCamel(eventType) + ` = "` + eventType + `"
/*
` + data + `
*/

type Event` + strcase.ToCamel(eventType) + ` struct {
` + noComment + `
}
`)
				}
			}
		}

	}
}
