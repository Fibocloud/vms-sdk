package marketing

type (
	Client struct {
		APIKey   string `json:"api_key"`
		Endpoint string `json:"endpoint"`
	}

	PromoFilterInput struct {
		System      *string   `json:"system"`
		MachineKey  *string   `json:"machine_key"`
		ProductCode []*string `json:"product_code"`
	}

	Promo struct {
		ProductCode string  `json:"product_code"`
		Percent     float64 `json:"percent"`
	}

	BaseResponse struct {
		Message string `json:"message"`
	}

	PromoListResponse struct {
		BaseResponse
		Body []Promo `json:"body"`
	}
)
