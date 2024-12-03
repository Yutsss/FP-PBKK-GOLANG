package dto

//type User struct {
//	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
//	Name         string `gorm:"type:varchar(255);not null" json:"name"`
//	CompleteName string `gorm:"type:varchar(255);not null" json:"complate_name"`
//	Email        string `gorm:"type:varchar(255);not null" json:"email"`
//	Password     string `gorm:"type:varchar(255);not null" json:"password"`
//	PhoneNumber  string `gorm:"type:varchar(255);not null" json:"phone_number"`
//	Address      string `gorm:"type:varchar(255);not null" json:"address"`
//	Role         string `gorm:"type:varchar(10);not null;default:'user'" json:"role"`
//}

type (
	UserRegisterRequest struct {
		Name         string `json:"name" form:"name" validate:"required,min=1,max=255"`
		CompleteName string `json:"complete_name" form:"complete_name" validate:"required,min=1,max=255"`
		Email        string `json:"email" form:"email" validate:"required,email"`
		Password     string `json:"password" form:"password" validate:"required,min=8,max=255"`
		PhoneNumber  string `json:"phone_number" form:"phone_number" validate:"required,min=10,max=255"`
		Address      string `json:"address" form:"address" validate:"required,min=1,max=255"`
	}

	UserRegisterResponse struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		CompleteName string `json:"complete_name"`
		Email        string `json:"email"`
		PhoneNumber  string `json:"phone_number"`
		Address      string `json:"address"`
		Role         string `json:"role"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
	}

	UserLoginResponse struct {
		Role        string `json:"role"`
		AccessToken string `json:"access_token"`
	}

	UserGetByIdRequest struct {
		UserID uint `json:"id" form:"id" validate:"required"`
	}

	UserGetByIdResponse struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		CompleteName string `json:"complete_name"`
		Email        string `json:"email"`
		PhoneNumber  string `json:"phone_number"`
		Address      string `json:"address"`
		Role         string `json:"role"`
	}

	UserGetAllResponse struct {
		Users []UserGetByIdResponse `json:"users"`
	}
)
