package database

import (
	"context"
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresUtil struct {
	Dsn    string
	Dbname string
	Ctx    context.Context
}

func NewPostgresUtil() *PostgresUtil {
	return &PostgresUtil{
		Dsn:    "host=localhost user=postgres password=jeruk.123 dbname=new_portal_data_kalsel port=5432 sslmode=disable TimeZone=Asia/Jakarta",
		Dbname: "new_portal_data_kalsel",
		Ctx:    context.Background(),
	}
}

func (d *PostgresUtil) Connect() (client *gorm.DB, err error) {
	fmt.Println(aurora.Green(d.Dsn))
	db, err := gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

var DB *gorm.DB

func Connect() {
	util := NewPostgresUtil()
	var err error
	DB, err = util.Connect()
	if err != nil {
		log.Println("Failed to connect to database:", err)
	}
}
