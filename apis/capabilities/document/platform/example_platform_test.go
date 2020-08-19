package platform_test

import (
	"fmt"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/platform"
)

func ExampleDocsMeta() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := platform.DocsMeta(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleSearchObject() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := platform.SearchObject(ctx, payload, accessToken)

	fmt.Println(resp, err)
}
