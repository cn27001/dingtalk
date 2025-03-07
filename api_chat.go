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
	"net/url"
	"strconv"

	"github.com/cn27001/dingtalk/constant"
	"github.com/cn27001/dingtalk/domain/message"
	"github.com/cn27001/dingtalk/request"
	"github.com/cn27001/dingtalk/response"
)

// CreateChat 创建群
func (ding *DingTalk) CreateChat(res *request.CreatChat) (rsp response.CreatChat, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateChatKey, nil, res, &rsp)
}

// GetChatInfo 获取群会话信息
func (ding *DingTalk) GetChatInfo(chatId string) (req response.GetChatInfo, err error) {
	query := url.Values{}
	query.Set("chatid", chatId)

	return req, ding.Request(http.MethodGet, constant.GetChatInfoKey, query, nil, &req)
}

// UpdateChat 修改群会话
func (ding *DingTalk) UpdateChat(res *request.UpdateChat) (req response.Response, err error) {
	return req, ding.Request(http.MethodPost, constant.UpdateChatKey, nil, res, &req)
}

// ChatSetSubAdmin 设置群管理员
func (ding *DingTalk) ChatSetSubAdmin(chatId, userId string, role int) (req response.Response, err error) {
	return req, ding.Request(http.MethodPost, constant.ChatSetSubAdminKey, nil,
		request.NewChatSetSubAdmin(chatId, userId, role), &req)
}

// ChatSetUserNick 设置群成员昵称
func (ding *DingTalk) ChatSetUserNick(chatId, userId, nick string) (req response.Response, err error) {
	return req, ding.Request(http.MethodPost, constant.ChatSetUserNickKey, nil,
		request.NewChatSetUserNick(chatId, userId, nick), &req)
}

// ChatFriendSwitch 设置禁止群成员私聊
func (ding *DingTalk) ChatFriendSwitch(chatId string, prohibit bool) (req response.Response, err error) {
	return req, ding.Request(http.MethodPost, constant.ChatFriendSwitchKey, nil,
		request.NewChatFriendSwitch(chatId, prohibit), &req)
}

// GetChatQRCode 获取入群二维码链接
func (ding *DingTalk) GetChatQRCode(chatId, userId string) (req response.ChatQRCode, err error) {
	return req, ding.Request(http.MethodPost, constant.GetChatQRCodeKey, nil,
		request.NewChatQRCode(chatId, userId), &req)
}

// SendChatMessage 发送消息到群
func (ding *DingTalk) SendChatMessage(chatId string, msg message.Message) (req response.SendChatMessage, err error) {
	return req, ding.Request(http.MethodPost, constant.SendChatMessageKey, nil,
		request.NewSendChatMessage(chatId, msg), &req)
}

// GetChatMsgReadUser 查询群消息已读人员列表
func (ding *DingTalk) GetChatMsgReadUser(messageId string, cursor, size int) (req response.ChatMsgReadUser, err error) {
	if size > 100 || size < 0 {
		size = 100
	}
	query := url.Values{}
	query.Set("messageId", messageId)
	query.Set("cursor", strconv.Itoa(cursor))
	query.Set("size", strconv.Itoa(size))

	return req, ding.Request(http.MethodGet, constant.GetChatReadUserKey, query, nil, &req)
}
