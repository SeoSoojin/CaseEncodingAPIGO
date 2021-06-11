package usecases

import (
	"testing"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type encRep struct {
}

func (r *encRep) Encode(phrase string, path string) (string, error) {

	return "", nil
}
func (i *encRep) Decode(path string) (string, error) {
	return "", nil
}
func (i *encRep) Get(path string) ([]byte, error) {
	return nil, nil
}
func (i *encRep) Upload(buffer []byte, path string) (string, error) {
	return "", nil
}

func TestEncode(t *testing.T) {

	uc := NewUCImageEncoder(services.NewImageService(new(encRep), new(encRep), new(encRep), new(encRep)))
	var tests = []struct {
		name        string
		expected    error
		given       string
		givenPhrase string
	}{
		{"success", nil, "../../../assets/raw/test.bmp", "test."},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := uc.Encode(tt.givenPhrase, tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
