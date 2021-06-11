package http

import (
	"encoding/json"
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

func (h *App) Decode(w http.ResponseWriter, r *http.Request) {

	jsonResponse := []JSONDecodeRes{}
	jsonErrorResponse := []JSONErrorRespose{}
	urlReq := r.URL.Path
	str, err := h.imageDecoder.Decode(urlReq)
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

func (h *App) Get(w http.ResponseWriter, r *http.Request) {

	jsonErrorResponse := []JSONErrorRespose{}
	urlReq := r.URL.Path
	data, err := h.imageGetter.Get(urlReq)
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
	path := header.Filename
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Body of error: %s", err)
		jsonErrorResponse = append(jsonErrorResponse, JSONErrorRespose{err.Error()})
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErrorResponse)
		return
	}
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
