package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/minio"
	"github.com/JIeeiroSst/store/pkg/snowflake"
	"github.com/google/uuid"
	"github.com/qor/media"
	"github.com/qor/media/oss"
)

type Medias interface {
	uploadFile(ctx context.Context, args *UploadObjectArgs) (*UploadObjectResponse, error)
	CreateMedia(ctx context.Context, args *CreateRequest) error
}

type MediaUsecase struct {
	Media       repository.Medias
	MinioClient minio.Client
	Snowflake   snowflake.SnowflakeData
}

type UploadObjectArgs struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadObjectResponse struct {
	URL      string
	FileName string
}

func NewMediaUsecase(Media repository.Medias, MinioClient minio.Client, Snowflake snowflake.SnowflakeData) *MediaUsecase {
	return &MediaUsecase{
		Media:       Media,
		MinioClient: MinioClient,
		Snowflake:   Snowflake,
	}
}

func (s *MediaUsecase) uploadFile(ctx context.Context, args *UploadObjectArgs) (*UploadObjectResponse, error) {
	userMetaData := map[string]string{
		"x-amz-acl": "public-read",
	}
	var fileExtension string
	splitedArr := strings.Split(args.FileHeader.Filename, ".")
	if len(splitedArr) > 0 {
		fileExtension = splitedArr[len(splitedArr)-1]
	}
	uuid := uuid.New().String()
	uuidFileName := fmt.Sprintf("%v.%v", uuid, fileExtension)
	res, err := s.MinioClient.UploadFile(ctx, &minio.UploadFileArgs{
		UserMetaData: userMetaData,
		File:         args.File,
		FileHeader:   args.FileHeader,
		FileName:     uuidFileName,
	})
	if err != nil {
		return nil, err
	}
	return &UploadObjectResponse{URL: res.URL, FileName: uuidFileName}, nil
}

func (s *MediaUsecase) CreateMedia(ctx context.Context, args *CreateRequest) error {
	var res *UploadObjectResponse

	pdfFile, err := args.FileHeader.Open()
	if err != nil {
		return err
	}

	res, err = s.uploadFile(ctx, &UploadObjectArgs{
		File:       pdfFile,
		FileHeader: args.FileHeader,
	})
	if err != nil {
		return err
	}

	siMedia := domain.Media{
		ID:  s.Snowflake.GearedID(),
		URL: res.URL,
		Thumbnail: oss.OSS{
			Base: media.Base{
				FileName: res.FileName,
				Url:      res.URL,
			},
		},
	}
	return s.Media.SaveMedia(ctx, siMedia)
}
