package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// データベースとの接続処理を記述している
func Connect() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	//.envファイル内の情報を元にしてdbとの接続を行う処理
	db, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASSWORD")+"@tcp(localhost:3306)/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err.Error())
	}
	//戻り値としてポインタが渡される。
	return db
}
