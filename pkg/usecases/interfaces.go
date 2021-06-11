//Interfaces for use cases
package usecases

//UCImageEncoder use case to encoder interface
type UCImageEncoder interface {
	Encode(phrase string, path string) (string, error)
}

//UCImageEncoder use case to decoder interface
type UCImageDecoder interface {
	Decode(path string) (string, error)
}

//UCImageEncoder use case to getter interface
type UCImageGetter interface {
	Get(path string) ([]byte, error)
}

//UCImageEncoder use case to uploader interface
type UCImageUploader interface {
	Upload(buffer []byte, path string) (string, error)
}
