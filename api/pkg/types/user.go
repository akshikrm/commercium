package types

import "time"

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserID    uint32 `json:"user_id"`
}

type UpdateProfileRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Pincode     string `json:"pincode"`
	AddressOne  string `json:"address_one"`
	AddressTwo  string `json:"address_two"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	ID        uint32     `json:"id"`
	Password  string     `json:"-"`
	Role      string     `json:"role_code"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Profile struct {
	ID          uint32     `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Pincode     string     `json:"pincode"`
	AddressOne  string     `json:"address_one"`
	AddressTwo  string     `json:"address_two"`
	PhoneNumber string     `json:"phone_number"`
	UserID      uint32     `json:"-"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
