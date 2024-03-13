package marketing

import "github.com/spf13/viper"

// InitClient initialize redis client
func Connect() *Client {
	client := Client{
		APIKey:   viper.GetString("MARKETING_API_KEY"),
		Endpoint: viper.GetString("MARKETING_API_ENDPOINT"),
	}
	return &client
}
