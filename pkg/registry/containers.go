package registry

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/gateway/controllers"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/usecases"
	"github.com/sarulabs/di"
)

//Struct to Container
type Container struct {
	ctn di.Container
}

//Get method implemented for this struct
func (containter *Container) Get(name string) interface{} {

	return containter.ctn.Get(name)
}

//Clean method implemented for this struct
func (container *Container) Clean() error {

	return container.ctn.Clean()
}

//Container creator, receives an []di.Def as param and returns an address of Container or an error
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

//Defs for ImageEncoder container, returns an address of container or an error
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

//Defs for ImageDecoder container, returns an address of container or an error
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

//Defs for ImageGetter container, returns an address of container or an error
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

//Defs for ImageUploader container, returns an address of container or an error
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
