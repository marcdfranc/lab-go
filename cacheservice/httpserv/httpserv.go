package httpserv

import (
	"cacheservice/logging"
	"context"
	"errors"
	"net/http"
	"regexp"
)

// Route representa uma única rota
type Route struct {
	pattern       *regexp.Regexp
	handler       http.HandlerFunc
	allowedMethod string
}

// HttpServer é um servidor HTTP personalizado
type HttpServer struct {
	Host   string
	routes []*Route
	logger *logging.Logger
}

// NewHttpServer cria uma nova instância do HttpServer
func NewHttpServer(host string, logger *logging.Logger) *HttpServer {
	return &HttpServer{
		Host:   host,
		logger: logger,
	}
}

// Start inicia o servidor HTTP
func (h *HttpServer) Start() {
	for _, route := range h.routes {
		http.HandleFunc(route.pattern.String(), route.handler)
	}
	h.logger.LogInfof("The server is running at http://%s", h.Host)
	http.ListenAndServe(h.Host, h)
}

// ServeHTTP implementa a Interface
func (h *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range h.routes {
		if matches := route.pattern.FindStringSubmatch(req.URL.Path); matches != nil {
			if route.allowedMethod != req.Method {
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
			route.handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// handleFunc adiciona uma nova rota ao servidor HTTP
func (h *HttpServer) handleFunc(allowedMethod, pattern string, handlerFunc http.HandlerFunc) error {
	if h.patternExists(pattern) {
		return errors.New("pattern already exists")
	}
	regexpattern, err := regexp.Compile("^" + h.convertPattern(pattern) + "$")
	if err != nil {
		return errors.New("invalid pattern")
	}
	h.routes = append(h.routes, &Route{
		pattern:       regexpattern,
		handler:       handlerFunc,
		allowedMethod: allowedMethod,
	})
	return nil
}

func (h *HttpServer) convertPattern(pattern string) string {
	re := regexp.MustCompile(`\{(\w+)}`)
	return re.ReplaceAllString(pattern, `(?P<$1>[^/]+)`)
}

// patternExists verifica se o padrão já está registrado
func (h *HttpServer) patternExists(pattern string) bool {
	for _, route := range h.routes {
		if route.pattern.String() == pattern {
			return true
		}
	}
	return false
}

// Param retorna o valor de um parâmetro da URL
func (h *HttpServer) Param(req *http.Request, name string) string {
	params, ok := req.Context().Value("params").(map[string]string)
	if !ok {
		return ""
	}
	return params[name]
}

// Query retorna o valor de um parâmetro de consulta
func (h *HttpServer) Query(req *http.Request, name string) string {
	return req.URL.Query().Get(name)
}

// Implementação dos métodos da interface HandleMethods

func (h *HttpServer) HandleGet(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodGet, pattern, handlerFunc)
}

func (h *HttpServer) HandleHead(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodHead, pattern, handlerFunc)
}

func (h *HttpServer) HandlePost(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodPost, pattern, handlerFunc)
}

func (h *HttpServer) HandlePut(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodPut, pattern, handlerFunc)
}

func (h *HttpServer) HandlePatch(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodPatch, pattern, handlerFunc)
}

func (h *HttpServer) HandleDelete(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodDelete, pattern, handlerFunc)
}

func (h *HttpServer) HandleConnect(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodConnect, pattern, handlerFunc)
}

func (h *HttpServer) HandleOptions(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodOptions, pattern, handlerFunc)
}

func (h *HttpServer) HandleTrace(pattern string, handlerFunc http.HandlerFunc) error {
	return h.handleFunc(http.MethodTrace, pattern, handlerFunc)
}
