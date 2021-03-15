# go-gcp-storage

Sample project to show how to use some of the function use.

The original function sample:

> https://github.com/GoogleCloudPlatform/golang-samples/blob/master/storage/objects/list_files.go

```go
func listFiles(w io.Writer, bucket string) error {
    // some code
}
```

To use this function we need to use an struct that implements _io.Writer_ inteface.

We can see in the interface document this function to implement:

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

So we can create an struct (in this example is named as 'W') and implement the _Write_ function:

```go
type W struct{} // create a struct

func (w W) Write(p []byte) (n int, err error) { // implement the Write function from Writer interface
	log.Print(string(p))
	return len(p), nil
}
```

In the main function we can use the struct as parameter to use the _listFiles_ function

```go
func main() {

	w := W{} // new W instance

	bucketName := os.Getenv("BUCKET_NAME")
	err := listFiles(w, bucketName) // using here as parameter
	if err != nil {
		log.Fatal(err)
	}
}
```

## Getting Started

Export variables:

```
export BUCKET_NAME=bucket-name
export GOOGLE_APPLICATION_CREDENTIALS=$(pwd)/resources/service-account-a53a55d10d53.json
```

Where:

| Name                             | Description                    |
|----------------------------------|--------------------------------|
| `BUCKET_NAME`                    | Bucket name                    |
| `GOOGLE_APPLICATION_CREDENTIALS` | Service Account json file path |

The _Service Account_ file can be downloaded when you create an service account access to your bucket. 

Execute:

```sh
go run main.go
```

Sample output

```sh
2021/03/15 14:58:20 test-dir/
2021/03/15 14:58:20 other-dir/2021/03/nonono.txt
2021/03/15 14:58:20 test-file.txt
```