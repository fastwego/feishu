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

package event_types

const (
	CallbackTypeUrlVerification = "url_verification"
	CallbackTypeEventCallback   = "event_callback"
)

type Callback struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type EventChallenge struct {
	Callback
	Challenge string `json:"challenge"`
}

type BaseEvent struct {
	Callback
	Ts    string `json:"ts"`
	UUID  string `json:"uuid"`
	Event struct {
		Type string `json:"type"`
	} `json:"event"`
}
