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
  "private_key_id": "5e657131513d54b03a4a5e1ca92c451f14c3c377",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCXYTMRLr1UBtAm\nQnV+cvsjPop3vhWZTyaet4brJpZoxh5fN1ieAc1hOkxS7GyX4WmbckBlIxcSmBJe\n2ozlSYI5HUf8xVCbL3/35X/o7XKs4QRpPLbubF323qphCn6xoBHGeWXJqWc4eKmQ\nJUCSpHlB3Dp6znjz9wUptCR70Trf/4ec5FHLilZPrutD6pr6JYTancwWktAh9wxY\ntH85vbc6Pa2UU8X3U7BAVs0i04lU/eYVipyTUUmfS6TYkjFnVct4Zf9OPowuxOum\nTEFqyHQ9VeQGRCd4fFnSibcdvkqJzBvGvaXJkKGylsmSNKY9zgGvNnqo/j5YYQm/\n0YQjfg1nAgMBAAECggEABTUHEFXplfoEr5w6SxbiGUyYb2quzC7ofC3XSpVxjQ8W\nPQjmrSC7aNccMv9bnj8sZngf9kVRENxW8sOqdGSlQcOxFeUk/aw88aZc/QKAnv5q\ngW2Ljq3+7slVVls0ZSBwHBdRmgXlaBA2fPUwaxrp6jW3c+v68Ha/IGJy8tNiprvS\nNBgnFJcSj7XjT1yMW44aWmnHZtO+scywudn3m3GLWxRY0y96jXGfrHMN2ftWG6jd\nhvI/p4HBk15hOvCek3MKaPUWwtPhs4pIePZNETNDWg5LQTfsMy/U+YGUXXXuvZDE\ngP4eAFt6Icqg0bvkV65Mhr50Pb5aZ4eP5J1fYX9QaQKBgQDVmkWt5ebwbmmGr1bl\nRssxpuhcbCzQz3dXDsowwX1WX8D0Ou4CTgLzkUGJ9rpLhvxBcvzh1TXVW2cMXsH3\nrajh+V+mCX7RhP4mmEjQUsEVZ3+XDWtA5HbeM9g3i4bepgAEseqq+hOkn1oUOUMs\nCvuxSOXSJieYfmepzlCSI4nIJQKBgQC1bTYSALlCgxRe8yZiX9WpG4sjNZrurkle\nA4shOgD/T6Nh/+qVHGgKig5NFe7pWuX9O4Fb6SMqs94EUhWz8eSSbzimRWUpmbDQ\nRO0ROx5vmxdyxFMxtDsQfsvIbwRpG5xPz/mDyNj0JjDd4n8ASNrYZpUPcukac7yl\n/T7uHaCzmwKBgQDFv+JcYShRVNbLQw7eipDZMd2Smu9OjZO9fErBopAn49QeB2SL\nq69d0Q9s64eDHr/w08S8SohhMxKOyDco1pQwSI89hzLqRAaKspdOjh6Fc2ES9fBi\njmBhzHuSyNCe1e3UauaHTafb008ww+alIrHrakKiMIKd0gCkmwXKla4cJQKBgHrO\nDuCGvuAHrB1NRvkl15FdRIQzv8i6Do1lpqJ6CbzH4lbujJIcJUylAGMraolobu9H\nfRuCprw7/hc8nNYt8UMxp83cgUp0btu7HCEivEpdV7TmsU2eqrBSs4vbocen2dmo\ndEPBCyHQTBmVO9QsfH4oidNoO+Mc3mAZGKTDR0jrAoGBALLzZBhoaa8gVPL2iL/o\nZ1Xxkh1zyGfvMn1vVf7rK6XEcITn8jIMl9/Vd2Q7ExquNURterDLWXA7ZwWNtXZ5\nNsHFC7OxgoCyA9zN9MGRGuzbtV91rXqJAheSK6M68rUdkkYN7B07hVxhn7LD3AS/\nMXssdY8bkow8G2AEtjsOqgAV\n-----END PRIVATE KEY-----\n",
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
