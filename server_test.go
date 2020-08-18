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
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fastwego/feishu/types/event_types"
)

var MockApp *App

func TestMain(m *testing.M) {
	MockApp = NewApp(AppConfig{
		AppId:      "AppId",
		AppSecret:  "AppSecret",
		EncryptKey: "kudryavka",
	})
	os.Exit(m.Run())
}

func TestServer_ParseEvent(t *testing.T) {

	tests := []struct {
		name      string
		wantEvent interface{}
		wantErr   bool
	}{
		{name: "case1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := `{
    "encrypt": "FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfLk4jx0k/U1deAqYK3AXOZ5zcHt/cC4ZNTqYwWUW/EoL+b2hW/C4zoAQQ5CeMtbxX2zHjm+E4nX/Aww+FHUL6iuIMaeL2KLxqdtbHRC50vgC2YI7xohnb3KuCNBMUzLiPeNIpVdnYaeteCmSaESb+AZpJB9PExzTpRDzCRv+T6o5vlzaE8UgIneC1sYu85BnPBEMTSuj1ZZzfdQi7ZW992Z4dmJxn9e8FL2VArNm99f5Io3c2O4AcNsQENNKtfAAxVjCqc3mg5jF0nKabA+u/5vrUD76flX1UOF5fzJ0sApG2OEn9wfyPDRBsApn9o+fceF9hNrYBGsdtZrZYyGG387CGOtKsuj8e2E8SNp+Pn4E9oYejOTR+ZNLNi+twxaXVlJhr6l+RXYwEiMGQE9zGFBD6h2dOhKh3W84p1GEYnSRIz1+9/Hp66arjC7RCrhuW5OjCj4QFEQJiwgL45XryxHtiZ7JdAlPmjVsL03CxxFZarzxzffryrWUG3VkRdHRHbTsC34+ScoL5MTDU1QAWdqUC1T7xT0lCvQELaIhBTXAYrznJl6PlA83oqlMxpHh0gZBB1jFbfoUr7OQbBs1xqzpYK6Yjux6diwpQB1zlZErYJUfCqK7G/zI9yK/60b4HW0k3M+AvzMcw=" 
}`
			/*{
				"uuid": "5226cd85b4d843dccee2e279d93f3ed3",
				"event": {
					"app_id": "cli_9e28cb7ba56a100e",
					"before_status": {
						"is_active": true,
						"is_frozen": true,
						"is_resigned": false
					},
					"change_time": "2020-05-20 18:33:25",
					"current_status": {
						"is_active": true,
						"is_frozen": false,
						"is_resigned": false
					},
					"employee_id": "75ge6c49",
					"open_id": "ou_2ef04637d933f798dcb92c99e845ed09",
					"tenant_key": "2d520d3b434f175e",
					"type": "user_status_change"
				},
				"token": "GzhQEyfUcx7eEungQFWtXgCbxSpUOJIb",
				"ts": "1589970805.376395",
				"type": "event_callback"
			}*/
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
			gotEvent, _ := MockApp.Server.ParseEvent(request)
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("ParseEvent() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			event, ok := gotEvent.(event_types.EventUserStatusChange)
			if !ok {
				t.Errorf("ParseEvent() error = %v", ok)
				return
			}

			if event.Event.OpenID != "ou_2ef04637d933f798dcb92c99e845ed09" {
				t.Errorf(`event.Event.OpenID != "ou_2ef04637d933f798dcb92c99e845ed09"`)
				return
			}
		})
	}
}
