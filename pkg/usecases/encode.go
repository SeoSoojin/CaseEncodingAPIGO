package usecases

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type imageEncoder struct {
	imageService *services.ImageService
}

func NewUCImageEncoder(imageService *services.ImageService) UCImageEncoder {

	return &imageEncoder{imageService: imageService}
}

func (imageEncoder *imageEncoder) Encode(phrase string, path string) (string, error) {

	return imageEncoder.imageService.Encode(phrase, path)

}
