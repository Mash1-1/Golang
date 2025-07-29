package usecases

import (
	"errors"
	"task_manager_ca/Domain"
)

type UserUseCase struct {
	PasswordService Domain.PasswordService
	UserRepo Domain.UserRepository
	JwtImp Domain.JwtService
}

type UserUseCaseI interface {
	Register(user Domain.User) error
	Login(user Domain.User) (string, error)
}

func NewUserUseCase(UsrRepo Domain.UserRepository, PassServ Domain.PasswordService, JwtServ Domain.JwtService) UserUseCase {
	return UserUseCase{
		PasswordService: PassServ,
		UserRepo: UsrRepo,
		JwtImp: JwtServ,
	}
} 

func (uc *UserUseCase) Register(user Domain.User) error {
	// Check if the username is unique
	found := uc.UserRepo.FindUserRepository(user.Username)
	if found {
		err := errors.New("username already in use")
		return err
	}

	// Encrypt the password
	hashedPass, err := uc.PasswordService.EncryptPassword(user.Password)
	if err != nil {
		// Handle hashing errors
		return errors.New("error while hashing using bcrypt")
	}

	// Insert the new_user information into the database
	user.Password = string(hashedPass)
	err = uc.UserRepo.Create(user)
	if err != nil {
		return errors.New("error while adding into database")
	}
	return nil
}

func (uc *UserUseCase) Login(user Domain.User) (string, error) {
	// Check if user is in the database 
	existingUser, err := uc.UserRepo.Login(user)
	if err != nil {
		return "", err 
	}
	
	// Check if the user entered the correct password
	if !uc.PasswordService.CheckPasswordHash(user.Password, existingUser.Password){
		return "", errors.New("invalid username or password")
	}

	// Generate Jwt Token
	token, err := uc.JwtImp.CreateJwtToken(existingUser)
	if err != nil {
		// Handle token generation failures
		return "", errors.New("error while generating jwt token")
	}
	return token, nil 
} 