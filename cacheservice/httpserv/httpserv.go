package httpserv

import (
	"cacheservice/logging"
	"context"
	"errors"
	"net/http"
	"regexp"
)

type Server interface {
	Start()
	Param(req *http.Request, name string) string
	Query(req *http.Request, name string) string
}

// Route representa uma única rota.
type Route struct {
	pattern  *regexp.Regexp
	handlers map[string]http.HandlerFunc
}

// HttpServer representa um servidor HTTP personalizado.
type HttpServer struct {
	Host   string
	routes []*Route
	logger *logging.Logger
}

// NewHttpServer cria uma nova instância de HttpServer.
func NewHttpServer(host string, logger *logging.Logger) *HttpServer {
	return &HttpServer{
		Host:   host,
		logger: logger,
	}
}

// Start inicia o servidor HTTP.
func (h *HttpServer) Start() {
	h.logger.LogInfof("The server is running at http://%s", h.Host)
	http.ListenAndServe(h.Host, h)
}

// ServeHTTP implementa a interface Handler do http.
func (h *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range h.routes {
		if matches := route.pattern.FindStringSubmatch(req.URL.Path); matches != nil {
			handler, exists := route.handlers[req.Method]
			if !exists {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			params := make(map[string]string)
			for i, name := range route.pattern.SubexpNames() {
				if i != 0 && name != "" {
					params[name] = matches[i]
				}
			}

			ctx := context.WithValue(req.Context(), "params", params)
			req = req.WithContext(ctx)
			handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// Handle adiciona uma nova rota ao servidor HTTP com suporte a múltiplos métodos.
func (h *HttpServer) Handle(pattern string, method string, handlerFunc http.HandlerFunc) error {
	regexpattern, err := regexp.Compile("^" + h.convertPattern(pattern) + "$")
	if err != nil {
		return errors.New("invalid pattern")
	}

	for _, route := range h.routes {
		if route.pattern.String() == regexpattern.String() {
			route.handlers[method] = handlerFunc
			return nil
		}
	}

	h.routes = append(h.routes, &Route{
		pattern:  regexpattern,
		handlers: map[string]http.HandlerFunc{method: handlerFunc},
	})
	return nil
}

func (h *HttpServer) convertPattern(pattern string) string {
	re := regexp.MustCompile(`\{(\w+)\}`)
	return re.ReplaceAllString(pattern, `(?P<$1>[^/]+)`)
}

// Param retorna o valor de um parâmetro da URL.
func (h *HttpServer) Param(req *http.Request, name string) string {
	params, ok := req.Context().Value("params").(map[string]string)
	if !ok {
		return ""
	}
	return params[name]
}

// Query retorna o valor de um parâmetro de consulta.
func (h *HttpServer) Query(req *http.Request, name string) string {
	return req.URL.Query().Get(name)
}

// patternExists verifica se o padrão já está registrado.
func (h *HttpServer) patternExists(pattern *regexp.Regexp) bool {
	for _, route := range h.routes {
		if route.pattern.String() == pattern.String() {
			return true
		}
	}
	return false
}

func (h *HttpServer) HandleGet(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodGet, handlerFunc)
}

func (h *HttpServer) HandleHead(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodHead, handlerFunc)
}

func (h *HttpServer) HandlePost(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodPost, handlerFunc)
}

func (h *HttpServer) HandlePut(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodPut, handlerFunc)
}

func (h *HttpServer) HandlePatch(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodPatch, handlerFunc)
}

func (h *HttpServer) HandleDelete(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodDelete, handlerFunc)
}

func (h *HttpServer) HandleConnect(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodConnect, handlerFunc)
}

func (h *HttpServer) HandleOptions(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodOptions, handlerFunc)
}

func (h *HttpServer) HandleTrace(pattern string, handlerFunc http.HandlerFunc) error {
	return h.Handle(pattern, http.MethodTrace, handlerFunc)
}
