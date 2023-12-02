package usecase

import (
	"context"
	"test-api/entities/models"
	"test-api/entities/payload/req"
	"test-api/entities/payload/res"
	"test-api/interfaces"
	"test-api/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUsecase struct {
	repo interfaces.UserRepository
	db   *gorm.DB
}

// GetUserById implements interfaces.UserUsecase.
func (usecase *userUsecase) GetUserById(ctx context.Context, id uint) (*res.UserResponse, error) {
	tx := usecase.db.Begin()
	user, err := usecase.repo.GetById(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	userResp := &res.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	tx.Commit()
	return userResp, nil
}

// GetAllUsers implements interfaces.UserUsecase.
func (usecase *userUsecase) GetAllUsers(ctx context.Context) ([]res.UserResponse, error) {
	tx := usecase.db.Begin()
	users, err := usecase.repo.Find(ctx, tx)
	if err != nil {
		return nil, err
	}
	var usersResp []res.UserResponse
	for _, user := range users {
		userResp := res.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		usersResp = append(usersResp, userResp)
	}
	tx.Commit()
	return usersResp, nil
}

// LoginUser implements interfaces.UserUsecase.
func (*userUsecase) LoginUser(ctx context.Context, login *req.UserLogin) (*res.UserResponse, error) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, register *req.UserRegister) (*res.UserResponse, error) {
	tx := usecase.db.Begin()
	pass, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: string(pass),
	}
	user, err = usecase.repo.Save(ctx, tx, user)
	if err != nil {
		return nil, err
	}
	userRes := &res.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	tx.Commit()
	return userRes, nil
}

func NewUserUsecase(db *gorm.DB) *userUsecase {
	return &userUsecase{
		db:   db,
		repo: repositories.NewUserRepository(),
	}
}

var _ interfaces.UserUsecase = &userUsecase{}
