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
	"testing"

	"github.com/cn27001/dingtalk/constant/file"
	"github.com/cn27001/dingtalk/constant/policy"
	"github.com/cn27001/dingtalk/constant/spaces"
	"github.com/cn27001/dingtalk/request"
	"github.com/stretchr/testify/assert"
)

func TestDingTalk_CreateDriveSpaces(t *testing.T) {
	t.Skip()
	res, err := client.CreateDriveSpaces("golang", "ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_DeleteDriveSpaces(t *testing.T) {
	t.Skip()
	res, err := client.DeleteDriveSpaces("5296669223", "ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDriveSpaces(t *testing.T) {
	res, err := client.GetDriveSpaces("ABNiSWeAolg5OETyYT60wdQiEiE", spaces.Org, "", 5)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Spaces), 1)
}

func TestDingTalk_GetDriveSpacesInfo(t *testing.T) {
	res, err := client.GetDriveSpacesInfo("3452011774", "ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Name, "全员文件夹")
	assert.Equal(t, res.PermissionMode, "acl")
	assert.Equal(t, res.Type, string(spaces.Org))
}

func TestDingTalk_GetDriveSpacesFiles(t *testing.T) {
	files := request.NewGetDriveSpacesFiles("3452011774", "ABNiSWeAolg5OETyYT60wdQiEiE", 10).
		Build()

	res, err := client.GetDriveSpacesFiles(files)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDriveSpacesFileInfo(t *testing.T) {
	res, err := client.GetDriveSpacesFileInfo("3452011774", "47453824236",
		"ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.FileExtension, "pdf")
}

func TestDingTalk_CreateDriveSpacesFiles_Folder(t *testing.T) {
	f := request.NewCreateDriveSpacesFiles("ABNiSWeAolg5OETyYT60wdQiEiE", "3452011774",
		"dingtalk", file.Folder).
		SetConflictPolicy(policy.Overwrite).
		Build()

	res, err := client.CreateDriveSpacesFiles(f)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.FileName, "dingtalk")
}

func TestDingTalk_CreateDriveSpacesFiles_File(t *testing.T) {
	t.Skip()
	f := request.NewCreateDriveSpacesFiles("ABNiSWeAolg5OETyYT60wdQiEiE", "3452011774",
		"dingtalk.png", file.File).
		SetMediaId("@lALPDeREWzc3zXBgzQJA").
		SetConflictPolicy(policy.Overwrite).
		Build()

	res, err := client.CreateDriveSpacesFiles(f)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.FileName, "golang")
}

func TestDingTalk_DeleteDriveSpacesFiles(t *testing.T) {
	t.Skip()
	res, err := client.DeleteDriveSpacesFiles("3452011774", "47459660818",
		"ABNiSWeAolg5OETyYT60wdQiEiE", policy.Completely)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_MoveDriveSpacesFiles(t *testing.T) {
	t.Skip()
	files := request.NewMoveDriveSpacesFiles("3452011774", "47494266429",
		"47494327156", "ABNiSWeAolg5OETyYT60wdQiEiE").
		SetTargetParentId("3452011774").
		SetConflictPolicy(policy.Overwrite).
		Build()

	res, err := client.MoveDriveSpacesFiles(files)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_RenameDriveSpacesFiles(t *testing.T) {
	t.Skip()
	res, err := client.RenameDriveSpacesFiles("3452011774", "47494266429",
		"newName", "ABNiSWeAolg5OETyYT60wdQiEiE")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.FileName, "newName")
}

func TestDingTalk_GetDriveSpacesFilesDownloadInfo(t *testing.T) {
	t.Skip()
	res, err := client.GetDriveSpacesFilesDownloadInfo("3452011774", "47494266429",
		"ABNiSWeAolg5OETyYT60wdQiEiE")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDriveSpacesFilesPermissions(t *testing.T) {
	res, err := client.GetDriveSpacesFilesPermissions("3452011774", "47494266429",
		"ABNiSWeAolg5OETyYT60wdQiEiE")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
