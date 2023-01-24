/*
* Written in 2021 by Nicholas S. Damuth
* V.1.0
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package main

import (
	"os"
	"net"
	"fmt"
	"github.com/spf13/viper"
	"damuth.nick/GraphPlane/internal/Logger"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var port = 4000
// var configs map[string]interface{}
var listen net.Listener
// var configs viper.New()

func main() {
	var configs = viper.New()
	logger.LogBCast("INFO", "Starting GraphPlane Server")
	if (os.Args != nil) {
		if (stringInSlice("test", os.Args)) {

		}
		if (stringInSlice("green", os.Args)) {

		}
	}
	/////////////////////////////////////////
	// Configuration Loading & Monitoring
	go LoadConfigs(configs, &listen)
	for (configs.Get("server") == nil) { }

	/////////////////////////////////////////
	// Begin threading DBs

	/////////////////////////////////////////
	// Begin threading Third Party API Services

	/////////////////////////////////////////
	// Begin threading Kafka Services
	
	/////////////////////////////////////////
	// Start up gRPC and Web Services, to includ UID
	ServeAndWait(port, configs, &listen)
	logger.LogBCast("INFO", "Completing GraphPlane Server")

	/////////////////////////////////////////
	// Report/Record Operational Statuses
}
