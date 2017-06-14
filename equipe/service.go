package equipe

import (
	"CCETEsportes/equipe/model"
)

func ListarEquipes() ([]model.Equipe, error) {
	return model.Equipe{}.GetAll()
}
