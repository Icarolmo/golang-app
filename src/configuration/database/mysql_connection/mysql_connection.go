package mysql_connection

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Função para realizar uma conexão com banco de dados MYSQL
func InitConnection() *sql.DB{

	if err := godotenv.Load(); err != nil {
		return nil
	}
	db, err := sql.Open("mysql",  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_IP"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB_NAME")))
	
	if err != nil {
		fmt.Println("Erro ao abrir conexão com MySQL")
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

// Fecha conexão com banco de dados
func CloseConnection(c *sql.DB){
	c.Close()
}