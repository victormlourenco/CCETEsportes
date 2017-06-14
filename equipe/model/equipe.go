package model

type Equipe struct {
}

// set User's table name to be `profiles`
func (User) TableName() string {
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

func (Equipe) Update() error {

}
