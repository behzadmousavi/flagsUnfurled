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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This is going to sum up two number per se",
	Long:  `Again I think the short description is enough!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		flag, _ := cmd.Flags().GetBool("float")
		if flag {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func addInt(args []string) {
	var sum int
	for _, value := range args {
		value, err := strconv.Atoi(value)
		if err == nil {
			sum = sum + value
		} else {
			fmt.Println(err)
		}
	}
	fmt.Printf("The addition of numbers %s is %d \n", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, value := range args {
		value, err := strconv.ParseFloat(value, 64)
		if err == nil {
			sum = sum + value
		} else {
			fmt.Println(err)
		}
	}
	fmt.Printf("The addition of numbers %s is %f \n", args, sum)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add floating numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
