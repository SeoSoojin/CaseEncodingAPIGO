package usecases

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type imageUploader struct {
	imageService *services.ImageService
}

func NewUCImageUploader(imageService *services.ImageService) UCImageUploader {

	return &imageUploader{imageService: imageService}

}
func (imageUploader *imageUploader) Upload(buffer []byte, path string) (string, error) {

	return imageUploader.imageService.Upload(buffer, path)

}
