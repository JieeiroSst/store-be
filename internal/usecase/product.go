package usecase

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/model"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Products interface {
	Create(product model.InputProduct) error
	Update(id string, product model.InputProduct) error
	Delete(id string) error
	ProductById(id string) (*domain.Product, error)
	Products(pagination domain.Pagination) ([]domain.Product, error)
}

type ProductUsecase struct {
	productRepo repository.Products
	snowflake   snowflake.SnowflakeData
}

func NewProductUsecase(productRepo repository.Products, snowflake snowflake.SnowflakeData) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
		snowflake:   snowflake,
	}
}

func (u *ProductUsecase) Create(input model.InputProduct) error {
	product := domain.Product{
		Id:          u.snowflake.GearedID(),
		Number:      input.Number,
		Description: input.Description,
		Type:        input.Type,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
	}
	if err := u.productRepo.Create(product); err != nil {
		return err
	}
	return nil
}

func (u *ProductUsecase) Update(id string, input model.InputProduct) error {
	product := domain.Product{
		Id:          u.snowflake.GearedID(),
		Number:      input.Number,
		Description: input.Description,
		Type:        input.Type,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
	}
	if err := u.productRepo.Update(id, product); err != nil {
		return err
	}
	return nil
}

func (u *ProductUsecase) Delete(id string) error {
	if err := u.productRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *ProductUsecase) ProductById(id string) (*domain.Product, error) {
	product, err := u.productRepo.ProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (u *ProductUsecase) Products(pagination domain.Pagination) ([]domain.Product, error) {
	products, err := u.productRepo.Products(pagination)
	if err != nil {
		return nil, err
	}
	return products, nil
}
