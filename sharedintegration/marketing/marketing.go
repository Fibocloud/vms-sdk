package marketing

import (
	"encoding/json"
)

func (q *Client) ListPromos(input *PromoFilterInput) (PromoListResponse, error) {
	res, err := q.httpRequest("/thirdparty/integration/promo/active", input)
	var response PromoListResponse

	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}
