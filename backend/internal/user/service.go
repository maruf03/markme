package user

import "github.com/google/uuid"

type UserService struct{}

func (u *UserService) GetProfile(id uuid.UUID) (ProfileResponseModel, error) {
	return ProfileResponseModel{}, nil
}

func (u *UserService) UpdateProfile(id uuid.UUID) (ProfileResponseModel, error) {
	return ProfileResponseModel{}, nil
}
