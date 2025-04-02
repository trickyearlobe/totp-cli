/*
Copyright © 2025 Richard Nixon <richard.nixon@btinternet.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"maps"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available TOTP tokens",
	Long: `
TOTP CLI supports multiple TOTP tokens.
This command produces a list of available tokens.`,
	Run: func(cmd *cobra.Command, args []string) {
		for k := range maps.Keys(viper.GetStringMap("secrets")) {
			println(k)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
