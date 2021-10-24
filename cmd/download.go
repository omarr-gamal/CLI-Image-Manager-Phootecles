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
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `Use this command to download one image or all the images 
in your collection. 

Example: download all 
This command will redownload all the images 
Example: download 177 
This command will downlaod the image with id 177

To change or view the download diretory, try "update -h" or "list -h" 
respectively.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: Invalid number of parameters. Try running download -h for help.")
			return
		}

		getConfigs()

		if args[0] == "all" {
			images := getImages()
			for _, img := range images {
				fmt.Printf("Downloading %v%v...\n", img.Id, img.Title)
				downloadAndSaveImage(img)
			}
		} else {
			images := getImages()
			for _, img := range images {
				if img.Id == args[0] {
					downloadAndSaveImage(img)
					return
				}
			}
			fmt.Printf("Error: Couldn't find image with the id: %v\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func downloadAndSaveImage(img Image) {
	// download the image
	response, err := http.Get(img.ImageOnlineUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {

		// create the image file
		out, err := os.Create(programVariables["imageSavePath"] + img.Id + "_" + img.Title + ".png")
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()

		// write the body to file
		_, err = io.Copy(out, response.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Successfully downloaded %v%v...\n", img.Id, img.Title)
	} else {
		fmt.Printf("Unexpected error happened")
	}
}
