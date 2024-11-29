package handlers

import (
	"cacheservice/domain"
	"cacheservice/httpserv"
	"cacheservice/logging"
	"encoding/json"
	"regexp"

	"fmt"
	"net/http"
)

type HandlerInterface interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	GetWithParamHandler(w http.ResponseWriter, r *http.Request)
	PostHandler(w http.ResponseWriter, r *http.Request)
	PutHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
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
	fmt.Fprintf(w, "GET request received. Query param 'Value': %s\n", k.httpserv.Query(r, "name"))
}

func (k *KeyValueHandler) GetWithParamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET request received. Variable 'id': %s\n", k.httpserv.Param(r, "id"))
}

func (k *KeyValueHandler) PostHandler(w http.ResponseWriter, r *http.Request) {

	var kv domain.Keyvalue

	if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !isValidGUID(kv.Key) {
		http.Error(w, "Invalid GUID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "POST request received. Key: %s, Value: %s\n", kv.Key, kv.Value)
}

func (k *KeyValueHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	id := k.httpserv.Param(r, "id")

	if !isValidGUID(id) {
		http.NotFound(w, r)
		return
	}

	var kv domain.Keyvalue

	if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !isValidGUID(kv.Key) {
		http.Error(w, "Invalid GUID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "PUT request received. Key: %s, Value: %s\n", kv.Key, kv.Value)
}

func (k *KeyValueHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := k.httpserv.Param(r, "id")

	if !isValidGUID(id) {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "DELETE request received. Variable 'id': %s\n", k.httpserv.Param(r, "id"))
}

func isValidGUID(guid string) bool {
	re := regexp.MustCompile(`^[a-fA-F0-9]{8}\-[a-fA-F0-9]{4}\-[a-fA-F0-9]{4}\-[a-fA-F0-9]{4}\-[a-fA-F0-9]{12}$`)
	return re.MatchString(guid)
}
