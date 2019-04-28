package app

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os/exec"
	"solarium-golang/internal/config_reader"
)

func AddFileToS3(file []byte, deviceId string) error {
	myConfig := config_reader.Config()

	s, err := session.NewSession()
	out, err := exec.Command("uuidgen").Output()

	bucketKey := deviceId + "/" + string(out)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(myConfig.S3BucketName),
		Key:                  aws.String(bucketKey),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(file),
		ContentType:          aws.String(http.DetectContentType(file)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
