package usecase

import (
	"context"
	"log"
	"test-api/auth"
	"test-api/entities/models"
	"test-api/entities/payload/req"
	"test-api/entities/payload/res"
	"test-api/interfaces"
	"test-api/repositories"
	"test-api/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type userUsecase struct {
	repo interfaces.UserRepository
	db   *gorm.DB
}

// GetUserById implements interfaces.UserUsecase.
func (usecase *userUsecase) GetUserById(ctx context.Context, id uint) (*res.UserResponse, error) {
	tx := usecase.db.Begin()
	defer utils.CommitRollback(tx)
	user, err := usecase.repo.GetById(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	userResp := &res.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return userResp, nil
}

// GetAllUsers implements interfaces.UserUsecase.
func (usecase *userUsecase) GetAllUsers(ctx context.Context) ([]res.UserResponse, error) {
	tx := usecase.db.Begin()
	defer utils.CommitRollback(tx)
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
	return usersResp, nil
}

// LoginUser implements interfaces.UserUsecase.
func (usecase *userUsecase) LoginUser(ctx context.Context, login *req.UserLogin) (*res.UserLoginResponse, error) {
	tx := usecase.db.Begin()
	defer utils.CommitRollback(tx)
	user, err := usecase.repo.GetUserByEmail(ctx, tx, login.Email)
	if err != nil {
		return nil, err
	}
	isPassValid := utils.IsPasswordValid(user.Password, login.Password)
	if !isPassValid {
		log.Println("password gak valid bleh :p")
		return nil, err
	}
	expTime := time.Now().Add(time.Minute * 15)
	claims := &auth.JWTClaim{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APPLICATION",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(auth.JWT_KEY)
	log.Println(string(auth.JWT_KEY))
	if err != nil {
		log.Println(err.Error())
		log.Println("jwt claim gak valid bleh :p")
		return nil, err
	}
	userResp := res.UserLoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	// tx.Commit()
	return &userResp, nil
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, register *req.UserRegister) (*res.UserResponse, error) {
	tx := usecase.db.Begin()
	pass, err := utils.HashPassword(register.Password)
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
