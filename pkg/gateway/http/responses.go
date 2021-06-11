//Http response struct
package http

//JSON response to write-message-on-image endpoint
type JSONDecodeRes struct {
	Path string `json:"message"`
}

//JSON response to decode-message-from-image endpoint
type JSONEncodeRes struct {
	Path string `json:"path"`
}

//JSON response to upload endpoint
type JSONUploadRes struct {
	Path string `json:"path"`
}

//JSON response to any error
type JSONErrorRespose struct {
	Error string `json:"error"`
}
