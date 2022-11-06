package repositories

import (
	"fundamental-golang/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(User models.Cart) (models.Cart, error)
	FindCarts() ([]models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart, ID int) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Status").Preload("ProductID").Preload("Product").Preload("Qty").Find(&cart).Error

	return cart, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var Cart models.Cart
	err := r.db.First(&Cart, ID).Error

	return Cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Raw("DELETE FROM cart WHERE id=?", ID).Scan(&cart).Error

	return cart, err
}
