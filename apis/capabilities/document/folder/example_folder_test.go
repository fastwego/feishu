package folder_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/folder"
)

func ExampleCreate() {
	var ctx *feishu.App

	payload := []byte("{}")
	params := url.Values{}
	accessToken := ""
	resp, err := folder.Create(ctx, payload, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleMeta() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := folder.Meta(ctx, params, accessToken)

	fmt.Println(resp, err)
}

func ExampleRootFolderMeta() {
	var ctx *feishu.App

	accessToken := ""
	resp, err := folder.RootFolderMeta(ctx, accessToken)

	fmt.Println(resp, err)
}

func ExampleChildren() {
	var ctx *feishu.App

	params := url.Values{}
	accessToken := ""
	resp, err := folder.Children(ctx, params, accessToken)

	fmt.Println(resp, err)
}
