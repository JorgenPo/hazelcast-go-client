// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proto

import (
	"github.com/jorgenpo/hazelcast-go-client/serialization"

	"github.com/jorgenpo/hazelcast-go-client/internal/proto/bufutil"
)

func mapTryRemoveCalculateSize(name string, key serialization.Data, threadId int64, timeout int64) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += stringCalculateSize(name)
	dataSize += dataCalculateSize(key)
	dataSize += bufutil.Int64SizeInBytes
	dataSize += bufutil.Int64SizeInBytes
	return dataSize
}

// MapTryRemoveEncodeRequest creates and encodes a client message
// with the given parameters.
// It returns the encoded client message.
func MapTryRemoveEncodeRequest(name string, key serialization.Data, threadId int64, timeout int64) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, mapTryRemoveCalculateSize(name, key, threadId, timeout))
	clientMessage.SetMessageType(mapTryRemove)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendData(key)
	clientMessage.AppendInt64(threadId)
	clientMessage.AppendInt64(timeout)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

// MapTryRemoveDecodeResponse decodes the given client message.
// It returns a function which returns the response parameters.
func MapTryRemoveDecodeResponse(clientMessage *ClientMessage) func() (response bool) {
	// Decode response from client message
	return func() (response bool) {
		response = clientMessage.ReadBool()
		return
	}
}
