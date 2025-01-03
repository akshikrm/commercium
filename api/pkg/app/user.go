package app

import (
	"akshidas/e-com/pkg/types"
	"context"
	"log"
	"net/http"
)

type UserApi struct {
	UserService types.UserServicer
}

func (u *UserApi) GetProfile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := ctx.Value("userID")
	userProfile, err := u.UserService.GetProfile(id.(uint32))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, userProfile)
}

func (u *UserApi) GetCustomerID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := ctx.Value("userID").(uint32)
	customerID, err := u.UserService.GetCustomerID(uint32(id))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, customerID)
}

func (u *UserApi) UpdateProfile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := ctx.Value("userID")
	a := types.UpdateProfileRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	user, err := u.UserService.Update(id.(uint32), &a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, user)
}

func (u *UserApi) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	users, err := u.UserService.Get()
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, users)
}
func (u *UserApi) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	foundUser, err := u.UserService.GetOne(uint32(id))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, foundUser)
}

func (u *UserApi) Login(w http.ResponseWriter, r *http.Request) error {
	a := types.LoginUserRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	token, err := u.UserService.Login(&a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, token)
}

func (u *UserApi) Create(w http.ResponseWriter, r *http.Request) error {
	a := &types.CreateUserRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	token, err := u.UserService.Create(*a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, token)
}

func (u *UserApi) Update(w http.ResponseWriter, r *http.Request) error {
	a := types.UpdateProfileRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	id, err := parseId(r.PathValue("id"))
	user, err := u.UserService.Update(uint32(id), &a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, user)
}

func (u *UserApi) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Println("deleting")
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	if err := u.UserService.Delete(uint32(id)); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "deleted successfully")
}

func NewUserApi(service types.UserServicer) *UserApi {
	return &UserApi{UserService: service}
}
