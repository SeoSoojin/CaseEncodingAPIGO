package models

import "errors"

var (
	ErrFormatNotSupported     = errors.New("Image should be an bmp")
	ErrEmptyString            = errors.New("Any field cant be empty")
	ErrInvalidOrCorruptedFile = errors.New("Invalid or corrupted file")
	ErrInvalidFormatFile      = errors.New("Invalid format file, file should be a bmp")
)
