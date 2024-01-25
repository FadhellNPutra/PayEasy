package usecase

import (
	"payeasy/entity/dto"
	"payeasy/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	// Logout ()
}

type authUseCase struct {
	userUC     UsersUseCase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.userUC.FindUsersForLogin(payload.User, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func NewAuthUseCase(userUC UsersUseCase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{userUC: userUC, jwtService: jwtService}
}
