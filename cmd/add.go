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
	"net/http"

	"github.com/spf13/cobra"
)

func getTextFromImage(Image Image) {
	url := "https://ocrly-image-to-text.p.rapidapi.com/?imageurl=https%3A%2F%2Fi.pinimg.com%2Foriginals%2F42%2F1b%2Fe6%2F421be6184e75937bb223c764ecbc2f2e.jpg&filename=sample.jpg"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "ocrly-image-to-text.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "ecfe9f08eamsh1a4d4986842b72cp1a60b5jsn497b0c7c1d2f")

	res, err := http.DefaultClient.Do(req)

	// check error
	if err != nil {
		fmt.Println(err)
	}

	// check for errors in the response
	if res.StatusCode != 200 {
		fmt.Println("Error: ", res.StatusCode)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

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

		if isIdUsed(newImage.Id) {
			panic("Error: Image id must be unique and a number")
		}

		getConfigs()

		downloadAndSaveImage(newImage)

		images := getImages()
		images = append(images, newImage)

		saveImages(images)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func isIdUsed(id string) bool {
	images := getImages()
	for _, img := range images {
		if img.Id == id {
			return true
		}
	}
	return false
}
