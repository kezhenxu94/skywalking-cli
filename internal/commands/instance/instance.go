// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package instance

import (
	"github.com/urfave/cli"

	"github.com/apache/skywalking-cli/pkg/graphql/metadata"

	"github.com/apache/skywalking-cli/internal/logger"
)

var Command = cli.Command{
	Name:      "instance",
	ShortName: "i",
	Usage:     "Instance related sub-command",
	Subcommands: cli.Commands{
		ListCommand,
		SearchCommand,
	},
}

func verifyAndSwitchServiceParameter(ctx *cli.Context) string {
	serviceID := ctx.String("service-id")
	serviceName := ctx.String("service-name")

	if serviceID == "" && serviceName == "" {
		logger.Log.Fatalf("flags \"service-id, service-name\" must set one")
	}

	if serviceID == "" && serviceName != "" {
		service, err := metadata.SearchService(ctx, serviceName)
		if err != nil {
			logger.Log.Fatalln(err)
		}
		serviceID = service.ID
	}
	return serviceID
}
