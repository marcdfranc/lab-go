// logformatter.go
package logger

import (
	"strings"
	"text/template"
)

type LogFormatter struct {
	template *template.Template
}

func NewLogFormatter(templateStr string) *LogFormatter {
	tmpl := template.Must(template.New("log").Parse(templateStr))
	return &LogFormatter{template: tmpl}
}

func (lf *LogFormatter) Format(data map[string]interface{}) string {
	var logMessage strings.Builder
	lf.template.Execute(&logMessage, data)
	return logMessage.String()
}
