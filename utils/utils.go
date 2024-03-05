package utils

import (
	"log"
	"os"
    "fmt"
//	"io"
	"strings"
	// "time"
//	"errors"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

)

// GetValue returns configuration value based on a given key from the .env file
func GetValue(key string) string {
	fmt.Println(os.Getenv("GO_ENV"))
	env := os.Getenv("GO_ENV")
    // load the .env file
	fmt.Printf("The env value is %s \n", env)

	if os.Getenv("GO_ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file!!\n")
		}
	}

    // return the value based on a given key
	return os.Getenv(key)
}

func UploadImage(filename string) (string , error) {

	var name string = strings.TrimSuffix(filename, ".jpg") 
	fmt.Println(name)
    key := os.Getenv("SPACES_KEY")
    secret := os.Getenv("SPACES_SECRET")
	endpoint := os.Getenv("SPACES_ENDPOINT")

    s3Config := &aws.Config{
        Credentials: credentials.NewStaticCredentials(key, secret, ""),
        Endpoint:    aws.String(endpoint),
        Region:      aws.String("nyc3"),
        S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
    }

    newSession := session.New(s3Config)
    s3Client := s3.New(newSession)
	fmt.Printf("%#v", s3Client)
	uploader := s3manager.NewUploader(newSession)

	pwd, _ := os.Getwd()

    filepath := pwd + "/assets/" + filename

	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("backend"),
		Key:    aws.String(name),
		Body:   f,
	})
	if err != nil {	
		fmt.Printf("error while uploading photos: %v \n", err)
		return "", err
	}
	fmt.Printf("Result: %#v", result)
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	
	return result.Location, nil
}