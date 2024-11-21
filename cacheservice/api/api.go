package api

import (
	"cacheservice/domain"
	"cacheservice/logging"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var _logger = &logging.Logger{Name: "api.go"}

var session = domain.UserSession{
	Id:       "8bceb58a-1f07-41af-94a8-73d0fa03f8db",
	Username: "janedoe1",
	Email:    "janedoe1@example.com",
	Preferences: domain.Preference{
		Language: "pt",
		Theme:    "light",
	},
	Expiration:   "2023-04-12T10:30:00Z",
	RequestCount: 0,
}

func StartServer() {
	http.HandleFunc("/api/create", create)
	http.HandleFunc("/api/", index)
	_logger.Log("Server Runnning at: http://localhost:8000")
	_logger.LogFatal("", http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/") // Verifica se o ID passado na URL é o mesmo da Session

	if id == session.Id {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(session)
	} else {
		http.Error(w, "Session not found", http.StatusNotFound)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Lê o corpo da requisição
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserializa o JSON no formato de Session
	var session domain.UserSession
	if err := json.Unmarshal(body, &session); err != nil {
		http.Error(w, "Invalid Session", http.StatusBadRequest)
		return
	}

	// Sucesso
	w.WriteHeader(http.StatusCreated)
	message := fmt.Sprintf("Sessão criada com sucesso: %+v\n", session)
	_logger.LogInfo(message)
	fmt.Fprintln(w, message)
}
