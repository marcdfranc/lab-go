package handlers

import (
	"cacheservice/httpserv"
	"cacheservice/logging"
	"fmt"
	"net/http"
)

type HandlerInterface interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	GetWithParamHandler(w http.ResponseWriter, r *http.Request)
	HeadHandler(w http.ResponseWriter, r *http.Request)
	PostHandler(w http.ResponseWriter, r *http.Request)
	PutHandler(w http.ResponseWriter, r *http.Request)
	PatchHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
	ConnectHandler(w http.ResponseWriter, r *http.Request)
	OptionsHandler(w http.ResponseWriter, r *http.Request)
	TraceHandler(w http.ResponseWriter, r *http.Request)
}

type KeyValueHandler struct {
	logger   *logging.Logger
	httpserv *httpserv.HttpServer
}

func NewKeyValueHandler(logger *logging.Logger, httpserv *httpserv.HttpServer) *KeyValueHandler {
	return &KeyValueHandler{
		logger:   logger,
		httpserv: httpserv,
	}
}

func (k *KeyValueHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET request received. Query param 'name': %s\n", k.httpserv.Query(r, "name"))
}

func (k *KeyValueHandler) GetWithParamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET request received. Query param 'name': %s\n", k.httpserv.Param(r, "id"))
}

func (k *KeyValueHandler) HeadHandler(w http.ResponseWriter, r *http.Request) {
	// HEAD requests typically do not have a response body
	w.WriteHeader(http.StatusOK)
}

func (k *KeyValueHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST request received.")
}

func (k *KeyValueHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PUT request received.")
}

func (k *KeyValueHandler) PatchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PATCH request received.")
}

func (k *KeyValueHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DELETE request received.")
}

func (k *KeyValueHandler) ConnectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CONNECT request received.")
}

func (k *KeyValueHandler) OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")
}

func (k *KeyValueHandler) TraceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TRACE request received.")
}
