package repository

import (
	"errors"
	"kmipn-2023/model"

	"gorm.io/gorm"
)

type SellerRepository interface {
	GetSellerByEmail(email string) (model.Seller, error)
	CreateSeller(seller model.Seller) (model.Seller, error)
	UpdateSeller(seller model.Seller) (model.Seller, error)
	DeleteSeller(seller model.Seller) (model.Seller, error)
}

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepo(db *gorm.DB) *sellerRepository {
	return &sellerRepository{db}
}

func (r *sellerRepository) GetSellerByEmail(email string) (model.Seller, error) {
	var seller model.Seller
	result := r.db.First(&seller, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Seller{}, errors.New("seller not found")
		}
		return model.Seller{}, result.Error
	}
	return seller, nil
}

func (r *sellerRepository) CreateSeller(seller model.Seller) (model.Seller, error) {
	err := r.db.Create(&seller).Error
	if err != nil {
		return seller, err
	}
	return seller, nil
}

func (r *sellerRepository) UpdateSeller(seller model.Seller) (model.Seller, error) {
	if err := r.db.Save(&seller).Error; err != nil {
		return model.Seller{}, err
	}
	return seller, nil
}

func (r *sellerRepository) DeleteSeller(seller model.Seller) (model.Seller, error) {
	if err := r.db.Delete(&seller).Error; err != nil {
		return model.Seller{}, err
	}
	return seller, nil
}
