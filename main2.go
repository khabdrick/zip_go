package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("read.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Create a new gzip writer
	gzipWriter, err := os.Create("read.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer gzipWriter.Close()

	// Use gzip.NewWriter to wrap the output file
	zipWriter := gzip.NewWriter(gzipWriter)
	defer zipWriter.Close()
// Copy the input file to the gzip writer
_, err = io.Copy(zipWriter, inputFile)
if err != nil {
	log.Fatal(err)
}

// Close the gzip writer
zipWriter.Close()
}

