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

	"github.com/stretchr/testify/assert"

	"github.com/cn27001/dingtalk/domain/message"
)

func TestNewCorpConvMessage(t *testing.T) {
	conv := NewCorpConvMessage(message.NewTextMessage("hello dubbo-go")).
		SetUserIds("1123", "12358", "788444").
		SetUsers([]string{"1123"}).
		SetAgentId(12345678).
		SetAllUser(false).
		SetDeptIds(123, 123, 456).
		Build()

	err := validate(conv)

	assert.Nil(t, err)
	assert.Equal(t, len(conv.UserIds), 1)
	assert.Equal(t, len(conv.DeptIds), 2)
}
