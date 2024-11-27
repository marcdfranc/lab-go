package httpserv

import (
	"cacheservice/logging"
	"log"
	"net/http"
	"strings"
)

type urlVars map[string]int

type Route struct {
	path    string
	handler http.HandlerFunc
	urlVars urlVars
}

type Httpserv struct {
	Host   string
	routes []*Route
}

var logger *logging.Logger

func NewHttpserv(host string, loggerIn *logging.Logger) *Httpserv {
	logger = loggerIn
	return &Httpserv{Host: host}
}

func (h *Httpserv) HandleFunc(path string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, &Route{
		path:    path,
		handler: handlerFunc,
		urlVars: loadVars(path),
	})
}

func loadVars(path string) urlVars {
	parts := strings.Split(path, "/")
	var result urlVars
	for i, v := range parts {
		if strings.HasPrefix(v, "{") && strings.HasSuffix(v, "}") {
			result[v[1:len(v)-1]] = i
		}
	}
	return result
}

func (h *Httpserv) Start() {
	for _, v := range h.routes {
		http.HandleFunc(v.path, v.handler)

	}

	logger.LogInfof("Server running at http://%s", h.Host)

	log.Fatal(http.ListenAndServe(h.Host, nil))
}
