package helper

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"thegraduate-server/model"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadFile(context context.Context, file multipart.File) string {

	// const projectId = "alpine-land-423208-d4"
	const bucketName = "thegraduate-bucket"

	client, err := storage.NewClient(context, option.WithCredentialsFile("./service.json"))

	if err != nil {
		fmt.Print(err)
		panic(&model.BadRequestError{Message: "error while uploading file"})
	}

	randomBytes := make([]byte, 16)
	_, err = rand.Read(randomBytes)
	if err != nil {
		fmt.Print("error di upload")
		panic(err)
	}

	objectName := fmt.Sprintf("%x", randomBytes)

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(context)

	if _, err := io.Copy(wc, file); err != nil {
		fmt.Print("error lagi ngecopy")
		panic(err)
	}
	if err := wc.Close(); err != nil {
		fmt.Print("error lagi ngeclose")
		panic(err)
	}

	return objectName
}
