//Http request structs
package http

//Json request of  write-message-on-image endpoint
type JSONEncode struct {
	Path   string `json:"path"`
	Phrase string `json:"phrase"`
}
