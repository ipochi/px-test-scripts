// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/ipochi/px-test-scripts/src"
	"github.com/spf13/cobra"
)

// snapCmd represents the snap command
var snapCmd = &cobra.Command{
	Use:   "snap",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: snapAction,
}

func init() {
	rootCmd.AddCommand(snapCmd)

	snapCmd.PersistentFlags().IntVarP(&wpnumber, "wordpress", "w", 2, "Number of wordpress instances for snapshot per namespace")
	snapCmd.MarkFlagRequired("wpnumber")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func snapAction(cmd *cobra.Command, args []string) {
	fmt.Println("Creating snapshots")
	src.CreateSnapshots(wpnumber)
}
