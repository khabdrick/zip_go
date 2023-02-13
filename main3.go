package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	// Open the zip file for reading
	zipReader, err := zip.OpenReader("archive.zip")
	if err != nil {
		panic(err)
	}
	defer zipReader.Close()

	// Open the zip file for writing
	zipfile, err := os.OpenFile("archive.zip", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()
// Create a new zip writer
zipWriter := zip.NewWriter(zipfile)
defer zipWriter.Close()

// Open the file you want to add to the zip
newfile, err := os.Open("newfile.txt")
if err != nil {
	panic(err)
}
defer newfile.Close()

// Create a new zip header for the file
fileInfo, err := newfile.Stat()
if err != nil {
	panic(err)
}

header, err := zip.FileInfoHeader(fileInfo)
if err != nil {
	panic(err)
}

// Set the file name in the zip header
header.Name = "newfile.txt"
header.Method = zip.Deflate

// Add the file to the zip
writer, err := zipWriter.CreateHeader(header)
if err != nil {
	panic(err)
}
_, err = io.Copy(writer, newfile)
if err != nil {
	panic(err)
}
for _, file := range zipReader.File {
	if file.Name != "newfile.txt" {
		writer, err := zipWriter.Create(file.Name)
		if err != nil {
			panic(err)
		}
		reader, err := file.Open()
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(writer, reader)
		if err != nil {
			panic(err)
		}
		reader.Close()
	}
}
}
