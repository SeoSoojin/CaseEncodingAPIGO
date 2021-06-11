package services

import (
	"io/ioutil"
	"testing"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/models"
)

type imgRep struct {
}

func (i *imgRep) Encode(phrase string, path string) (string, error) {
	return "", nil
}
func (i *imgRep) Decode(path string) (string, error) {
	return "", nil
}
func (i *imgRep) Get(path string) ([]byte, error) {
	return nil, nil
}
func (i *imgRep) Upload(buffer []byte, path string) (string, error) {
	return "", nil
}

func TestEncode(t *testing.T) {

	serv := NewImageService(new(imgRep), new(imgRep), new(imgRep), new(imgRep))
	var tests = []struct {
		name        string
		expected    error
		givenPhrase string
		givenPath   string
	}{
		{"Sucess", nil, "test", "test.bmp"},
		{"Not a bmp", models.ErrInvalidFormatFile, "test", "test"},
		{"Empty path", models.ErrEmptyString, "test", ""},
		{"Empty phrase", models.ErrEmptyString, "", "test.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := serv.Encode(tt.givenPhrase, tt.givenPath)
			if actual != tt.expected {
				t.Errorf("(%+v)/(%+v): expected %s, actual %s", tt.givenPhrase, tt.givenPath, tt.expected, actual)
			}

		})
	}
}

func TestDecode(t *testing.T) {

	serv := NewImageService(new(imgRep), new(imgRep), new(imgRep), new(imgRep))
	var tests = []struct {
		name     string
		expected error
		given    string
	}{
		{"Sucess", nil, "test.bmp"},
		{"Not a bmp", models.ErrInvalidFormatFile, "test"},
		{"Empty path", models.ErrEmptyString, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := serv.Decode(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestGet(t *testing.T) {

	serv := NewImageService(new(imgRep), new(imgRep), new(imgRep), new(imgRep))
	var tests = []struct {
		name     string
		expected error
		given    string
	}{
		{"Sucess", nil, "test.bmp"},
		{"Not a bmp", models.ErrInvalidFormatFile, "test"},
		{"Empty path", models.ErrEmptyString, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := serv.Get(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestUpload(t *testing.T) {

	file, _ := ioutil.ReadFile("../../../assets/raw/test.bmp")
	emptyBuffer := []byte{}
	serv := NewImageService(new(imgRep), new(imgRep), new(imgRep), new(imgRep))
	var tests = []struct {
		name        string
		expected    error
		givenBuffer []byte
		given       string
	}{
		{"Sucess", nil, file, "test.bmp"},
		{"Not a bmp", models.ErrInvalidFormatFile, file, "test"},
		{"Empty path", models.ErrEmptyString, file, ""},
		{"Empty buffer", models.ErrInvalidOrCorruptedFile, emptyBuffer, "test.bmp"},
		{"Nil buffer", models.ErrInvalidOrCorruptedFile, nil, "test.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := serv.Upload(tt.givenBuffer, tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
