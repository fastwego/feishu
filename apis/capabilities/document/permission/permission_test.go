package permission

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

func TestMemberCreate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMemberCreate, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := MemberCreate(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MemberCreate() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMemberTransfer(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMemberTransfer, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := MemberTransfer(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MemberTransfer() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestPublicUpdate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiPublicUpdate, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := PublicUpdate(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PublicUpdate() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMemberList(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMemberList, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := MemberList(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MemberList() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMemberDelete(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMemberDelete, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := MemberDelete(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MemberDelete() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMemberUpdate(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMemberUpdate, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := MemberUpdate(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MemberUpdate() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
func TestMemberPermitted(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockRouter.HandleFunc(apiMemberPermitted, func(w http.ResponseWriter, r *http.Request) {
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
			gotResp, err := MemberPermitted(tt.args.ctx, tt.args.payload, tt.args.accessToken)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberPermitted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("MemberPermitted() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
