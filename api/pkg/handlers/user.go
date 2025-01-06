package handlers

import (
	"akshidas/e-com/pkg/types"
	"context"
	"log"
	"net/http"
)

type user struct {
	service types.UserServicer
}

func (u *user) GetProfile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := ctx.Value("userID")
	userProfile, err := u.service.GetProfile(id.(uint32))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, userProfile)
}

func (u *user) GetCustomerID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := ctx.Value("userID").(uint32)
	customerID, err := u.service.GetCustomerID(uint32(id))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, customerID)
}

func (u *user) UpdateProfile(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := ctx.Value("userID")
	a := types.UpdateProfileRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	user, err := u.service.Update(id.(uint32), &a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, user)
}

func (u *user) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	users, err := u.service.Get()
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, users)
}
func (u *user) GetOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	foundUser, err := u.service.GetOne(uint32(id))
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, foundUser)
}

func (u *user) Login(w http.ResponseWriter, r *http.Request) error {
	a := types.LoginUserRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	token, err := u.service.Login(&a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, token)
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) error {
	a := &types.CreateUserRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	token, err := u.service.Create(*a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, token)
}

func (u *user) Update(w http.ResponseWriter, r *http.Request) error {
	a := types.UpdateProfileRequest{}
	if err := DecodeBody(r.Body, &a); err != nil {
		return err
	}
	id, err := parseId(r.PathValue("id"))
	user, err := u.service.Update(uint32(id), &a)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, user)
}

func (u *user) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	log.Println("deleting")
	id, err := parseId(r.PathValue("id"))
	if err != nil {
		return err
	}
	if err := u.service.Delete(uint32(id)); err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, "deleted successfully")
}

func newUser(service types.UserServicer) types.UserHandler {
	handler := new(user)
	handler.service = service
	return handler
}
