// aws_s3_iam_example project main.go
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Uploads content to a user specific S3 content
//
//	usage:   go run aws_s3_example

func main() {

	fmt.Println("trying to access the s3")
	if len(os.Args) < 4 {
		fmt.Println("too little arguments")
		return
	}

	fmt.Println("-----")
	sess := session.New(

		&aws.Config{
			Region:      aws.String(os.Args[1]),
			Credentials: credentials.NewStaticCredentials(os.Args[2], os.Args[3], ""),
		})

	svc := s3.New(sess)

	i := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: aws.String(os.Args[4]),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return

	}

}
