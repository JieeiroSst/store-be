package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Carts interface {
	Create(cart domain.Cart, cartProduct domain.CartProduct) error
	Update(idCart string, cart domain.Cart, cartProduct domain.CartProduct) error
	Delete(id string) error
	Carts(pagination domain.Pagination) ([]domain.Cart, error)
	CartById(userId string) (*domain.Cart, error)
}

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

func (r *CartRepo) Create(cart domain.Cart, cartProduct domain.CartProduct) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(&cart).Error; err != nil {
			return err
		}
		if err := r.db.Create(&cartProduct).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *CartRepo) Update(idCart string, cart domain.Cart, cartProduct domain.CartProduct) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Model(domain.Cart{}).Where("id = ?", idCart).Updates(&cart).Error; err != nil {
			return err
		}
		if err := r.db.Model(domain.CartProduct{}).Where("id = ?", idCart).Updates(&cartProduct).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

func (r *CartRepo) Delete(id string) error {
	if err := r.db.Model(&domain.CartProduct{}).Where("cart_id = ?", id).Update("active", false).Error; err != nil {
		return err
	}
	return nil
}

func (r *CartRepo) Carts(pagination domain.Pagination) ([]domain.Cart, error) {
	var carts []domain.Cart
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&carts)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return carts, nil
}

func (r *CartRepo) CartById(userId string) (*domain.Cart, error) {
	var cart domain.Cart
	query := r.db.Where("user_id = ?", userId).Preload("Products").Find(&cart)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return &cart, nil
}
