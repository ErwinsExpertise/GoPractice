package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/minio/minio-go"
)

func UploadFile(client *minio.Client, spaceName string) {

	id, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
	}
	file, err := os.Create("tmp/file-" + id.String())
	if err != nil {
		log.Fatal(err)
	}
	size := rand.Int63n(2000)
	if err := file.Truncate(size); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Attempting: " + file.Name())
	n, err := client.PutObject(spaceName, file.Name(), file, -1, minio.PutObjectOptions{ContentType: "text/plain"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nUploaded %v size in bytes: %v", file.Name(), n)
	file.Close()

}

func main() {
	accessKey := ""
	secKey := ""
	endpoint := "nyc3.digitaloceanspaces.com"
	spaceName := "" // Space names must be globally unique
	ssl := false

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Going to upload")
	var waitgroup sync.WaitGroup
	waitgroup.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			UploadFile(client, spaceName)
			waitgroup.Done()
		}()

	}
	waitgroup.Wait()
}
