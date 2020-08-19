package file_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/file"
)

func ExampleCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := file.Create(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleCopy() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := file.Copy(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleDeleteDoc() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := file.DeleteDoc(ctx, params, accessToken)

	fmt.Println(resp, err)
}
