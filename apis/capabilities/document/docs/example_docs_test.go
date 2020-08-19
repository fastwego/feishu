package docs_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/docs"
)

func ExampleRawContent() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := docs.RawContent(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleSheetMeta() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := docs.SheetMeta(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleDocMeta() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := docs.DocMeta(ctx, params, accessToken)

	fmt.Println(resp, err)
}
