package model

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
	SaldoGols  int64
	Posicao    int64
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
	result := resource.Raw("SELECT cod_equipe, nome, pontuacao, partidas, vitorias, empates, derrotas, gols_pro, gols_contra, gols_pro - gols_contra AS saldo_gols, row_number() OVER (ORDER BY pontuacao DESC, gols_pro - gols_contra DESC) as posicao FROM equipes").Order("posicao").Find(&equipes)
	if result.Error != nil {
		return nil, result.Error
	}

	return equipes, nil
}
