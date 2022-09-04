package app

import (
	"fmt"
	"github.com/EmirShimshir/crud-books/internal/config"
	"github.com/EmirShimshir/crud-books/internal/repository/psql"
	"github.com/EmirShimshir/crud-books/internal/service"
	"github.com/EmirShimshir/crud-books/internal/transport/rest"
	"github.com/EmirShimshir/crud-books/pkg/database"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func Run(configDir, configFile string) {
	log.Info("application startup...")
	log.Info("logger initialized")

	// init config
	cfg, err := config.New(configDir, configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "app.Run()",
			"problem": "can't initialize config",
		}).Fatal(err.Error())
	}
	log.Info("config created")

	// init db
	db, err := database.NewPosgresqlConnection(
		database.ConnectionSettings{
			Host:     cfg.DB.Host,
			Port:     cfg.DB.Port,
			Username: cfg.DB.Username,
			DBName:   cfg.DB.Name,
			SSLMode:  cfg.DB.SSLMode,
			Password: cfg.DB.Password,
		})
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "app.Run()",
			"problem": "can't connect to db",
		}).Fatal(err.Error())
	}
	defer db.Close()
	log.Info("db is connected")

	// init deps
	repositories := psql.NewRepositories(db)
	services := service.NewServices(repositories)
	handler := rest.NewHandler(services)
	log.Info("repositorys, services  and handler initialized")

	// init server
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      handler.InitRouter(),
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
	}
	log.Info("SERVER STARTED")

	// TODO graceful shutdown
	err = server.ListenAndServe()
	if err != nil {
		log.WithFields(log.Fields{
			"from":    "app.Run()",
			"problem": "shutdown server",
		}).Fatal(err.Error())
	}
}
