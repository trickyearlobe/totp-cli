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
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Generate a TOTP authentication code",
	Long: `
Generate a TOTP authentication code using the specified token name.
The token is time based and changes every 30 seconds to prevent reuse.

The --name flag can be used to select a previously added TOTP token
The --secret flag can be used to directly use a TOTP secret
The TOTP_SECRET environment variable can be used to directly use a TOTP secret

Examples:-

	# Named token
	totp-cli auth --name barclays-totp

	# Direct secret
	totp-cli auth --secret KRUGS4ZANFZSAYJAONUGC4TFMQQHGZLDOJSXIIDFPBQW24DMMU=====

	# Secret in environment variable
	export TOTP_SECRET=KRUGS4ZANFZSAYJAONUGC4TFMQQHGZLDOJSXIIDFPBQW24DMMU=====
	totp-cli auth
`,
	Run: func(cmd *cobra.Command, args []string) {

		if os.Getenv("TOTP_SECRET") != "" {
			displayCode(os.Getenv("TOTP_SECRET"))
		}

		if secret, _ := cmd.Flags().GetString("secret"); secret != "" {
			displayCode(secret)
		}

		if name, _ := cmd.Flags().GetString("name"); name != "" {
			secret := viper.GetString("secrets." + name)
			if secret == "" {
				fmt.Printf("Token '%v' not found\n", name)
				os.Exit(1)
			} else {
				displayCode(secret)
			}
		}

		cmd.Help()
	},
}

func displayCode(secret string) {
	code, err := totp.GenerateCode(secret, time.Now())
	cobra.CheckErr(err)
	fmt.Println(code)
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringP("name", "n", "", "Name of the TOTP token to be used")
	authCmd.Flags().StringP("secret", "s", "", "The secret which is used to generate the token")
}
