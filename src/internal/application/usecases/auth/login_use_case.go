package usecases

type LoginUseCase struct {
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (uc *LoginUseCase) Handle() string {
	return "hello"
}
