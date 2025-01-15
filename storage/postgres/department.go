package postgres

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"app/api/models"
	"app/pkg/helper"
)

type departmentRepo struct {
	db *pgxpool.Pool
}

func NewDepartmentRepo(db *pgxpool.Pool) *departmentRepo {
	return &departmentRepo{
		db: db,
	}
}

func (r *departmentRepo) Create(ctx context.Context, req *models.CreateDepartment) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO departments(id, name, email,password,region,address,phone_number,updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Email,
		req.Password,
		req.Region,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *departmentRepo) GetByID(ctx context.Context, req *models.DepartmentPrimaryKey) (*models.Department, error) {

	var whereField = "id"
	if len(req.Email) > 0 {
		whereField = "email"
		req.Id = req.Email
	}
	var (
		query string

		id           sql.NullString
		name         sql.NullString
		email        sql.NullString
		password     sql.NullString
		region       sql.NullString
		address      sql.NullString
		phone_number sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			email,
			password,
			region,
			address,
			phone_number,
			created_at,
			updated_at
		FROM departments
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&email,
		&password,
		&region,
		&address,
		&phone_number,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Department{
		Id:          id.String,
		Name:        name.String,
		Email:       email.String,
		Password:    password.String,
		Region:      region.String,
		Address:     address.String,
		PhoneNumber: phone_number.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *departmentRepo) GetList(ctx context.Context, req *models.DepartmentGetListRequest) (*models.DepartmentGetListResponse, error) {

	var (
		resp   = &models.DepartmentGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			email,
			password,
			region,
			address,
			phone_number,
			created_at,
			updated_at
		FROM departments
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id           sql.NullString
			name         sql.NullString
			email        sql.NullString
			password     sql.NullString
			region       sql.NullString
			address      sql.NullString
			phone_number sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&email,
			&password,
			&region,
			&address,
			&phone_number,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Departments = append(resp.Departments, &models.Department{
			Id:          id.String,
			Name:        name.String,
			Email:       email.String,
			Password:    password.String,
			Region:      region.String,
			Address:     address.String,
			PhoneNumber: phone_number.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *departmentRepo) Update(ctx context.Context, req *models.UpdateDepartment) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			departments
		SET
		id = :id,
		name = :name,
		email = :email,
		password = :password,
		region = :region,
		address = :address,
		phone_number = :phone_number,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"email":        req.Email,
		"password":     req.Password,
		"region":       req.Region,
		"address":      req.Address,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *departmentRepo) Delete(ctx context.Context, req *models.DepartmentPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM departments WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
