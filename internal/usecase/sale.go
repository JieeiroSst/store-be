package usecase

import (
	"fmt"
	"time"

	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/model"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Sales interface {
	Create(sale model.InputSale) error
	Update(id string, model model.InputSale) error
	Delete(id string) error
	Sales(pagination domain.Pagination) ([]domain.Sale, error)
	SaleById(id string) (*domain.Sale, error)
	IsExpireById(id string) error
}

type SaleUsecase struct {
	saleRepo  repository.Sales
	snowflake snowflake.SnowflakeData
}

func NewSaleUsecase(saleRepo repository.Sales, snowflake snowflake.SnowflakeData) *SaleUsecase {
	return &SaleUsecase{
		saleRepo:  saleRepo,
		snowflake: snowflake,
	}
}

func (u *SaleUsecase) Create(model model.InputSale) error {
	unixTimeUTC := time.Unix(int64(model.Expire), 0)
	sale := domain.Sale{
		Id:          u.snowflake.GearedID(),
		Amount:      model.Amount,
		Description: model.Description,
		Type:        model.Type,
		CustomerId:  model.CustomerId,
		Expire:      unixTimeUTC,
	}
	if err := u.saleRepo.Create(sale); err != nil {
		return err
	}
	return nil
}

func (u *SaleUsecase) Update(id string, model model.InputSale) error {
	unixTimeUTC := time.Unix(int64(model.Expire), 0)
	sale := domain.Sale{
		Amount:      model.Amount,
		Description: model.Description,
		Type:        model.Type,
		CustomerId:  model.CustomerId,
		Expire:      unixTimeUTC,
	}
	if err := u.saleRepo.Update(id, sale); err != nil {
		return err
	}
	return nil
}

func (u *SaleUsecase) Delete(id string) error {
	if err := u.saleRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *SaleUsecase) Sales(pagination domain.Pagination) ([]domain.Sale, error) {
	sales, err := u.saleRepo.Sales(pagination)
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (u *SaleUsecase) SaleById(id string) (*domain.Sale, error) {
	sale, err := u.saleRepo.SaleById(id)
	if err != nil {
		return nil, err
	}
	return sale, nil
}

func (u *SaleUsecase) IsExpireById(id string) error {
	now := time.Now()
	expire, err := u.saleRepo.ExpireById(id)
	if err != nil {
		return err
	}
	if ok := expire.Equal(now); !ok {
		return fmt.Errorf("sale time is over")
	}
	return nil
}
