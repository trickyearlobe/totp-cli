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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a TOTP token and it's associated secret",
	Long: `
Delete a TOTP token and it's associated secret.
There is no undelete command or confirmation so be careful.
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		cobra.CheckErr(err)

		if name == "" {
			cmd.Help()
			os.Exit(1)
		}

		key := "secrets." + name
		if viper.IsSet(key) {
			delete(viper.Get("secrets").(map[string]interface{}), name)
			viper.WriteConfig()
		} else {
			fmt.Printf("Key '%v' not found", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name", "n", "", "Name of the TOTP 2FA token to be deleted")
}
