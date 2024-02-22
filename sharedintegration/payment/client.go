package payment

import "github.com/spf13/viper"

// InitClient initialize redis client
func Connect() *Client {
	client := Client{
		APIKey:   viper.GetString("PAYMENT_API_KEY"),
		Endpoint: viper.GetString("PAYMENT_API_ENDPOINT"),
	}
	return &client
}
