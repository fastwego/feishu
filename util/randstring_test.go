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

package util

import (
	"fmt"
	"testing"
)

func TestGetRandString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantLen int
	}{
		{name: "case1", args: args{length: 6}, wantLen: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandString(tt.args.length)
			if len(got) != tt.wantLen {
				t.Errorf("GetRandString() = %v, want %v", got, tt.want)
			}

			fmt.Println(got)
		})
	}
}

func TestGetRandStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		charset string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantLen int
	}{
		{name: "case1", args: args{length: 6, charset: "0x0x0x"}, wantLen: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandStringWithCharset(tt.args.length, tt.args.charset)
			if len(got) != tt.wantLen {
				t.Errorf("GetRandStringWithCharset() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
