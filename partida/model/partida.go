package model

import (
	equipeModel "CCETEsportes/equipe/model"
	"CCETEsportes/lib/database"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Partida : Definição do objeto de tipo partida.
type Partida struct {
	CodPartida  uint64 `gorm:"primary_key:true"`
	Local       string
	CodEquipe1  uint64
	CodEquipe2  uint64
	GolsEquipe1 uint64
	GolsEquipe2 uint64
	Data        time.Time
	Equipe1     equipeModel.Equipe `gorm:"ForeignKey:CodEquipe1"`
	Equipe2     equipeModel.Equipe `gorm:"ForeignKey:CodEquipe2"`
}

// Validate : Realiza a validação dos campos do objeto do tipo partida antes da sua inserção do banco
func (p *Partida) Validate() error {
	if p == nil {
		return errors.New("Partida inválida")
	}
	if p.Local == "" {
		return errors.New("O local deve ser preenchido")
	}
	if p.CodEquipe1 == 0 || p.CodEquipe2 == 0 {
		return errors.New("Selecione duas equipes")
	}
	if p.Data.IsZero() {
		return errors.New("A data da partida deve ser preenchida")
	}

	if p.CodEquipe1 == p.CodEquipe2 {
		return errors.New("Selecione duas equipes diferentes")
	}
	return nil
}

// GetAll : Consulta todas as partidas presentes no banco ordenando por data
func (Partida) GetAll() ([]Partida, error) {
	var partidas []Partida
	preloads := database.Get().Preload("Equipe1").Preload("Equipe2")
	result := preloads.Order("Data DESC").Find(&partidas)
	if result.Error != nil {
		return nil, result.Error
	}

	return partidas, nil
}

// Save : Armazena partida no banco
func (p *Partida) Save(tx *gorm.DB) error {
	err := p.Validate()
	if err != nil {
		return err
	}

	result := tx.Create(&p)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return nil
}
