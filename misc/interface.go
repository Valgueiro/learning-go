package main

import "fmt"

type DataProvider interface {
	backup(string) bool
}

type GoogleDrive struct {
	url string
}

func (g *GoogleDrive) backup(folder string) bool {
	fmt.Printf("Backing up from Google Drive folder %s\n", folder)
	return true
}

type DropBox struct {
	url   string
	token string
}

func (d *DropBox) backup(folder string) bool {
	fmt.Printf("Backing up from Drop Box folder %s\n", folder)
	return true
}

func callBackupFromProvider(d DataProvider, f string) bool {
	return d.backup(f)
}

func main() {
	folder := "/path/to/folder"
	callBackupFromProvider(&GoogleDrive{url: "googledrive.com"}, folder)
	callBackupFromProvider(&DropBox{url: "dropbox.com", token: "blabla"}, folder)
}
