package services

import (
	"strings"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/interfaces"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/models"
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

	if strings.TrimSpace(phrase) == "" {

		return "", models.ErrEmptyString
	}
	if strings.TrimSpace(path) == "" {

		return "", models.ErrEmptyString

	}
	err := checkIfBmp(path)
	if err != nil {
		return "", err
	}

	return imageService.encoder.Encode(phrase, path)
}

func (imageService *ImageService) Decode(path string) (string, error) {

	if strings.TrimSpace(path) == "" {

		return "", models.ErrEmptyString

	}
	err := checkIfBmp(path)
	if err != nil {
		return "", err
	}

	return imageService.decoder.Decode(path)
}

func (imageService *ImageService) Get(path string) ([]byte, error) {

	if strings.TrimSpace(path) == "" {

		return []byte{}, models.ErrEmptyString

	}
	err := checkIfBmp(path)
	if err != nil {
		return []byte{}, err
	}

	return imageService.getter.Get(path)
}

func (imageService *ImageService) Upload(buffer []byte, path string) (string, error) {

	if strings.TrimSpace(path) == "" {

		return "", models.ErrEmptyString

	}
	if len(buffer) < 1 {

		return "", models.ErrInvalidOrCorruptedFile

	}
	return imageService.uploader.Upload(buffer, path)
}

func checkIfBmp(path string) error {

	aux := strings.LastIndex(path, ".") + 1
	if path[aux:] != "bmp" {
		return models.ErrInvalidFormatFile
	}
	return nil

}
