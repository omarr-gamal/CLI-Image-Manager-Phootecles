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
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <all | IMAGE_ID>",
	Short: "Use this command to delete an image",
	Long: `This command lets you remove an image from the list of images.
You can either delete all images by running:

delete all

or delete an image with an id by running:

delete IMAGE_ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: Invalid number of parameters. Try running delete -h for help.")
			return
		}

		if args[0] == "all" {
			imageSlice := make([]Image, 0)
			saveImages(imageSlice)
		} else {
			if !isNumber(args[0]) {
				fmt.Println("Error: Invalid image id.")
				return
			}
			images := getImages()
			for i := 0; i < len(images); i++ {
				if images[i].Id == args[0] {
					images[i] = images[len(images)-1]
					images = images[:len(images)-1]

					saveImages(images)

					fmt.Printf("Successfully deleted image with id \"%v\"", args[0])
					return
				}
			}
			fmt.Printf("Error: Couldn't find image with id \"%v\"", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
