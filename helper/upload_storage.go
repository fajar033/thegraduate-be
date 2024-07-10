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

	serviceJSON := `
{
  "type": "service_account",
  "project_id": "alpine-land-423208-d4",
  "private_key_id": "ebd714371470a8ae8320cc37eb8a9245b26255a8",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCVepmUhqcoHCh9\nofHFdPoaqjI+Qe7mMS1RuaG8YcKL0FROEIv8Jr17xWVcyv8Q5+J6HlygV9IdSbdJ\nAtI/uvDjrCunhjN6Cj7XeBp/5eM24d9CmduzH29h2D6kE+Hniq6pLFUSnxbdPQNG\nC4y7/v7MCT39ShxH9hEQ/n09P+afIIa8OCt4hKWQD3nuz2C4qZAP7GG75o7DTZLJ\nbcupDhISK3Bee18IhL9xdvpxo/lNj0kWOKfuQNntN23sYUMXDX19yEGt32E3F2xg\nPx7mTrC0xc6k6v5wl/F7LAREiUsdUzydolOFHOjW6hMMtYkv/+BRotOjOkIFTcEI\nWiWpFyR1AgMBAAECggEABWMCa1uR5QxkuihOn+zIAZN+FYfHPfTqOu57YWBNhD/Y\n9eo7PTZBMFdsAHTRTdDpdLio4l9f003fhMpM/N52BWJkPViE/OInIxkxLUhB0K3B\nn7MPO4yHtUk3cq7DvpYdFrwiAbaJeqYS+uB6LSEymfRAT2MvEznPIiNlrIIjw7q4\ndGIMoFWywxuAoDH1ealIJs9AFIQqCIUE5cWKwLp1p1bcBAUexq9UE/xEhVuGwhjO\nam2S3/PyMFtqAP8Jm7zex9HpTkpjdB6dW67bLr2Zlg1J5EOp6ZUwUp9koyM/FaNk\n4Nc+0G26yLDgVXvlmjzvUyiADWMO0FcMsnrKiLlgyQKBgQDTIf9DghqpvQPSVPIM\n+2Dd/USIBNNdJ2p1rv/1tne5Q8PWTOCYnorOLZJTf4RjXlMy97FS3Hx6ICLHGxhA\nWRaWJ0MxQp0QPR0mh6YSAfWYRb8m+OJEg09/YiPNV4uxWpH6axXMLzmFUy25h4H6\n7bkwmQXqNW/OFASdueNPRNkwrQKBgQC1PoU28LcYfRZcRdKhvjtNVnCU1C3vR2Nh\nTk1ed8feZZYJkLWLHcp90Ybj9UntnK9/bQjGvrPNOCLI+K6XgqXnc+2R0dw1qQbg\nRopgpKLEoOahbwCM6qS5+8vkYVZZ9Mf6K9CqKlYL98EdyH06Lz+3G+TV7h4ciJ5p\nu4KR1KET6QKBgB7vAjlf3Iw191Nfwr8ILZ8Ytmu/WUByv/RSGYpkm5H3TcAeL2Ht\nCKVApAm7yyfL8CCtjbt4NnymLLJDVABJjTeetQeInP5+FcFJapmE0/jlbyZEnNIw\n8vrU5C33v63SeUTZX6401RalozNlmTOslM97/BPelGz7HdoGHDnG+pNZAoGAUUXz\n3qsar7SH7mCxDy6K0SdN3K842e+dNOkz8ISt76CGI0vs3LFDOCBY6Kf5ur9kQzPV\nl6m0XvYpQw2g3PgNBTkK8zS0FqvxFWkMbEHna1zrjsCD8qzVqUCsMYEkg7osYpZ7\nP5M0erFiu8RHw2ukmqu7deMNRaZNwwP+jP5YLOECgYBktfk1G/eiyMq4G4icU7jv\nyFdcUv1q/MKVxJH/YyWxdkcz0PeX+j8ZXl64BG9Z+8Wsc+9HRbmgbyyiTq/JpAzH\nFRhJV6AHoir/RooPPKk7i5c1HyTPtil/TsoCDGI9cF15VQ/GB4hl3cFRwWPw1Sgl\nBadKvPfhmrQGIbZ5dHDMqQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "thegraduate-storage@alpine-land-423208-d4.iam.gserviceaccount.com",
  "client_id": "108010815204839051860",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/thegraduate-storage%40alpine-land-423208-d4.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}
`

	client, err := storage.NewClient(context, option.WithCredentialsJSON([]byte(serviceJSON)))

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
