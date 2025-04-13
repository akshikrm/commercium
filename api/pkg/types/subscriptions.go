package types

import (
	"context"
	"net/http"
)

type SubscriptionHandler interface {
	GetAllSubscriptions(context.Context, http.ResponseWriter, *http.Request) error
}
