package shareddb

type (
	Payment struct {
		BaseOrg
		InvoiceUUID   string  `gorm:"column:invoice_uuid;primary_key" json:"invoice_uuid"`
		UPointKey     string  `gorm:"column:upoint_key" json:"upoint_key"` // Банкнаас явуулсан гүйлгээний дугаар
		Total         float64 `gorm:"column:total" json:"total"`
		PaymentType   string  `gorm:"column:payment_type" json:"payment_type"`
		PaymentMethod string  `gorm:"column:payment_method" json:"payment_method"`
		BankInvioce   string  `gorm:"column:bank_invoice" json:"bank_invoice"` // Банкнаас явуулсан гүйлгээний дугаар
		BankQrText    string  `gorm:"column:bank_qr_text" json:"bank_qr_text"` // Qr хадгалах талбар
		Phone         string  `gorm:"column:phone" json:"phone"`
		IsPaid        bool    `gorm:"column:is_paid" json:"is_paid"` // амжилттай дууссан
		Callback      string  `gorm:"column:callback" json:"callback"`
		PaymentError  string  `gorm:"column:payment_error" json:"payment_error"`
	}

	EbarimtStock struct {
		Base
		BaseOrg
		InvoiceUUID string `gorm:"column:invoice_uuid" json:"invoice_uuid"`
		Code        string `gorm:"column:code" json:"code"`
		Name        string `gorm:"column:name" json:"name"`
		MeasureUnit string `gorm:"column:measureUnit" json:"measureUnit"`
		Qty         string `gorm:"column:qty" json:"qty"`
		UnitPrice   string `gorm:"column:unitPrice" json:"unitPrice"`
		TotalAmount string `gorm:"column:totalAmount" json:"totalAmount"`
		CityTax     string `gorm:"column:cityTax" json:"cityTax"`
		Vat         string `gorm:"column:vat" json:"vat"`
		BarCode     string `gorm:"column:barCode" json:"barCode"`
	}

	Ebarimt struct {
		BaseOrg
		InvoiceUUID   string         `gorm:"column:invoice_uuid;primary_key" json:"invoice_uuid"`
		Stocks        []EbarimtStock `gorm:"foreignKey:InvoiceUUID" json:"stocks"`
		Amount        string         `gorm:"column:amount" json:"amount"`
		Vat           string         `gorm:"column:vat" json:"vat"`
		CashAmount    string         `gorm:"column:cash_amount" json:"cash_amount"`
		NonCashAmount string         `gorm:"column:non_cash_amount" json:"non_cash_amount"`
		CityTax       string         `gorm:"column:city_tax" json:"city_tax"`
		CustomerNo    string         `gorm:"column:customer_no" json:"customer_no"`
		BillType      string         `gorm:"column:bill_type" json:"bill_type"`
		BranchNo      string         `gorm:"column:branch_no" json:"branch_no"`
		DistrictCode  string         `gorm:"column:district_code" json:"district_code"`
		TaxType       string         `gorm:"column:tax_type" json:"tax_type"`
		RegisterNo    string         `gorm:"column:register_no" json:"register_no"`
		BillId        string         `gorm:"column:bill_id" json:"bill_id"`
		MacAddress    string         `gorm:"column:mac_address" json:"mac_address"`
		Date          string         `gorm:"column:date" json:"date"`
		Lottery       string         `gorm:"column:lottery" json:"lottery"`
		InternalCode  string         `gorm:"column:internal_code" json:"internal_code"`
		QrData        string         `gorm:"column:qr_data" json:"qr_data"`
		MerchantId    string         `gorm:"column:merchant_id" json:"merchant_id"`
		Success       bool           `gorm:"column:success" json:"success"`
		IsCancelled   bool           `gorm:"column:is_cancelled" json:"is_cancelled"`
		Image         string         `gorm:"column:image" json:"image"`
		Message       string         `gorm:"column:message" json:"message"`
		IsReturn      bool           `gorm:"is_return" json:"is_return"`
	}
)
