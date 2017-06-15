package equipe

import (
	"CCETEsportes/equipe/model"
)

// ListarEquipes : Recebe da model uma lista de equipes
func ListarEquipes() ([]model.Equipe, error) {
	return model.Equipe{}.GetAll()
}
