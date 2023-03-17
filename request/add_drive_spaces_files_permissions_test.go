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

package request

import (
	"testing"

	"github.com/cn27001/dingtalk/constant/member"
	"github.com/cn27001/dingtalk/constant/role"
	"github.com/stretchr/testify/assert"
)

func TestNewAddDriveSpacesFilesPermissions(t *testing.T) {
	permissions := NewAddDriveSpacesFilesPermissions().
		SetSpaceId("spaceId").
		SetFileId("fileId").
		SetUnionId("unionId").
		SetRole(role.Admin).
		SetSpacesFileMember("corpId", "memberId", member.User).
		SetSpacesFileMember("corpId", "memberId", member.User).
		Build()

	assert.NotNil(t, permissions)
}

// newDriveSpacesFileMember
func TestNewDriveSpacesFileMember(t *testing.T) {
	fm := newDriveSpacesFileMember("corpId", "memberId", member.Conversation)

	assert.NotNil(t, fm)
}
