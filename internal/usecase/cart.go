package usecase

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Carts interface {
	Create(cart domain.Cart, cartProduct domain.CartProduct) error
	Update(idCart string, cart domain.Cart, cartProduct domain.CartProduct) error
	Delete(id string) error
	Carts(pagination domain.Pagination) ([]domain.Cart, error)
	CartById(userId string) (*domain.Cart, error)
}

type CartUsecase struct {
	cartRepo  repository.Carts
	snowflake snowflake.SnowflakeData
}

func NewCartUsecase(cartRepo repository.Carts, snowflake snowflake.SnowflakeData) *CartUsecase {
	return &CartUsecase{
		cartRepo:  cartRepo,
		snowflake: snowflake,
	}
}

func (u *CartUsecase) Create(cart domain.Cart, cartProduct domain.CartProduct) error {
	cartId := u.snowflake.GearedID()
	argsCart := domain.Cart{
		Id:       cartId,
		UserId:   cart.UserId,
		Products: cart.Products,
	}

	argsProduct := domain.CartProduct{
		ProductId: cartProduct.CartId,
		CartId:    cartId,
		Active:    cartProduct.Active,
	}

	if err := u.cartRepo.Create(argsCart, argsProduct); err != nil {
		return err
	}
	return nil

}

func (u *CartUsecase) Update(idCart string, cart domain.Cart, cartProduct domain.CartProduct) error {
	if err := u.cartRepo.Update(idCart, cart, cartProduct); err != nil {
		return err
	}
	return nil
}

func (u *CartUsecase) Delete(id string) error {
	if err := u.cartRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *CartUsecase) Carts(pagination domain.Pagination) ([]domain.Cart, error) {
	carts, err := u.cartRepo.Carts(pagination)
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (u *CartUsecase) CartById(userId string) (*domain.Cart, error) {
	cart, err := u.cartRepo.CartById(userId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
