/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 9:48:39 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package manager

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rizkyfazri23/dripay/config"
)

type InfraManager interface {
	DbConn() *sql.DB
}

type infraManager struct {
	db *sql.DB
	cfg config.AppConfig
}

func (i *infraManager) initDb() {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name, i.cfg.SslMode)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer func () {
		if err := recover(); err != nil {
			log.Println("Application filed to run", err)
			db.Close()
		}
	}()

	i.db = db
	fmt.Println("DB Connected")
}

func (i *infraManager) DbConn() *sql.DB {
	return i.db
}

func NewInfraManager(cfg config.AppConfig) InfraManager {
	infra := infraManager {
		cfg: cfg,
	}
	infra.initDb()
	return &infra
}
