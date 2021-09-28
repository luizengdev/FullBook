package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// Insere os templates html ba variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))

}

// Renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
