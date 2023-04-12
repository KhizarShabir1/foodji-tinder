package main

import (
	"log"

	"github.com/KhizarShabir1/foodji-tinder/database"
	"github.com/KhizarShabir1/foodji-tinder/http"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//load configuration

	configFileAddr := "../../internal/config/config.yaml"

	if err := loadConfig(configFileAddr); err != nil {
		log.Fatalf("failed to load configuration: %s", err)
	}

	// Initialize database
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	} else {
		log.Println("connected database")
	}

	defer db.Close()

	dbp := database.NewProvider(db)

	server := http.NewServer(dbp)

	if err := server.Start(); err != nil {
		log.Fatal("Failed to run HTTP server", zap.Error(err))
	}
	log.Println("Exited")

}

func loadConfig(configFile string) error {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.AutomaticEnv()
	return nil
}
