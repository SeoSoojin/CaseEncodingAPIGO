package usecases

import (
	"testing"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/services"
)

type getRep struct {
}

func (r *getRep) Encode(phrase string, path string) (string, error) {

	return "", nil
}
func (i *getRep) Decode(path string) (string, error) {
	return "", nil
}
func (i *getRep) Get(path string) ([]byte, error) {
	return nil, nil
}
func (i *getRep) Upload(buffer []byte, path string) (string, error) {
	return "", nil
}

func TestGet(t *testing.T) {

	uc := NewUCImageGetter(services.NewImageService(new(getRep), new(getRep), new(getRep), new(getRep)))
	var tests = []struct {
		name     string
		expected error
		given    string
	}{
		{"success", nil, "../../../assets/raw/test.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := uc.Get(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
