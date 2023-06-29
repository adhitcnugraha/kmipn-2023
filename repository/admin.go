package repository

import (
	"kmipn-2023/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin model.Admin) (model.Admin, error)
	UpdateAdmin(admin model.Admin) (model.Admin, error)
	DeleteAdmin(id string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) CreateAdmin(admin model.Admin) (model.Admin, error) {
	if err := r.db.Create(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *adminRepository) UpdateAdmin(admin model.Admin) (model.Admin, error) {
	if err := r.db.Save(&admin).Error; err != nil {
		return model.Admin{}, err
	}
	return admin, nil
}

func (r *adminRepository) DeleteAdmin(id string) error {
	result := r.db.Delete(&model.Admin{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
