// main.go
package main

import (
	"cacheservice/logger"
	"log"
	"net/http"
	"os"
)

func main() {
	// Template de log corrigido e dinâmico
	templateStr := `[{{.Name}}] Level: {{.Level }} Time: {{.Time }} Msg: {{.Message }}
    {{range $key, $value := .ExtraData}} {{$key}}: {{ $value}}
    {{end}}
    `

	formatter := logger.NewLogFormatter(templateStr)
	logInstance := logger.NewLogger(os.Stdout, "main", formatter)

	// Uso do Logger na aplicação
	logInstance.LogInfo(map[string]interface{}{
		"Message": "Iniciando servidor...",
		"ExtraData": map[string]interface{}{
			"SessionID": "abc123",
			"IP":        "192.168.1.1",
		},
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logInstance.LogInfo(map[string]interface{}{
			"Message": "Recebida requisição para /",
			"ExtraData": map[string]interface{}{
				"Request": r.URL.Path,
				"Method":  r.Method,
			},
		})
		w.Write([]byte("Hello, World!"))
	})

	// Exemplos de Log com diferentes níveis
	logInstance.LogDebug(map[string]interface{}{
		"Message": "Debugando aplicação",
		"ExtraData": map[string]interface{}{
			"DebugData": "Detalhes de depuração",
		},
	})
	logInstance.LogWarning(map[string]interface{}{
		"Message": "Aviso: algo pode estar errado",
		"ExtraData": map[string]interface{}{
			"WarningCode": 1001,
		},
	})
	logInstance.LogError(map[string]interface{}{
		"Message": "Erro encontrado",
		"ExtraData": map[string]interface{}{
			"Error": "could not process data",
		},
	})
	logInstance.LogFatal(map[string]interface{}{
		"Message": "Erro fatal: o sistema irá parar",
		"ExtraData": map[string]interface{}{
			"Error": "Critical failure",
		},
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
