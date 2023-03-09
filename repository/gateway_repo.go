package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rizkyfazri23/dripay/model/entity"
)

type gatewayRepo struct {
	db *sqlx.DB
}

type GatewayRepo interface {
	CreateGateway(gateway *entity.Gateway) (*entity.Gateway, error)
	ReadGateway() ([]*entity.Gateway, error)
	UpdateGateway(gateway *entity.Gateway) (*entity.Gateway, error)
	DeleteGateway(id int) error
}

func NewGatewayRepo(db *sqlx.DB) GatewayRepo {
	return &gatewayRepo{db: db}
}

func (r *gatewayRepo) CreateGateway(gateway *entity.Gateway) (*entity.Gateway, error) {
	query := "INSERT INTO gateway (gateway_name, type, status) VALUES ($1, $2, $3)"
	_, err := r.db.Query(query, gateway.Gateway_Name, gateway.Type, gateway.Status)
	if err != nil {
		return nil, err
	}
	return gateway, nil
}

func (r *gatewayRepo) ReadGateway() ([]*entity.Gateway, error) {
	var gateways []*entity.Gateway
	query := "SELECT * FROM gateway"
	err := r.db.Select(&gateways, query)
	if err != nil {
		return nil, err
	}
	return gateways, nil
}

func (r *gatewayRepo) UpdateGateway(gateway *entity.Gateway) (*entity.Gateway, error) {
	var data entity.Gateway
	query := "UPDATE gateway SET gateway_name = $1, type = $2, status = $3 WHERE gateway_id = $4"
	err := r.db.QueryRow(query, gateway.Gateway_Name, gateway.Type, gateway.Status, gateway.Gateway_Id).Scan(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *gatewayRepo) DeleteGateway(id int) error {
	query := "DELETE FROM gateway WHERE gateway_id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
