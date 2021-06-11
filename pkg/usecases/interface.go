package usecases

type UCImageEncoder interface {
	Encode(phrase string, path string) (string, error)
}

type UCImageDecoder interface {
	Decode(path string) (string, error)
}

type UCImageGetter interface {
	Get(path string) ([]byte, error)
}

type UCImageUploader interface {
	Upload(buffer []byte, path string) (string, error)
}
