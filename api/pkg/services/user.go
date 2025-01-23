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
	user, ok := u.userRepository.GetPasswordByEmail(payload.Email)
	if !ok {
		return "", utils.ServerError
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
	user, ok := u.userRepository.Get()
	if !ok {
		return nil, utils.ServerError
	}
	return user, nil
}

func (u *user) GetCustomerID(id uint32) (*string, error) {
	customerID := u.userRepository.GetCustomerID(id)
	if customerID == nil {
		return nil, utils.ServerError
	}
	return customerID, nil
}

func (u *user) GetProfile(userId uint32) (*types.Profile, error) {
	user, ok := u.profileRepository.GetByUserId(userId)
	if !ok {
		return nil, utils.ServerError
	}
	return user, nil
}

func (u *user) GetOne(id uint32) (*types.User, error) {
	user, ok := u.userRepository.GetOne(id)
	if !ok {
		return nil, utils.ServerError
	}
	return user, nil
}

func (u *user) Exists(email string) bool {
	exists := u.profileRepository.CheckIfUserExists(email)
	return exists
}

func (u *user) Create(user types.CreateUserRequest) (string, error) {
	hashedPassword, ok := utils.HashPassword([]byte(user.Password))
	if !ok {
		return "", utils.ServerError
	}

	user.Password = hashedPassword
	savedUser, ok := u.userRepository.Create(user)
	if !ok {
		return "", utils.ServerError
	}

	newUserProfile := types.NewProfileRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		UserID:    savedUser.ID,
	}

	if _, ok := u.profileRepository.Create(&newUserProfile); !ok {
		return "", utils.ServerError
	}
	token, err := utils.CreateJwt(savedUser.ID, savedUser.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *user) Update(id uint32, user *types.UpdateProfileRequest) (*types.Profile, error) {
	ok := u.profileRepository.UpdateProfileByUserID(id, user)
	if !ok {
		return nil, utils.ServerError
	}
	return u.GetProfile(id)
}

func (u *user) Delete(id uint32) error {
	ok := u.userRepository.Delete(id)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newUserService(userRepository types.UserRepository, profileRepository types.ProfileRepository) *user {
	return &user{userRepository: userRepository, profileRepository: profileRepository}
}
