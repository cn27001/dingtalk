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
	"time"

	"github.com/cn27001/dingtalk/constant"
	"github.com/cn27001/dingtalk/crypto"
	"github.com/cn27001/dingtalk/request"
	"github.com/cn27001/dingtalk/response"
	"github.com/pkg/errors"
)

// GetAuthInfo 获取企业授权信息
func (ding *DingTalk) GetAuthInfo(corpId string) (rsp response.CorpAuthInfo, err error) {
	if !ding.isv() {
		return response.CorpAuthInfo{}, errors.New("ticket or corpId is null")
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := crypto.GetSignature(timestamp, ding.secret, ding.ticket)

	query := url.Values{}
	query.Set("accessKey", ding.key)
	query.Set("suiteTicket", ding.ticket)
	query.Set("timestamp", timestamp)
	query.Set("signature", sign)

	return rsp, ding.Request(http.MethodPost, constant.GetAuthInfo, query,
		request.NewAuthInfo(ding.key, corpId), &rsp)
}

// ActivateSuite 激活应用
func (ding *DingTalk) ActivateSuite(corpId, code string) (rsp response.CorpAuthInfo, err error) {
	token, err := ding.GetSuiteAccessToken()
	if err != nil {
		return response.CorpAuthInfo{}, err
	}

	query := url.Values{}
	query.Set("suite_access_token", token)

	return rsp, ding.Request(http.MethodPost, constant.ActivateSuiteKey, query,
		request.NewActivateSuite(ding.key, corpId, code), &rsp)
}

// GetAgentInfo 获取授权应用的基本信息
func (ding *DingTalk) GetAgentInfo(agentId int, corpId string) (rsp response.AgentInfo, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := crypto.GetSignature(timestamp, ding.secret, ding.ticket)

	token, err := ding.GetSuiteAccessToken()
	if err != nil {
		return response.AgentInfo{}, err
	}

	query := url.Values{}
	query.Set("suite_access_token", token)
	query.Set("timestamp", timestamp)
	query.Set("accessKey", ding.key)
	query.Set("signature", sign)
	query.Set("suiteTicket", ding.ticket)

	return rsp, ding.Request(http.MethodPost, constant.GetAgentKey, query,
		request.NewAgentInfo(agentId, ding.key, corpId), &rsp)
}

// GetCorpPermanentCode 获取授权企业的永久授权码
func (ding *DingTalk) GetCorpPermanentCode(code string) (rsp response.CorpPermanentCode, err error) {
	token, err := ding.GetSuiteAccessToken()
	if err != nil {
		return response.CorpPermanentCode{}, err
	}
	query := url.Values{}
	query.Set("suite_access_token", token)

	return rsp, ding.Request(http.MethodPost, constant.GetCorpPermanentCodeKey, query,
		request.NewCorpPermanentCode(code), &rsp)
}

// GetUnactiveCorp 获取应用未激活的企业列表
func (ding *DingTalk) GetUnactiveCorp(appId int) (rsp response.UnactiveCorp, err error) {
	token, err := ding.GetSuiteAccessToken()
	if err != nil {
		return response.UnactiveCorp{}, err
	}
	query := url.Values{}
	query.Set("suite_access_token", token)

	return rsp, ding.Request(http.MethodPost, constant.GetUnactiveCorpKey, query, request.NewUnactiveCorp(appId), &rsp)
}

// ReauthCorp 重新授权未激活应用的企业
func (ding *DingTalk) ReauthCorp(appId int, corpId string, corpIds ...string) (rsp response.Response, err error) {
	token, err := ding.GetSuiteAccessToken()
	if err != nil {
		return response.Response{}, err
	}
	query := url.Values{}
	query.Set("suite_access_token", token)

	return rsp, ding.Request(http.MethodPost, constant.ReauthCorpKey, query,
		request.NewReauthCorp(appId, append(corpIds, corpId)), &rsp)
}
