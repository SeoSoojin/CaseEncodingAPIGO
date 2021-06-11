package http

type JSONEncode struct {
	Path   string `json:"path"`
	Phrase string `json:"phrase"`
}

type JSONDecode struct {
	Path   string `json:"path"`
	Phrase string `json:"phrase"`
}
