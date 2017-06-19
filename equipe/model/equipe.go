package model

import (
	"CCETEsportes/lib/database"
	"errors"

	"github.com/jinzhu/gorm"
)

// Equipe : Definição do objeto de tipo equipe.
type Equipe struct {
	CodEquipe  uint64 `gorm:"primary_key:true"`
	Nome       string
	Estado     string
	Pontuacao  uint64
	GolsPro    uint64
	GolsContra uint64
	Partidas   uint64
	Vitorias   uint64
	Derrotas   uint64
	Empates    uint64
	SaldoGols  int64 `gorm:"-"`
	Posicao    int64 `gorm:"-"`
	Tecnico    string
}

// GetAll : Consulta todas as partidas presentes no banco ordenando por colocação no campeonato
func (Equipe) GetAll() ([]Equipe, error) {
	var equipes []Equipe
	result := database.Get().Raw("SELECT cod_equipe, nome, pontuacao, partidas, vitorias, empates, derrotas, gols_pro, gols_contra, gols_pro - gols_contra AS saldo_gols, row_number() OVER (ORDER BY pontuacao DESC, gols_pro - gols_contra DESC) as posicao FROM equipes").Order("posicao").Find(&equipes)
	if result.Error != nil {
		return nil, result.Error
	}

	return equipes, nil
}

// Get : Consulta uma determinada equipe
func (e *Equipe) Get(tx *gorm.DB) error {
	get := tx.Find(&e)
	if get.Error != nil {
		return get.Error
	}
	if get.RowsAffected <= 0 {
		return errors.New("Equipe não encontrada")
	}

	return nil
}

// Save : Atualiza equipe no banco
func (e *Equipe) Save(tx *gorm.DB) error {
	result := tx.Model(Equipe{}).Save(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
