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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		panic("Unexpected error happened")
	}
}
