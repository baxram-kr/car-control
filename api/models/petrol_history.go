package models

type PetrolHistoryPrimaryKey struct {
	Id string `json:"id"`
}

type CreatePetrolHistory struct {
	CarID              string  `json:"car_id"`
	CarName            string  `json:"car_name"`
	CarModel           string  `json:"car_model"`
	CarStateNumber     string  `json:"car_state_number"`
	CarYear            string  `json:"car_year"`
	CarBanStatus       string  `json:"car_ban_status"`
	CarTechCondition   string  `json:"car_tech_condition"`
	CarDefect          string  `json:"car_defect"`
	CarAddress         string  `json:"car_address"`
	CarDepartmentID    string  `json:"car_department_id"`
	CarPetrolType      string  `json:"car_petrol_name"`
	CarRemainingPetrol float64 `json:"car_remaining_petrol"`
	CarAddedPetrol     float64 `json:"car_added_petrol"`
}

type PetrolHistory struct {
	Id                 string  `json:"id"`
	CarID              string  `json:"car_id"`
	CarName            string  `json:"car_name"`
	CarModel           string  `json:"car_model"`
	CarStateNumber     string  `json:"car_state_number"`
	CarYear            string  `json:"car_year"`
	CarBanStatus       string  `json:"car_ban_status"`
	CarTechCondition   string  `json:"car_tech_condition"`
	CarDefect          string  `json:"car_defect"`
	CarAddress         string  `json:"car_address"`
	CarDepartmentID    string  `json:"car_department_id"`
	CarPetrolType      string  `json:"car_petrol_name"`
	CarRemainingPetrol float64 `json:"car_remaining_petrol"`
	CarAddedPetrol     float64 `json:"car_added_petrol"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

type UpdatePetrolHistory struct {
	Id                 string  `json:"id"`
	CarID              string  `json:"car_id"`
	CarName            string  `json:"car_name"`
	CarModel           string  `json:"car_model"`
	CarStateNumber     string  `json:"car_state_number"`
	CarYear            string  `json:"car_year"`
	CarBanStatus       string  `json:"car_ban_status"`
	CarTechCondition   string  `json:"car_tech_condition"`
	CarDefect          string  `json:"car_defect"`
	CarAddress         string  `json:"car_address"`
	CarDepartmentID    string  `json:"car_department_id"`
	CarPetrolType      string  `json:"car_petrol_name"`
	CarRemainingPetrol float64 `json:"car_remaining_petrol"`
	CarAddedPetrol     float64 `json:"car_added_petrol"`
}

type PetrolHistoryGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type PetrolHistoryGetListResponse struct {
	Count           int              `json:"count"`
	PetrolHistories []*PetrolHistory `json:"petrol_histories"`
}
