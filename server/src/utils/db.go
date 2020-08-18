package utils

import(
	"github.com/jinzhu/gorm"
	"os"
	"log"
	_ "github.com/lib/pq"
)

var DSN = ""
type Database struct{
	Conn *gorm.DB
}

func Connect() Database {
	if DSN == ""{
		dsn, provided := os.LookupEnv("DSN")
		
		if !provided {
			log.Fatalf("Error loading password DataBase.")
		}
		DSN = dsn
	}
	db, err := gorm.Open("postgres", DSN)
	if err != nil{
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}
	return Database{Conn: db}
}

func (db Database) Close(){
	if db.Conn != nil{
		db.Conn.Close()
	}
}