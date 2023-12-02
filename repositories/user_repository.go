package repositories

import (
	"context"
	"test-api/entities/models"
	"test-api/enums"
	"test-api/interfaces"
	"test-api/utils"

	"gorm.io/gorm"
)

type userRepository struct{}

// Update implements interfaces.UserRepository.
func (*userRepository) Update(ctx context.Context, tx *gorm.DB, entity *models.User) (*models.User, error) {
	err := tx.WithContext(ctx).Updates(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Find implements interfaces.UserRepository.
func (*userRepository) Find(ctx context.Context, tx *gorm.DB) ([]models.User, error) {
	return utils.SelectQuery[models.User](ctx, tx)
}

// GetById implements interfaces.UserRepository.
func (*userRepository) GetById(ctx context.Context, tx *gorm.DB, id uint) (*models.User, error) {
	return utils.GetById[models.User](ctx, tx, id)
}

// GetUserByEmail implements interfaces.UserRepository.
func (*userRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (*models.User, error) {
	panic("unimplemented")
}

// Save implements interfaces.UserRepository.
func (*userRepository) Save(ctx context.Context, tx *gorm.DB, entity *models.User) (*models.User, error) {
	return utils.SaveQuery[models.User](ctx, tx, entity, enums.Create)
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

var _ interfaces.UserRepository = &userRepository{}
