package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotapaginaPrincipal = []Rota{
	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: true,
	},
}
