package model

import (
	libDB "CCETEsportes/lib/database"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //gorm dependency
)

var resource *gorm.DB

func init() {
	resource, _ = gorm.Open(libDB.ConnProvider, libDB.ConnString)
	resource.LogMode(true)
	/* Método para não salvar associações das estruturas quando for realizar qualquer alteração no banco
	* Teve que ser implementado pois, devido a um bug no gorm, os campos de uma estrutura que faziam
	* parte de uma associação eram atualizados ao realizar uma operação de inserção ou atualização dessa
	* estrutura no banco. Devido a isso as associações devem ser inseridas manualmente. */
	resource.Callback().Create().After("gorm:begin_transaction").Register("ignore_associations", ignoreAssociations)
	resource.Callback().Update().After("gorm:begin_transaction").Register("ignore_associations", ignoreAssociations)
}

func ignoreAssociations(s *gorm.Scope) {
	s.Set("gorm:save_associations", false)
}
