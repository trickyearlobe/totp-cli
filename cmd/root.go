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
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "totp-cli",
	Short: "A brief description of your application",
	Long: `
TOTP CLI is a tool for generating one time TOTP codes for 2FA (aka. MFA).
It supports multiple TOTP tokens so you can authenticate to multiple services.
Token secrets can be extracted from your favourite TOTP app to add to TOTP CLI`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".totp-cli")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("No config file found. Creating empty config file\n")
		viper.SetDefault("secrets", map[string]string{"example-token": "KRUGS4ZANFZSAYJAONUGC4TFMQQHGZLDOJSXIIDFPBQW24DMMU====="})
		cobra.CheckErr(viper.WriteConfigAs(filepath.Join(home, ".totp-cli")))
	}
}
