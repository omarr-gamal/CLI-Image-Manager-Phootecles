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
	"io"
	"io/ioutil"
	"net/http"
	"os"

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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 4 {
			fmt.Println("invalid add command")
			return
		}

		newImage := Image{
			Id:             args[0],
			Title:          args[1],
			Description:    args[2],
			ImageOnlineUrl: args[3],
		}

		// download the image
		response, err := http.Get(newImage.ImageOnlineUrl)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// create the image file
			out, err := os.Create(newImage.Id + "_" + newImage.Title + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			// write the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			images := getImages()
			images = append(images, newImage)

			saveImages(images)

			fmt.Println("Successfully created image file")
		} else {
			fmt.Println("Unexpected error happened")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
