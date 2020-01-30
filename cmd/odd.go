/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "Add odd numbers",
	Long: `This is same as 'even' command. Instead of adding even numbers, it will add odd numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		var oddSum int

		for _, intVal := range args {
			intTemp, _ := strconv.Atoi(intVal)
			if intTemp%2 != 0 {
				oddSum = oddSum + intTemp
			}
		}

		fmt.Printf("The odd addition of %s is %d \n", args, oddSum)
	},
}

func init() {
	addCmd.AddCommand(oddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
