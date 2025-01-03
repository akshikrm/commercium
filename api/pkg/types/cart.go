package types

import "time"

type CreateCartRequest struct {
	UserID    uint32 `json:"user_id"`
	ProductID uint32 `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity uint `json:"quantity"`
}

type Cart struct {
	ID        uint32     `json:"id"`
	UserID    uint32     `json:"user_id"`
	ProductID uint32     `json:"product_id"`
	Quantity  uint       `json:"quantity"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CartList struct {
	ID       uint32 `json:"id"`
	Quantity uint   `json:"quantity"`
	PriceID  string `json:"price_id"`
	Product  struct {
		ID          uint32 `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Price       uint   `json:"price"`
		Description string `json:"description"`
		Image       string `json:"image"`
	} `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type CartRepository interface {
	GetAll(uint32) ([]*CartList, error)
	GetOne(uint32) (*CartList, error)
	Create(*CreateCartRequest) (*Cart, error)
	Update(uint32, *UpdateCartRequest) (*CartList, error)
	Delete(uint32) error
	CheckIfEntryExist(uint32, uint32) (bool, error)
	UpdateQuantity(uint32, uint32, uint) error
	HardDeleteByUserID(string) error
}

type CartServicer interface {
	GetAll(uint32) ([]*CartList, error)
	GetOne(uint32) (*CartList, error)
	Create(*CreateCartRequest) error
	Update(uint32, *UpdateCartRequest) (*CartList, error)
	Delete(uint32) error
	HardDeleteByUserID(string) error
}
