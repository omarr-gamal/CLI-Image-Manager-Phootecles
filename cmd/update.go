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

var programVariables map[string]string = map[string]string{
	"imageSavePath": "",
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Use this command to configure the program variables",
	Long: `This command lets you change the values that the program 
uses such as the path that new images are downloaded in.

Use it like this: update <varName> <value>

Example: update imageSavePath "C:/Users/Hp/Desktop/"

Note: it's important not to forget the "/" at the end.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Error: Invalid number of parameters. Try running update -h for help.")
			return
		}

		for key := range programVariables {
			if key == args[0] {
				// fmt.Printf("1: %v, 2: %v\n", args[0], args[1])
				// fmt.Println(programVariables[args[0]])
				programVariables[args[0]] = args[1]

				encodeFile, err := os.Create("config.gob")
				if err != nil {
					panic(err)
				}
				encoder := gob.NewEncoder(encodeFile)

				if err := encoder.Encode(programVariables); err != nil {
					panic(err)
				}
				encodeFile.Close()
			}
			fmt.Printf("Successfully updated %v to become %v\n", key, args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
