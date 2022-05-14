package domain

import "strings"

type TableNameType string

var (
	CART_PRODUCT TableNameType = "cart_products"
	CART         TableNameType = "carts"
	CASBIN       TableNameType = "casbins"
	CATEGORY     TableNameType = "categories"
	DISCOUNT     TableNameType = "discounts"
	MEDIA        TableNameType = "medias"
	PAYMENT      TableNameType = "payments"
	PERMISSION   TableNameType = "permissions"
	PRODUCT      TableNameType = "products"
	ROLE         TableNameType = "roles"
	SALE         TableNameType = "sales"
	USER         TableNameType = "users"
)

func (t TableNameType) String() string {
	switch t {
	case CART_PRODUCT:
		return "cart_products"
	case CART:
		return "carts"
	case CASBIN:
		return "casbins"
	case CATEGORY:
		return "categories"
	case DISCOUNT:
		return "discounts"
	case MEDIA:
		return "medias"
	case PAYMENT:
		return "payments"
	case PERMISSION:
		return "permissions"
	case ROLE:
		return "roles"
	case SALE:
		return "sales"
	case USER:
		return "users"
	default:
	}
	return "unknown"
}

func (t TableNameType) ParseString(str string) (TableNameType, bool) {
	capabilitiesMap := map[string]TableNameType{
		"cart_product": CART_PRODUCT,
		"cart":         CART,
		"casbin":       CASBIN,
		"category":     CATEGORY,
		"discount":     DISCOUNT,
		"media":        MEDIA,
		"payment":      PAYMENT,
		"permission":   PERMISSION,
		"product":      PRODUCT,
		"role":         ROLE,
		"sale":         SALE,
		"user":         USER,
	}
	c, ok := capabilitiesMap[strings.ToUpper(str)]
	return c, ok
}
