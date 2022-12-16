package repository

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

type DAO interface {
}

type dao struct{}

var DB *sql.DB

func pgQl() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.RunWith(DB)
}

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("cannot read from config file")
	}
	fmt.Println(viper.Get("database"))
	host := viper.Get("database.host").(string)
	port := viper.Get("database.port").(int)
	user := viper.Get("database.user").(string)
	dbname := viper.Get("database.dbname").(string)
	password := viper.Get("database.password").(string)

	// Starting a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
