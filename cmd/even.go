/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "Sum up even numbers",
	Long:  `This is going to sum up even numbers among inputs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("even called")
		var evenSum int
		for _, value := range args {
			value, err := strconv.Atoi(value)
			if err == nil {
				if value%2 == 0 {
					evenSum = evenSum + value
				}
			} else {
				fmt.Println(err)
			}
		}
		fmt.Printf("The addition of numbers %s is %d \n", args, evenSum)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
