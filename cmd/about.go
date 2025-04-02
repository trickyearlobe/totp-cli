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

	"github.com/spf13/cobra"
)

var BuildTimestamp = "unknown"
var GitCommit = "unknown"
var GitRepo = "unknown"
var BuiltBy = "unknown"

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Information about totp-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nTOTP CLI (C)2025 Richard Nixon\n\n")
		fmt.Printf("Git Repo:           %v\n", GitRepo)
		fmt.Printf("Git Commit:         %v\n", GitCommit)
		fmt.Printf("Build timestamp:    %v\n", BuildTimestamp)
		fmt.Printf("Built By:           %v\n", BuiltBy)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
