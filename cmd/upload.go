/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

func NewUploadCmd(fileService *drive.FilesService) *cobra.Command {
	// uploadCmd represents the upload command
	var uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "uploads the folder in google cloud bucket",
		Long:  `ploads the folder in google cloud bucket`,
		Run: func(cmd *cobra.Command, args []string) {
			uploadRun()
		},
	}

	return uploadCmd
}

func init() {

}

func uploadRun() {
	fmt.Println("upload called")
}
