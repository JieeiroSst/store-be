package usecase

import "mime/multipart"

type CreateRequest struct {
	FileHeader *multipart.FileHeader
}