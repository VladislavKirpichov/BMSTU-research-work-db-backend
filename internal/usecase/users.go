package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
	"github.com/v.kirpichov/admin/pkg/errorHandler"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg               *configs.Config
	usersRepository   repository.UserR
	sessionRepository repository.SessionR
}

const (
	salt        = "sks23dmc[zpdoi6"
	hashingCost = 14
	uuidBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func NewUserUsecase(usersRepo repository.UserR, sessionRepo repository.SessionR, cfg *configs.Config) *UserUsecase {
	return &UserUsecase{
		cfg:               cfg,
		usersRepository:   usersRepo,
		sessionRepository: sessionRepo,
	}
}

func (u *UserUsecase) SignIn(ctx context.Context, user *models.SignInUser) (*models.User, error) {
	fullUserInfo, err := u.usersRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, errorHandler.ErrNoUser
	}

	passwordWithoutSalt := strings.TrimPrefix(fullUserInfo.Password, salt)

	err = bcrypt.CompareHashAndPassword([]byte(passwordWithoutSalt), []byte(user.Password))
	if err != nil {
		return nil, errorHandler.ErrWrongPassword
	}

	return fullUserInfo, nil
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

func (u *UserUsecase) GetSessionToken(ctx context.Context, email string) (string, error) {
	token := uuid.NewString()

	err := u.sessionRepository.CreateSession(email, token, time.Duration(u.cfg.SessionConfig.ExpiresAt*int(time.Hour.Nanoseconds())))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return u.usersRepository.GetUsers(ctx)
}

func (u *UserUsecase) Logout(ctx context.Context, email string) error {
	return u.sessionRepository.DeleteSession(email)
}
