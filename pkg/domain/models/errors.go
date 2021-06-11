package models

import "errors"

//Custom error messages
var (
	ErrFormatNotSupported     = errors.New("Image should be an bmp")
	ErrEmptyString            = errors.New("Any field cant be empty")
	ErrInvalidOrCorruptedFile = errors.New("Invalid or corrupted file")
	ErrInvalidFormatFile      = errors.New("Invalid format file, file should be a bmp")
	ErrSize                   = errors.New("Message dosnt fit in image")
)
