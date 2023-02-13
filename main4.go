package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Specify the destination directory
	dst := "output"

	// Open the zip file
	fmt.Println("open zip archive...")
	archive, err := zip.OpenReader("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
// Extract the files from the zip
for _, f := range archive.File {
	// Create the destination file path
	filePath := filepath.Join(dst, f.Name)

	// Print the file path
	fmt.Println("extracting file ", filePath)

	// Check if the file is a directory
	if f.FileInfo().IsDir() {
		// Create the directory
		fmt.Println("creating directory...")
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			panic(err)
		}
		continue
	}

	// Create the parent directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		panic(err)
	}

	// Create an empty destination file
	dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		panic(err)
	}

	// Open the file in the zip and copy its contents to the destination file
	srcFile, err := f.Open()
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		panic(err)
	}

	// Close the files
	dstFile.Close()
	srcFile.Close()
}
}
