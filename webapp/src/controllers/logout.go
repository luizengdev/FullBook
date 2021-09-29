package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Remove os dados de autenticação salvos no brouser do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
