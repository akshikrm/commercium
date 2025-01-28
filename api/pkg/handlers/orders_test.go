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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AllOrdersResponse struct {
	Data []*types.OrderList `json:"data"`
}

func TestPurchase(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store)
	handlers := handlers.New(services)

	t.Run("create a new transaction", func(t *testing.T) {
		reader := bytes.NewReader(transactionCreatedEvent)
		req, _ := http.NewRequest(http.MethodPost, "/transactions", reader)
		res := httptest.NewRecorder()
		handlers.Purchase.HandleTransactionHook(res, req)

		got := res.Result().StatusCode
		want := 200
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestOrderStatus(t *testing.T) {
	config := config.NewTestConfig()
	store := repository.New(config)
	services := services.New(store)
	handlers := handlers.New(services)

	ctx := context.Background()

	// t.Run("get my orders", func(t *testing.T) {
	// 	req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	// 	res := httptest.NewRecorder()
	// 	ctx = context.WithValue(ctx, "userID", uint32(2))
	// 	ctx = context.WithValue(ctx, "role", "user")
	// 	handlers.Purchase.GetAllOrders(ctx, res, req)
	// 	got := res.Body.String()
	// 	fmt.Println(got)
	// 	want := "admin route"
	// 	if got != want {
	// 		t.Errorf("got %s, want %s", got, want)
	// 	}
	// })

	t.Run("get orders of user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
		res := httptest.NewRecorder()
		ctx = context.WithValue(ctx, "userID", uint32(1))
		ctx = context.WithValue(ctx, "role", "admin")
		handlers.Purchase.GetAllOrders(ctx, res, req)
		orders := &AllOrdersResponse{}
		if err := json.NewDecoder(res.Body).Decode(orders); err != nil {
			panic("failed to decode orders")
		}
		got := res.Result().StatusCode
		want := 200
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("get all orders", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
		res := httptest.NewRecorder()
		ctx = context.WithValue(ctx, "userID", uint32(1))
		ctx = context.WithValue(ctx, "role", "admin")
		handlers.Purchase.GetAllOrders(ctx, res, req)
		orders := &AllOrdersResponse{}
		if err := json.NewDecoder(res.Body).Decode(orders); err != nil {
			panic("failed to decode orders")
		}
		got := res.Result().StatusCode
		want := 200
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

var transactionCreatedEvent []byte = []byte(`{
	  "event_id": "evt_01hv8wpv7nkct05j2zxxvpz6dm",
	  "event_type": "transaction.created",
	  "occurred_at": "2024-04-12T10:12:33.653418Z",
	  "notification_id": "ntf_01hv8wpva216nz1ykxf3m3vwyr",
	  "data": {
	    "id": "txn_01hv8wptq8987qeep44cyrewp9",
	    "items": [
	      {
	        "price": {
	          "id": "pri_01gsz8x8sawmvhz1pv30nge1ke",
	          "name": "Monthly (per seat)",
	          "type": "standard",
	          "status": "active",
	          "quantity": {
	            "maximum": 999,
	            "minimum": 1
	          },
	          "tax_mode": "account_setting",
	          "created_at": "2023-02-23T13:55:22.538367Z",
	          "product_id": "pro_01gsz4t5hdjse780zja8vvr7jg",
	          "unit_price": {
	            "amount": "3000",
	            "currency_code": "USD"
	          },
	          "updated_at": "2024-04-11T13:54:52.254748Z",
	          "custom_data": null,
	          "description": "Monthly",
	          "trial_period": null,
	          "billing_cycle": {
	            "interval": "month",
	            "frequency": 1
	          },
	          "unit_price_overrides": [],
	          "import_meta": null
	        },
	        "quantity": 10,
	        "proration": null
	      },
	      {
	        "price": {
	          "id": "pri_01h1vjfevh5etwq3rb416a23h2",
	          "name": "Monthly (recurring addon)",
	          "type": "standard",
	          "status": "active",
	          "quantity": {
	            "maximum": 100,
	            "minimum": 1
	          },
	          "tax_mode": "account_setting",
	          "created_at": "2023-06-01T13:31:12.625056Z",
	          "product_id": "pro_01h1vjes1y163xfj1rh1tkfb65",
	          "unit_price": {
	            "amount": "10000",
	            "currency_code": "USD"
	          },
	          "updated_at": "2024-04-09T07:23:00.907834Z",
	          "custom_data": null,
	          "description": "Monthly",
	          "trial_period": null,
	          "billing_cycle": {
	            "interval": "month",
	            "frequency": 1
	          },
	          "unit_price_overrides": [],
	          "import_meta": null
	        },
	        "quantity": 1,
	        "proration": null
	      },
	      {
	        "price": {
	          "id": "pri_01gsz98e27ak2tyhexptwc58yk",
	          "name": "One-time addon",
	          "type": "standard",
	          "status": "active",
	          "quantity": {
	            "maximum": 1,
	            "minimum": 1
	          },
	          "tax_mode": "account_setting",
	          "created_at": "2023-02-23T14:01:28.391712Z",
	          "product_id": "pro_01gsz97mq9pa4fkyy0wqenepkz",
	          "unit_price": {
	            "amount": "19900",
	            "currency_code": "USD"
	          },
	          "updated_at": "2024-04-09T07:23:10.921392Z",
	          "custom_data": null,
	          "description": "One-time addon",
	          "trial_period": null,
	          "billing_cycle": null,
	          "unit_price_overrides": [],
	          "import_meta": null
	        },
	        "quantity": 1,
	        "proration": null
	      }
	    ],
	    "origin": "web",
	    "status": "draft",
	    "details": {
	      "totals": {
	        "fee": null,
	        "tax": "9585",
	        "total": "57509",
	        "credit": "0",
	        "balance": "57509",
	        "discount": "0",
	        "earnings": null,
	        "subtotal": "47924",
	        "grand_total": "57509",
	        "currency_code": "GBP",
	        "credit_to_balance": "0"
	      },
	      "line_items": [
	        {
	          "id": "txnitm_01hv8wpts6n6wkcr973fmt6gt7",
	          "totals": {
	            "tax": "4800",
	            "total": "28800",
	            "discount": "0",
	            "subtotal": "24000"
	          },
	          "product": {
	            "id": "pro_01gsz4t5hdjse780zja8vvr7jg",
	            "name": "AeroEdit Pro",
	            "type": "standard",
	            "status": "active",
	            "image_url": "https://paddle.s3.amazonaws.com/user/165798/bT1XUOJAQhOUxGs83cbk_pro.png",
	            "created_at": "2023-02-23T12:43:46.605Z",
	            "updated_at": "2024-04-05T15:53:44.687Z",
	            "custom_data": {
	              "features": {
	                "sso": false,
	                "route_planning": true,
	                "payment_by_invoice": false,
	                "aircraft_performance": true,
	                "compliance_monitoring": true,
	                "flight_log_management": true
	              },
	              "suggested_addons": [
	                "pro_01h1vjes1y163xfj1rh1tkfb65",
	                "pro_01gsz97mq9pa4fkyy0wqenepkz"
	              ],
	              "upgrade_description": "Move from Basic to Pro to take advantage of aircraft performance, advanced route planning, and compliance monitoring."
	            },
	            "description": "Designed for professional pilots, including all features plus in Basic plus compliance monitoring, route optimization, and third-party integrations.",
	            "tax_category": "standard",
	            "import_meta": null
	          },
	          "price_id": "pri_01gsz8x8sawmvhz1pv30nge1ke",
	          "quantity": 10,
	          "tax_rate": "0.2",
	          "unit_totals": {
	            "tax": "480",
	            "total": "2880",
	            "discount": "0",
	            "subtotal": "2400"
	          }
	        },
	        {
	          "id": "txnitm_01hv8wpts6n6wkcr973jywgjzx",
	          "totals": {
	            "tax": "1600",
	            "total": "9601",
	            "discount": "0",
	            "subtotal": "8001"
	          },
	          "product": {
	            "id": "pro_01h1vjes1y163xfj1rh1tkfb65",
	            "name": "Analytics addon",
	            "type": "standard",
	            "status": "active",
	            "image_url": "https://paddle.s3.amazonaws.com/user/165798/97dRpA6SXzcE6ekK9CAr_analytics.png",
	            "created_at": "2023-06-01T13:30:50.302Z",
	            "updated_at": "2024-04-05T15:47:17.163Z",
	            "custom_data": null,
	            "description": "Unlock advanced insights into your flight data with enhanced analytics and reporting features. Includes customizable reporting templates and trend analysis across flights.",
	            "tax_category": "standard",
	            "import_meta": null
	          },
	          "price_id": "pri_01h1vjfevh5etwq3rb416a23h2",
	          "quantity": 1,
	          "tax_rate": "0.2",
	          "unit_totals": {
	            "tax": "1600",
	            "total": "9601",
	            "discount": "0",
	            "subtotal": "8001"
	          }
	        },
	        {
	          "id": "txnitm_01hv8wpts6n6wkcr973nc9r7dm",
	          "totals": {
	            "tax": "3185",
	            "total": "19108",
	            "discount": "0",
	            "subtotal": "15923"
	          },
	          "product": {
	            "id": "pro_01gsz97mq9pa4fkyy0wqenepkz",
	            "name": "Custom domains",
	            "type": "standard",
	            "status": "active",
	            "image_url": "https://paddle.s3.amazonaws.com/user/165798/XIG7UXoJQHmlIAiKcnkA_custom-domains.png",
	            "created_at": "2023-02-23T14:01:02.441Z",
	            "updated_at": "2024-04-05T15:43:28.971Z",
	            "custom_data": null,
	            "description": "Make AeroEdit truly your own with custom domains. Custom domains reinforce your brand identity and make it easy for your team to access your account.",
	            "tax_category": "standard",
	            "import_meta": null
	          },
	          "price_id": "pri_01gsz98e27ak2tyhexptwc58yk",
	          "quantity": 1,
	          "tax_rate": "0.2",
	          "unit_totals": {
	            "tax": "3185",
	            "total": "19108",
	            "discount": "0",
	            "subtotal": "15923"
	          }
	        }
	      ],
	      "payout_totals": null,
	      "tax_rates_used": [
	        {
	          "totals": {
	            "tax": "9585",
	            "total": "57509",
	            "discount": "0",
	            "subtotal": "47924"
	          },
	          "tax_rate": "0.2"
	        }
	      ],
	      "adjusted_totals": {
	        "fee": "0",
	        "tax": "9585",
	        "total": "57509",
	        "earnings": "0",
	        "subtotal": "47924",
	        "grand_total": "57509",
	        "currency_code": "GBP"
	      }
	    },
	    "checkout": {
	      "url": "https://aeroedit.com/pay?_ptxn=txn_01hv8wptq8987qeep44cyrewp9"
	    },
	    "payments": [],
	    "billed_at": null,
	    "address_id": null,
	    "created_at": "2024-04-12T10:12:33.201400507Z",
	    "invoice_id": null,
	    "updated_at": "2024-04-12T10:12:33.201400507Z",
	    "revised_at": null,
	    "business_id": null,
	    "custom_data": null,
	    "customer_id": null,
	    "discount_id": null,
	    "currency_code": "GBP",
	    "billing_period": null,
	    "invoice_number": null,
	    "billing_details": null,
	    "collection_mode": "automatic",
	    "subscription_id": null
	  }
	}`)
