package model

type Equipe struct {
	CodEquipe  uint64
	Nome       string
	Estado     string
	Pontuacao  uint64
	GolsPro    uint64
	GolsContra uint64
	Partidas   uint64
	Vitorias   uint64
	Derrotas   uint64
	Empates    uint64
	Tecnico    string
}

// set User's table name to be `profiles`
func (Equipe) TableName() string {
	return "equipes"
}

func (Equipe) migrateTables() {
}

func (Equipe) migrateConstraints() {
}

func (Equipe) GetAll() ([]Equipe, error) {
	var equipes []Equipe
	result := resource.Find(&equipes)
	if result.Error != nil {
		return nil, result.Error
	}

	return equipes, nil
}
