package payment

import "github.com/Fibocloud/payment-sdks/ebarimt"

type (
	Client struct {
		APIKey   string `json:"api_key"`
		Endpoint string `json:"endpoint"`
	}

	PaymentCreateInput struct {
		OrgID         uint    `json:"org_id"`
		InvoiceUUID   string  `json:"invoice_uuid"`
		UPointKey     string  `json:"upoint_key"` // Банкнаас явуулсан гүйлгээний дугаар
		Total         float64 `json:"total"`
		PaymentType   string  `json:"payment_type"`
		PaymentMethod string  `json:"payment_method"`
		Phone         string  `json:"phone"`
		Callback      string  `json:"callback"`
	}

	BaseResponse struct {
		Message string `json:"message"`
	}

	BaseSuccessBody struct {
		Success bool `json:"success"`
	}

	PaymentCreateResponse struct {
		BaseResponse
		Body BaseSuccessBody `json:"body"`
	}

	EbarimtStockBody struct {
		Items []ebarimt.StockInput `json:"items"`
	}

	EbarimtStockResponse struct {
		BaseResponse
		Body EbarimtStockBody `json:"body"`
	}
)
