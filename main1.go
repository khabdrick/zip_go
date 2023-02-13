package main
import(
	"archive/zip"
	"fmt"
	"log"
	)
func main(){

zipListing, err := zip.OpenReader("archive.zip")
if err != nil {
  log.Fatal(err)
}
defer zipListing.Close()
for _, file := range zipListing.File {
  fmt.Println(file.Name)
}
}
