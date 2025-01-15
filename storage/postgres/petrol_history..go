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

type petrol_historyRepo struct {
	db *pgxpool.Pool
}

func NewPetrolHistoryRepo(db *pgxpool.Pool) *petrol_historyRepo {
	return &petrol_historyRepo{
		db: db,
	}
}

func (r *petrol_historyRepo) Create(ctx context.Context, req *models.CreatePetrolHistory) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO petrol_history(id, car_id,car_name,car_model,car_state_number,car_year,car_ban_status,car_tech_condition,car_defect,car_address,car_department_id,car_petrol_name,car_remaining_petrol,car_added_petrol, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12,$13, $14,NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.CarID,
		req.CarName,
		req.CarModel,
		req.CarStateNumber,
		req.CarYear,
		req.CarBanStatus,
		req.CarTechCondition,
		req.CarDefect,
		req.CarAddress,
		req.CarDepartmentID,
		req.CarPetrolType,
		req.CarRemainingPetrol,
		req.CarAddedPetrol,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *petrol_historyRepo) GetByID(ctx context.Context, req *models.PetrolHistoryPrimaryKey) (*models.PetrolHistory, error) {

	var whereField = "id"
	var (
		query string

		id                   sql.NullString
		car_id               sql.NullString
		car_name             sql.NullString
		car_model            sql.NullString
		car_state_number     sql.NullString
		car_year             sql.NullString
		car_ban_status       sql.NullString
		car_tech_condition   sql.NullString
		car_defect           sql.NullString
		car_address          sql.NullString
		car_department_id    sql.NullString
		car_petrol_name      sql.NullString
		car_remaining_petrol sql.NullFloat64
		car_added_petrol     sql.NullFloat64
		createdAt            sql.NullString
		updatedAt            sql.NullString
	)

	query = `
		SELECT
			id,
			car_id,
			car_name,
			car_model,
			car_state_number,
			car_year,
			car_ban_status,
			car_tech_condition,
			car_defect,
			car_address,
			car_department_id,
			car_petrol_name,
			car_remaining_petrol,
			car_addedd_petrol,
			created_at,
			updated_at
		FROM petrol_history
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&car_id,
		&car_name,
		&car_model,
		&car_state_number,
		&car_year,
		&car_ban_status,
		&car_tech_condition,
		&car_defect,
		&car_address,
		&car_department_id,
		&car_petrol_name,
		&car_remaining_petrol,
		&car_added_petrol,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.PetrolHistory{
		Id:                 id.String,
		CarID:              car_id.String,
		CarName:            car_name.String,
		CarModel:           car_model.String,
		CarStateNumber:     car_state_number.String,
		CarYear:            car_year.String,
		CarBanStatus:       car_ban_status.String,
		CarTechCondition:   car_tech_condition.String,
		CarDefect:          car_defect.String,
		CarAddress:         car_address.String,
		CarDepartmentID:    car_department_id.String,
		CarPetrolType:      car_petrol_name.String,
		CarRemainingPetrol: car_remaining_petrol.Float64,
		CarAddedPetrol:     car_added_petrol.Float64,
		CreatedAt:          createdAt.String,
		UpdatedAt:          updatedAt.String,
	}, nil
}

func (r *petrol_historyRepo) GetList(ctx context.Context, req *models.PetrolHistoryGetListRequest) (*models.PetrolHistoryGetListResponse, error) {

	var (
		resp   = &models.PetrolHistoryGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			car_id,
			car_name,
			car_model,
			car_state_number,
			car_year,
			car_ban_status,
			car_tech_condition,
			car_defect,
			car_address,
			car_department_id,
			car_petrol_name,
			car_remaining_petrol,
			car_added_petrol,
			created_at,
			updated_at
		FROM petrol_history
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
			id                   sql.NullString
			car_id               sql.NullString
			car_name             sql.NullString
			car_model            sql.NullString
			car_state_number     sql.NullString
			car_year             sql.NullString
			car_ban_status       sql.NullString
			car_tech_condition   sql.NullString
			car_defect           sql.NullString
			car_address          sql.NullString
			car_department_id    sql.NullString
			car_petrol_name      sql.NullString
			car_remaining_petrol sql.NullFloat64
			car_added_petrol     sql.NullFloat64
			createdAt            sql.NullString
			updatedAt            sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&car_id,
			&car_name,
			&car_model,
			&car_state_number,
			&car_year,
			&car_ban_status,
			&car_tech_condition,
			&car_defect,
			&car_address,
			&car_department_id,
			&car_petrol_name,
			&car_remaining_petrol,
			&car_added_petrol,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.PetrolHistories = append(resp.PetrolHistories, &models.PetrolHistory{
			Id:                 id.String,
			CarID:              car_id.String,
			CarName:            car_name.String,
			CarModel:           car_model.String,
			CarStateNumber:     car_state_number.String,
			CarYear:            car_year.String,
			CarBanStatus:       car_ban_status.String,
			CarTechCondition:   car_tech_condition.String,
			CarDefect:          car_defect.String,
			CarAddress:         car_address.String,
			CarDepartmentID:    car_department_id.String,
			CarPetrolType:      car_petrol_name.String,
			CarRemainingPetrol: car_remaining_petrol.Float64,
			CarAddedPetrol:     car_added_petrol.Float64,
			CreatedAt:          createdAt.String,
			UpdatedAt:          updatedAt.String,
		})
	}

	return resp, nil
}

func (r *petrol_historyRepo) Update(ctx context.Context, req *models.UpdatePetrolHistory) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			petrol_history
		SET
		id = :id,
		car_id = :car_id,
		car_name = :car_name,
		car_model = :car_model,
		car_state_number = :car_state_number,
		car_year = :car_year,
		car_ban_status = :car_ban_status,
		car_tech_condition = :car_tech_condition,
		car_defect = :car_defect,
		car_address = :car_address,
		car_department_id = :car_department_id,
		car_petrol_name = :car_petrol_name,
		car_remaining_petrol = :car_remaining_petrol,
		car_added_petrol = :car_added_petrol,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":                   req.Id,
		"car_id":               req.Id,
		"car_name":             req.CarName,
		"car_model":            req.CarModel,
		"car_state_number":     req.CarStateNumber,
		"car_year":             req.CarYear,
		"car_ban_status":       req.CarBanStatus,
		"car_tech_condition":   req.CarTechCondition,
		"car_defect":           req.CarDefect,
		"car_address":          req.CarAddress,
		"car_department_id":    req.CarDepartmentID,
		"car_petrol_name":      req.CarPetrolType,
		"car_remaining_petrol": req.CarRemainingPetrol,
		"car_added_petrol":     req.CarAddedPetrol,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *petrol_historyRepo) Delete(ctx context.Context, req *models.PetrolHistoryPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM petrol_history WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
