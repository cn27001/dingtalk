/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"net/http"

	"github.com/cn27001/dingtalk/constant"
	"github.com/cn27001/dingtalk/domain/message"
	"github.com/cn27001/dingtalk/request"
	"github.com/cn27001/dingtalk/response"
)

// SendMessage 发送普通消息
func (ding *DingTalk) SendMessage(senderId, chatId string, msg message.Message) (req response.SendMessage, err error) {
	return req, ding.Request(http.MethodPost, constant.SendToConversationKey, nil,
		request.NewSendMessage(senderId, chatId, msg), &req)
}
