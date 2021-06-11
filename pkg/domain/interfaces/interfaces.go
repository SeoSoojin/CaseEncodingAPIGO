package interfaces

type ImageEncoder interface {
	Encode(phrase string, path string) (string, error)
}

type ImageDecoder interface {
	Decode(path string) (string, error)
}

type ImageGetter interface {
	Get(path string) ([]byte, error)
}

type ImageUploader interface {
	Upload(buffer []byte, path string) (string, error)
}
