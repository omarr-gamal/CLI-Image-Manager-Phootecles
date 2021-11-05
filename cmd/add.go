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
	"sync"

	"github.com/spf13/cobra"
)

var waitGroup = sync.WaitGroup{}

func setInsideTextForImage(image Image) {
	imageUrl := image.ImageOnlineUrl

	// change / to %2F and : to %3A in s
	// loop through and replace all instances of "/" with "%2F"
	// loop through and replace all instances of ":" with "%3A"
	// loop through and replace all instances of "@" with "%40"
	// loop through and replace all instances of "#" with "%23"
	for i := 0; i < len(imageUrl); i++ {
		if imageUrl[i] == '/' {
			imageUrl = imageUrl[:i] + "%2F" + imageUrl[i+1:]
		}
		if imageUrl[i] == ':' {
			imageUrl = imageUrl[:i] + "%3A" + imageUrl[i+1:]
		}
		if imageUrl[i] == '@' {
			imageUrl = imageUrl[:i] + "%40" + imageUrl[i+1:]
		}
		if imageUrl[i] == '#' {
			imageUrl = imageUrl[:i] + "%23" + imageUrl[i+1:]
		}
	}

	url := "https://ocrly-image-to-text.p.rapidapi.com/?imageurl=" + imageUrl + "&filename=sample.jpg"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "ocrly-image-to-text.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "ecfe9f08eamsh1a4d4986842b72cp1a60b5jsn497b0c7c1d2f")

	res, err := http.DefaultClient.Do(req)

	// check error
	if err != nil {
		fmt.Println(err)
		image.InsideText = ""
	}

	// check for errors in the response
	if res.StatusCode != 200 {
		image.InsideText = ""
		if res.Header.Get("x-ratelimit-requests-remaining") == "0" {
			fmt.Println("Warning: Phootecles is out of api requests to extract text from images.")
		}
		fmt.Println("Error: ", res.StatusCode)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	image.InsideText = string(body)

	// loop through images and replace the image with the new one
	images := getImages()
	for i, img := range images {
		if img.Id == image.Id {
			images[i] = image
		}
	}
	saveImages(images)

	waitGroup.Done()
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

		images := getImages()
		images = append(images, newImage)

		saveImages(images)

		waitGroup.Add(1)
		go downloadAndSaveImage(newImage)
		waitGroup.Wait()

		waitGroup.Add(1)
		go setInsideTextForImage(newImage)
		waitGroup.Wait()

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
