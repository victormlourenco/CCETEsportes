package database

import (
	"fmt"

	"CCETEsportes/lib/database/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //database driver
)

var resource *gorm.DB

func init() {
	ConnProvider := config.GetString("DBProvider")

	dbName := config.GetString("DBName")
	dbUser := config.GetString("DBUser")
	dbPass := config.GetString("DBPass")
	dbHost := config.GetString("DBHost")
	dbPort := config.GetString("DBPort")
	dbSslMode := config.GetString("DBSslMode")
	ConnString := fmt.Sprintf("dbname='%s' user='%s' password='%s' host='%s' port='%s' sslmode='%s'",
		dbName, dbUser, dbPass, dbHost, dbPort, dbSslMode)

	resource, _ = gorm.Open(ConnProvider, ConnString)
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

// Get : Retorna a conexão com o banco
func Get() *gorm.DB {
	return resource
}
