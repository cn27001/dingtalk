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
	"github.com/cn27001/dingtalk/request"
	"github.com/cn27001/dingtalk/response"
)

// GetAttendanceGroups 批量获取考勤组详情
func (ding *DingTalk) GetAttendanceGroups(offset, size int) (rsp response.GetAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupsKey, nil,
		request.NewGetAttendanceGroup(offset, size), &rsp)
}

// GetAttendanceUserGroup 获取用户考勤组
func (ding *DingTalk) GetAttendanceUserGroup(userId string) (rsp response.GetAttendanceUserGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceUserGroupKey, nil,
		request.NewGetAttendanceUserGroup(userId), &rsp)
}

// GetAttendanceGroupMinimalism 批量获取考勤组摘要
func (ding *DingTalk) GetAttendanceGroupMinimalism(userId string, cursor int) (rsp response.GetAttendanceGroupMinimalism,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupMinimalismKey, nil,
		request.NewGetAttendanceGroupMinimalism(userId, cursor), &rsp)
}

// GetAttendanceGroupDetail 获取考勤组详情
func (ding *DingTalk) GetAttendanceGroupDetail(userId string, groupId int) (rsp response.GetAttendanceGroupDetail,
	err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupDetailKey, nil,
		request.NewGetAttendanceGroupDetail(userId, groupId), &rsp)
}

// SearchAttendanceGroup 搜索考勤组摘要
func (ding *DingTalk) SearchAttendanceGroup(userId, groupName string) (rsp response.SearchAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.SearchAttendanceGroupKey, nil,
		request.NewSearchAttendanceGroup(userId, groupName), &rsp)
}

// CreateAttendanceGroup 创建考勤组
func (ding *DingTalk) CreateAttendanceGroup(res *request.CreateAttendanceGroup) (rsp response.CreateAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.CreateAttendanceGroupKey, nil, res, &rsp)
}
