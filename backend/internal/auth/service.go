package auth

type AuthService struct{}

func (a *AuthService) Login(email string, password string) (LoginResponseModel, error) {
	return LoginResponseModel{}, nil
}

func (a *AuthService) Refresh(refresgToken string) (LoginResponseModel, error) {
	return LoginResponseModel{}, nil
}
