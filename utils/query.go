package utils

import (
	"context"
	"test-api/enums"

	"gorm.io/gorm"
)

// func InsertQuery[T any](ctx context.Context, tx *gorm.DB, entity *T) (*T, error) {
// 	err := tx.WithContext(ctx).Create(&entity).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return entity, nil
// }

func SelectQuery[T any](ctx context.Context, tx *gorm.DB) ([]T, error) {
	var entities []T
	err := tx.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// func UpdateQuery[T any](ctx context.Context, tx *gorm.DB, entity *T) (*T, error) {
// 	err := tx.WithContext(ctx).Updates(&entity).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return entity, nil
// }

func SaveQuery[T any](ctx context.Context, tx *gorm.DB, entity *T, action string) (*T, error) {
	query := tx.WithContext(ctx)
	switch action {
	case enums.Create:
		query.Create(&entity)
	case enums.Update:
		query.Updates(&entity)
	}
	err := query.Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func GetById[T any](ctx context.Context, tx *gorm.DB, id uint) (entity *T, err error) {
	err = tx.WithContext(ctx).Model(&entity).Where("id = ?", id).Scan(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
