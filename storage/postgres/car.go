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

type carRepo struct {
	db *pgxpool.Pool
}

func NewCarRepo(db *pgxpool.Pool) *carRepo {
	return &carRepo{
		db: db,
	}
}

func (r *carRepo) Create(ctx context.Context, req *models.CreateCar) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO cars(id, name,model,state_number,year,ban_status,tech_condition,defect,address,department_id,petrol_name,petrol, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Model,
		req.StateNumber,
		req.Year,
		req.BanStatus,
		req.TechCondition,
		req.Defect,
		req.Address,
		req.DepartmentID,
		req.PetrolType,
		req.Petrol,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *carRepo) GetByID(ctx context.Context, req *models.CarPrimaryKey) (*models.Car, error) {

	var whereField = "id"
	var (
		query string

		id             sql.NullString
		name           sql.NullString
		model          sql.NullString
		state_number   sql.NullString
		year           sql.NullString
		ban_status     sql.NullString
		tech_condition sql.NullString
		defect         sql.NullString
		address        sql.NullString
		department_id  sql.NullString
		petrol_name    sql.NullString
		petrol         sql.NullFloat64
		createdAt      sql.NullString
		updatedAt      sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			model,
			state_number,
			year,
			ban_status,
			tech_condition,
			defect,
			address,
			department_id,
			petrol_name,
			petrol,
			created_at,
			updated_at
		FROM cars
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&model,
		&state_number,
		&year,
		&ban_status,
		&tech_condition,
		&defect,
		&address,
		&department_id,
		&petrol_name,
		&petrol,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Car{
		Id:            id.String,
		Name:          name.String,
		Model:         model.String,
		StateNumber:   state_number.String,
		Year:          year.String,
		BanStatus:     ban_status.String,
		TechCondition: tech_condition.String,
		Defect:        defect.String,
		Address:       address.String,
		DepartmentID:  department_id.String,
		PetrolType:    petrol_name.String,
		Petrol:        petrol.Float64,
		CreatedAt:     createdAt.String,
		UpdatedAt:     updatedAt.String,
	}, nil
}

func (r *carRepo) GetList(ctx context.Context, req *models.CarGetListRequest) (*models.CarGetListResponse, error) {

	var (
		resp   = &models.CarGetListResponse{}
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
			model,
			state_number,
			year,
			ban_status,
			tech_condition,
			defect,
			address,
			department_id,
			petrol_name,
			petrol,
			created_at,
			updated_at
		FROM cars
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
			id             sql.NullString
			name           sql.NullString
			model          sql.NullString
			state_number   sql.NullString
			year           sql.NullString
			ban_status     sql.NullString
			tech_condition sql.NullString
			defect         sql.NullString
			address        sql.NullString
			department_id  sql.NullString
			petrol_name    sql.NullString
			petrol         sql.NullFloat64
			createdAt      sql.NullString
			updatedAt      sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&model,
			&state_number,
			&year,
			&ban_status,
			&tech_condition,
			&defect,
			&address,
			&department_id,
			&petrol_name,
			&petrol,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Cars = append(resp.Cars, &models.Car{
			Id:            id.String,
			Name:          name.String,
			Model:         model.String,
			StateNumber:   state_number.String,
			Year:          year.String,
			BanStatus:     ban_status.String,
			TechCondition: tech_condition.String,
			Defect:        defect.String,
			Address:       address.String,
			DepartmentID:  department_id.String,
			PetrolType:    petrol_name.String,
			Petrol:        petrol.Float64,
			CreatedAt:     createdAt.String,
			UpdatedAt:     updatedAt.String,
		})
	}

	return resp, nil
}

func (r *carRepo) Update(ctx context.Context, req *models.UpdateCar) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			cars
		SET
		id = :id,
		name = :name,
		model = :model,
		state_number = :state_number,
		year = :year,
		ban_status = :ban_status,
		tech_condition = :tech_condition,
		defect = :defect,
		address = :address,
		department_id = :department_id,
		petrol_name = :petrol_name,
		petrol = :petrol,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":             req.Id,
		"name":           req.Name,
		"model":          req.Model,
		"state_number":   req.StateNumber,
		"year":           req.Year,
		"ban_status":     req.BanStatus,
		"tech_condition": req.TechCondition,
		"defect":         req.Defect,
		"address":        req.Address,
		"department_id":  req.DepartmentID,
		"petrol_name":    req.PetrolType,
		"petrol":         req.Petrol,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *carRepo) Delete(ctx context.Context, req *models.CarPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM cars WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
