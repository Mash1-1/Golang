package infrastructure

import "golang.org/x/crypto/bcrypt"

type BcryptPasswordService struct {} 
// Implement the PasswordService interface
func (b BcryptPasswordService) EncryptPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, err 
	}
	return hashedPassword, nil 
}

func (b BcryptPasswordService) CheckPasswordHash(password, hashedPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}