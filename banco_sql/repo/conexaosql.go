package repo

import (
	"database/sql"

	/*
		github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
		_ "github.com/go-sql-driver/mysql"
	*/
	_ "github.com/mattn/go-oci8"
)

//Db armazena a conexão com o banco de dados
var Db *sql.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	//Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/cursodego")
	Db, err = sql.Open("oci8", "ricardogs/ricardogs@10.113.0.227:1536/db2.uninove.br")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

//GetDB get database connection
func GetDB() (db *sql.DB, err error) {
	err = nil
	if Db != nil {
		db = Db
		return
	}
	//Db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/cursodego")
	Db, err = sql.Open("oci8", "ricardogs/ricardogs@10.113.0.227:1536/db2.uninove.br")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	db = Db
	return
}
