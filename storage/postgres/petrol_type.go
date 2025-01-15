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

type petrol_typeRepo struct {
	db *pgxpool.Pool
}

func NewPetrolTypeRepo(db *pgxpool.Pool) *petrol_typeRepo {
	return &petrol_typeRepo{
		db: db,
	}
}

func (r *petrol_typeRepo) Create(ctx context.Context, req *models.CreatePetrolType) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO petrol_types(id, name, updated_at)
		VALUES ($1, $2, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *petrol_typeRepo) GetByID(ctx context.Context, req *models.PetrolTypePrimaryKey) (*models.PetrolType, error) {

	var whereField = "id"
	var (
		query string

		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM petrol_types
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.PetrolType{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *petrol_typeRepo) GetList(ctx context.Context, req *models.PetrolTypeGetListRequest) (*models.PetrolTypeGetListResponse, error) {

	var (
		resp   = &models.PetrolTypeGetListResponse{}
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
			created_at,
			updated_at
		FROM petrol_types
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
			id        sql.NullString
			name      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.PetrolTypes = append(resp.PetrolTypes, &models.PetrolType{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *petrol_typeRepo) Update(ctx context.Context, req *models.UpdatePetrolType) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			petrol_types
		SET
		id = :id,
		name = :name,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":   req.Id,
		"name": req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *petrol_typeRepo) Delete(ctx context.Context, req *models.PetrolTypePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM petrol_types WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
