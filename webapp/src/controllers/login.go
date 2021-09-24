package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// Carrega a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
