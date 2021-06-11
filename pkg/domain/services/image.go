//Services handle params
package services

import (
	"strings"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/interfaces"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/models"
)

//Struct to give registry access to controllers
type ImageService struct {
	encoder  interfaces.ImageEncoder
	decoder  interfaces.ImageDecoder
	getter   interfaces.ImageGetter
	uploader interfaces.ImageUploader
}

//Creator of ImageService, recieves the interfaces as params and return an address of ImageService struct
func NewImageService(
	encoder interfaces.ImageEncoder,
	decoder interfaces.ImageDecoder,
	getter interfaces.ImageGetter,
	uploader interfaces.ImageUploader) *ImageService {

	return &ImageService{encoder: encoder, decoder: decoder, getter: getter, uploader: uploader}

}

//Error of params handler for Encode()
func (imageService *ImageService) Encode(phrase string, path string) (string, error) {
	//TrimSpace corta a string
	if strings.TrimSpace(phrase) == "" {

		return "", models.ErrEmptyString
	}
	if strings.TrimSpace(path) == "" {

		return "", models.ErrEmptyString

	}
	// declarada no próprio documento
	err := checkIfBmp(path)
	if err != nil {
		return "", err
	}

	return imageService.encoder.Encode(phrase, path)
}

//Error of params handler for Decode()
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

//Error of params handler for Get()
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

//Error of params handler for Upload()
func (imageService *ImageService) Upload(buffer []byte, path string) (string, error) {

	if strings.TrimSpace(path) == "" {

		return "", models.ErrEmptyString

	}
	err := checkIfBmp(path)
	if err != nil {
		return "", err
	}
	if len(buffer) < 1 {

		return "", models.ErrInvalidOrCorruptedFile

	}
	if buffer == nil {

		return "", models.ErrInvalidOrCorruptedFile

	}
	return imageService.uploader.Upload(buffer, path)
}

//Function to check if file extension equals .bmp
func checkIfBmp(path string) error {

	aux := strings.LastIndex(path, ".") + 1
	if path[aux:] != "bmp" {
		return models.ErrInvalidFormatFile
	}
	return nil

}
