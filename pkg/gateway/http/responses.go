package http

type JSONDecodeRes struct {
	Path string `json:"message"`
}

type JSONEncodeRes struct {
	Path string `json:"path"`
}

type JSONUploadRes struct {
	Path string `json:"path"`
}

type JSONErrorRespose struct {
	Error string `json:"error"`
}
