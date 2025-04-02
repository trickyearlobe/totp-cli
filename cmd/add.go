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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [flags]",
	Short: "Add a new TOTP token and it's associated secret",
	Long: `
TOTP CLI supports multiple TOTP tokens.
This command allows you to add tokens and their associated secrets
which can be exported from many TOTP apps that can read QR codes.

=======================
Example using 1Password
=======================

- Navigate to the website your need a TOTP token for and set up your basic credentials.

- Follow the website instructions to add a TOTP (2FA or MFA) token.
    - If a choice of apps is offered, select Google Authenticator
    - Do not select RSA token or Microsoft Authenticator as they don't use TOTP tokens 
    - When the QR code is shown, scan it to obtain the secret
    - Perform the post scan activities (usually entering 1 or more codes to test the token)

- In the 1Password desktop app
    - select the credential we just configured
    - click edit and copy the secret string (eg. otpauth://totp/bob@thebuilders.com?secret=KRUGS4ZANFZSAYJAONUGC4TFMQQHGZLDOJSXIIDFPBQW24DMMU=====&issuer=JumpCloud%20User)
    - strip off the fluff to obtain the secret field (in this case KRUGS4ZANFZSAYJAONUGC4TFMQQHGZLDOJSXIIDFPBQW24DMMU=====)

========================= WARNING ============================
    The add command stores the secret configs in CLEARTEXT
    Make sure you are the only person that can access the
    config file (chmod 700 ~/.totp-cli.yaml)
==============================================================
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		cobra.CheckErr(err)
		secret, err := cmd.Flags().GetString("secret")
		cobra.CheckErr(err)

		// {
		// 	cmd.Help()
		// 	os.Exit(1)
		// }
		viper.Set("secrets."+name, secret)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("name", "n", "", "Name of the TOTP 2FA token")
	addCmd.Flags().StringP("secret", "s", "", "The secret which is used to generate the token")
	// _ = viper.BindPFlags(addCmd.PersistentFlags())
}
