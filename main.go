package main
import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)
func main(){
    fmt.Println("creating zip archive")
//Create a new zip archive and named archive.zip
	archive,err:=os.Create("archive.zip")
	if err!=nil{
		panic(err)
// this is to catch errors if any
	}
defer archive.Close()
fmt.Println("archive file created successfully....")
//we use the defer key to close it, once we create an archive we need to close it using the defer keyword
defer archive.Close()
fmt.Println("archive file created successfully")
//Create a new zip writer
zipWriter:=zip.NewWriter(archive)
fmt.Println("opening first file")
//Add files to the zip archive
f1, err:=os.Open("test.csv")
if err!=nil{
	panic(err)
}
defer f1.Close()
fmt.Println("adding file to archive..")
w1,err:=zipWriter.Create ("test.csv")
if err!=nil{
	panic(err)
}
if _,err:=io.Copy(w1,f1); err != nil{
panic(err)
}
fmt.Println("closing archive")
zipWriter.Close()
}
