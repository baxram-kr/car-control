package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type store struct {
	db             *pgxpool.Pool
	admin          *adminRepo
	petrol_type    *petrol_typeRepo
	region         *regionRepo
	department     *departmentRepo
	car            *carRepo
	petrol_history *petrol_historyRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Admin() storage.AdminRepoI {

	if s.admin == nil {
		s.admin = NewAdminRepo(s.db)
	}

	return s.admin
}

func (s *store) PetrolType() storage.PetrolTypeRepoI {

	if s.petrol_type == nil {
		s.petrol_type = NewPetrolTypeRepo(s.db)
	}

	return s.petrol_type
}

func (s *store) Region() storage.RegionRepoI {

	if s.region == nil {
		s.region = NewRegionRepo(s.db)
	}

	return s.region
}

func (s *store) Department() storage.DepartmentRepoI {

	if s.department == nil {
		s.department = NewDepartmentRepo(s.db)
	}

	return s.department
}

func (s *store) Car() storage.CarRepoI {

	if s.car == nil {
		s.car = NewCarRepo(s.db)
	}

	return s.car
}

func (s *store) PetrolHistory() storage.PetrolHistoryRepoI {

	if s.petrol_history == nil {
		s.petrol_history = NewPetrolHistoryRepo(s.db)
	}

	return s.petrol_history
}
