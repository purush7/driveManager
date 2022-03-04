/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/purush7/project/client"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
)

func NewDownloadCmd(fileService *drive.FilesService) *cobra.Command {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("error while getting current working directory")
	}
	cwd = filepath.Join(cwd, client.FolderName)

	if _, err := os.Stat(cwd); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(cwd, 0755)
			if err != nil {
				log.Printf("error while creating folder %s: %v\n", client.FolderName, err)
			}
		} else {
			log.Fatalf("error while creating folder %s: %v\n", client.FolderName, err)
		}
	}

	var opts = newDownloadOpts(fileService, client.FolderID, cwd)

	// downloadCmd represents the download command
	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "command to download the folder in google drive",
		Long:  `command to download the folder in google drive`,
		Run: func(cmd *cobra.Command, args []string) {
			opts.downloadFolder()
		},
	}

	return downloadCmd

}

func newDownloadOpts(fileService *drive.FilesService, folderID, cwd string) *downloadOpts {
	listOpts := newListOpts(fileService, folderID)
	return &downloadOpts{listOptions: listOpts, cwd: cwd}
}

type downloadOpts struct {
	*listOptions
	cwd string
}

func init() {}

func (opts *downloadOpts) downloadFolder() {

	opts.listRun()
	opts.makeFolderStruct()
	opts.downloadFiles()
	for _, folder := range opts.folders {
		path := filepath.Join(opts.cwd, folder.Name)
		childOpts := newDownloadOpts(opts.fileService, folder.Id, path)
		childOpts.downloadFolder()
	}
}

func (opts *downloadOpts) downloadFiles() {
	for _, file := range opts.files {
		path := filepath.Join(opts.cwd, file.Name)
		childOpts := newDownloadOpts(opts.fileService, file.Id, path)
		childOpts.downloadFileCall()
	}
}

func (opts *downloadOpts) downloadFileCall() {

	fid := opts.folderID
	resp, err := opts.fileService.Export(fid, `text/plain`).Download()
	if err != nil {
		log.Fatalf("Unable to download files: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to download files2: %v", err)
	}

	file, err := os.OpenFile(
		opts.cwd,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	bytesWritten, err := file.Write(body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes into %s.\n", bytesWritten, opts.cwd)

}

func (opts *downloadOpts) makeFolderStruct() {
	for _, folder := range opts.folders {
		path := filepath.Join(opts.cwd, folder.Name)
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(path, 0755)
				if err != nil {
					log.Printf("error while creating folder %s: %v\n", folder.Name, err)
				}
			} else {
				log.Fatalf("error while creating folder %s: %v\n", folder.Name, err)
			}
		}
	}

	for _, file := range opts.files {
		path := filepath.Join(opts.cwd, file.Name)
		newFile, err := os.Create(path)
		if err != nil {
			log.Fatalf("error while creating file %s\n", file.Name)
		}
		newFile.Close()
	}
}
