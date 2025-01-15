package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	Admin() AdminRepoI
	PetrolType() PetrolTypeRepoI
	Region() RegionRepoI
	Department() DepartmentRepoI
	Car() CarRepoI
	PetrolHistory() PetrolHistoryRepoI
}

type AdminRepoI interface {
	Create(context.Context, *models.CreateAdmin) (string, error)
	GetByID(context.Context, *models.AdminPrimaryKey) (*models.Admin, error)
	GetList(context.Context, *models.AdminGetListRequest) (*models.AdminGetListResponse, error)
	Update(context.Context, *models.UpdateAdmin) (int64, error)
	Delete(context.Context, *models.AdminPrimaryKey) error
}

type PetrolTypeRepoI interface {
	Create(context.Context, *models.CreatePetrolType) (string, error)
	GetByID(context.Context, *models.PetrolTypePrimaryKey) (*models.PetrolType, error)
	GetList(context.Context, *models.PetrolTypeGetListRequest) (*models.PetrolTypeGetListResponse, error)
	Update(context.Context, *models.UpdatePetrolType) (int64, error)
	Delete(context.Context, *models.PetrolTypePrimaryKey) error
}

type RegionRepoI interface {
	Create(context.Context, *models.CreateRegion) (string, error)
	GetByID(context.Context, *models.RegionPrimaryKey) (*models.Region, error)
	GetList(context.Context, *models.RegionGetListRequest) (*models.RegionGetListResponse, error)
	Update(context.Context, *models.UpdateRegion) (int64, error)
	Delete(context.Context, *models.RegionPrimaryKey) error
}

type DepartmentRepoI interface {
	Create(context.Context, *models.CreateDepartment) (string, error)
	GetByID(context.Context, *models.DepartmentPrimaryKey) (*models.Department, error)
	GetList(context.Context, *models.DepartmentGetListRequest) (*models.DepartmentGetListResponse, error)
	Update(context.Context, *models.UpdateDepartment) (int64, error)
	Delete(context.Context, *models.DepartmentPrimaryKey) error
}

type CarRepoI interface {
	Create(context.Context, *models.CreateCar) (string, error)
	GetByID(context.Context, *models.CarPrimaryKey) (*models.Car, error)
	GetList(context.Context, *models.CarGetListRequest) (*models.CarGetListResponse, error)
	Update(context.Context, *models.UpdateCar) (int64, error)
	Delete(context.Context, *models.CarPrimaryKey) error
}
type PetrolHistoryRepoI interface {
	Create(context.Context, *models.CreatePetrolHistory) (string, error)
	GetByID(context.Context, *models.PetrolHistoryPrimaryKey) (*models.PetrolHistory, error)
	GetList(context.Context, *models.PetrolHistoryGetListRequest) (*models.PetrolHistoryGetListResponse, error)
	Update(context.Context, *models.UpdatePetrolHistory) (int64, error)
	Delete(context.Context, *models.PetrolHistoryPrimaryKey) error
}
