package models

type CarPrimaryKey struct {
	Id string `json:"id"`
}

type CreateCar struct {
	Name          string  `json:"name"`
	Model         string  `json:"model"`
	StateNumber   string  `json:"state_number"`
	Year          string  `json:"year"`
	BanStatus     string  `json:"ban_status"`
	TechCondition string  `json:"tech_condition"`
	Defect        string  `json:"defect"`
	Address       string  `json:"address"`
	DepartmentID  string  `json:"department_id"`
	PetrolType    string  `json:"petrol_name"`
	Petrol        float64 `json:"petrol"`
}

type Car struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Model         string  `json:"model"`
	StateNumber   string  `json:"state_number"`
	Year          string  `json:"year"`
	BanStatus     string  `json:"ban_status"`
	TechCondition string  `json:"tech_condition"`
	Defect        string  `json:"defect"`
	Address       string  `json:"address"`
	DepartmentID  string  `json:"department_id"`
	PetrolType    string  `json:"petrol_name"`
	Petrol        float64 `json:"petrol"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type UpdateCar struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Model         string  `json:"model"`
	StateNumber   string  `json:"state_number"`
	Year          string  `json:"year"`
	BanStatus     string  `json:"ban_status"`
	TechCondition string  `json:"tech_condition"`
	Defect        string  `json:"defect"`
	Address       string  `json:"address"`
	DepartmentID  string  `json:"department_id"`
	PetrolType    string  `json:"petrol_name"`
	Petrol        float64 `json:"petrol"`
}

type CarGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type CarGetListResponse struct {
	Count int    `json:"count"`
	Cars  []*Car `json:"cars"`
}
