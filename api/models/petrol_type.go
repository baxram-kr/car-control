package models

type PetrolTypePrimaryKey struct {
	Id string `json:"id"`
}

type CreatePetrolType struct {
	Name string `json:"name"`
}

type PetrolType struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdatePetrolType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PetrolTypeGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type PetrolTypeGetListResponse struct {
	Count       int           `json:"count"`
	PetrolTypes []*PetrolType `json:"petrol_types"`
}
