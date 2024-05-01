package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arif-rizal1122/go-online-shop/app/models"
	"github.com/arif-rizal1122/go-online-shop/database/seeders"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func (server *Server) Initialize(appConfig *AppConfig, dbConfig *DBConfig) {
	fmt.Println("welcome to go online shop " + appConfig.AppName)
	server.initializeDB(*dbConfig)
	server.initializeRoutes()
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error

	if dbConfig.DBDriver == "mysql" {

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed on connection database server mysql")
		}

	} else {

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed on connection database server postgreSQL")
		}
	}

}

func (server *Server) Run(addr string) {
	fmt.Println("listening to on port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) dbMigrate() {
	// looping migrate from interface model.registermodels
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("database migration succesfully")
}





func (server *Server) InitCommands(dbconfig DBConfig) {
	// seeders.DBSeed(server.DB)
	server.initializeDB(dbconfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
