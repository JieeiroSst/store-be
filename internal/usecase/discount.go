package usecase

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Discounts interface {
	Create(discount domain.Discount) error
	Update(id string, discount domain.Discount) error
	Delete(id string) error
	Discounts(pagination domain.Pagination) ([]domain.Discount, error)
	DiscountById(id string) (*domain.Discount, error)
}

type DiscountUsecase struct {
	discountRepo repository.Discounts
	snowflake    snowflake.SnowflakeData
}

func NewDiscountUsecase(discountRepo repository.Discounts, snowflake snowflake.SnowflakeData) *DiscountUsecase {
	return &DiscountUsecase{
		discountRepo: discountRepo,
		snowflake:    snowflake,
	}
}

func (u *DiscountUsecase) Create(discount domain.Discount) error {
	args := domain.Discount{
		Id:          u.snowflake.GearedID(),
		Name:        discount.Name,
		Description: discount.Description,
		Type:        discount.Type,
		ProductId:   discount.ProductId,
		StartTime:   discount.StartTime,
		EndTime:     discount.EndTime,
	}
	if err := u.discountRepo.Create(args); err != nil {
		return err
	}
	return nil
}

func (u *DiscountUsecase) Update(id string, discount domain.Discount) error {
	if err := u.discountRepo.Update(id, discount); err != nil {
		return err
	}
	return nil
}

func (u *DiscountUsecase) Delete(id string) error {
	if err := u.discountRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *DiscountUsecase) Discounts(pagination domain.Pagination) ([]domain.Discount, error) {
	discounts, err := u.discountRepo.Discounts(pagination)
	if err != nil {
		return nil, err
	}
	return discounts, nil
}

func (u *DiscountUsecase) DiscountById(id string) (*domain.Discount, error) {
	discount, err := u.discountRepo.DiscountById(id)
	if err != nil {
		return nil, err
	}
	return discount, nil
}
