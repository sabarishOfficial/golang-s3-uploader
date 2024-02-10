package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"s3fileupload/initFunction"
)

func S3SingleFileUploadFunction(w http.ResponseWriter, r *http.Request, ses *session.Session) {

	file, fileHander, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(fileHander.Filename)

	defer file.Close()

	read, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	_, err = s3.New(ses).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("BUCKET_NAME")),
		Key:         aws.String(fileHander.Filename),
		Body:        bytes.NewReader(read),
		ContentType: aws.String(http.DetectContentType(read)),
	})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	status, err := w.Write([]byte("File uploaded successfully"))

	if err != nil {
		panic(err)
	}
	fmt.Println(status)

}

func MultipleFileUploading(w http.ResponseWriter, r *http.Request, ses *session.Session) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, filesReading := range r.MultipartForm.File {
		for _, files := range filesReading {
			fileOpen, err := files.Open()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer fileOpen.Close()

			readValues, err := ioutil.ReadAll(fileOpen)
			if err != nil {
				panic("The file data not reading")
			}

			_, err = s3.New(ses).PutObject(&s3.PutObjectInput{
				Bucket:      aws.String(os.Getenv("BUCKET_NAME")),
				Key:         aws.String(files.Filename),
				Body:        bytes.NewReader(readValues),
				ContentType: aws.String(http.DetectContentType(readValues)),
			})
			if err != nil {
				panic("the file not uploading")
				return
			}

		}

	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Multiple file uploading completed"))
	if err != nil {
		panic("Uploading error")
	}

}
func main() {
	awsConfigure := initFunction.Credentials()

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		S3SingleFileUploadFunction(w, r, awsConfigure)
	})
	http.HandleFunc("/multiple", func(w http.ResponseWriter, r *http.Request) {
		MultipleFileUploading(w, r, awsConfigure)
	})

	err := http.ListenAndServe("0.0.0.0:8000", nil)

	if err != nil {
		log.Fatalf("server stoped")
	}
	fmt.Println("The server start")

}
