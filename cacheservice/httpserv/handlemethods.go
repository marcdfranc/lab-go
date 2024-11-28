package httpserv

import "net/http"

// HandleMethods define uma interface para os métodos de manipulação de rota.
type HandleMethods interface {
	HandleGet(pattern string, handlerFunc http.HandlerFunc) error
	HandleHead(pattern string, handlerFunc http.HandlerFunc) error
	HandlePost(pattern string, handlerFunc http.HandlerFunc) error
	HandlePut(pattern string, handlerFunc http.HandlerFunc) error
	HandlePatch(pattern string, handlerFunc http.HandlerFunc) error
	HandleDelete(pattern string, handlerFunc http.HandlerFunc) error
	HandleConnect(pattern string, handlerFunc http.HandlerFunc) error
	HandleOptions(pattern string, handlerFunc http.HandlerFunc) error
	HandleTrace(pattern string, handlerFunc http.HandlerFunc) error
	patternExists(pattern string) bool
}
