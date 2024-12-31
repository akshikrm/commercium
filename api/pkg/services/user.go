package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"log"
)

type UserModeler interface {
	Get() ([]*types.User, error)
	GetOne(id uint32) (*types.User, error)
	GetPasswordByEmail(string) (*types.User, error)
	GetUserByEmail(string) (*types.User, error)
	Create(types.CreateUserRequest) (*types.User, error)
	Update(uint32, types.UpdateUserRequest) error
	Delete(uint32) error
	GetCustomerID(uint32) *string
}

type ProfileModeler interface {
	GetByUserId(uint32) (*types.Profile, error)
	Create(*types.NewProfileRequest) (uint32, error)
	UpdateProfileByUserID(uint32, *types.UpdateProfileRequest) error
	CheckIfUserExists(string) bool
}

type UserService struct {
	userModel    UserModeler
	profileModel ProfileModeler
}

func (u *UserService) Login(payload *types.LoginUserRequest) (string, error) {
	user, err := u.userModel.GetPasswordByEmail(payload.Email)
	if err != nil {
		return "", err
	}

	if err := utils.ValidateHash([]byte(user.Password), payload.Password); err != nil {
		log.Printf("invalid password for user %s", payload.Email)
		return "", utils.Unauthorized
	}

	token, err := utils.CreateJwt(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) Get() ([]*types.User, error) {
	return u.userModel.Get()
}

func (u *UserService) GetCustomerID(id uint32) (*string, error) {
	customerID := u.userModel.GetCustomerID(id)
	if customerID == nil {
		return nil, utils.ServerError
	}
	return customerID, nil
}

func (u *UserService) GetProfile(userId uint32) (*types.Profile, error) {
	return u.profileModel.GetByUserId(userId)
}

func (u *UserService) GetOne(id uint32) (*types.User, error) {
	user, err := u.userModel.GetOne(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Create(user types.CreateUserRequest) (string, error) {
	hashedPassword, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		return "", err
	}

	exists := u.profileModel.CheckIfUserExists(user.Email)
	if exists {
		return "", utils.Conflict
	}
	user.Password = hashedPassword
	paddlePayment := new(PaddlePayment)
	if err := paddlePayment.Init(); err != nil {
		return "", err
	}

	if err := paddlePayment.CreateCustomer(&user); err != nil {
		return "", err
	}
	savedUser, err := u.userModel.Create(user)
	if err != nil {
		return "", err
	}
	newUserProfile := types.NewProfileRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserID:    savedUser.ID,
	}
	if _, err := u.profileModel.Create(&newUserProfile); err != nil {
		return "", err
	}
	token, err := utils.CreateJwt(savedUser.ID, savedUser.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserService) Update(id uint32, user *types.UpdateProfileRequest) (*types.Profile, error) {
	err := u.profileModel.UpdateProfileByUserID(id, user)
	if err != nil {
		return nil, err
	}
	return u.GetProfile(id)
}

func (u *UserService) Delete(id uint32) error {
	return u.userModel.Delete(id)
}

func NewUserService(userModel UserModeler, profileModel ProfileModeler) *UserService {
	return &UserService{userModel: userModel, profileModel: profileModel}
}
