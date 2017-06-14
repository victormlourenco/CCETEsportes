package controller

import (
	"net/http"

	equipeApi "CCETEsportes/equipe"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

var r *render.Render

func Create(router *gin.RouterGroup) {
	equipe := router.Group("")
	{
		equipe.GET("/", getEquipes)
	}

	r = render.New(render.Options{
		IndentJSON: true,
		Directory:  "equipe/template",
		Extensions: []string{".html"},
	})
}

func getEquipes(c *gin.Context) {
	equipes, _ := equipeApi.ListarEquipes()
	r.HTML(c.Writer, http.StatusOK, "equipe", equipes)
}
