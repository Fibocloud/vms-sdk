package payment

import (
	"encoding/json"
	"fmt"
)

func (q *Client) CreatePayment(input *PaymentCreateInput) (PaymentCreateResponse, error) {
	res, err := q.httpRequest("/integration/payment/create", input)
	var response PaymentCreateResponse

	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}

func getPaymentCompleteURL(url, uuid string) string {
	return fmt.Sprintf("%s/complete?invoice_uuid=%s", url, uuid)
}

func (q *Client) CompletePayment(url, uuid string) (PaymentCreateResponse, error) {
	res, err := q.httpRequestRaw(getPaymentCompleteURL(url, uuid), nil)
	var response PaymentCreateResponse

	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}

func getPaymentEbarimtURL(url, uuid string) string {
	return fmt.Sprintf("%s/ebarimt?invoice_uuid=%s", url, uuid)
}

func (q *Client) GetEbarimtItems(url, uuid string) (EbarimtStockResponse, error) {
	res, err := q.httpRequestRaw(getPaymentEbarimtURL(url, uuid), nil)
	var response EbarimtStockResponse

	if err != nil {
		return response, err
	}

	json.Unmarshal(res, &response)

	return response, nil
}
