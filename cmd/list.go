/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/gob"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use this command to list the variables that the program use",
	Long: `Use this command to get a list of the variables that the program
uses if you need to configure them. Such variables include the path 
which new images are downloaded to.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("Error: Invalid number of parameters. Try running list -h for help.")
			return
		}

		getConfigs()

		for key, value := range programVariables {
			fmt.Printf("%v: %v\n", key, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getConfigs() {
	decodeFile, err := os.Open("config.gob")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	decoder.Decode(&programVariables)
}
