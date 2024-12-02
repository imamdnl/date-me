package request

type (
	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserRegisterRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Name     string `json:"name" validate:"required"`
		Age      uint   `json:"age" validate:"required,min=17"`
		Gender   uint   `json:"gender" validate:"required"`
		Bio      string `json:"bio,omitempty"`
		Photo    string `json:"photo" validate:"required"`
	}
)
