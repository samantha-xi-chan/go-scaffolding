package repo

import (
	"context"
	"fmt"
	"go-scaffolding/internal/app01/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateItem(item model.User) (*model.User, error) {
	result := r.db.Create(&item)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user item: %w", result.Error)
	}
	return &item, nil
}

func (r *UserRepo) FindByID(id string) (*model.User, error) {
	var user model.User
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user by ID: %w", result.Error)
	}
	return &user, nil
}

func (r *UserRepo) buildQueryByWhereConditions(ctx context.Context, modelType interface{}, whereConditions map[string]interface{}) *gorm.DB {
	query := r.db.WithContext(ctx).Model(modelType)
	for key, value := range whereConditions {
		query = query.Where(key, value)
	}
	return query
}

func (r *UserRepo) UpdateItemsByWhereConditions(ctx context.Context, modelType interface{}, whereConditions map[string]interface{}, fieldsToUpdate map[string]interface{}) (int64, error) {
	result := r.buildQueryByWhereConditions(ctx, modelType, whereConditions).Updates(fieldsToUpdate)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to update items by where conditions: %w", result.Error)
	}
	return result.RowsAffected, nil
}
