package auth

type RegisterRequest struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
	PhoneNumber string `json:"phone_number" validate:"required,containsNumber"`
	Address     string `json:"address" validate:"required,containsNumber"`
	ProvinceID  uint   `json:"province_id" validate:"required"`
	CityID      uint   `json:"city_id" validate:"required"`
	DistrictID  uint   `json:"district_id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
