//Test controllers functions
package controllers

import (
	"io/ioutil"
	"testing"
)

//Test if Encode dosn't return a error
func TestEncode(t *testing.T) {
	encoder := NewControllers()
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
			_, actual := encoder.Encode(tt.givenPhrase, tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

//Test if Decode dosn't return a error
func TestDecode(t *testing.T) {
	encoder := NewControllers()
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
			_, actual := encoder.Decode(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

//Test if Getter dosn't return a error
func TestGetter(t *testing.T) {
	encoder := NewControllers()
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
			_, actual := encoder.Get(tt.given)
			if actual != tt.expected {
				t.Errorf("(%+v): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}

//Test if Uploader dosn't return a error
func TestUpload(t *testing.T) {
	file, _ := ioutil.ReadFile("../../../test.bmp")
	encoder := NewControllers()
	var tests = []struct {
		name      string
		expected  error
		given     []byte
		givenPath string
	}{
		{"success", nil, file, "../../../assets/raw/test.bmp"},
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
