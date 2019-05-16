// Copyright © 2019 LLeo <lleoem@gmail.com>
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

	"github.com/spf13/cobra"
)

var typ string
var size uint32

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <filename>",
	Short: "Create an empty datafile.",
	Long: `Create an empty datafile establishing the basic settings of the
datafile Type and pageSize.

This is the template:
$ df create [-s <uint32>] <filename>

Examples:
$ df create myfile.dat
$ df create -s 4096 mydata.df
$ df create -s 32 mydata.df`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("create called args = %s\n", args)
		fmt.Printf("typ = %q; size = %d\n", typ, size)
		if len(args) < 1 {
			usage(cmd, 1, "must provide a <filename>")
		} else if len(args) > 2 {
			usage(cmd, 1, "to many arguments")
		}
		filename := args[0]
		fmt.Printf("calling datafile.Create(%q, %d, %d)\n",
			filename, size)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().Uint32VarP(&size, "pageSize", "s", 256, "page size in bytes.")
}
