package controller

import (
	"net/http"

	equipeApi "CCETEsportes/equipe"
	partidaApi "CCETEsportes/partida"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

var r *render.Render

// Resposta : Definição do objeto esperado pelo template
type Resposta struct {
	Sucesso  bool
	Erro     bool
	Mensagem string
	Equipes  interface{}
}

// Create : Criação das rotas HTTP
func Create(router *gin.RouterGroup) {
	equipe := router.Group("/partidas")
	{
		equipe.GET("/", getPartidas)
		equipe.GET("/novo", adicionarPartida)
		equipe.POST("/novo", inserirPartida)
	}

	r = render.New(render.Options{
		IndentJSON: true,
		Directory:  "partida/template",
		Extensions: []string{".html"},
	})
}

func getPartidas(c *gin.Context) {
	partidas, _ := partidaApi.ListarPartidas()
	r.HTML(c.Writer, http.StatusOK, "partida", partidas)
}

func adicionarPartida(c *gin.Context) {
	equipes, err := equipeApi.ListarEquipes()
	if err != nil {
		r.HTML(c.Writer, http.StatusOK, "adicionar_partida", Resposta{Sucesso: false, Erro: true, Mensagem: err.Error(), Equipes: equipes})
	} else {
		r.HTML(c.Writer, http.StatusOK, "adicionar_partida", Resposta{Sucesso: false, Erro: false, Mensagem: "", Equipes: equipes})
	}
}

func inserirPartida(c *gin.Context) {
	equipes, _ := equipeApi.ListarEquipes()
	err := partidaApi.AdicionarPartida(c)
	if err != nil {
		println(err.Error())
		r.HTML(c.Writer, http.StatusBadRequest, "adicionar_partida", Resposta{Sucesso: false, Erro: true, Mensagem: err.Error(), Equipes: equipes})
	} else {
		r.HTML(c.Writer, http.StatusOK, "adicionar_partida", Resposta{Sucesso: true, Erro: false, Mensagem: "Partida inserida com suceso", Equipes: equipes})
	}
}
