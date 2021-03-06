// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
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

package internal_test

import (
	"testing"

	"github.com/jorgenpo/hazelcast-go-client/internal"
	"github.com/jorgenpo/hazelcast-go-client/serialization"
	"github.com/jorgenpo/hazelcast-go-client/serialization/spi"
	"github.com/jorgenpo/hazelcast-go-client/test/testutil"
	"github.com/stretchr/testify/assert"
)

func TestLazyReadResultSet_GetNonDeserializableDataSlice(t *testing.T) {
	service, _ := spi.NewSerializationService(serialization.NewConfig())
	lazyset := internal.NewLazyReadResultSet(0, testutil.NewNonDeserializableDataSlice(), nil, service)
	_, err := lazyset.Get(0)
	assert.Error(t, err)
}

func TestLazyReadResultSet_SequenceWhenNotReady(t *testing.T) {
	service, _ := spi.NewSerializationService(serialization.NewConfig())
	lazyset := internal.NewLazyReadResultSet(0, nil, nil, service)
	_, err := lazyset.Sequence(0)
	assert.Error(t, err)
}
