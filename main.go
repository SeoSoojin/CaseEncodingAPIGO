package main

import (
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/gateway/http"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/registry"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/usecases"
)

//Main function
//Creates all the containers
//Initializes the App struct using all the Containers as params with the proper usecases
func main() {

	imageEncoderCtn, err := registry.NewImageEncoderContainer()
	if err != nil {
		panic(err)
	}
	defer imageEncoderCtn.Clean()

	imageDecoderCtn, err := registry.NewImageDecoderContainer()
	if err != nil {
		panic(err)
	}
	defer imageDecoderCtn.Clean()

	imageGetterCtn, err := registry.NewImageGetterContainer()
	if err != nil {
		panic(err)
	}
	defer imageGetterCtn.Clean()

	imageUploaderCtn, err := registry.NewImageUploaderContainer()
	if err != nil {
		panic(err)
	}
	defer imageUploaderCtn.Clean()

	app := http.NewApp(
		imageEncoderCtn.Get("ImageEncoder").(usecases.UCImageEncoder),
		imageDecoderCtn.Get("ImageDecoder").(usecases.UCImageDecoder),
		imageGetterCtn.Get("ImageGetter").(usecases.UCImageGetter),
		imageUploaderCtn.Get("ImageUploader").(usecases.UCImageUploader),
	)

	app.Run(":3030")
}
