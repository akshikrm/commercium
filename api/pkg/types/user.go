package types

import (
	"context"
	"net/http"
	"time"
)

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

type UserProfile struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Profile   *Profile  `json:"profile"`
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

type UserRepository interface {
	Get() ([]*User, bool)
	GetOne(id uint32) (*User, bool)
	GetPasswordByEmail(string) (*User, bool)
	GetUserByEmail(string) (*User, bool)
	Create(CreateUserRequest) (*User, bool)
	Update(uint32, UpdateUserRequest) bool
	Delete(uint32) bool
	GetCustomerID(uint32) *string
}

type ProfileRepository interface {
	GetByUserId(uint32) (*Profile, bool)
	Create(*NewProfileRequest) (uint32, bool)
	UpdateProfileByUserID(uint32, *UpdateProfileRequest) bool
	CheckIfUserExists(string) bool
}

type UserServicer interface {
	Get() ([]*User, error)
	GetProfile(uint32) (*Profile, error)
	GetOne(uint32) (*User, error)
	Login(*LoginUserRequest) (string, error)
	Create(CreateUserRequest) (string, error)
	Update(uint32, *UpdateProfileRequest) (*Profile, error)
	Delete(uint32) error
	GetCustomerID(uint32) (*string, error)
	Exists(string) bool
}

type UserHandler interface {
	GetProfile(context.Context, http.ResponseWriter, *http.Request) error
	GetCustomerID(context.Context, http.ResponseWriter, *http.Request) error
	UpdateProfile(context.Context, http.ResponseWriter, *http.Request) error
	GetAll(context.Context, http.ResponseWriter, *http.Request) error
	GetOne(context.Context, http.ResponseWriter, *http.Request) error
	Login(http.ResponseWriter, *http.Request) error
	Create(http.ResponseWriter, *http.Request) error
	Update(http.ResponseWriter, *http.Request) error
	Delete(context.Context, http.ResponseWriter, *http.Request) error
}
