package model

import (
	equipeModel "CCETEsportes/equipe/model"
	"errors"
	"time"
)

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

func (Partida) GetAll() ([]Partida, error) {
	var partidas []Partida
	preloads := resource.Preload("Equipe1").Preload("Equipe2")
	result := preloads.Find(&partidas).Order("Data DESC")
	if result.Error != nil {
		return nil, result.Error
	}

	return partidas, nil
}

func (p *Partida) Save() error {
	err := p.Validate()
	if err != nil {
		return err
	}
	result := resource.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
