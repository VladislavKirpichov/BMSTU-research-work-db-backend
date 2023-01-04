package usecase

import (
	"context"
	"strings"

	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg             *configs.Config
	usersRepository *repository.UsersRepository
}

const salt = "sks23dmc[zpdoi6"
const hashingCost = 14

func NewUserUsecase(repository *repository.UsersRepository, cfg *configs.Config) *UserUsecase {
	return &UserUsecase{
		cfg:             cfg,
		usersRepository: repository,
	}
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{salt, string(hash)}, ""), nil
}

func (u *UserUsecase) SignUp(ctx context.Context, user *models.InputUser) (int64, error) {
	hashedPassword, err := generatePassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hashedPassword

	return u.usersRepository.CreateUser(ctx, user)
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return u.usersRepository.GetUsers(ctx)
}
