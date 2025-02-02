package uploadService

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nnniyaz/blog/internal/domain/base/uuid"
	"github.com/nnniyaz/blog/pkg/core"
	"io"
	"mime/multipart"
)

// MaxFileSize is 5 MB:
const MaxFileSize = 5 * (1 << 20)

var ErrMaxFileSizeIs1MB = core.NewI18NError(core.EINVALID, core.TXT_MAX_FILE_SIZE_IS_5MB)

type UploadService interface {
	UploadImage(folderName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type uploadService struct {
	s3Client    *s3.S3
	spaceBucket string
}

func NewUploadService(s3Client *s3.S3, spaceBucket string) UploadService {
	return &uploadService{s3Client: s3Client, spaceBucket: spaceBucket}
}

func (s *uploadService) UploadImage(folderName string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	var buf bytes.Buffer
	io.Copy(&buf, file)
	if buf.Len() > MaxFileSize {
		return "", ErrMaxFileSizeIs1MB
	}
	fileName := fileHeader.Filename + "_" + uuid.NewUUID().String()
	_, err := s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:       aws.String(s.spaceBucket),
		Key:          aws.String("/blog/" + folderName + "/" + fileName),
		Body:         bytes.NewReader(buf.Bytes()),
		ACL:          aws.String("public-read"),
		CacheControl: aws.String("max-age=21600000"),
		ContentType:  aws.String(fileHeader.Header.Get("Content-Type")),
	})
	buf.Reset()
	if err != nil {
		return "", err
	}
	return fileName, nil
}
