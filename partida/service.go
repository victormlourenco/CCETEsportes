package partida

import (
	equipeModel "CCETEsportes/equipe/model"
	"CCETEsportes/lib/database"
	"CCETEsportes/partida/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ListarPartidas : Chama a model para consultar as partidas do banco
func ListarPartidas() ([]model.Partida, error) {
	return model.Partida{}.GetAll()
}

// AdicionarPartida : Recebe um formulario e converte para um objeto do tipo Partida e atualiza a tabela do Campeonato
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

	tx := database.Get().Begin()

	err = partida.Save(tx)
	if err != nil {
		return err
	}

	eqp1 := equipeModel.Equipe{CodEquipe: partida.CodEquipe1}
	err = eqp1.Get()
	if err != nil {
		tx.Rollback()
		return err
	}

	eqp2 := equipeModel.Equipe{CodEquipe: partida.CodEquipe2}
	err = eqp2.Get()
	if err != nil {
		tx.Rollback()
		return err
	}

	eqp1.GolsPro = eqp1.GolsPro + partida.GolsEquipe1
	eqp1.GolsContra = eqp1.GolsContra + partida.GolsEquipe2
	eqp2.GolsPro = eqp2.GolsPro + partida.GolsEquipe2
	eqp2.GolsContra = eqp2.GolsContra + partida.GolsEquipe1

	if partida.GolsEquipe1 > partida.GolsEquipe2 {
		eqp1.Vitorias++
		eqp2.Derrotas++
		eqp1.Pontuacao = eqp1.Pontuacao + 3
	} else if partida.GolsEquipe2 > partida.GolsEquipe1 {
		eqp2.Vitorias++
		eqp1.Derrotas++
		eqp2.Pontuacao = eqp2.Pontuacao + 3
	} else {
		eqp2.Empates++
		eqp1.Empates++
		eqp1.Pontuacao = eqp1.Pontuacao + 1
		eqp2.Pontuacao = eqp2.Pontuacao + 1
	}

	err = eqp1.Save(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = eqp2.Save(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
