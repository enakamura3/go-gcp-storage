package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type W struct{}

func (w W) Write(p []byte) (n int, err error) {
	log.Print(string(p))
	return len(p), nil
}

func main() {

	w := W{}

	bucketName := os.Getenv("BUCKET_NAME")
	err := listFiles(w, bucketName)
	if err != nil {
		log.Fatal(err)
	}
}

func listFiles(w io.Writer, bucket string) error {
	// bucket := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}
		fmt.Fprintln(w, attrs.Name)
	}
	return nil
}
