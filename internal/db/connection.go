package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	//conf := configs.GetDB()

	/*dsn := conf.User +
	":" + conf.Pass +
	"@tcp(" + conf.Host +
	":" + conf.Port +
	"/" + conf.DBName +
	"?charset=utf8mb4&parseTime=True&loc=Local"*/
	dsn := "user:password@tcp(host:port)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Falha ao conectar o banco de dados")
	}

	log.Println("Conexão ao Banco de Dados estabelecida...")
	return db

}
