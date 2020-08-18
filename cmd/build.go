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
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	var pkgFlag string
	flag.StringVar(&pkgFlag, "package", "user", "")
	flag.Parse()
	for _, group := range apiConfig {
		if group.Package == "oauth" { // 单独处理 oauth 模块
			continue
		}

		if group.Package == pkgFlag {
			build(group)
		}
	}

	if pkgFlag == "apilist" {
		apilist()
	}

}

func apilist() {
	for _, group := range apiConfig2 {
		fmt.Printf("- %s(%s)\n", group.Name, group.Package)
		for _, api := range group.Apis {
			split := strings.Split(api.Request, " ")
			parse, _ := url.Parse(split[1])

			if api.FuncName == "" {
				api.FuncName = strcase.ToCamel(path.Base(parse.Path))
			}

			godocLink := fmt.Sprintf("https://pkg.go.dev/github.com/fastwego/feishu/apis/%s?tab=doc#%s", group.Package, api.FuncName)
			fmt.Printf("\t- [%s](%s) \n\t\t- [%s (%s)](%s)\n", api.Name, api.See, api.FuncName, parse.Path, godocLink)
		}
	}
}

func build(group ApiGroup) {
	var funcs []string
	var consts []string
	var testFuncs []string
	var exampleFuncs []string

	for _, api := range group.Apis {
		tpl := postFuncTpl
		_FUNC_NAME_ := ""
		_GET_PARAMS_ := ""
		_GET_SUFFIX_PARAMS_ := ""
		_PAYLOAD_ := ""
		switch {
		case strings.Contains(api.Request, "GET http"):
			tpl = getFuncTpl
		case strings.Contains(api.Request, "DELETE http"):
			tpl = deleteFuncTpl
		case strings.Contains(api.Request, "POST http"):
			tpl = postFuncTpl
		}
		if len(api.GetParams) > 0 || len(api.PathParams) > 0 {
			_GET_PARAMS_ = `, params url.Values`
			//if strings.Contains(api.Request, "POST") {
			//	_GET_PARAMS_ = `, ` + _GET_PARAMS_
			//}
			if len(api.GetParams) > 0 {
				_GET_SUFFIX_PARAMS_ = `+ "?" + params.Encode()`
			}
		}

		split := strings.Split(api.Request, " ")
		parseUrl, _ := url.Parse(split[1])

		if api.FuncName == "" {
			_FUNC_NAME_ = strcase.ToCamel(path.Base(parseUrl.Path))
		} else {
			_FUNC_NAME_ = api.FuncName
		}

		_APIVARNAME_ := "api" + _FUNC_NAME_
		_HEADER_ := tenantAccessTokenTpl
		_HEADERPARAM_ := ""
		if api.AccessToken == "app" {
			_HEADER_ = appAccessTokenTpl
		} else if api.AccessToken == "user" {
			_HEADERPARAM_ = ", header http.Header"
			_HEADER_ = userAccessTokenTpl
		}

		_PATH_PARAMS_ := ""
		if strings.Contains(parseUrl.Path, ":") {
			_PATH_PARAMS_ = `
			api := api` + _FUNC_NAME_ + `
			for paramName, paramValues := range params {
				api = strings.ReplaceAll(api, ":"+paramName, paramValues[0])
			}`
			_APIVARNAME_ = "api"
		}

		tpl = strings.ReplaceAll(tpl, "_HEADER_", _HEADER_)
		tpl = strings.ReplaceAll(tpl, "_HEADERPARAM_", _HEADERPARAM_)
		tpl = strings.ReplaceAll(tpl, "_TITLE_", api.Name)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", api.Description)
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", api.Request)
		tpl = strings.ReplaceAll(tpl, "_SEE_", api.See)
		tpl = strings.ReplaceAll(tpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)
		tpl = strings.ReplaceAll(tpl, "_GET_SUFFIX_PARAMS_", _GET_SUFFIX_PARAMS_)
		tpl = strings.ReplaceAll(tpl, "_PAYLOAD_", _PAYLOAD_)
		tpl = strings.ReplaceAll(tpl, "_PATH_PARAMS_", _PATH_PARAMS_)
		tpl = strings.ReplaceAll(tpl, "_APIVARNAME_", _APIVARNAME_)

		funcs = append(funcs, tpl)

		tpl = strings.ReplaceAll(constTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_API_PATH_", parseUrl.Path)

		consts = append(consts, tpl)

		// TestFunc
		_TEST_ARGS_STRUCT_ := ""
		_EXAMPLE_HEADER_STMT_ := ""
		_HEADERINIT_ := ""
		switch {
		case strings.Contains(api.Request, "GET http") || strings.Contains(api.Request, "DELETE http"):
			_TEST_ARGS_STRUCT_ = `ctx *feishu.App, ` + _GET_PARAMS_
			if api.AccessToken == "user" {
				_TEST_ARGS_STRUCT_ += `, header http.Header`
				_EXAMPLE_HEADER_STMT_ = `header := http.Header{}`
				_HEADERINIT_ = `, header: http.Header{}`
			}
		case strings.Contains(api.Request, "POST http"):
			_TEST_ARGS_STRUCT_ = `ctx *feishu.App, payload []byte`

			if _GET_PARAMS_ != "" {
				_TEST_ARGS_STRUCT_ += `,` + _GET_PARAMS_
			}
			if api.AccessToken == "user" {
				_TEST_ARGS_STRUCT_ += `, header http.Header`
				_EXAMPLE_HEADER_STMT_ = `header := http.Header{}`
				_HEADERINIT_ = `, header: http.Header{}`
			}
		}
		_TEST_ARGS_STRUCT_ = strings.ReplaceAll(_TEST_ARGS_STRUCT_, ",", "\n")

		_TEST_FUNC_SIGNATURE_ := ""
		_EXAMPLE_ARGS_STMT_ := ""

		if strings.TrimSpace(_TEST_ARGS_STRUCT_) != "" {
			signatures := strings.Split(_TEST_ARGS_STRUCT_, "\n")
			paramNames := []string{}
			exampleStmt := []string{}
			for _, signature := range signatures {
				signature = strings.TrimSpace(signature)
				tmp := strings.Split(signature, " ")
				//fmt.Println(tmp)
				if len(tmp[0]) > 0 {
					paramNames = append(paramNames, "tt.args."+tmp[0])

					switch tmp[1] {
					case `[]byte`:
						exampleStmt = append(exampleStmt, tmp[0]+" := []byte(\"{}\")")
					case `string`:
						exampleStmt = append(exampleStmt, tmp[0]+" := \"\"")
					case `url.Values`:
						exampleStmt = append(exampleStmt, tmp[0]+" := url.Values{}")
					}
				}
			}
			_TEST_FUNC_SIGNATURE_ = strings.Join(paramNames, ",")
			_EXAMPLE_ARGS_STMT_ = strings.Join(exampleStmt, "\n")
		}

		tpl = strings.ReplaceAll(testFuncTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_TEST_ARGS_STRUCT_", _TEST_ARGS_STRUCT_)
		tpl = strings.ReplaceAll(tpl, "_TEST_FUNC_SIGNATURE_", _TEST_FUNC_SIGNATURE_)
		tpl = strings.ReplaceAll(tpl, "_HEADERINIT_", _HEADERINIT_)
		testFuncs = append(testFuncs, tpl)

		//Example

		tpl = strings.ReplaceAll(exampleFuncTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_PACKAGE_", path.Base(group.Package))
		tpl = strings.ReplaceAll(tpl, "_TEST_FUNC_SIGNATURE_", strings.ReplaceAll(_TEST_FUNC_SIGNATURE_, "tt.args.", ""))
		tpl = strings.ReplaceAll(tpl, "_EXAMPLE_ARGS_STMT_", _EXAMPLE_ARGS_STMT_)
		tpl = strings.ReplaceAll(tpl, "_EXAMPLE_HEADER_STMT_", _EXAMPLE_HEADER_STMT_)
		exampleFuncs = append(exampleFuncs, tpl)

	}

	fileContent := fmt.Sprintf(fileTpl, path.Base(group.Package), group.Name, path.Base(group.Package), strings.Join(consts, ``), strings.Join(funcs, ``))
	filename := "./../apis/" + group.Package + "/" + path.Base(group.Package) + ".go"
	_ = os.MkdirAll(path.Dir(filename), 0644)
	ioutil.WriteFile(filename, []byte(fileContent), 0644)

	// output Test
	testFileContent := fmt.Sprintf(testFileTpl, path.Base(group.Package), strings.Join(testFuncs, ``))
	//fmt.Println(testFileContent)
	ioutil.WriteFile("./../apis/"+group.Package+"/"+path.Base(group.Package)+"_test.go", []byte(testFileContent), 0644)

	// output example
	exampleFileContent := fmt.Sprintf(exampleFileTpl, path.Base(group.Package), strings.Join(exampleFuncs, ``))
	//fmt.Println(testFileContent)
	ioutil.WriteFile("./../apis/"+group.Package+"/example_"+path.Base(group.Package)+"_test.go", []byte(exampleFileContent), 0644)

}

var constTpl = `
	api_FUNC_NAME_ = "_API_PATH_"`
var commentTpl = `
/*
_TITLE_

_DESCRIPTION_

See: _SEE_

_REQUEST_
*/`
var postFuncTpl = commentTpl + `
func _FUNC_NAME_(ctx *feishu.App, payload []byte_GET_PARAMS__HEADERPARAM_) (resp []byte, err error) {
	_PATH_PARAMS_
	_HEADER_
	return ctx.Client.HTTPPost(_APIVARNAME__GET_SUFFIX_PARAMS_, bytes.NewReader(payload), header)
}
`
var getFuncTpl = commentTpl + `
func _FUNC_NAME_(ctx *feishu.App_GET_PARAMS__HEADERPARAM_) (resp []byte, err error) {
	_PATH_PARAMS_
	_HEADER_

	return ctx.Client.HTTPGet(_APIVARNAME__GET_SUFFIX_PARAMS_,header)
}
`
var deleteFuncTpl = commentTpl + `
func _FUNC_NAME_(ctx *feishu.App_GET_PARAMS__HEADERPARAM_) (resp []byte, err error) {
	_PATH_PARAMS_
	_HEADER_

	return ctx.Client.HTTPDelete(_APIVARNAME__GET_SUFFIX_PARAMS_,header)
}
`
var tenantAccessTokenTpl = `
	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")
`

var appAccessTokenTpl = `
	accessToken, err := ctx.GetAppAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-appType", "application/json")
`

var userAccessTokenTpl = `
	header.Set("Content-appType", "application/json")
`

var fieldTpl = `
		// field
		err = m.WriteField("_FIELD_NAME_", string(payload))
		if err != nil {
			return
		}
`

var fileTpl = `// Package %s %s
package %s

const (
	%s
)
%s`

var testFileTpl = `package %s

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

%s
`

var testFuncTpl = `
func Test_FUNC_NAME_(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(api_FUNC_NAME_, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		_TEST_ARGS_STRUCT_
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp_HEADERINIT_}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := _FUNC_NAME_(_TEST_FUNC_SIGNATURE_)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("_FUNC_NAME_() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("_FUNC_NAME_() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}`

var exampleFileTpl = `package %s_test

%s
`
var exampleFuncTpl = `
func Example_FUNC_NAME_() {
	var ctx *feishu.App

	_EXAMPLE_ARGS_STMT_
	_EXAMPLE_HEADER_STMT_
	resp, err := _PACKAGE_._FUNC_NAME_(_TEST_FUNC_SIGNATURE_)

	fmt.Println(resp, err)
}

`
