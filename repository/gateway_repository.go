package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rizkyfazri23/dripay/model/entity"
)

type GatewayRepo interface {
	FindAll() ([]entity.Gateway, error)
	FindOne(id int) (entity.Gateway, error)
	Create(newGateway *entity.Gateway) (entity.Gateway, error)
	Update(gateway *entity.Gateway) (entity.Gateway, error)
	Delete(id int) error
}

type gatewayRepo struct {
	db *sql.DB
}

func (r *gatewayRepo) FindAll() ([]entity.Gateway, error) {
	var gateways []entity.Gateway

	query := "SELECT gateway_id, gateway_name, status FROM m_gateway ORDER BY gateway_id"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var gateway entity.Gateway
		if err := rows.Scan(&gateway.Gateway_Id, &gateway.Gateway_Name, &gateway.Status); err != nil {
			return nil, err
		}
		gateways = append(gateways, gateway)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return gateways, nil
}

func (r *gatewayRepo) FindOne(id int) (entity.Gateway, error) {
	var gatewayInDb entity.Gateway

	query := "SELECT gateway_id, gateway_name, status FROM m_gateway WHERE gateway_id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&gatewayInDb.Gateway_Id, &gatewayInDb.Gateway_Name, &gatewayInDb.Status)

	if err == sql.ErrNoRows {
		return entity.Gateway{}, fmt.Errorf("gateway with id %d not found", id)
	} else if err != nil {
		return entity.Gateway{}, err
	}

	return gatewayInDb, nil
}

func (r *gatewayRepo) Create(newGateway *entity.Gateway) (entity.Gateway, error) {
	query := "INSERT INTO m_gateway (gateway_name) VALUES ($1) RETURNING gateway_id"
	var gatewayID int
	err := r.db.QueryRow(query, newGateway.Gateway_Name).Scan(&gatewayID)
	if err != nil {
		log.Println(err)
		return entity.Gateway{}, err
	}

	newGateway.Gateway_Id = gatewayID

	return *newGateway, nil
}

func (r *gatewayRepo) Update(gateway *entity.Gateway) (entity.Gateway, error) {
	var status int
	if gateway.Status {
		status = 1
	} else {
		status = 0
	}
	query := "UPDATE m_gateway SET gateway_name = $1, status = $2 WHERE gateway_id = $3 RETURNING gateway_id, gateway_name, status"
	row := r.db.QueryRow(query, gateway.Gateway_Name, status, gateway.Gateway_Id)

	var updatedGateway entity.Gateway
	err := row.Scan(&updatedGateway.Gateway_Id, &updatedGateway.Gateway_Name, &updatedGateway.Status)
	if err != nil {
		log.Println(err)
		return entity.Gateway{}, err
	}

	return updatedGateway, nil
}

func (r *gatewayRepo) Delete(id int) error {
	query := "DELETE FROM m_gateway WHERE gateway_id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("gateway with id %d not found", id)
	}
	return nil
}

func NewGatewayRepository(db *sql.DB) GatewayRepo {
	repo := new(gatewayRepo)
	repo.db = db
	return repo
}
