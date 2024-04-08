package repo

import (
	"context"
	"fmt"
	"go-scaffolding/internal/app01/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) buildQueryByWhereConditions(ctx context.Context, model interface{}, whereConditions map[string]interface{}) *gorm.DB {
	query := r.db.Model(model)
	for key, value := range whereConditions {
		query = query.Where(key, value)
	}

	return query
}

func (r *UserRepository) UpdateItemsByWhereConditions(ctx context.Context, model interface{}, whereConditions map[string]interface{}, fieldsToUpdate map[string]interface{}) (affected int64, e error) {
	result := r.buildQueryByWhereConditions(ctx, model, whereConditions).Updates(fieldsToUpdate)
	if result.Error != nil {
		return 0, fmt.Errorf("query.Updates(fieldsToUpdate) : %w", result.Error)
	}

	return result.RowsAffected, nil
}
