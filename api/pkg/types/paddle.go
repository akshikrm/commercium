package types

import "time"

type Body struct {
	EventID        string    `json:"event_id"`
	EventType      string    `json:"event_type"`
	OccurredAt     time.Time `json:"occurred_at"`
	NotificationID string    `json:"notification_id"`
	Data           Data      `json:"data"`
}

type Quantity struct {
	Maximum int `json:"maximum"`
	Minimum int `json:"minimum"`
}
type UnitPrice struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}
type BillingCycle struct {
	Interval  string `json:"interval"`
	Frequency int    `json:"frequency"`
}

type Price struct {
	ID                 string       `json:"id"`
	Name               string       `json:"name"`
	Type               string       `json:"type"`
	Status             string       `json:"status"`
	Quantity           Quantity     `json:"quantity"`
	TaxMode            string       `json:"tax_mode"`
	CreatedAt          time.Time    `json:"created_at"`
	ProductID          string       `json:"product_id"`
	UnitPrice          UnitPrice    `json:"unit_price"`
	UpdatedAt          time.Time    `json:"updated_at"`
	CustomData         any          `json:"custom_data"`
	Description        string       `json:"description"`
	ImportMeta         any          `json:"import_meta"`
	TrialPeriod        any          `json:"trial_period"`
	BillingCycle       BillingCycle `json:"billing_cycle"`
	UnitPriceOverrides []any        `json:"unit_price_overrides"`
}

type Items struct {
	Price     Price `json:"price"`
	Quantity  uint  `json:"quantity"`
	Proration any   `json:"proration"`
}

type Totals struct {
	Fee             string `json:"fee"`
	Tax             string `json:"tax"`
	Total           string `json:"total"`
	Credit          string `json:"credit"`
	Balance         string `json:"balance"`
	Discount        string `json:"discount"`
	Earnings        string `json:"earnings"`
	Subtotal        string `json:"subtotal"`
	GrandTotal      string `json:"grand_total"`
	CurrencyCode    string `json:"currency_code"`
	CreditToBalance string `json:"credit_to_balance"`
}

// type Totals struct {
// 	Tax      string `json:"tax"`
// 	Total    string `json:"total"`
// 	Discount string `json:"discount"`
// 	Subtotal string `json:"subtotal"`
// }

type Features struct {
	Sso                  bool `json:"sso"`
	RoutePlanning        bool `json:"route_planning"`
	PaymentByInvoice     bool `json:"payment_by_invoice"`
	AircraftPerformance  bool `json:"aircraft_performance"`
	ComplianceMonitoring bool `json:"compliance_monitoring"`
	FlightLogManagement  bool `json:"flight_log_management"`
}
type CustomData struct {
	Features           Features `json:"features"`
	SuggestedAddons    []string `json:"suggested_addons"`
	UpgradeDescription string   `json:"upgrade_description"`
}

type PaddleProduct struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Status      string     `json:"status"`
	ImageURL    string     `json:"image_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CustomData  CustomData `json:"custom_data"`
	Description string     `json:"description"`
	ImportMeta  any        `json:"import_meta"`
	TaxCategory string     `json:"tax_category"`
}
type UnitTotals struct {
	Tax      string `json:"tax"`
	Total    string `json:"total"`
	Discount string `json:"discount"`
	Subtotal string `json:"subtotal"`
}
type LineItems struct {
	ID         string        `json:"id"`
	Totals     Totals        `json:"totals"`
	Product    PaddleProduct `json:"product"`
	PriceID    string        `json:"price_id"`
	Quantity   int           `json:"quantity"`
	TaxRate    string        `json:"tax_rate"`
	UnitTotals UnitTotals    `json:"unit_totals"`
}
type PayoutTotals struct {
	Fee             string `json:"fee"`
	Tax             string `json:"tax"`
	Total           string `json:"total"`
	Credit          string `json:"credit"`
	Balance         string `json:"balance"`
	Discount        string `json:"discount"`
	Earnings        string `json:"earnings"`
	Subtotal        string `json:"subtotal"`
	GrandTotal      string `json:"grand_total"`
	CurrencyCode    string `json:"currency_code"`
	CreditToBalance string `json:"credit_to_balance"`
}
type TaxRatesUsed struct {
	Totals  Totals `json:"totals"`
	TaxRate string `json:"tax_rate"`
}
type AdjustedTotals struct {
	Fee          string `json:"fee"`
	Tax          string `json:"tax"`
	Total        string `json:"total"`
	Earnings     string `json:"earnings"`
	Subtotal     string `json:"subtotal"`
	GrandTotal   string `json:"grand_total"`
	CurrencyCode string `json:"currency_code"`
}
type Details struct {
	Totals         Totals         `json:"totals"`
	LineItems      []LineItems    `json:"line_items"`
	PayoutTotals   PayoutTotals   `json:"payout_totals"`
	TaxRatesUsed   []TaxRatesUsed `json:"tax_rates_used"`
	AdjustedTotals AdjustedTotals `json:"adjusted_totals"`
}
type Checkout struct {
	URL string `json:"url"`
}
type Card struct {
	Type           string `json:"type"`
	Last4          string `json:"last4"`
	ExpiryYear     int    `json:"expiry_year"`
	ExpiryMonth    int    `json:"expiry_month"`
	CardholderName string `json:"cardholder_name"`
}
type MethodDetails struct {
	Card Card   `json:"card"`
	Type string `json:"type"`
}
type Payments struct {
	Amount                string        `json:"amount"`
	Status                string        `json:"status"`
	CreatedAt             time.Time     `json:"created_at"`
	ErrorCode             any           `json:"error_code"`
	CapturedAt            time.Time     `json:"captured_at"`
	MethodDetails         MethodDetails `json:"method_details"`
	PaymentMethodID       string        `json:"payment_method_id"`
	PaymentAttemptID      string        `json:"payment_attempt_id"`
	StoredPaymentMethodID string        `json:"stored_payment_method_id"`
}
type BillingPeriod struct {
	EndsAt   time.Time `json:"ends_at"`
	StartsAt time.Time `json:"starts_at"`
}
type Data struct {
	ID             string        `json:"id"`
	Items          []Items       `json:"items"`
	Origin         string        `json:"origin"`
	Status         string        `json:"status"`
	Details        Details       `json:"details"`
	Checkout       Checkout      `json:"checkout"`
	Payments       []Payments    `json:"payments"`
	BilledAt       time.Time     `json:"billed_at"`
	AddressID      string        `json:"address_id"`
	CreatedAt      time.Time     `json:"created_at"`
	InvoiceID      string        `json:"invoice_id"`
	UpdatedAt      time.Time     `json:"updated_at"`
	BusinessID     any           `json:"business_id"`
	CustomData     any           `json:"custom_data"`
	CustomerID     string        `json:"customer_id"`
	DiscountID     any           `json:"discount_id"`
	CurrencyCode   string        `json:"currency_code"`
	BillingPeriod  BillingPeriod `json:"billing_period"`
	InvoiceNumber  string        `json:"invoice_number"`
	BillingDetails any           `json:"billing_details"`
	CollectionMode string        `json:"collection_mode"`
	SubscriptionID string        `json:"subscription_id"`
}
