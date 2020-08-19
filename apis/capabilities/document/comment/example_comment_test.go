package comment_test

import (
	"fmt"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/apis/capabilities/document/comment"
)

func ExampleAddWhole() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := comment.AddWhole(ctx, payload, accessToken)

	fmt.Println(resp, err)
}
