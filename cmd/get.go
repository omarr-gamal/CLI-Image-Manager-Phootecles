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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"unicode"

	"github.com/spf13/cobra"
)

type Image struct {
	Id             string
	Title          string
	Description    string
	ImageOnlineUrl string
}

func getImages() (images []Image) {

	fileBytes, err := ioutil.ReadFile("./images.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &images)

	if err != nil {
		panic(err)
	}

	return images
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Use this command to get images.",
	Long: `‘Use this command to get a list of all images in your image collection, 
use get all to get a list of all images. 

Alternatively, use get <image_id> to get that specific image.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: Invalid number of parameters. Try running get -h for help.")
			return
		}

		if args[0] == "all" {
			images := getImages()
			for _, img := range images {
				formatImage(img)
			}
			fmt.Printf("Number of images is %v", len(images))
		} else {
			if !isNumber(args[0]) {
				fmt.Println("Error: Invalid image id.")
				return
			}
			images := getImages()
			for _, img := range images {
				if img.Id == args[0] {
					formatImage(img)
					return
				}
			}
			fmt.Printf("Error: Couldn't find image with id \"%v\"", args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func formatImage(img Image) {
	fmt.Printf("id:%v\n", img.Id)
	fmt.Printf("title:%v\n", img.Title)
	fmt.Printf("description:%v\n", img.Description)
	fmt.Printf("url:%v\n", img.ImageOnlineUrl)
	println("---------------")
}

func isNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
