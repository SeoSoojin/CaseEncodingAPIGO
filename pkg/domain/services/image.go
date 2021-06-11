package services

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/interfaces"
)

type ImageService struct {
	encoder  interfaces.ImageEncoder
	decoder  interfaces.ImageDecoder
	getter   interfaces.ImageGetter
	uploader interfaces.ImageUploader
}

func NewImageService(
	encoder interfaces.ImageEncoder,
	decoder interfaces.ImageDecoder,
	getter interfaces.ImageGetter,
	uploader interfaces.ImageUploader) *ImageService {

	return &ImageService{encoder: encoder, decoder: decoder, getter: getter, uploader: uploader}

}

func (imageService *ImageService) Encode(phrase string, path string) (string, error) {

	return imageService.encoder.Encode(phrase, path)
}

func (imageService *ImageService) Decode(path string) (string, error) {

	return imageService.decoder.Decode(path)
}

func (imageService *ImageService) Get(path string) ([]byte, error) {

	return imageService.getter.Get(path)
}

func (imageService *ImageService) Upload(buffer []byte, path string) (string, error) {

	return imageService.uploader.Upload(buffer, path)
}
