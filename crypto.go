// Copyright 2021 FastWeGo
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

package feishu

import (
	"github.com/fastwego/feishu/util"
)

type Crypto struct {
	EncryptKey        string
	VerificationToken string
}

// NewCrypto 创建加解密处理器
func NewCrypto(EncryptKey string) *Crypto {
	return &Crypto{
		EncryptKey: EncryptKey,
	}
}

// GetDecryptMsg 解密消息
func (c *Crypto) GetDecryptMsg(encryptMsg string) (decryptMsg []byte, err error) {

	// 解密
	decryptMsg, err = util.AESDecrypt(encryptMsg, []byte(c.EncryptKey))
	if err != nil {
		return
	}

	if Logger != nil {
		Logger.Printf("AESDecryptMsg %s => %s", encryptMsg, string(decryptMsg))
	}

	return
}
