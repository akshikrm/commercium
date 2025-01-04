package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"log"
)

type user struct {
	userRepository    types.UserRepository
	profileRepository types.ProfileRepository
}

func (u *user) Login(payload *types.LoginUserRequest) (string, error) {
	user, err := u.userRepository.GetPasswordByEmail(payload.Email)
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

func (u *user) Get() ([]*types.User, error) {
	return u.userRepository.Get()
}

func (u *user) GetCustomerID(id uint32) (*string, error) {
	customerID := u.userRepository.GetCustomerID(id)
	if customerID == nil {
		return nil, utils.ServerError
	}
	return customerID, nil
}

func (u *user) GetProfile(userId uint32) (*types.Profile, error) {
	return u.profileRepository.GetByUserId(userId)
}

func (u *user) GetOne(id uint32) (*types.User, error) {
	user, err := u.userRepository.GetOne(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) Create(user types.CreateUserRequest) (string, error) {
	hashedPassword, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		return "", err
	}

	exists := u.profileRepository.CheckIfUserExists(user.Email)
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
	savedUser, err := u.userRepository.Create(user)
	if err != nil {
		return "", err
	}
	newUserProfile := types.NewProfileRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserID:    savedUser.ID,
	}
	if _, err := u.profileRepository.Create(&newUserProfile); err != nil {
		return "", err
	}
	token, err := utils.CreateJwt(savedUser.ID, savedUser.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *user) Update(id uint32, user *types.UpdateProfileRequest) (*types.Profile, error) {
	err := u.profileRepository.UpdateProfileByUserID(id, user)
	if err != nil {
		return nil, err
	}
	return u.GetProfile(id)
}

func (u *user) Delete(id uint32) error {
	return u.userRepository.Delete(id)
}

func newUserService(userRepository types.UserRepository, profileRepository types.ProfileRepository) *user {
	return &user{userRepository: userRepository, profileRepository: profileRepository}
}
