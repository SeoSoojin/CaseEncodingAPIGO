//Use case for decoder
package usecases

import "github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"

type imageDecoder struct {
	imageService *services.ImageService
}

func NewUCImageDecoder(imageService *services.ImageService) UCImageDecoder {

	return &imageDecoder{imageService: imageService}
}

func (imageDecoder *imageDecoder) Decode(path string) (string, error) {

	return imageDecoder.imageService.Decode(path)

}
