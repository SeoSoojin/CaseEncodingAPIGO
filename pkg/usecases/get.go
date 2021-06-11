//Use case for getter
package usecases

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type imageGetter struct {
	imageService *services.ImageService
}

func NewUCImageGetter(imageService *services.ImageService) UCImageGetter {

	return &imageGetter{imageService: imageService}

}

func (imageGetter *imageGetter) Get(path string) ([]byte, error) {

	return imageGetter.imageService.Get(path)
}
