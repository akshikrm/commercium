package services

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"net/url"
)

type product struct {
	repository      types.ProductRepository
	paymentProvider types.PaymentProvider
}

func (r *product) GetAll(filter url.Values) ([]*types.ProductsList, error) {
	products, ok := r.repository.GetAll(filter)
	if !ok {
		return nil, utils.ServerError
	}
	return products, nil
}

func (r *product) Create(newProduct *types.NewProductRequest) error {
	if err := r.paymentProvider.CreateProduct(newProduct); err != nil {
		return err
	}

	savedProduct, ok := r.repository.InsertOne(newProduct)
	if !ok {
		return utils.ServerError
	}

	if savedProduct.Type == types.OneTimeProduct {
		newPrice := types.NewPricePayload{
			ProductID: savedProduct.ProductID,
			Name:      "one time price",
			Price:     newProduct.Price,
		}

		if price := r.paymentProvider.CreatePrice(newPrice); price != nil {
			price.ProductID = savedProduct.ID
			if ok := r.repository.InsertPrice(nil, price); ok {
				return nil
			}
		}
		return utils.ServerError
	}

	for key, priceItem := range newProduct.SubscriptionPrice {
		newPrice := types.NewPricePayload{
			ProductID: savedProduct.ProductID,
			Name:      priceItem.Label,
			Price:     priceItem.Price,
			BillingCycle: &types.BillingCycle{
				Interval:  string(priceItem.Interval),
				Frequency: 12,
			},
		}
		if price := r.paymentProvider.CreatePrice(newPrice); price != nil {
			price.ProductID = savedProduct.ID
			if ok := r.repository.InsertPrice(&key, price); !ok {
				return utils.ServerError
			}
		}
	}
	return nil
}

func (r *product) Update(id int, newProduct *types.NewProductRequest) (*types.OneProduct, error) {
	updatedProduct, ok := r.repository.Update(id, newProduct)
	if !ok {
		return nil, utils.ServerError
	}
	return updatedProduct, nil
}

func (r *product) UpdatePrice(priceID string, updatePrice *types.UpdatePriceRequest) error {
	updatedPrice := r.paymentProvider.UpdatePrice(priceID, updatePrice)
	if updatePrice == nil {
		return utils.PaddleError
	}
	if ok := r.repository.UpdatePrice(updatedPrice); !ok {
		return utils.ServerError
	}
	return nil
}

func (r *product) GetOne(id int) (*types.OneProduct, error) {
	product, ok := r.repository.GetOne(id)
	if !ok {
		return nil, utils.ServerError
	}
	if product.Type == types.OneTimeProduct {
		product.Price = product.Prices[0].Price
	} else {
		for _, price := range product.Prices {
			product.SubscriptionPrice[price.Interval] = types.ProductPrice{
				ID:      price.ID,
				PriceID: price.PriceID,
				Label:   price.Label,
				Price:   price.Price,
			}
		}
	}
	return product, nil
}

func (r *product) Delete(id int) error {
	ok := r.repository.Delete(id)
	if !ok {
		return utils.ServerError
	}
	return nil
}

func newProductService(repository types.ProductRepository, paymentProvider types.PaymentProvider) *product {
	return &product{
		repository:      repository,
		paymentProvider: paymentProvider,
	}
}
