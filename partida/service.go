package partida

import (
	"CCETEsportes/partida/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ListarPartidas : Chama a model para consultar as partidas do banco
func ListarPartidas() ([]model.Partida, error) {
	return model.Partida{}.GetAll()
}

// AdicionarPartida : Recebe um formulario e converte para um objeto do tipo Partida
func AdicionarPartida(c *gin.Context) error {
	partida := model.Partida{}

	equipe1, err := strconv.ParseUint(c.PostForm("equipe1"), 10, 64)
	if err != nil {
		return err
	}
	partida.CodEquipe1 = equipe1

	golsEquipe1, err := strconv.ParseUint(c.PostForm("equipe1Gols"), 10, 64)
	if err != nil {
		return err
	}
	partida.GolsEquipe1 = golsEquipe1

	equipe2, err := strconv.ParseUint(c.PostForm("equipe2"), 10, 64)
	if err != nil {
		return err
	}
	partida.CodEquipe2 = equipe2

	golsEquipe2, err := strconv.ParseUint(c.PostForm("equipe2Gols"), 10, 64)
	if err != nil {
		return err
	}
	partida.GolsEquipe2 = golsEquipe2

	partida.Local = c.PostForm("local")

	data := c.PostForm("data")
	t, err := time.Parse("2006-01-02T15:04", data)
	if err != nil {
		return err
	}
	partida.Data = t

	err = partida.Save()
	if err != nil {
		return err
	}

	return nil
}
