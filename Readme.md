
# S3 File Upload with AWS SDK for Go

This Go application demonstrates how to upload files to Amazon S3 (Simple Storage Service) using the AWS SDK for Go.

## Features

- Single file upload to S3
- Multiple file upload to S3
- Error handling for file uploads

## Dependencies
- AWS SDK for Go
- initFunction (Please replace this with the appropriate link or description of the package)

## Setup
Ensure you have Go installed on your machine.
Clone this repository to your local machine:
```
https://github.com/sabarishOfficial/golang-s3-uploader.git
```
## Install dependencies:

```
 go mod tidy
```
Set up your AWS credentials. You can do this by configuring your AWS credentials file or by setting environment variables. Refer to the AWS documentation for more information.

## create .env file
```
ACCESS_KEY=""
SECRET_KEY=""
REGION=""
BUCKET_NAME=""
```
past the above values

## run the application:
```
go run main.go
```
## Usage
## Single File Upload

To upload a single file to S3, send a POST request to /upload endpoint with the file attached. Example using curl:
```
curl -X POST -F "file=@<path_to_file>" http://localhost:8000/upload
```
Replace <path_to_file> with the path to the file you want to upload.

## Multiple File Upload

To upload multiple files to S3, send a POST request to /multiple endpoint with the files attached. Example using curl:
```
curl -X POST -F "file=@<path_to_file_1>" -F "file=@<path_to_file_2>" http://localhost:8000/multiple
```