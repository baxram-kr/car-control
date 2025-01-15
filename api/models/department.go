package models

type DepartmentPrimaryKey struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type CreateDepartment struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Region      string `json:"region"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type Department struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Region      string `json:"region"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateDepartment struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Region      string `json:"region"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type DepartmentGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type DepartmentGetListResponse struct {
	Count       int           `json:"count"`
	Departments []*Department `json:"departments"`
}
