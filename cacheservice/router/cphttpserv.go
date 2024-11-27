package router

/*import (
	"cacheservice/httpserv"
	"context"
	"net/http"
	"regexp"
)

type Router struct {
	routes []*httpserv.Route
}

func NewRouter() *Router {
	return &Router{}
}

// HandleFunc adiciona uma rota ao Router
func (r *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	regexPattern := regexp.MustCompile("^" + pattern + "$")
	r.routes = append(r.routes, &httpserv.Route{pattern: regexPattern, handler: handlerFunc})
}

// ServeHTTP implementa o manipulador padrão de HTTP para o Router
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if matches := route.pattern.FindStringSubmatch(req.URL.Path); matches != nil {
			params := make(map[string]string)
			for i, name := range route.pattern.SubexpNames() {
				if i != 0 && name != "" {
					params[name] = matches[i]
				}
			}
			// Adiciona os parâmetros de URL ao contexto da requisição
			ctx := context.WithValue(req.Context(), "params", params)
			req = req.WithContext(ctx)
			route.handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// Param retorna o valor de um parâmetro da URL
func Param(req *http.Request, name string) string {
	params, ok := req.Context().Value("params").(map[string]string)
	if !ok {
		return ""
	}
	return params[name]
}

// Query retorna o valor de um parâmetro de consulta
func Query(req *http.Request, name string) string {
	return req.URL.Query().Get(name)
}
*/
