package main

import (
	"CCETEsportes/config"
	equipeCtrl "CCETEsportes/equipe/controller"
	partidaCtrl "CCETEsportes/partida/controller"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	publicRoute := router.Group("/")
	equipeCtrl.Create(publicRoute)
	partidaCtrl.Create(publicRoute)
	s := &http.Server{
		Addr:    config.GetString("ServerAddress"),
		Handler: router,
	}

	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
