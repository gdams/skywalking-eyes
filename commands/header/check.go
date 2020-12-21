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
//
package header

import (
	"github.com/spf13/cobra"
	"license-checker/internal/logger"
	"license-checker/pkg/header"
)

var CheckCommand = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Long:    "`check` command walks the specified paths recursively and checks if the specified files have the license header in the config file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var config header.Config
		var result header.Result

		if err := config.Parse(cfgFile); err != nil {
			return err
		}

		if err := header.Check(&config, &result); err != nil {
			return err
		}

		logger.Log.Infoln(result.String())

		if result.HasFailure() {
			return result.Error()
		}

		return nil
	},
}
