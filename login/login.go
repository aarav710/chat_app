package login

type Register struct {
  Username string `json:"username" validate:"max=25,required"`
  Password string `json:"password" validate:"min=25,required"`
  Email string `json:"email" validate:"required,email"`
}