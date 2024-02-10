package initFunction

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"os"
)

func Credentials() *session.Session {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	awsConfigure := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY"), ""),
	}))
	return awsConfigure
}
