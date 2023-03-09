package manager

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyfazri23/dripay/config"
)

type InfraManager interface {
	DbConn() *sqlx.DB
}

type infraManager struct {
	db  *sqlx.DB
	cfg config.AppConfig
}

func (i *infraManager) initDb() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", i.cfg.DBHost, i.cfg.DBPort, i.cfg.DBUser, i.cfg.DBPass, i.cfg.DBName, i.cfg.DBSSLMode)
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application filed to run", err)
			db.Close()
		}
	}()

	i.db = db
	fmt.Println("DB Connected")
}

func (i *infraManager) DbConn() *sqlx.DB {
	return i.db
}

func NewInfraManager(cfg config.AppConfig) InfraManager {
	infra := infraManager{
		cfg: cfg,
	}
	infra.initDb()
	return &infra
}
