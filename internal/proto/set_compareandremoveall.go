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

func setCompareAndRemoveAllCalculateSize(name string, values []serialization.Data) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += stringCalculateSize(name)
	dataSize += bufutil.Int32SizeInBytes
	for _, valuesItem := range values {
		dataSize += dataCalculateSize(valuesItem)
	}
	return dataSize
}

// SetCompareAndRemoveAllEncodeRequest creates and encodes a client message
// with the given parameters.
// It returns the encoded client message.
func SetCompareAndRemoveAllEncodeRequest(name string, values []serialization.Data) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, setCompareAndRemoveAllCalculateSize(name, values))
	clientMessage.SetMessageType(setCompareAndRemoveAll)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendInt32(int32(len(values)))
	for _, valuesItem := range values {
		clientMessage.AppendData(valuesItem)
	}
	clientMessage.UpdateFrameLength()
	return clientMessage
}

// SetCompareAndRemoveAllDecodeResponse decodes the given client message.
// It returns a function which returns the response parameters.
func SetCompareAndRemoveAllDecodeResponse(clientMessage *ClientMessage) func() (response bool) {
	// Decode response from client message
	return func() (response bool) {
		response = clientMessage.ReadBool()
		return
	}
}
