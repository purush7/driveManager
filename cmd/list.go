/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/purush7/project/client"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

func NewListCmd(fileService *drive.FilesService) *cobra.Command {
	// listCmd represents the list command
	var opts = newListOpts(fileService, client.FolderID)
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists the contents of the folder",
		Long:  `ists the contents of the folder`,
		Run: func(cmd *cobra.Command, args []string) {
			opts.listRun()
			opts.printList()
		},
	}

	return listCmd
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type listOptions struct {
	fileService    *drive.FilesService
	folders, files []*drive.File
	folderID       string
}

func newListOpts(fileService *drive.FilesService, folderID string) *listOptions {
	return &listOptions{fileService: fileService, folderID: folderID}
}

func (opts *listOptions) listRun() {

	filesListCall := opts.fileService.List()
	listQuery := fmt.Sprintf(`'%s' in parents`, opts.folderID)

	fileList, err := filesListCall.Q(listQuery).Fields("files(name,mimeType,id)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	if len(fileList.Files) == 0 {
		fmt.Println("No files or subfolders are found")
		return
	}

	var folderMimeType = "application/vnd.google-apps.folder"
	// var documentMimeType = "application/vnd.google-apps.document"

	for _, file := range fileList.Files {
		switch file.MimeType {
		case folderMimeType:
			opts.folders = append(opts.folders, file)
		default:
			opts.files = append(opts.files, file)
		}
	}
}

func (opts *listOptions) printList() {
	fmt.Println("Files and sub folders:")
	fmt.Println("Name\t\tType")

	for _, folder := range opts.folders {
		fmt.Printf("%s\t\t%s", folder.Name, "folder\n")
	}

	for _, file := range opts.files {
		fmt.Printf("%s\t%s", file.Name, "file\n")
	}
}
