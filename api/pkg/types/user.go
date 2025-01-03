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
	Get() ([]*User, error)
	GetOne(id uint32) (*User, error)
	GetPasswordByEmail(string) (*User, error)
	GetUserByEmail(string) (*User, error)
	Create(CreateUserRequest) (*User, error)
	Update(uint32, UpdateUserRequest) error
	Delete(uint32) error
	GetCustomerID(uint32) *string
}

type ProfileRepository interface {
	GetByUserId(uint32) (*Profile, error)
	Create(*NewProfileRequest) (uint32, error)
	UpdateProfileByUserID(uint32, *UpdateProfileRequest) error
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
	GetCustomerID(id uint32) (*string, error)
}
