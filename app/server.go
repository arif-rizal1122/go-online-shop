package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arif-rizal1122/go-online-shop/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)





type Server struct {
	  DB             *gorm.DB
	  Router		 *mux.Router
}



type AppConfig struct {
	AppName    string
	AppEnv	   string
	AppPort	   string
}



type DBConfig struct {
	DBHost         string
	DBUser		   string
	DBPassword     string
	DBName		   string
	DBPort		   string
	DBDriver 	   string
}


func (server *Server) Initialize(appConfig *AppConfig, dbConfig *DBConfig) {
	fmt.Println("welcome to go online shop " + appConfig.AppName)
	server.initializeRoutes()
}


func (server *Server) initializeDB(dbConfig DBConfig)  {
		var err error

		if (dbConfig.DBDriver == "mysql") {

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
	fmt.Println("listening to on port",  addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}



func Run()  {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "go-online-shop")
	appConfig.AppEnv  = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "goToko")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")


	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(dbConfig)
	} else {
		server.Initialize(&appConfig, &dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}


// nilai default env
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}


func (server *Server) dbMigrate() {
	// looping migrate from interface model.registermodels
	for _, model := range RegisterModels(){
	  err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
			}
	}
	fmt.Println("database migration succesfully")
}


func (server *Server )initCommands(dbconfig DBConfig)  {
	// seeders.DBSeed(server.DB)
	server.initializeDB(dbconfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func (c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func (c *cli.Context) error {
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