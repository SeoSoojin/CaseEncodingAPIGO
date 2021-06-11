package controllers

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/interfaces"
	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/domain/models"
)

type Controllers struct {
	imageEncoder  interfaces.ImageEncoder
	imageDecoder  interfaces.ImageDecoder
	imageGetter   interfaces.ImageGetter
	imageUploader interfaces.ImageUploader
}

func NewControllers() *Controllers {

	return &Controllers{}

}

func (s *Controllers) Encode(phrase string, path string) (string, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	count := 0
	for i := 0; i < len(phrase); i++ {

		auxForByte := byte(phrase[i])
		dataAux := strconv.FormatInt(int64(auxForByte), 2)
		if n := utf8.RuneCountInString(dataAux); n < 8 {

			dataAux = strings.Repeat("0", 8-n) + dataAux

		}
		for j := 0; j < 8; j++ {

			bit := -(48 - dataAux[j])
			operation := -(-2*int(bit) + 1)
			if data[54+count]%2 != bit {

				data[54+count] = data[54+count] + byte(operation)

			}
			count++

		}
	}
	pathAux := strings.Replace(path, ".bmp", "-encoded.bmp", 1)
	newPath := strings.Replace(pathAux, "raw", "encoded", 1)
	ioutil.WriteFile(newPath, data, 0644)
	return newPath, nil

}

func (s *Controllers) Decode(path string) (string, error) {

	aux := strings.LastIndex(path, "/")
	pathFinal := "./assets/encoded/" + path[aux:]
	data, err := ioutil.ReadFile(pathFinal)
	count := 0
	byte := 0
	str := ""
	if err != nil {

		return "", err
	}
	if len(data) < 54 {

		return "", models.ErrFormatNotSupported

	}
	for i := 54; i < len(data); i++ {

		byte = byte + int(data[i]%2)*int(math.Pow(2, float64(7-count)))
		if count == 7 {
			strAux := string(rune(byte))
			str += strAux
			byte = 0
			count = 0
			if strAux == "." {
				break
			}
		} else {
			count++
		}
	}
	return str, nil

}

func (s *Controllers) Get(path string) ([]byte, error) {

	aux := strings.LastIndex(path, "/")
	pathFinal := "./assets/encoded/" + path[aux:]
	data, err := ioutil.ReadFile(pathFinal)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Controllers) Upload(buffer []byte, path string) (string, error) {

	newPath := "./assets/raw/" + path
	err := ioutil.WriteFile(newPath, buffer, 0644)
	if err != nil {
		return "", err
	}

	return newPath, nil
}
