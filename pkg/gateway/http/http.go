package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SeoSoojin/CaseEncodingAPIGO/pkg/usecases"
	"github.com/gorilla/mux"
)

type App struct {
	Router        *mux.Router
	imageEncoder  usecases.UCImageEncoder
	imageDecoder  usecases.UCImageDecoder
	imageGetter   usecases.UCImageGetter
	imageUploader usecases.UCImageUploader
}

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

func (h *App) Encode(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Cu http")
	jsonResponse := []JSONEncodeRes{}
	body := r.Body
	jsonRequest := JSONEncode{}
	b, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(b), &jsonRequest)
	phrase := jsonRequest.Phrase
	path := jsonRequest.Path
	newPath, _ := h.imageEncoder.Encode(phrase, path)
	jsonResponse = append(jsonResponse, JSONEncodeRes{newPath})
	json.NewEncoder(w).Encode(jsonResponse)

}

func (h *App) Decode(w http.ResponseWriter, r *http.Request) {

	jsonResponse := []JSONDecodeRes{}
	urlReq := r.URL.Path
	str, _ := h.imageDecoder.Decode(urlReq)
	jsonResponse = append(jsonResponse, JSONDecodeRes{str})
	json.NewEncoder(w).Encode(jsonResponse)

}

func (h *App) Get(w http.ResponseWriter, r *http.Request) {

	urlReq := r.URL.Path
	data, _ := h.imageGetter.Get(urlReq)
	w.Header().Set("Content-Type", "image/bmp")
	w.Write(data)

}

func (h *App) Upload(w http.ResponseWriter, r *http.Request) {

	jsonResponse := []JSONUploadRes{}
	r.ParseMultipartForm(1024 * 16)
	file, header, _ := r.FormFile("file")
	path := header.Filename
	buffer, _ := ioutil.ReadAll(file)
	newPath, _ := h.imageUploader.Upload(buffer, path)
	jsonResponse = append(jsonResponse, JSONUploadRes{newPath})
	json.NewEncoder(w).Encode(jsonResponse)

}

func (h *App) initializeRoutes() {

	h.Router.HandleFunc("/write-message-on-image", h.Encode).Methods("POST")
	h.Router.HandleFunc("/decode-message-from-image/{file}", h.Decode).Methods("GET")
	h.Router.HandleFunc("/get-image/{file}", h.Get).Methods("GET")
	h.Router.HandleFunc("/upload", h.Upload).Methods("POST")

}

func (h *App) Run(addr string) {

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, h.Router))
}
