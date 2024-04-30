package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
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



// nilai default env
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}




func (server *Server) Initialize(appConfig *AppConfig, dbConfig *DBConfig) {
	fmt.Println("welcome to go online shop " + appConfig.AppName)
	server.initializeDB(*dbConfig)
	// server.Router = mux.NewRouter()
	server.InitializeRoutes()
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


		// looping migrate from interface model.registermodels
		for _, model := range RegisterModels(){
			err := server.DB.Debug().AutoMigrate(model.Model)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("database migration succesfully")

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

	server.Initialize(&appConfig, &dbConfig)
	server.Run(":" + appConfig.AppPort)

}


