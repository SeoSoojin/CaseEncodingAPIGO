package usecases

import (
	"testing"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type decRep struct {
}

func (r *decRep) Encode(phrase string, path string) (string, error) {

	return "", nil
}
func (i *decRep) Decode(path string) (string, error) {
	return "", nil
}
func (i *decRep) Get(path string) ([]byte, error) {
	return nil, nil
}
func (i *decRep) Upload(buffer []byte, path string) (string, error) {
	return "", nil
}

func TestDecode(t *testing.T) {
	uc := NewUCImageDecoder(services.NewImageService(new(decRep), new(decRep), new(decRep), new(decRep)))
	var tests = []struct {
		name     string
		expected error
		given    string
	}{
		{"success", nil, "../../../assets/encoded/test-encoded.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := uc.Decode(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
