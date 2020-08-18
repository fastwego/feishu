// Copyright 2014 chanxuehong(chanxuehong@gmail.com)
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
	"reflect"
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	type args struct {
		base64CipherText string
		key              []byte
	}

	tests := []struct {
		name               string
		args               args
		wantUnpadDecrypted []byte
		wantErr            bool
	}{
		{
			name: "case1",
			args: args{
				base64CipherText: "E/SxIO6PaRHpTXfcS+/PTKQcQhjJ8oFhBIHfrvuOi40lry+db6WAI9BGZFfPsLyRPYffrnUHJWvLcuPavtpiAgtB3j9AX3GBWD/n4twxP7JsqgZQ+0xWk0BOdJ7MSzt3rUmzdOQEgYQZyI5EKlDpU2yRWqLG0uk+i+aquPLzQpRvMlKWfhLXhBdNgqbztE4M036Jci/wByhAbLfO2WsLpPMRvdSdAkYK22Hr2rKkQnqfvb4oZNZmqpwUxPMJGWr/e7krKWwPc7PBAyMunl00XQogOqkTX5uBs6i1HQvUNPwa6rG3ABk14o2j4Qig0kBS3foUWwAEAd6DpsW7kmH+SFjJFsCY7pawVEYf3NMY+Gb7zLfTtuydCUG4ZUtqTPF1SrZIBtHarzy6FlT458lZdmN6BlbyDM1M9NrjLKwmLJqhBGT6JAk4/IVdGaypdFWkGIbjWXskBN4Il7+E1KwTAV9mpNvOAHyooiYrJp4qNLB5ts9uPs0JT0U23Em0Ra/GTrlaftwB9Cu7e4ePm6s5Ok0bdu/u51j+FX1UocduArfS2TIfwQ8hJ7iiqhg3Zlce0SKOwS+OfUkydxLRW637bninfb7oOuGJmaojFihz1moguVtRulOxL/1jSjbHV8bdFma9dIbVSJTNoy2cTqZowPf6Cv2WkMuE3ewdaq2SZAyOVt4tWoBBMk+fyDiPbKPoCzsdSRvmn1FQHU0HX/xPdnAYB9TnpuVpyXC5Kp/ZTekExkb+zaAuW5AI1EztzayN4PNEXNg64mz7tstHEcEti63dVULgwsDml7RNy1YDKsJm3r05XmJYTmnheXpAyPMgYt0mmxjBxlDsoOwISQ0vAEGrU1tnzIMhl/3ZEy40g/zcLzLc8cbu9KuZ2109jyiV8z9YqrI5QCe6rYcuE850HMNP00OWXkxu16jG39yWSQ0GlSr1Q562uKG+nLzvXASIFJ0huVCTrEzogxJgWEGBY3cp6t4GK7VXJdZZSlTqC0wgRNPcDxMRPMhzNGXL5bHnENrZI8ycfta+McvKBUadVkVBAvSFXmcPxhXXdGTMAYpcg8gt3j4S4gpp+2T693QRp68ks7HvJnXAMBUhX+tP5dTZ4HRNwupaZhuzf/e0Sxdi8UkcFCXIqySgwBCnmzoF",
				key:              []byte("lkNa48m2mvOPeMXCQU201UfoDnd3WtrR"),
			},
			wantErr:            false,
			wantUnpadDecrypted: []byte(`{"uuid":"e5b933050c4045fce8f09a5f612c53a0","event":{"app_id":"cli_9fbc93aab56f500c","chat_type":"private","is_mention":false,"lark_version":"lark/3.29.4","message_id":"","msg_type":"text","open_chat_id":"oc_a9c45bf55957776216b5e3a998ea8a96","open_id":"ou_3eb7fcb13e6d756f46e4d910df760e85","open_message_id":"om_ea53bac0b8eaf73933fc4a640906e38c","parent_id":"","root_id":"","tenant_key":"2d54f4e2f8cf575d","text":"哈哈","text_without_at_bot":"哈哈","type":"message","user_agent":"Mozilla/5.0 (Linux; Android 10; LIO-AN00 Build/HUAWEILIO-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/75.0.3770.156 Mobile Safari/537.36 Lark/3.29.4 LarkLocale/zh_CN SDK-Version/3.29.13","user_open_id":"ou_3eb7fcb13e6d756f46e4d910df760e85"},"token":"jz51vCdQz04lxVsvA7I1zbF15fexduh2","ts":"1597670902.943681","type":"event_callback"}`),
		},
		{
			name: "case2",
			args: args{
				base64CipherText: "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk=",
				key:              []byte("test key"),
			},
			wantErr:            false,
			wantUnpadDecrypted: []byte(`hello world`),
		},
		{
			name: "case3",
			args: args{
				base64CipherText: "FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfLk4jx0k/U1deAqYK3AXOZ5zcHt/cC4ZNTqYwWUW/EoL+b2hW/C4zoAQQ5CeMtbxX2zHjm+E4nX/Aww+FHUL6iuIMaeL2KLxqdtbHRC50vgC2YI7xohnb3KuCNBMUzLiPeNIpVdnYaeteCmSaESb+AZpJB9PExzTpRDzCRv+T6o5vlzaE8UgIneC1sYu85BnPBEMTSuj1ZZzfdQi7ZW992Z4dmJxn9e8FL2VArNm99f5Io3c2O4AcNsQENNKtfAAxVjCqc3mg5jF0nKabA+u/5vrUD76flX1UOF5fzJ0sApG2OEn9wfyPDRBsApn9o+fceF9hNrYBGsdtZrZYyGG387CGOtKsuj8e2E8SNp+Pn4E9oYejOTR+ZNLNi+twxaXVlJhr6l+RXYwEiMGQE9zGFBD6h2dOhKh3W84p1GEYnSRIz1+9/Hp66arjC7RCrhuW5OjCj4QFEQJiwgL45XryxHtiZ7JdAlPmjVsL03CxxFZarzxzffryrWUG3VkRdHRHbTsC34+ScoL5MTDU1QAWdqUC1T7xT0lCvQELaIhBTXAYrznJl6PlA83oqlMxpHh0gZBB1jFbfoUr7OQbBs1xqzpYK6Yjux6diwpQB1zlZErYJUfCqK7G/zI9yK/60b4HW0k3M+AvzMcw=",
				key:              []byte("kudryavka"),
			},
			wantErr:            false,
			wantUnpadDecrypted: []byte(`{"uuid":"5226cd85b4d843dccee2e279d93f3ed3","event":{"app_id":"cli_9e28cb7ba56a100e","before_status":{"is_active":true,"is_frozen":true,"is_resigned":false},"change_time":"2020-05-20 18:33:25","current_status":{"is_active":true,"is_frozen":false,"is_resigned":false},"employee_id":"75ge6c49","open_id":"ou_2ef04637d933f798dcb92c99e845ed09","tenant_key":"2d520d3b434f175e","type":"user_status_change"},"token":"GzhQEyfUcx7eEungQFWtXgCbxSpUOJIb","ts":"1589970805.376395","type":"event_callback"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnpadDecrypted, err := AESDecrypt(tt.args.base64CipherText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUnpadDecrypted, tt.wantUnpadDecrypted) {
				t.Errorf("AESDecrypt() gotUnpadDecrypted = %v, \n want %v", string(gotUnpadDecrypted), tt.wantUnpadDecrypted)
			}
		})
	}
}
