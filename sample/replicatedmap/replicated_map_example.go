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

package main

import (
	"fmt"

	"github.com/jorgenpo/hazelcast-go-client"
)

func main() {

	config := hazelcast.NewConfig()
	config.NetworkConfig().AddAddress("127.0.0.1:5701")

	client, err := hazelcast.NewClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	replicatedMap, _ := client.GetReplicatedMap("myReplicatedMap")

	replicatedMap.Put("key", "value")
	replicatedMap.Put("key2", "value2")

	value, _ := replicatedMap.Get("key")
	fmt.Println("Get: ", value)

	values, _ := replicatedMap.Values()
	fmt.Println("Values: ", values)

	keySet, _ := replicatedMap.KeySet()
	fmt.Println("KeySet: ", keySet)

	size, _ := replicatedMap.Size()
	fmt.Println("Size: ", size)

	entrySet, _ := replicatedMap.EntrySet()
	fmt.Println("EntrySet: ")
	for _, entry := range entrySet {
		fmt.Println(entry.Key(), entry.Value())
	}

	found, _ := replicatedMap.ContainsKey("key")
	fmt.Println("ContainsKey: ", found)

	found, _ = replicatedMap.ContainsValue("value")
	fmt.Println("ContainsValue: ", found)

	replicatedMap.Destroy()
	client.Shutdown()
}
