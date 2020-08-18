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

package document_test

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document"
)

func ExampleCreateFolder() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.CreateFolder(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleRootFolderMeta() {
	var ctx *feishu.App

	header := http.Header{}
	resp, err := document.RootFolderMeta(ctx, header)

	fmt.Println(resp, err)
}

func ExampleFolderChildren() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.FolderChildren(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleCreateFile() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.CreateFile(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleCopyFile() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.CopyFile(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleDeleteFile() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.DeleteFile(ctx, params, header)

	fmt.Println(resp, err)
}

func ExamplePermissionMemberCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionMemberCreate(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExamplePermissionMemberTransfer() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionMemberTransfer(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExamplePermissionPublicUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionPublicUpdate(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExamplePermissionMemberList() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionMemberList(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExamplePermissionMemberDelete() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionMemberDelete(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExamplePermissionMemberUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionMemberUpdate(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExamplePermissionMemberPermitted() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.PermissionMemberPermitted(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExampleRawContent() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.RawContent(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleSheetMeta() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.SheetMeta(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleDocMeta() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.DocMeta(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsMetainfo() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsMetainfo(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleSheetsBatchUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SheetsBatchUpdate(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsValuesPrepend() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsValuesPrepend(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsValuesAppend() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsValuesAppend(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsInsertDimensionRange() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsInsertDimensionRange(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsDimensionRange() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsDimensionRange(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsProtectedDimension() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsProtectedDimension(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsMergeCells() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsMergeCells(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsUnmergeCells() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsUnmergeCells(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsRange() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsRange(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsValuesBatchGet() {
	var ctx *feishu.App

	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsValuesBatchGet(ctx, params, header)

	fmt.Println(resp, err)
}

func ExampleSpreadSheetsValuesBatchUpdate() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	header := http.Header{}
	resp, err := document.SpreadSheetsValuesBatchUpdate(ctx, payload, params, header)

	fmt.Println(resp, err)
}

func ExampleDocsMeta() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.DocsMeta(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExampleSearchObject() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.SearchObject(ctx, payload, header)

	fmt.Println(resp, err)
}

func ExampleCommentAddWhole() {
	var ctx *feishu.App

	payload := []byte("{}")
	header := http.Header{}
	resp, err := document.CommentAddWhole(ctx, payload, header)

	fmt.Println(resp, err)
}
