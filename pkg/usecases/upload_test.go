package usecases

import (
	"io/ioutil"
	"testing"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type upRep struct {
}

func (r *upRep) Encode(phrase string, path string) (string, error) {

	return "", nil
}
func (i *upRep) Decode(path string) (string, error) {
	return "", nil
}
func (i *upRep) Get(path string) ([]byte, error) {
	return nil, nil
}
func (i *upRep) Upload(buffer []byte, path string) (string, error) {
	return "", nil
}

func TestUpload(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test.bmp")
	uc := NewUCImageUploader(services.NewImageService(new(upRep), new(upRep), new(upRep), new(upRep)))
	var tests = []struct {
		name      string
		expected  error
		given     []byte
		givenPath string
	}{
		{"success", nil, file, "../../assets/raw/test.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := uc.Upload(tt.given, tt.givenPath)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
