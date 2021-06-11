//External communication, handle with server related functions
package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/usecases"
	"github.com/gorilla/mux"
)

//App struct with router and usecases
type App struct {
	Router        *mux.Router
	imageEncoder  usecases.UCImageEncoder
	imageDecoder  usecases.UCImageDecoder
	imageGetter   usecases.UCImageGetter
	imageUploader usecases.UCImageUploader
}

//Creator of App, receives all the usecases as params and return a addres to App
func NewApp(
	imageEncoder usecases.UCImageEncoder,
	imageDecoder usecases.UCImageDecoder,
	imageGetter usecases.UCImageGetter,
	imageUploader usecases.UCImageUploader) *App {

	app := App{imageEncoder: imageEncoder, imageDecoder: imageDecoder, imageGetter: imageGetter, imageUploader: imageUploader}
	app.Router = mux.NewRouter()
	app.initializeRoutes()

	return &app

}

//Handler to write-message-on-image endpoint
//Minor error handlers
//Response handler
func (h *App) Encode(w http.ResponseWriter, r *http.Request) {

	jsonResponse := []JSONEncodeRes{}
	jsonErrorResponse := []JSONErrorRespose{}
	body := r.Body
	jsonRequest := JSONEncode{}
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(jsonErrorResponse)
	}
	//pass json to struct in golang
	json.Unmarshal([]byte(b), &jsonRequest)
	phrase := jsonRequest.Phrase
	path := jsonRequest.Path
	newPath, err := h.imageEncoder.Encode(phrase, path)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return

	}

	jsonResponse = append(jsonResponse, JSONEncodeRes{newPath})
	json.NewEncoder(w).Encode(jsonResponse)

}

//Handler to decode-message-on-image endpoint
//Minor error handlers
//Response handler
func (h *App) Decode(w http.ResponseWriter, r *http.Request) {

	jsonResponse := []JSONDecodeRes{}
	jsonErrorResponse := []JSONErrorRespose{}
	urlReq := r.URL.Path
	aux := strings.LastIndex(urlReq, "/")
	pathFinal := "./assets/encoded/" + urlReq[aux:]
	str, err := h.imageDecoder.Decode(pathFinal)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return
	}
	jsonResponse = append(jsonResponse, JSONDecodeRes{str})
	json.NewEncoder(w).Encode(jsonResponse)

}

//Handler to get endpoint
//Minor error handlers
//Response handler
func (h *App) Get(w http.ResponseWriter, r *http.Request) {

	jsonErrorResponse := []JSONErrorRespose{}
	urlReq := r.URL.Path
	aux := strings.LastIndex(urlReq, "/")
	pathFinal := "./assets/encoded/" + urlReq[aux:]
	data, err := h.imageGetter.Get(pathFinal)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return
	}
	w.Header().Set("Content-Type", "image/bmp")
	w.Write(data)

}

//Handler to upload endpoint
//Minor error handlers
//Response handler
func (h *App) Upload(w http.ResponseWriter, r *http.Request) {

	jsonResponse := []JSONUploadRes{}
	jsonErrorResponse := []JSONErrorRespose{}
	r.ParseMultipartForm(1024 * 16)
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return
	}
	filename := header.Filename
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return
	}
	path := "./assets/raw/" + filename
	newPath, err := h.imageUploader.Upload(buffer, path)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return
	}
	jsonResponse = append(jsonResponse, JSONUploadRes{newPath})
	json.NewEncoder(w).Encode(jsonResponse)

}

//Function to initialize all the endpoint routes
func (h *App) initializeRoutes() {

	h.Router.HandleFunc("/write-message-on-image", h.Encode).Methods("POST")
	h.Router.HandleFunc("/decode-message-from-image/{file}", h.Decode).Methods("GET")
	h.Router.HandleFunc("/get-image/{file}", h.Get).Methods("GET")
	h.Router.HandleFunc("/upload", h.Upload).Methods("POST")

}

//Function to start the server
func (h *App) Run(addr string) {

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, h.Router))
}
