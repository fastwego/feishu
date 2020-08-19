package comment

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/feishu"
	"github.com/fastwego/feishu/test"
)

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

func TestAddWhole(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiAddWhole, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}).Methods("POST")

	type args struct {
		ctx         *feishu.App
		payload     []byte
		accessToken string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{ctx: test.MockApp, accessToken: "USER_ACCESS_TOKEN"}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := AddWhole(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWhole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AddWhole() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
