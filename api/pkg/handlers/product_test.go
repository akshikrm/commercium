package handlers_test

import (
	config "akshidas/e-com"
	"akshidas/e-com/pkg/handlers"
	"akshidas/e-com/pkg/repository"
	"akshidas/e-com/pkg/services"
	"akshidas/e-com/pkg/types"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AllProductResponse struct {
	Data []types.ProductsList `json:"data"`
}

func TestProduct(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store, config)
	handlers := handlers.New(services)

	t.Run("create a normal product", func(t *testing.T) {
		payload := types.NewProductRequest{
			Name: "test product",
		}
		request, _ := http.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()
		ctx := context.Background()
		ctx = context.WithValue(ctx, "userID", 1)
		ctx = context.WithValue(ctx, "role", "admin")
	})

	t.Run("get all products", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()
		ctx := context.Background()
		ctx = context.WithValue(ctx, "userID", 2)
		ctx = context.WithValue(ctx, "role", "user")
		handlers.Product.GetAll(ctx, response, request)

		// products := &AllProductResponse{}
		// if err := json.NewDecoder(response.Body).Decode(products); err != nil {
		// 	fmt.Println(err)
		// }
		got := response.Result().StatusCode
		want := 200
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	// t.Run("Add new cart entry for user 2", func(t *testing.T) {
	// 	payload := types.CreateCartRequest{
	// 		ProductID: 1,
	// 		Quantity:  2,
	// 	}
	// 	test, err := json.Marshal(payload)
	// 	if err != nil {
	// 		panic("test failed")
	// 	}
	//
	// 	reader := bytes.NewReader(test)
	// 	request, _ := http.NewRequest(http.MethodPost, "/carts", reader)
	// 	response := httptest.NewRecorder()
	// 	ctx := context.Background()
	// 	ctx = context.WithValue(ctx, "userID", 2)
	// 	ctx = context.WithValue(ctx, "role", "user")
	// 	handlers.Cart.Create(ctx, response, request)
	// 	got := response.Result().StatusCode
	// 	want := 201
	// 	if got != want {
	// 		t.Errorf("got %d, want %d", got, want)
	// 	}
	// })

}
