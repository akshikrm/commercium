package handlers_test

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/handlers"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCart(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store)
	handlers := handlers.New(services)

	t.Run("return user with id 2's cart", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/carts", nil)
		response := httptest.NewRecorder()
		ctx := context.Background()
		ctx = context.WithValue(ctx, "userID", uint32(2))
		ctx = context.WithValue(ctx, "role", "user")
		handlers.Cart.GetAll(ctx, response, request)
		got := response.Result().StatusCode
		want := 200
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Add new cart entry for user 2", func(t *testing.T) {
		payload := types.CreateCartRequest{
			ProductID: 1,
			Quantity:  2,
		}
		test, err := json.Marshal(payload)
		if err != nil {
			panic("test failed")
		}

		reader := bytes.NewReader(test)
		request, _ := http.NewRequest(http.MethodPost, "/carts", reader)
		response := httptest.NewRecorder()
		ctx := context.Background()
		ctx = context.WithValue(ctx, "userID", uint32(2))
		ctx = context.WithValue(ctx, "role", "user")
		handlers.Cart.Create(ctx, response, request)
		got := response.Result().StatusCode
		want := 201
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
