//Interface to business rules
package interfaces

//ImageEncoder
type ImageEncoder interface {
	//Encode encodes message on an image
	Encode(phrase string, path string) (string, error)
}

//ImageDecoder
type ImageDecoder interface {
	//Decode decodes a message from a image
	Decode(path string) (string, error)
}

//ImageGetter
type ImageGetter interface {
	//Get get a path then return the file in this path
	Get(path string) ([]byte, error)
}

//ImageUploader
type ImageUploader interface {
	//Upload receives an image buffer and writes it on a given path
	Upload(buffer []byte, path string) (string, error)
}
