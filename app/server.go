package app

import (
	"flag"
	"log"
	"os"

	"github.com/arif-rizal1122/go-online-shop/app/controllers"
	"github.com/joho/godotenv"
)

// nilai default env 
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "go-online-shop")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "goToko")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	// Parse argumen dari baris perintah
	flag.Parse()
	// Ambil argumen pertama dari baris perintah
	arg := flag.Arg(0)
	// Periksa apakah argumen tidak kosong
	if arg != "" {
		// Jika argumen tidak kosong, inisialisasi perintah berdasarkan argumen
		server.InitCommands(dbConfig)
	} else {
		// Jika argumen kosong, inisialisasi aplikasi secara normal
		// dan jalankan server menggunakan port yang ditentukan dalam appConfig
		server.Initialize(&appConfig, &dbConfig)
		server.Run(":" + appConfig.AppPort)
	}


}
