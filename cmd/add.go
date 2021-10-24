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

	"github.com/spf13/cobra"
)

func saveImages(images []Image) {

	imageBytes, err := json.Marshal(images)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./images.json", imageBytes, 0644)
	if err != nil {
		panic(err)
	}

}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use this command to create and add images.",
	Long: `This command lets you create and add a new image to the 
image collection. Example:

add "001" "Mount Fuji" "this active volcano is a very distinctive feature of the geography 
of Japan...." "https://upload.wikimedia.org/wikipedia/commons/1/1b/080103_hakkai_fuji.jpg"

There needs to be four arguments that are ordered like this: image id, image title, 
image descreption, and image url. `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 4 {
			fmt.Println("Error: Invalid number of parameters. Try running add -h for help.")
			return
		}

		newImage := Image{
			Id:             args[0],
			Title:          args[1],
			Description:    args[2],
			ImageOnlineUrl: args[3],
		}

		images := getImages()
		for _, img := range images {
			if img.Id == newImage.Id {
				fmt.Println("Error: Image id must be a unique number")
				return
			}
		}

		getConfigs()

		downloadAndSaveImage(newImage)

		images = append(images, newImage)

		saveImages(images)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
