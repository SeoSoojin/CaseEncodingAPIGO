package controllers

import (
	"io/ioutil"
	"testing"
)

func TestEncode(t *testing.T) {
	encoder := NewControllers().imageEncoder
	var tests = []struct {
		name        string
		expected    error
		given       string
		givenPhrase string
	}{
		{"success", nil, "./assets/raw/test.bmp", "test."},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := encoder.Encode(tt.given, tt.givenPhrase)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestDecode(t *testing.T) {
	encoder := NewControllers().imageDecoder
	var tests = []struct {
		name     string
		expected error
		given    string
	}{
		{"success", nil, "/test-encoded.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := encoder.Decode(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestGetter(t *testing.T) {
	encoder := NewControllers().imageGetter
	var tests = []struct {
		name     string
		expected error
		given    string
	}{
		{"success", nil, "/test-encoded.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := encoder.Get(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

func TestUpload(t *testing.T) {
	file, _ := ioutil.ReadFile("test")
	encoder := NewControllers().imageUploader
	var tests = []struct {
		name      string
		expected  error
		given     []byte
		givenPath string
	}{
		{"success", nil, file, "/test-encoded.bmp"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, actual := encoder.Upload(tt.given, tt.givenPath)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
