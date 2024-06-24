package upload

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	ClientBucket UploadBucket = "client-bucket"
)

type (
	UploadBucket string

	UploadFile struct {
		Bucket   UploadBucket
		FileName string
		FileId   string
		FileBody io.Reader
	}

	BlobService struct {
		svc *s3.S3
	}
)

func NewBlobService() *BlobService {
	env := config.GetCofig()
	sess := session.Must(session.NewSession())

	svc := s3.New(sess, &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(env.BLOB_REGION),
		Endpoint:         aws.String(env.BLOB_S3_URL),
		Credentials:      credentials.NewStaticCredentials(env.BLOB_ACCESS_KEY, env.BLOB_PROJECT_KEY, ""),
	})

	return &BlobService{
		svc: svc,
	}
}

func (uf *UploadFile) GetFilePath() string {
	fileExtension := filepath.Ext(uf.FileName)

	return fmt.Sprintf("%s%s", uf.FileId, fileExtension)
}

func (uf *UploadFile) GetContentType() (*string, error) {
	var buffer []byte = make([]byte, 512)

	seeker, ok := uf.FileBody.(io.Seeker)

	if !ok {
		return nil, exec.FILE_NOT_ABLE_UPLOAD
	}

	n, err := uf.FileBody.Read(buffer)

	if err != nil {
		return nil, err
	}

	seeker.Seek(0, 0)
	return aws.String(http.DetectContentType(buffer[:n])), nil
}

func (blob *BlobService) UploadFile(input *UploadFile) (string, error) {
	var err error

	contentType, err := input.GetContentType()

	if err != nil {
		return "", exec.FILE_NOT_ABLE_UPLOAD
	}

	_, err = blob.svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(string(input.Bucket)),
		Key:         aws.String(input.GetFilePath()),
		Body:        aws.ReadSeekCloser(input.FileBody),
		ACL:         aws.String("public-read"),
		ContentType: contentType,
	})

	if err != nil {
		fmt.Println("Error uploading file:", err)
		return "", err
	}

	return blob.GetUrl(string(input.Bucket), input.GetFilePath()), nil
}

func (blob *BlobService) GetUrl(bucket, file string) string {
	env := config.GetCofig()

	return fmt.Sprintf("%s/%s/%s", env.BLOB_PUBLIC_URL, bucket, file)
}
