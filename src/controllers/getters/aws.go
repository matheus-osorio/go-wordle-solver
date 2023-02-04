package getters

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Getter struct {
	WordPicker
	Language string
}

func (getter S3Getter) GetWords() []string {
	sess := session.Must(session.NewSession())

	downloader := s3manager.NewDownloader(sess)
	localFile := "/tmp/words.txt"
	f, err := os.Create(localFile)
	if err != nil {
		panic("Could not create file")
	}

	bucketName := os.Getenv("BUCKET_NAME")
	fmt.Println("bucket:", bucketName)
	filename := getter.Language + "-words.txt"
	size, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})

	fmt.Println("Got file of size:", size)
	if err != nil {
		fmt.Println(err)
		panic("Error while downloading file")
	}
	f.Close()

	textByte, err := os.ReadFile(localFile)
	if err != nil {
		panic("Could not fetch file")
	}

	text := string(textByte)

	textArr := strings.Split(text, "\n")

	getter.WordList = textArr
	getter.SelectWordsBySize()
	getter.ReplaceLatinCharacters()

	return getter.WordList
}
