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

//Controllers struct
type Controllers struct {
	imageEncoder  interfaces.ImageEncoder
	imageDecoder  interfaces.ImageDecoder
	imageGetter   interfaces.ImageGetter
	imageUploader interfaces.ImageUploader
}

//Creator of controllers, returns an address of controller
func NewControllers() *Controllers {

	return &Controllers{}

}

//Encoding function
//Starts encoding at byte 54 to prevent writing on bmp header
//How it works is commented during the function, to make it easier to understand
//Returns string with the new path or an error
func (s *Controllers) Encode(phrase string, path string) (string, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	count := 0
	if strings.HasSuffix(phrase, ".") == false {

		phrase = phrase + "."
	}
	for i := 0; i < len(phrase); i++ {

		//Transforms the char in a byte in decimal
		auxForByte := byte(phrase[i])
		//Transforms this decimal to binary
		dataAux := strconv.FormatInt(int64(auxForByte), 2)
		//Assure that binary byte has 8 digits
		if n := utf8.RuneCountInString(dataAux); n < 8 {

			dataAux = strings.Repeat("0", 8-n) + dataAux

		}
		//Save each bit in last bit of image byte
		for j := 0; j < 8; j++ {

			//Convert the decimal representation of the bit to binary
			bit := -(48 - dataAux[j])
			//Math to properly do an operation in the last bit of image byte
			operation := -(-2*int(bit) + 1)
			//All of above are to save processing in this part, reducing a comparasion for each image byte
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

//Decoding function
//Decode the message by reading last bite of each byte on image
//Save this bits in groups of 8, forming a byte
//Decode this bytes and push it on str string
//Returns the decoded message or an error
func (s *Controllers) Decode(path string) (string, error) {

	data, err := ioutil.ReadFile(path)
	count := 0
	byte := 0
	//Receives decoded message
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

//Function to get a image
//Returns the []byte of the image
func (s *Controllers) Get(path string) ([]byte, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//Function to upload a image
//Write the image on server
//Returns the path to the uploaded image
func (s *Controllers) Upload(buffer []byte, path string) (string, error) {

	err := ioutil.WriteFile(path, buffer, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}
