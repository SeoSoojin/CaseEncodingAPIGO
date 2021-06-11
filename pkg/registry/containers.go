package registry

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/gateway/controllers"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/usecases"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func (containter *Container) Get(name string) interface{} {

	return containter.ctn.Get(name)
}

func (container *Container) Clean() error {

	return container.ctn.Clean()
}

func NewContainer(defs []di.Def) (*Container, error) {

	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	err = builder.Add(defs...)
	if err != nil {
		return nil, err
	}
	return &Container{ctn: builder.Build()}, nil
}

func NewImageEncoderContainer() (*Container, error) {

	defs := []di.Def{
		{
			Name: "ImageEncoder",
			Build: func(ctn di.Container) (interface{}, error) {

				return usecases.NewUCImageEncoder(services.NewImageService(controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers())), nil

			},
		},
	}
	return NewContainer(defs)
}

func NewImageDecoderContainer() (*Container, error) {

	defs := []di.Def{
		{
			Name: "ImageDecoder",
			Build: func(ctn di.Container) (interface{}, error) {

				return usecases.NewUCImageDecoder(services.NewImageService(controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers())), nil

			},
		},
	}
	return NewContainer(defs)

}

func NewImageGetterContainer() (*Container, error) {

	defs := []di.Def{
		{
			Name: "ImageGetter",
			Build: func(ctn di.Container) (interface{}, error) {

				return usecases.NewUCImageGetter(services.NewImageService(controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers())), nil

			},
		},
	}
	return NewContainer(defs)

}

func NewImageUploaderContainer() (*Container, error) {

	defs := []di.Def{
		{
			Name: "ImageUploader",
			Build: func(ctn di.Container) (interface{}, error) {

				return usecases.NewUCImageUploader(services.NewImageService(controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers(), controllers.NewControllers())), nil

			},
		},
	}
	return NewContainer(defs)

}
